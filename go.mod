module github.com/egsam98/voting/cli

go 1.16

require (
	github.com/Shopify/sarama v1.29.1
	github.com/egsam98/voting/proto v0.2.1-0.20210704193943-52015c4dbb31
	github.com/pkg/errors v0.9.1
	github.com/urfave/cli v1.22.5
	google.golang.org/protobuf v1.27.1
)

replace github.com/egsam98/voting/proto => github.com/egsam98/voting__proto v0.2.1-0.20210704193943-52015c4dbb31
