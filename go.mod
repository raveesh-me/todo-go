module github.com/raveesh-me/todo

go 1.24.0

require (
	connectrpc.com/connect v1.19.0
	github.com/google/uuid v1.6.0
	github.com/rs/zerolog v1.34.0
	google.golang.org/grpc v1.75.1
	google.golang.org/protobuf v1.36.9
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
)

replace github.com/raveesh-me/todo => ./
