package features

import "github.com/golang-jwt/jwt/v5"

func GenerateJWT() string {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte("secret")
	t = jwt.New(jwt.SigningMethodHS256)
	s, _ = t.SignedString(key)
	_ = s

	return s
}
