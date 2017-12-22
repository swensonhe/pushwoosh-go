package pushwoosh_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSceneService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pushwoosh Suite")
}
