module github.com/conductorone/conductorone-sdk-go/v2

go 1.21

toolchain go1.21.3

require (
	github.com/cenkalti/backoff/v4 v4.2.1
	github.com/conductorone/conductorone-sdk-go v1.14.0
	github.com/ericlagergren/decimal v0.0.0-20221120152707-495c53812d05
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/pquerna/xjwt v0.2.0
	go.uber.org/zap v1.26.0
	golang.org/x/net v0.19.0
	golang.org/x/oauth2 v0.15.0
	google.golang.org/grpc v1.59.0
	gopkg.in/square/go-jose.v2 v2.6.0
)

require (
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

retract [v2.1.8, v2.1.11] // Accidentally bumped the major number. v1.X.X is the correct major version, please use that instead (go get github.com/conductorone/conductorone-sdk-go)
