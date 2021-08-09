package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func setupJWT() *JWTAuthenticator {
	return &JWTAuthenticator{
		Secret: "-------secretconnectkeytogenerateJWT",
	}
}

func TestJWTCreationAndValidating(t *testing.T) {
	j := setupJWT()
	iat := time.Now().Unix()
	token, err := j.MakeToken(map[string]interface{}{
		"sub":  "userid",
		"name": "Amirreza",
		"iat":  iat,
	})
	assert.NoError(t, err)

	claims, err := j.ClaimsOf(token)
	assert.NoError(t, err)
	mapClaims, ok := claims.(jwt.MapClaims)
	assert.True(t, ok)
	assert.Equal(t, "Amirreza", mapClaims["name"])
	assert.Equal(t, float64(iat), mapClaims["iat"])
	assert.Equal(t, "userid", mapClaims["sub"])
}
