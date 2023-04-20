package interceptor

import (
	"context"
	"github.com/hankeyyh/mxshop-srv/user-srv/log"
	"google.golang.org/grpc"
	"time"
)

// LogReqRsp 废弃
func LogReqRsp(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	startTime := time.Now()
	resp, err = handler(ctx, req)

	log.Info(ctx, "", log.Any("server", info.Server),
		log.Any("method", info.FullMethod),
		log.Any("start_time", startTime),
		log.Any("duration", time.Since(startTime)),
		log.Any("request", req),
		log.Any("response", resp),
		log.Any("err", err),
	)
	return resp, err
}
