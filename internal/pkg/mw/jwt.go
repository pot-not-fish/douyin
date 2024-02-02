package mw

import (
	"context"
	"douyin/internal/pkg/kitex_client"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	webtoken "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "douyin"
)

type LoginReq struct {
	Username string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 12); msg:'Illegal format'"`
	Password string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 12); msg:'Illegal format'"`
}

type Payload struct {
	Name    string
	User_id int64
}

func InitJwt() {
	var err error

	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("ZhaZrDBQb7MYdJWaPf5gJmGbYyVjLYgz"),
		Timeout:     time.Hour * 24 * 7,
		IdentityKey: IdentityKey,
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginReq LoginReq
			if err := c.BindAndValidate(&loginReq); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userActionResp, err := kitex_client.UserActionRpc(ctx, 2, loginReq.Username, loginReq.Password)
			if err != nil {
				return "", err
			}

			// 后续相应会用到user_id，所以要存入上下文
			c.Set("user_id", userActionResp.User.Id)

			return &Payload{
				Name:    loginReq.Username,
				User_id: userActionResp.User.Id,
			}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Payload); ok {
				return jwt.MapClaims{
					"name":    v.Name,
					"user_id": v.User_id,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)

			return &Payload{
				Name:    claims["name"].(string),
				User_id: int64(claims["user_id"].(float64)),
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			var status_msg string
			var status_code int
			row_user_id, ok := c.Get("user_id")
			if !ok {
				status_code = 1
				status_msg = "user_id is empty"
			}

			user_id, ok := row_user_id.(uint)
			if !ok {
				status_code = 1
				status_msg = "user_id not uint"
			}

			status_msg = "Welcome"
			c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": status_code,
				"status_msg":  status_msg,
				"user_id":     user_id,
				"token":       token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": 1,
				"status_msg":  message,
			})
		},
		TokenLookup: "query: token, form: token",
	})

	if err != nil {
		panic(err)
	}
}

/**
 * @function
 * @description 用于Middlewarefunc后，也就是必须要token的场景，保存user_id
 * @param
 * @return
 */
func SaveUserId(ctx context.Context, c *app.RequestContext) {
	user_id, err := getUserId(ctx, c)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 1,
			"status_msg":  err.Error(),
		})
		return
	}

	c.Set("user_id", user_id)

	c.Next(ctx)
}

func getUserId(ctx context.Context, c *app.RequestContext) (int64, error) {
	value, ok := c.Get("JWT_PAYLOAD") // jwt.MapClaims
	if !ok {
		return 0, fmt.Errorf("could not find JWT_PAYLOAD")
	}

	claims, ok := value.(jwt.MapClaims) // 在这里不能直接用map[string]interface{}
	if !ok {
		return 0, fmt.Errorf("could not find claims")
	}

	user_id, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("could not find user_id")
	}

	return int64(user_id), nil
}

/**
 * @function
 * @description 用于user id可选的场景
 * @param
 * @return
 */
func TokenGetUserId(token *string) (int64, error) {
	if token == nil {
		return 0, nil
	}

	jwt_token, err := JwtMiddleware.ParseTokenString(*token)
	if err != nil {
		return 0, err
	}

	claims := jwt_token.Claims.(webtoken.MapClaims)

	user_id := claims["user_id"].(float64)

	return int64(user_id), nil
}
