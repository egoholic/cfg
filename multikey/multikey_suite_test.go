package multikey_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultikey(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Multikey Suite")
}
