package session

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/request"
	"gopkg.in/guregu/null.v4"
	_ "reflect"
	"testing"
)

func TestNewAwsSession(t *testing.T) {
	type args struct {
		opts []Option
	}

	tests := []struct {
		name string
		args args
		//want    *session.Session
		wantErr bool
	}{
		{
			name: "AwsProfile 테스트",
			args: args{
				opts: []Option{
					WithAwsProfile(null.StringFrom("Test")),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAwsSession(tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAwsSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got.Handlers.Send.PushFront(func(r *request.Request) {
				// Log every request made and its payload
				fmt.Printf("Request: %s/%v, Params: %s",
					r.ClientInfo.ServiceName, r.Operation, r.Params)
			})

			fmt.Println(got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewAwsSession() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
