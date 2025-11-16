module awesomeProject

go 1.25.3

require (
	github.com/ydb-platform/ydb-go-sdk/v3 v3.117.1
	github.com/ydb-platform/ydb-go-yc v0.12.3
)

require (
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jonboulle/clockwork v0.5.0 // indirect
	github.com/yandex-cloud/go-genproto v0.0.0-20240819112322-98a264d392f6 // indirect
	github.com/ydb-platform/ydb-go-genproto v0.0.0-20250911135631-b3beddd517d9 // indirect
	github.com/ydb-platform/ydb-go-yc-metadata v0.6.1 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b // indirect
	google.golang.org/grpc v1.76.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace github.com/ydb-platform/ydb-go-sdk/v3 v3.117.1 => ../ydb-go-sdk

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20250804133106-a7a43d27e69b
