package request

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/validator"
	"fmt"
	"mime/multipart"
)

type UploadAssetsRequest struct {
	Type        string                `form:"type"`
	Filename    string                `form:"filename"`
	Content     *multipart.FileHeader `form:"content"`
	ContentType string                `form:"content_type"`
}

func (req *UploadAssetsRequest) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors

	if req.Type == "" {
		errors = append(errors, validator.ValidationError{Field: "type", Message: "asset type is required"})
	} else {
		if _, err := entities.ParseAssetType(req.Type); err != nil {
			errors = append(errors, validator.ValidationError{Field: "type", Message: err.Error()})
		}
	}

	if req.Filename == "" {
		errors = append(errors, validator.ValidationError{Field: "filename", Message: "filename is required"})
	}

	if req.Content == nil {
		errors = append(errors, validator.ValidationError{Field: "content", Message: "content file is required"})
	}

	if req.Content.Size == 0 {
		errors = append(errors, validator.ValidationError{Field: "content", Message: "content size must not be zero"})
	}

	if req.ContentType != "" && len(req.ContentType) == 0 {
		errors = append(errors, validator.ValidationError{Field: "content_type", Message: "content type is invalid"})
	}

	return errors
}

func (req *UploadAssetsRequest) ToUploadAssetsCommand() (*command.UploadAssetsCommand, error) {
	if ve := req.Validate(); ve.HasErrors() {
		return nil, ve
	}
	file, err := req.Content.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %w", err)
	}

	ctype := req.ContentType
	if ctype == "" && req.Content != nil && req.Content.Header != nil {
		if v := req.Content.Header.Get("Content-Type"); v != "" {
			ctype = v
		}
	}

	assets := []command.AssetDetail{
		{
			Type:        entities.AssetType(req.Type),
			URL:         "",
			Filename:    req.Filename,
			Content:     file,
			ContentType: ctype,
		},
	}

	return command.NewUploadAssetsCommand(assets), nil
}
