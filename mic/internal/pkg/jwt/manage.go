package jwt

import (
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/config"
	"google.golang.org/grpc/codes"
	"time"
	"github.com/qinxiaogit/k8sToGo/mic/internal/pkg/log"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/status"
)

type Manage struct {
	secret  string
	expires time.Duration
	logger  *log.Logger
}

func NewManage(logger *log.Logger, config *config.Config) *Manage {
	return &Manage{secret: config.JWT.Secret, expires: config.JWT.Expires, logger: logger}
}

type UserClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}
//生成token
func (m *Manage) Generate(id uint64) (string, error) {
	claims := &UserClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(m.expires).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.secret))
}
//验证token
func (m *Manage) Validate(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	return claims, nil
}
