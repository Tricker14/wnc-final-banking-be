package middleware

import (
	"net/http"

	httpcommon "github.com/VuKhoa23/advanced-web-be/internal/domain/http_common"
	"github.com/VuKhoa23/advanced-web-be/internal/repository"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/env"
	"github.com/VuKhoa23/advanced-web-be/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	customerRepository repository.CustomerRepository
}

func NewAuthMiddleware(customerRepository repository.CustomerRepository) *AuthMiddleware {
	return &AuthMiddleware{customerRepository: customerRepository}
}

func getAccessToken(c *gin.Context) (token string) {
	token = c.Request.Header.Get("Authorization")
	if token == "" {
		var err error
		token, err = c.Cookie("access_token")
		if err != nil {
			return ""
		}
	}

	return token
}

func getRefreshToken(c *gin.Context) (token string) {
	token, err := c.Cookie("refresh_token")
	if err != nil {
		return ""
	}
	return token
}

func (a *AuthMiddleware) VerifyToken(c *gin.Context) {
	jwtSecret, err := env.GetEnv("JWT_SECRET")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: err.Error(),
				Code:    httpcommon.ErrorResponseCode.InternalServerError,
			},
		))
		return
	}

	accessToken := getAccessToken(c)
	claims, err := jwt.VerifyToken(accessToken, jwtSecret)
	if err != nil {
		if err.Error() != httpcommon.ErrorMessage.TokenExpired {
			c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
				httpcommon.Error{
					Message: httpcommon.ErrorMessage.BadCredential,
					Code:    httpcommon.ErrorResponseCode.InvalidRequest,
				},
			))
			return
		} else {
			// if access token expired, check refresh token
			refreshToken := getRefreshToken(c)
			refreshClaims, errRf := jwt.VerifyToken(refreshToken, jwtSecret)
			if errRf != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
					httpcommon.Error{
						Message: httpcommon.ErrorMessage.BadCredential,
						Code:    httpcommon.ErrorResponseCode.InvalidRequest,
					},
				))
				return
			}

			payload, ok := refreshClaims.Payload.(map[string]interface{})
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
					httpcommon.Error{
						Message: httpcommon.ErrorMessage.BadCredential,
						Code:    httpcommon.ErrorResponseCode.InvalidRequest,
					},
				))
				return
			}

			customerId := int64(payload["id"].(float64))

			isValid := a.customerRepository.ValidateRefreshToken(c, customerId, refreshToken)
			if !isValid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
					httpcommon.Error{
						Message: httpcommon.ErrorMessage.BadCredential,
						Code:    httpcommon.ErrorResponseCode.InvalidRequest,
					},
				))
				return
			}
			c.Set("customerId", customerId)
			c.Next()
			return
		}
	}

	payload, ok := claims.Payload.(map[string]interface{})
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, httpcommon.NewErrorResponse(
			httpcommon.Error{
				Message: httpcommon.ErrorMessage.BadCredential,
				Code:    httpcommon.ErrorResponseCode.InvalidRequest,
			},
		))
		return
	}
	customerId := int64(payload["id"].(float64))
	c.Set("customerId", customerId)
	c.Next()
}
