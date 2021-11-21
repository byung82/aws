package session

import "gopkg.in/guregu/null.v4"

type options struct {
	awsRegion       null.String
	awsProfile      null.String
	awsAccessKey    null.String
	awsSecretKey    null.String
	awsSessionToken null.String
}

type Option interface {
	apply(*options)
}

type awsRegionOption null.String
type awsProfileOption null.String
type awsAccessKeyOption null.String
type awsSecretKeyOption null.String
type awsSessionTokenOption null.String

func (c awsRegionOption) apply(opts *options) {
	opts.awsRegion = null.String(c)
}

func (c awsProfileOption) apply(opts *options) {
	opts.awsProfile = null.String(c)
}

func (c awsAccessKeyOption) apply(opts *options) {
	opts.awsAccessKey = null.String(c)
}

func (c awsSecretKeyOption) apply(opts *options) {
	opts.awsSecretKey = null.String(c)
}

func (c awsSessionTokenOption) apply(opts *options) {
	opts.awsSessionToken = null.String(c)
}

//goland:noinspection GoUnusedExportedFunction
func WithAwsRegion(c null.String) Option {
	return awsRegionOption(c)
}

func WithAwsProfile(c null.String) Option {
	return awsProfileOption(c)
}

//goland:noinspection GoUnusedExportedFunction
func WithAwsAccessKey(c null.String) Option {
	return awsAccessKeyOption(c)
}

//goland:noinspection GoUnusedExportedFunction
func WithAwsSecretKey(c null.String) Option {
	return awsSecretKeyOption(c)
}

//goland:noinspection GoUnusedExportedFunction
func WithAwsSessionToken(c null.String) Option {
	return awsSessionTokenOption(c)
}
