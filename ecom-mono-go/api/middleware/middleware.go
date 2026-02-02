package middleware

import (
	auth_utils "ecom-mono-go/api/auth/utils"
	"ecom-mono-go/api/base"
	"ecom-mono-go/domain/types"
	"ecom-mono-go/infrastructure"
	"ecom-mono-go/utils/apperror"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware interface {
	HandleClient() func(ctx *gin.Context)
	HandleEmailVerification(token string) (*types.ID,error)
	verfiyToken(token string, t auth_utils.TokenType) (jwt.MapClaims,*apperror.AppError)
}

type authMiddleware struct {
	env *infrastructure.Env
	*base.Handler
}

func NewAuthMiddleware(env *infrastructure.Env, base *base.Handler) AuthMiddleware {
	return &authMiddleware{
		env: env,
		Handler: base,
	}
}

func (am *authMiddleware) HandleClient() func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_,err := am.verfiyToken(token, auth_utils.ACCESS_TOKEN)

		if err!=nil{
			am.HandleError(ctx,err)
			return
		}
	}
}

func (am *authMiddleware) HandleEmailVerification(token string) (*types.ID,error) {
	claims,err := am.verfiyToken(token, auth_utils.EMAIL_VERIFICATION_TOKEN)
	if err!=nil{
		return nil,err
	}
	id,ok := claims["id"].(string)
	if !ok {
		return nil, apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid userID"))
	}

	userID:= types.ID(id)
	return &userID,nil
}

func (am *authMiddleware) verfiyToken(token string, t auth_utils.TokenType) (jwt.MapClaims,*apperror.AppError) {
	if token == ""{
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid access token"))
	}
	jwtToken,err := jwt.Parse(token,func(t *jwt.Token) (any, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil,fmt.Errorf("invalid sigining method")
		}
		return []byte(am.env.AUTH_TOKEN_KEY_SECRET),nil
	})

	if err!=nil {
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid access token"))
	}

	if !jwtToken.Valid{
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid access token"))
	}

	claims,ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid access token couldnt extract claims"))
	}

	providedTokenType := claims["type"].(string)
	if auth_utils.TokenType(providedTokenType) != t {
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("invalid token type"))
	}

	exp,err := claims.GetExpirationTime()

	if err!=nil {
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("couldnt extract expiration time"))
	}

	if time.Now().After(exp.Time) {
		return nil,apperror.New(http.StatusUnauthorized, fmt.Errorf("token expired"))
	}

	return claims,nil
}