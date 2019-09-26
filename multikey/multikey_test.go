package multikey_test

import (
	. "github.com/egoholic/cfg/multikey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("multikey", func() {
	Describe("New()", func() {
		It("creates multiple keys by convention", func() {
			Expect(New("some_Key")).To(Equal(MK{
				Key:    "some_key",
				ENVVar: "SOME_KEY",
				Flag:   "-some_key",
			}))
		})
	})
})
