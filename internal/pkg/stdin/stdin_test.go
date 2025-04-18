package stdin_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tmsmr/jwtee/internal/pkg/stdin"
	"os"
	"path"
	"testing"
)

var _ = Describe("stdin pkg:", func() {

	Context("reading from a pipe", func() {
		It("should read the input correctly", func() {
			a := "string in pipe"
			r, w, _ := os.Pipe()
			_, _ = w.WriteString(a)
			_ = w.Close()
			os.Stdin = r
			b, err := stdin.Read()
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(Equal(a))
		})
	})

	Context("reading from a file", func() {
		It("should read the input correctly", func() {
			p := path.Join(GinkgoT().TempDir(), "testinput.txt")
			f, _ := os.Create(p)
			a := "string in file"
			_, _ = f.WriteString(a)
			_ = f.Close()
			f, _ = os.Open(p)
			os.Stdin = f
			b, err := stdin.Read()
			Expect(err).ToNot(HaveOccurred())
			_ = f.Close()
			Expect(b).To(Equal(a))
		})
	})
})

func TestStdin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "stdin")
}
