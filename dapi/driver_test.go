package dapi

import (
	"database/sql"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewOpen(t *testing.T) {

	var sess *session.Session
	var err error

	if v, ok := os.LookupEnv("AWS_PROFILE"); ok {
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewSharedCredentials("", v),
		})
	} else {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				os.Getenv("AWS_SESSION_TOKEN"),
			),
		})
	}

	require.NoError(t, err)

	opts := []Option{
		WithDataServiceApiOption(rdsdataservice.New(sess)),
		WithDatabase(os.Getenv("DB_NAME")),
		WithDbEngine(DbEnginePostgres),
		WithResourceArn(os.Getenv("AWS_RESOURCE_ARN")),
		WithSecretArn(os.Getenv("AWS_SECRET_ARN")),
		WithSchema("public"),
		WithContinueAfterTimeout(true),
	}

	type args struct {
		opts []Option
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		{
			name:    "OpenTest",
			args:    args{opts: opts},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewOpen(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOpen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewOpen() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
