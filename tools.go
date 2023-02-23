//go:build tools

//go:generate sh -c "GOBIN=$PWD/.gobincache go install $(sed -n 's/.*_ \"\\(.*\\)\".*/\\1/p' <$GOFILE)"
package tools

import (
	_ "ariga.io/entimport/cmd/entimport"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/cheekybits/genny"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/golang/mock/mockgen"

	_ "github.com/cosmtrek/air"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/mattn/goveralls"
	_ "github.com/powerman/dockerize"
	_ "golang.org/x/tools/cmd/stringer"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
