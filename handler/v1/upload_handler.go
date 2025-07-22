package v1

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"
)

type UploadHandler struct {
	BaseHandler
	S3Client *s3.Client
	Bucket   string
}

func NewUploadHandler(e *echo.Echo, bucket string) (*UploadHandler, error) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(cfg)
	h := &UploadHandler{
		S3Client: s3Client,
		Bucket:   bucket,
	}
	e.POST("/api/v1/upload", h.Upload)
	return h, nil
}

func (h *UploadHandler) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return h.NewResponseWithError(c, "missing file", err)
	}
	src, err := file.Open()
	if err != nil {
		return h.NewResponseWithError(c, "open file failed", err)
	}
	defer src.Close()

	key := file.Filename
	_, err = h.S3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(h.Bucket),
		Key:    aws.String(key),
		Body:   src,
	})
	if err != nil {
		return h.NewResponseWithError(c, fmt.Sprintf("upload failed: %v", err), err)
	}
	return h.NewResponseWithData(c, map[string]string{"message": "upload success", "key": key})
}
