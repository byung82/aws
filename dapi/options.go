package dapi

import (
	"github.com/aws/aws-sdk-go/service/rdsdataservice/rdsdataserviceiface"
	"gopkg.in/guregu/null.v4"
)

type Option interface {
	apply(*options)
}

type options struct {
	dataServiceApi       rdsdataserviceiface.RDSDataServiceAPI
	database             null.String
	resourceArn          null.String
	secretArn            null.String
	schema               null.String
	dbEngine             DbEngine
	continueAfterTimeout null.Bool
}

type dataServiceApiOption struct {
	DataServiceApi rdsdataserviceiface.RDSDataServiceAPI
}

type databaseOption string
type resourceArnOption string
type secretArnOption string
type schemaOption string
type dbEngineOption DbEngine
type continueAfterTimeoutOption bool

func (c dataServiceApiOption) apply(opts *options) {
	opts.dataServiceApi = c.DataServiceApi
}

func (c databaseOption) apply(opts *options) {
	opts.database = null.StringFrom(string(c))
}

func (c resourceArnOption) apply(opts *options) {
	opts.resourceArn = null.StringFrom(string(c))
}

func (c secretArnOption) apply(opts *options) {
	opts.secretArn = null.StringFrom(string(c))
}

func (c schemaOption) apply(opts *options) {
	opts.schema = null.StringFrom(string(c))
}

func (c dbEngineOption) apply(opts *options) {
	opts.dbEngine = DbEngine(c)
}

func (c continueAfterTimeoutOption) apply(opts *options) {
	opts.continueAfterTimeout = null.BoolFrom(bool(c))
}

func WithDataServiceApiOption(dataServiceApi rdsdataserviceiface.RDSDataServiceAPI) Option {
	return dataServiceApiOption{
		DataServiceApi: dataServiceApi,
	}
}

func WithDatabase(c string) Option {
	return databaseOption(c)
}

func WithResourceArn(c string) Option {
	return resourceArnOption(c)
}

func WithSecretArn(c string) Option {
	return secretArnOption(c)
}

func WithSchema(c string) Option {
	return schemaOption(c)
}

func WithDbEngine(c DbEngine) Option {
	return dbEngineOption(c)
}

func WithContinueAfterTimeout(c bool) Option {
	return continueAfterTimeoutOption(c)
}
