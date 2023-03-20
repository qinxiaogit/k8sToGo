package jwt

import (
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManage_Generate(t *testing.T) {
	t.Parallel()
	logger := log.New()
	conf := &config.Config{JWT: config.JWT{Secret: "secret", Expires: 3600}}

	jwtManage := NewManage(logger, conf)
	id := uint64(1)
	tokenStr, err := jwtManage.Generate(id)

	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
}

func TestManage_Validate(t *testing.T) {
	t.Parallel()

	logger := log.New()
	conf := &config.Config{JWT: config.JWT{Secret: "secret", Expires: 3600}}

	jwtManage := NewManage(logger, conf)
	id := uint64(2)
	tokenStr, err := jwtManage.Generate(id)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)

	claims, err := jwtManage.Validate(tokenStr)
	require.NoError(t, err)
	require.Equal(t, id, claims.ID)

}
