package yaml

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestInvalidYaml(t *testing.T) {

	t.Run("NoSuchFile", func(t *testing.T) {
		g := NewWithT(t)

		data, err := Read("no_such_file.yaml")

		g.Expect(err).NotTo(BeNil())
		g.Expect(data).To(BeNil())
	})
}
