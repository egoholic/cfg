package doc_test

import (
	. "github.com/egoholic/cfg/doc"
	"github.com/egoholic/cfg/multikey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	name = "Help"
	typ  = Command
	desc = "Presents help documentation about the program and its arguments."
	mk   = multikey.New("help")
)

var _ = Describe("doc", func() {
	Describe("New()", func() {
		It("returns Doc", func() {
			Expect(New(name, typ, desc, mk)).To(Equal(Doc{
				Name:          "Help",
				Documentation: "\n\tName: Help (of type: 'Command')\n\n\t\tDescription: Presents help documentation about the program and its arguments.\n\n\t\tKey:        help\n\n\n",
			}))
		})
	})
})
