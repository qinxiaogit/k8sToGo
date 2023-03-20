package auth

import (
	"context"
	v1 "github.com/qinxiaogit/k8sToGo/mic/api/protobuf/auth/v1"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/jwt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewServer(t *testing.T) {
	logger := log.New()
	path := config.GetPath()
	conf, err := config.Load(path)
	require.NoError(t, err)
	jwtManager := jwt.NewManage(logger, conf)
	server := NewServer(logger, jwtManager)

	genResp, err := server.GenerateToken(context.Background(), &v1.GenerateTokenRequest{
		UserId: 1,
	})
	if err != nil {
		return
	}
	require.NoError(t, err)
	require.NotEmpty(t, genResp.GetToken())

	validateResp, err := server.ValidateToken(context.Background(), &v1.ValidateTokenRequest{
		Token: genResp.GetToken(),
	})
	require.NoError(t, err)
	require.True(t, validateResp.GetValid())

	refreshResp, err := server.RefreshToken(context.Background(), &v1.RefreshTokenRequest{
		Token: genResp.GetToken(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, refreshResp.GetToken())

	refreshResp, err = server.RefreshToken(context.Background(), &v1.RefreshTokenRequest{
		Token: "a.b.c",
	})
	require.Error(t, err)
	require.Nil(t, refreshResp)
}
