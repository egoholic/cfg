package doc_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Doc Suite")
}
