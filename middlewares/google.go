package middlewares

import (
	"github.com/ANPez/gogsi/types"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/oauth2/v2"
	"log"
	"net/http"
)

type Google struct {
	*oauth2.Service
}

func (google *Google) VerifyToken(token string) (*types.User, error) {
	tokenInfo, err := google.Service.Tokeninfo().IdToken(token).Do()
	if nil != err {
		return nil, err
	}

	user := &types.User{
		UserID: tokenInfo.UserId,
		Email:  tokenInfo.Email,
	}
	return user, nil
}

func NewGoogleMiddlewareWithRoundTripper(roundTripper http.RoundTripper) gin.HandlerFunc {
	oauth2Service, err := oauth2.New(&http.Client{Transport: roundTripper})
	if nil != err {
		log.Fatal(err)
	}
	google := &Google{oauth2Service}
	return func(c *gin.Context) {
		c.Set("google", google)
		c.Next()
	}
}

func NewGoogleMiddleware() gin.HandlerFunc {
	return NewGoogleMiddlewareWithRoundTripper(http.DefaultTransport)
}
