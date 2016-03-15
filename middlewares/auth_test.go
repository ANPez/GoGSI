package middlewares

import (
	"github.com/anpez/gogsi/mocks"
	"github.com/anpez/gogsi/types"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthUserNoAuthFail(t *testing.T) {
	Convey("Given a valid Gin environment", t, func() {
		googleMock := new(mocks.Google)
		dbMock := new(mocks.Database)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("google", googleMock)
			c.Set("mongodb_db", dbMock)
			c.Next()
		})

		router.GET("/test", authUser, func(c *gin.Context) {
			c.JSON(200, "test data")
		})

		Convey("and given a valid request", func() {
			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/test", nil)
			So(err, ShouldBeNil)

			Convey("when calling to test function", func() {
				router.ServeHTTP(response, request)

				Convey("should return a forbidden response", func() {
					So(response.Code, ShouldEqual, http.StatusForbidden)
				})
			})
		})
	})
}

func TestAuthUserNoBearerFail(t *testing.T) {
	Convey("Given a valid Gin environment", t, func() {
		googleMock := new(mocks.Google)
		dbMock := new(mocks.Database)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("google", googleMock)
			c.Set("mongodb_db", dbMock)
			c.Next()
		})

		router.GET("/test", authUser, func(c *gin.Context) {
			c.JSON(200, "test data")
		})

		Convey("and given a valid request", func() {
			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/test", nil)
			So(err, ShouldBeNil)
			request.Header.Set("Authorization", "Invalid format header")

			Convey("when calling to test function", func() {
				router.ServeHTTP(response, request)

				Convey("should return a forbidden response", func() {
					So(response.Code, ShouldEqual, http.StatusForbidden)
				})
			})
		})
	})
}

func TestAuthUserInvalidToken(t *testing.T) {
	Convey("Given a valid Gin environment", t, func() {
		googleMock := new(mocks.Google)
		googleMock.On("VerifyToken", mock.Anything).Once().Return(nil, fmt.Errorf("Some error"))
		dbMock := new(mocks.Database)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("google", googleMock)
			c.Set("mongodb_db", dbMock)
			c.Next()
		})

		router.GET("/test", authUser, func(c *gin.Context) {
			c.JSON(200, "test data")
		})

		Convey("and given a valid request", func() {
			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/test", nil)
			So(err, ShouldBeNil)
			request.Header.Set("Authorization", "Bearer InvalidToken")

			Convey("when calling to test function", func() {
				router.ServeHTTP(response, request)

				Convey("should return a forbidden response", func() {
					So(response.Code, ShouldEqual, http.StatusForbidden)
				})
			})
		})
	})
}

func TestAuthUserNewUserOk(t *testing.T) {
	const EMAIL string = "test@gmail.com"
	const USER_ID string = "a user id"
	const GOOGLE_TOKEN string = "a google token"

	Convey("Given a valid Gin environment", t, func() {
		user := &types.User{UserID: USER_ID, Email: EMAIL}

		googleMock := new(mocks.Google)
		googleMock.On("VerifyToken", GOOGLE_TOKEN).Once().Return(user, nil)

		dbCollMock := new(mocks.DBCollection)
		dbCollMock.On("FindOne", mock.Anything, mock.Anything).Return(false, nil)
		dbCollMock.On("Insert", user).Run(func(args mock.Arguments) {
			assert.IsType(t, new(types.User), args.Get(0), "Inserting not an User struct")
			u := args.Get(0).(*types.User)
			assert.Equal(t, "g:"+USER_ID, u.UserID)
			assert.Equal(t, EMAIL, u.Email)
		}).Once().Return(nil)
		dbMock := new(mocks.Database)
		dbMock.On("C", "users").Return(dbCollMock)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("google", googleMock)
			c.Set("mongodb_db", dbMock)
			c.Next()
		})

		router.GET("/test", authUser, func(c *gin.Context) {
			user := c.MustGet("user").(*types.User)
			c.JSON(200, map[string]interface{}{"email": user.Email})
		})

		Convey("and given a valid request", func() {
			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/test", nil)
			So(err, ShouldBeNil)
			request.Header.Set("Authorization", "Bearer "+GOOGLE_TOKEN)

			Convey("when calling to test function", func() {
				router.ServeHTTP(response, request)

				Convey("should return an ok response", func() {
					So(response.Code, ShouldEqual, http.StatusOK)

					var ret map[string]interface{}
					So(json.Unmarshal(response.Body.Bytes(), &ret), ShouldBeNil)

					So(ret, ShouldNotBeNil)
					So(ret, ShouldHaveLength, 1)
					So(ret, ShouldContainKey, "email")
					So(ret["email"], ShouldEqual, EMAIL)
				})
			})
		})

		dbMock.AssertExpectations(t)
		dbCollMock.AssertExpectations(t)
	})
}

func TestAuthUserReturningUserOk(t *testing.T) {
	const EMAIL string = "test@gmail.com"
	const USER_ID string = "a user id"
	const GOOGLE_TOKEN string = "a google token"

	Convey("Given a valid Gin environment", t, func() {
		user := &types.User{UserID: USER_ID, Email: EMAIL}

		googleMock := new(mocks.Google)
		googleMock.On("VerifyToken", GOOGLE_TOKEN).Once().Return(user, nil)

		dbCollMock := new(mocks.DBCollection)
		dbCollMock.On("FindOne", mock.Anything, mock.AnythingOfType("*types.User")).Run(func(args mock.Arguments) {
			u := args.Get(1).(*types.User)
			u.UserID = USER_ID
			u.Email = EMAIL
		}).Return(true, nil)
		dbCollMock.AssertNotCalled(t, "Insert", mock.Anything)
		dbMock := new(mocks.Database)
		dbMock.On("C", "users").Return(dbCollMock)

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("google", googleMock)
			c.Set("mongodb_db", dbMock)
			c.Next()
		})

		router.GET("/test", authUser, func(c *gin.Context) {
			user := c.MustGet("user").(*types.User)
			c.JSON(200, map[string]interface{}{"email": user.Email})
		})

		Convey("and given a valid request", func() {
			response := httptest.NewRecorder()
			request, err := http.NewRequest("GET", "/test", nil)
			So(err, ShouldBeNil)
			request.Header.Set("Authorization", "Bearer "+GOOGLE_TOKEN)

			Convey("when calling to test function", func() {
				router.ServeHTTP(response, request)

				Convey("should return an ok response", func() {
					So(response.Code, ShouldEqual, http.StatusOK)

					var ret map[string]interface{}
					So(json.Unmarshal(response.Body.Bytes(), &ret), ShouldBeNil)

					So(ret, ShouldNotBeNil)
					So(ret, ShouldHaveLength, 1)
					So(ret, ShouldContainKey, "email")
					So(ret["email"], ShouldEqual, EMAIL)
				})
			})
		})

		dbMock.AssertExpectations(t)
		dbCollMock.AssertExpectations(t)
	})
}
