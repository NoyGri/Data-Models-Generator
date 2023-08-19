package yaml_python_source

import (
	"github.com/onsi/gomega"
	"os/exec"
	"testing"
)

func TestYamlPythonSource(t *testing.T) {
	g := gomega.NewWithT(t)

	command := exec.Command("dmgen", "-f", "yaml_python_source/dmgen.yaml", "-o", "yaml_python_source/dmgen.py")
	err := command.Run()

	g.Expect(err).To(gomega.BeNil())
}
