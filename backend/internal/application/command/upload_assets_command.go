package command

import "bitnix-backend/internal/application/common"

type UploadAssetsCommand struct {
	Assets []AssetDetail
}

func NewUploadAssetsCommand(assets []AssetDetail) *UploadAssetsCommand {
	return &UploadAssetsCommand{
		Assets: assets,
	}
}

type UploadAssetsCommandResult struct {
	Result []*common.AssetResult
}
