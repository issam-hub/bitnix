package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/interface/api/rest/dto/response"
)

func ToUploadAssetsResponse(assetsResult []*common.AssetResult) *response.UploadAssetsResponse {
	var urls []string

	for _, asset := range assetsResult {
		urls = append(urls, asset.URL)
	}
	return &response.UploadAssetsResponse{
		URLs: urls,
	}
}
