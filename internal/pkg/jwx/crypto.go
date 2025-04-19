package jwx

import "github.com/go-jose/go-jose/v4"

var (
	supportedSignatureAlgs = []jose.SignatureAlgorithm{
		jose.EdDSA,
		jose.HS256,
		jose.HS384,
		jose.HS512,
		jose.RS256,
		jose.RS384,
		jose.RS512,
		jose.ES256,
		jose.ES384,
		jose.ES512,
		jose.PS256,
		jose.PS384,
		jose.PS512,
	}
)
