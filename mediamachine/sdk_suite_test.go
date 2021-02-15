package mediamachine_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMediamachine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "mediamachine Suite")
}
