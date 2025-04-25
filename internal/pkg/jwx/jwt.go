package jwx

import (
	"errors"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"reflect"
	"strings"
)

var (
	ErrMissingHeader = errors.New("missing header in token")
)

type Claims struct {
	Registered jwt.Claims
	Custom     map[string]interface{}
}

type Jwt struct {
	Header jose.Header
	Claims Claims
}

func ParseUnsafe(token string) (*Jwt, error) {
	parsed, err := jwt.ParseSigned(token, supportedSignatureAlgs)
	if err != nil {
		return nil, err
	}

	if len(parsed.Headers) == 0 {
		return nil, ErrMissingHeader
	}
	header := parsed.Headers[0]

	registered := jwt.Claims{}
	custom := map[string]interface{}{}

	err = parsed.UnsafeClaimsWithoutVerification(&registered, &custom)
	if err != nil {
		return nil, err
	}

	fields := reflect.VisibleFields(reflect.TypeOf(registered))
	for _, field := range fields {
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		delete(custom, name)
	}

	return &Jwt{
		header,
		Claims{
			Registered: registered,
			Custom:     custom,
		},
	}, nil
}
