package jwx

import (
	"errors"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

var (
	ErrMissingHeader = errors.New("missing header in token")
)

func GetTokenHeader(token string) (*jose.Header, error) {
	parsed, err := jwt.ParseSigned(token, supportedSignatureAlgs)
	if err != nil {
		return nil, err
	}
	if len(parsed.Headers) == 0 {
		return nil, ErrMissingHeader
	}
	header := parsed.Headers[0]
	return &header, nil
}
