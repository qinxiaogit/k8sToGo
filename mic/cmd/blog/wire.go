//go:build wireinject
// +build wireinject
package main

import (
	"github.com/google/wire"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/blog/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/auth"
	"github.com/qinxiaogit/k8sToGo/mic/internal/blog"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/qinxiaogit/k8sToGo/mic/internal/user"
)

func InitServer(logger *log.Logger, conf *config.Config) (v1.BlogServiceServer, error) {
	wire.Build(
		auth.NewClient,
		blog.NewServer,
		user.NewClient,
	)
	return &blog.Server{}, nil
}
