package session

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/wire"
	"gopkg.in/guregu/null.v4"
)

func NewSession(awsRegion string,
	awsProfile string,
	awsAccessKeyID string,
	awsSecretKeyID string,
	awsSessionToken string,
) (*session.Session, error) {
	sess, err := NewAwsSession(
		WithAwsRegion(null.StringFrom(awsRegion)),
		WithAwsProfile(null.StringFrom(awsProfile)),
		WithAwsAccessKey(null.StringFrom(awsAccessKeyID)),
		WithAwsSecretKey(null.StringFrom(awsSecretKeyID)),
		WithAwsSessionToken(null.StringFrom(awsSessionToken)),
	)

	return sess, err
}

// AwsSessionProviderSet AWS Session
var AwsSessionProviderSet = wire.NewSet(
	NewSession,
)
