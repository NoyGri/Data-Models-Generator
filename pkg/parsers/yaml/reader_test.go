package yaml

import (
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestInvalidYaml(t *testing.T) {

	t.Run("NoSuchFile", func(t *testing.T) {
		g := NewWithT(t)

		// Act
		data, err := Read("no_such_file.yaml")

		// Assert
		g.Expect(err).NotTo(BeNil())
		g.Expect(data).To(BeNil())
	})

	t.Run("EmptyFile", func(t *testing.T) {
		g := NewWithT(t)

		// Arrange
		err := os.WriteFile("empty_file.yaml", []byte{}, 0666)
		g.Expect(err).To(BeNil())
		defer func() {
			err := os.Remove("empty_file.yaml")
			g.Expect(err).To(BeNil())
		}()

		// Act
		data, err := Read("empty_file.yaml")

		// Assert
		g.Expect(err).NotTo(BeNil())
		g.Expect(data).To(BeNil())
	})
}
