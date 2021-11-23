package dapi

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/google/wire"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func NewAwsDApi(sess *session.Session,
	dbName string,
	dbEngine DbEngine,
	awsResourceArn string,
	awsSecretArn string,
) (db.Session, error) {
	sqlDb, err := NewOpen(
		WithDataServiceApi(rdsdataservice.New(sess)),
		WithDatabase(dbName),
		WithDbEngine(dbEngine),
		WithResourceArn(awsResourceArn),
		WithSecretArn(awsSecretArn),
	)

	if err != nil {
		return nil, err
	}

	return postgresql.New(sqlDb)
}

// AwsDApiProviderSet AWS DAPI Provider Set
var AwsDApiProviderSet = wire.NewSet(
	NewAwsDApi,
)
