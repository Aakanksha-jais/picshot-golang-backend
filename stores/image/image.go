package image

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type image struct{}

func New() image {
	return image{}
}

func (i image) Upload(c *app.Context, fileHeader *multipart.FileHeader, name string) error {
	size := fileHeader.Size

	buffer := make([]byte, size)

	file, err := fileHeader.Open()
	if err != nil {
		return errors.DBError{Err: err}
	}

	_, err = file.Read(buffer)
	if err != nil {
		return errors.DBError{Err: err}
	}

	_, err = c.S3.Service().PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(c.Config.Get("AWS_BUCKET")),
		Key:                  aws.String(name),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}
