package request

import "mime/multipart"

type UploadAssetsRequest struct {
    Type        string                  `form:"type"`
    Filename    string                  `form:"filename"`
    Content     *multipart.FileHeader   `form:"content"`
    ContentType string                  `form:"content_type"`
}
