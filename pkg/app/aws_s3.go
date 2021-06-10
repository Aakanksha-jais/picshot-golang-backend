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
	Session() *session.Session
}

type awsS3 struct {
	*s3.S3
	session *session.Session
	config  *aws.Config
}

func (a awsS3) Service() *s3.S3 {
	return a.S3
}

func (a awsS3) Session() *session.Session {
	return a.session
}

func GetNewS3(logger log.Logger, config configs.Config) (AWSS3, error) {
	awsConfigs := &aws.Config{Region: aws.String(config.GetOrDefault("AWS_REGION", "ap-south-1")), Logger: logger}

	if config.Get("AWS_LOG_LEVEL") == log.DEBUG.String() {
		awsConfigs.WithLogLevel(aws.LogDebug)
	}

	sess, err := session.NewSession(awsConfigs)
	if err != nil {
		logger.Errorf("cannot create aws session: %s", err.Error())

		return nil, err
	}

	svc := s3.New(sess)

	return awsS3{S3: svc, session: sess, config: awsConfigs}, nil
}
