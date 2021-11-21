package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewAwsSession 신규 Aws Session 생성
func NewAwsSession(
	opts ...Option,
) (*session.Session, error) {
	options := options{}

	awsRegion := "us-east-1"

	for _, o := range opts {
		o.apply(&options)
	}

	if options.awsRegion.Valid {
		awsRegion = options.awsRegion.String
	}

	if options.awsProfile.Valid {
		return session.NewSession(&aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewSharedCredentials("", options.awsProfile.String),
		})
	} else if options.awsAccessKey.Valid && options.awsSecretKey.Valid {
		return session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(options.awsAccessKey.String,
				options.awsSecretKey.String,
				options.awsSessionToken.String),
		})
	} else {
		return session.NewSession(&aws.Config{
			Region: aws.String(awsRegion),
		})
	}
}
