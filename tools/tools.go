// +build tools

package tools

// tool dependencies
import (
	_ "github.com/CircleCI-Public/circleci-cli"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/frapposelli/wwhrd"
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/sqs/goreturns"
	_ "github.com/srikrsna/protoc-gen-mock"
	_ "github.com/uber/prototool/cmd/prototool"
	_ "go.zenithar.org/protoc-gen-cobra"
	_ "gotest.tools/gotestsum"
	_ "mvdan.cc/gofumpt"
)

