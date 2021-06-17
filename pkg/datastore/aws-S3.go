package datastore

import (
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/configs"
	"github.com/Aakanksha-jais/picshot-golang-backend/pkg/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSS3 struct {
	*s3.S3
	Session *session.Session
	config  *aws.Config
}

func GetNewS3(logger log.Logger, config configs.Config) (AWSS3, error) {
	awsConfigs := &aws.Config{Region: aws.String(config.GetOrDefault("AWS_REGION", "ap-south-1")), Logger: logger}

	if config.Get("AWS_LOG_LEVEL") == log.DEBUG.String() {
		awsConfigs.WithLogLevel(aws.LogDebug)
	}

	sess, err := session.NewSession(awsConfigs)
	if err != nil {
		logger.Errorf("cannot create aws Session: %s", err.Error())

		return AWSS3{}, err
	}

	svc := s3.New(sess)

	return AWSS3{S3: svc, Session: sess, config: awsConfigs}, nil
}
