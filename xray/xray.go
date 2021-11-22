package xray

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-xray-sdk-go/header"
	"github.com/aws/aws-xray-sdk-go/xray"
	"strings"
)

func getTraceHeaderFromContext(ctx context.Context) *header.Header {
	var traceHeader string

	if traceHeaderValue := ctx.Value(xray.LambdaTraceHeaderKey); traceHeaderValue != nil {
		traceHeader = traceHeaderValue.(string)
		return header.FromString(traceHeader)
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func NewSegment(ctx context.Context, name string, fn func(context.Context) error) (err error) {
	var funcName string

	lc, _ := lambdacontext.FromContext(ctx)

	names := strings.Split(lc.InvokedFunctionArn, ":")
	funcName = names[len(names)-1]

	h := getTraceHeaderFromContext(ctx)

	ctx, seg := xray.BeginFacadeSegment(ctx, name, h)

	if h.TraceID != "" {
		seg.TraceID = h.TraceID
	}
	if h.ParentID != "" {
		seg.ParentID = h.ParentID
	}

	seg.IncomingHeader = h
	seg.RequestWasTraced = true

	seg.Lock()

	seg.GetHTTP().GetRequest().Method = "POST"
	seg.GetHTTP().GetRequest().URL = fmt.Sprintf("//%s", funcName)

	seg.Unlock()

	err = fn(ctx)

	defer func(err error) {
		if err != nil {
			seg.GetHTTP().GetResponse().Status = 500
		} else {
			seg.GetHTTP().GetResponse().Status = 200
		}

		seg.Close(err)
	}(err)

	return err
}
