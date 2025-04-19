package jwx

import (
	"github.com/go-jose/go-jose/v4/jwt"
	"reflect"
	"strings"
)

type Claims struct {
	Registered jwt.Claims
	Custom     map[string]interface{}
}

func ParseClaimsUnsafe(token string) (*Claims, error) {
	parsed, err := jwt.ParseSigned(token, supportedSignatureAlgs)
	if err != nil {
		return nil, err
	}

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

	return &Claims{
		Registered: registered,
		Custom:     custom,
	}, nil
}
