// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/qinxiaogit/k8sToGo/mic/api/protobuf/blog/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/auth"
	"github.com/qinxiaogit/k8sToGo/mic/internal/blog"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/qinxiaogit/k8sToGo/mic/internal/user"
)

// Injectors from wire.go:

func InitServer(logger *log.Logger, conf *config.Config) (v1.BlogServiceServer, error) {
	authServiceClient, err := auth.NewClient(logger, conf)
	if err != nil {
		return nil, err
	}
	userServiceClient, err := user.NewClient(logger, conf)
	if err != nil {
		return nil, err
	}
	blogServiceServer := blog.NewServer(logger, conf, authServiceClient, userServiceClient)
	return blogServiceServer, nil
}