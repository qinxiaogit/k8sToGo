//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/user/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/dbcontext"
	"github.com/qinxiaogit/k8sToGo/mic/internal/user"
)

func InitServer(logger *log.Logger, conf *config.Config) (v1.UserServiceServer, error) {
	wire.Build(
		dbcontext.NewUserDB,
		user.NewRepository,
		user.NewServer,
	)
	return &user.Server{}, nil
}
