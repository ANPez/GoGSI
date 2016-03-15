package middlewares

import (
	"github.com/ANPez/gogsi/interfaces"
	"github.com/ANPez/gogsi/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Auth handles the user authorization in the application.
func Auth() gin.HandlerFunc {
	return authUser
}

func authUser(c *gin.Context) {
	google := c.MustGet("google").(interfaces.Google)
	db := c.MustGet("mongodb_db").(interfaces.Database)

	auth := c.Request.Header.Get("Authorization")
	if "" == auth {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	auth = strings.TrimPrefix(auth, "Bearer ")

	user, err := google.VerifyToken(auth)
	if nil != err {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	var dbUser types.User
	uid := fmt.Sprintf("g:%s", user.UserID)
	userFound, err := db.C("users").FindOne(map[string]interface{}{"_id": uid}, &dbUser)
	if nil != err {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if userFound {
		user = &dbUser
	} else {
		user.UserID = uid
		if err = db.C("users").Insert(user); nil != err {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	c.Set("user", user)

	c.Next()
}
