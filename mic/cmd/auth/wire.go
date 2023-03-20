//go:build wireinject
// +build wireinject
package main

import (
	"github.com/google/wire"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/auth/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/auth"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/jwt"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
)
//初始化服务
func InitServer(logger *log.Logger,conf *config.Config)(v1.AuthServiceServer,error){
	wire.Build(
		jwt.NewManage,
		auth.NewServer,
	)
	return &auth.Server{},nil
}
