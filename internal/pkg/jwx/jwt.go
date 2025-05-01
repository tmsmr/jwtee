package jwx

import (
	"errors"
	"fmt"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"reflect"
	"strconv"
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

func (c Claims) TableDataRegistered() [][]string {
	var data [][]string

	data = append(data, []string{"Registered claim", "Value"})
	if c.Registered.Issuer != "" {
		data = append(data, []string{"Issuer", c.Registered.Issuer})
	}
	if c.Registered.Subject != "" {
		data = append(data, []string{"Subject", c.Registered.Subject})
	}
	if len(c.Registered.Audience) > 0 {
		data = append(data, []string{"Audience", strings.Join(c.Registered.Audience, ", ")})
	}
	if c.Registered.Expiry != nil {
		data = append(data, []string{"Expiry", c.Registered.Expiry.Time().String()})
	}
	if c.Registered.NotBefore != nil {
		data = append(data, []string{"NotBefore", c.Registered.NotBefore.Time().String()})
	}
	if c.Registered.IssuedAt != nil {
		data = append(data, []string{"IssuedAt", c.Registered.IssuedAt.Time().String()})
	}
	if c.Registered.ID != "" {
		data = append(data, []string{"ID", c.Registered.ID})
	}
	/*
		Issuer    string       `json:"iss,omitempty"`
		Subject   string       `json:"sub,omitempty"`
		Audience  Audience     `json:"aud,omitempty"`
		Expiry    *NumericDate `json:"exp,omitempty"`
		NotBefore *NumericDate `json:"nbf,omitempty"`
		IssuedAt  *NumericDate `json:"iat,omitempty"`
		ID        string       `json:"jti,omitempty"`
	*/

	return data
}

func (c Claims) TableDataCustom() [][]string {
	var data [][]string

	data = append(data, []string{"Custom claim", "Value"})
	for k, v := range c.Custom {
		switch v := v.(type) {
		case string:
			data = append(data, []string{k, v})
		case int:
			data = append(data, []string{k, strconv.Itoa(v)})
		case float64:
			data = append(data, []string{k, strconv.FormatFloat(v, 'f', -1, 64)})
		case bool:
			data = append(data, []string{k, strconv.FormatBool(v)})
		default:
			data = append(data, []string{k, fmt.Sprintf("%v", v)})
		}
	}

	return data
}
