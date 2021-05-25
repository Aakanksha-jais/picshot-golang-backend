package image

import (
	"bytes"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/request"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/errors"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/app"

	"github.com/aws/aws-sdk-go/aws"
)

type image struct{}

func New() image {
	return image{}
}

func (i image) Upload(ctx *app.Context, fileHeader *multipart.FileHeader, name string) error {
	svc := ctx.S3.Service()

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

	input := &s3.PutObjectInput{
		Bucket:        aws.String(ctx.Get("AWS_BUCKET")),
		Key:           aws.String(name),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	}

	if err := input.Validate(); err != nil {
		err := err.(request.ErrInvalidParams)

		return errors.DBError{Err: err.OrigErr()}
	}

	_, err = svc.PutObjectWithContext(ctx, input)
	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}

func (i image) DeleteBulk(ctx *app.Context, names []string) error {
	svc := ctx.S3.Service()

	objects := make([]*s3.ObjectIdentifier, 0)

	for _, name := range names {
		objects = append(objects, &s3.ObjectIdentifier{Key: aws.String(name)})
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(ctx.Get("AWS_BUCKET")),
		Delete: &s3.Delete{Objects: objects},
	}

	if err := input.Validate(); err != nil {
		err := err.(request.ErrInvalidParams)

		return errors.DBError{Err: err.OrigErr()}
	}

	_, err := svc.DeleteObjectsWithContext(ctx, input)
	if err != nil {
		return errors.DBError{Err: err}
	}

	return nil
}
