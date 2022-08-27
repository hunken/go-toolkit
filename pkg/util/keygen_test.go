package utilpackage

import (
	"github.com/golang-jwt/jwt"
	"testing"
)
util

import (
"github.com/dgrijalva/jwt-go"
"github.com/stretchr/testify/assert"
"testing"
)

func TestGenerateRSAKey(t *testing.T) {
	t.Parallel()
	var tests = []int{64, 128, 256, 512, 1024, 2048, 4096}
	for _, keySize := range tests {
		pubKeyStr, prvKeyStr, err := GenerateRSAKey(keySize)
		if err != nil {
			t.Error(err)
		}
		prvKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(prvKeyStr))
		assert.Nil(t, err)
		assert.Equal(t, keySize, prvKey.Size()*8)
		pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKeyStr))
		assert.Nil(t, err)
		assert.Equal(t, keySize, pubKey.Size()*8)
	}
}

func TestGenerateRSAKey_Failure(t *testing.T) {
	t.Parallel()
	var tests = []int{0, 1, 2, 8, 11}
	for _, keySize := range tests {
		pubKeyStr, prvKeyStr, err := GenerateRSAKey(keySize)
		assert.NotNil(t, err)
		assert.Equal(t, "", pubKeyStr)
		assert.Equal(t, "", prvKeyStr)
	}
}
