package auth

import (
	ssov1 "github.com/kms-qwe/protos/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
