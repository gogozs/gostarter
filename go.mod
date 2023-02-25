module github.com/gogozs/gostarter

go 1.14

require (
	github.com/gogozs/zlib v0.0.2-0.20221218025126-7c551299d8f9
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/spf13/viper v1.15.0
	go.uber.org/zap v1.24.0
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

replace github.com/gogozs/zlib => ../zlib
