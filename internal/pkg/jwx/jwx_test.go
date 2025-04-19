package jwx_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tmsmr/jwtee/internal/pkg/jwx"
	"testing"
	"time"
)

var _ = Describe("jwx pkg:", func() {

	var signer jose.Signer
	var claims jwx.Claims

	BeforeEach(func() {
		signingKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		signer, _ = jose.NewSigner(jose.SigningKey{Algorithm: jose.ES384, Key: signingKey}, nil)
		claims = jwx.Claims{
			Registered: jwt.Claims{
				Issuer:    faker.DomainName(),
				Subject:   faker.Username(),
				Audience:  jwt.Audience{faker.DomainName()},
				Expiry:    jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
				NotBefore: jwt.NewNumericDate(time.Now().Add(-1 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ID:        faker.UUIDHyphenated(),
			},
			Custom: map[string]interface{}{
				"name":  faker.Name(),
				"email": faker.Email(),
			},
		}
	})

	Context("ParseClaimsUnsafe", func() {
		It("should extract registered claims correctly", func() {
			data, _ := json.Marshal(claims.Custom)
			obj, _ := signer.Sign(data)
			token, _ := obj.CompactSerialize()

			res, err := jwx.ParseClaimsUnsafe(token)
			Expect(err).ToNot(HaveOccurred())
			b := res.Custom
			Expect(b).To(Equal(claims.Custom))
		})

		It("should extract custom claims correctly", func() {
			data, _ := json.Marshal(claims.Registered)
			obj, _ := signer.Sign(data)
			token, _ := obj.CompactSerialize()

			res, err := jwx.ParseClaimsUnsafe(token)
			Expect(err).ToNot(HaveOccurred())
			b := res.Registered
			Expect(b).To(Equal(claims.Registered))
		})

		It("should filter custom claims to not contain registered claims", func() {
			registered, _ := json.Marshal(claims.Registered)
			all := map[string]interface{}{}
			_ = json.Unmarshal(registered, &all)
			for k, v := range claims.Custom {
				all[k] = v
			}
			data, _ := json.Marshal(all)
			obj, _ := signer.Sign(data)
			token, _ := obj.CompactSerialize()

			res, err := jwx.ParseClaimsUnsafe(token)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Registered).To(Equal(claims.Registered))
			Expect(res.Custom).To(Equal(claims.Custom))
		})

		It("should return an error if the token is invalid", func() {
			invalidToken := "not.a.jwt"
			res, err := jwx.ParseClaimsUnsafe(invalidToken)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})

	Context("GetTokenHeader", func() {
		It("should return header for a valid token", func() {
			data, _ := json.Marshal(claims.Registered)
			obj, _ := signer.Sign(data)
			token, _ := obj.CompactSerialize()

			res, err := jwx.GetTokenHeader(token)
			Expect(err).ToNot(HaveOccurred())
			Expect(res.Algorithm).ToNot(BeEmpty())
		})

		It("should return an error if the token is invalid", func() {
			invalidToken := "not.a.jwt"
			res, err := jwx.GetTokenHeader(invalidToken)
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})
	})
})

func TestStdin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "jwx")
}
