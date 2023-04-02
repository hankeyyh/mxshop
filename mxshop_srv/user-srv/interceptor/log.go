package interceptor

import (
	"context"
	"github.com/hankeyyh/mxshop_user_srv/logger"
	"google.golang.org/grpc"
	"time"
)

// LogReqRsp 废弃
func LogReqRsp(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	startTime := time.Now()
	resp, err = handler(ctx, req)

	logger.Info(ctx, "", logger.Any("server", info.Server),
		logger.Any("method", info.FullMethod),
		logger.Any("start_time", startTime),
		logger.Any("duration", time.Since(startTime)),
		logger.Any("request", req),
		logger.Any("response", resp),
		logger.Any("err", err),
	)
	return resp, err
}
