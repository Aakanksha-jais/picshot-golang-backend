package app

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSS3 interface {
	Service() *s3.S3
}

type awsS3 struct {
	*s3.S3
	config *S3Config
}

func (a awsS3) Service() *s3.S3 {
	return a.S3
}

type S3Config struct {
	Region string
}

func GetNewS3(logger log.Logger, config configs.Config) (AWSS3, error) {
	awsConfigs := &aws.Config{Region: aws.String("ap-south-1"), Logger: logger}
	awsConfigs.WithLogLevel(aws.LogDebug)

	sess, err := session.NewSession(awsConfigs)
	if err != nil {
		logger.Errorf("cannot create aws session: %s", err.Error())

		return nil, err
	}

	svc := s3.New(sess)

	return awsS3{S3: svc}, nil
}
