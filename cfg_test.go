package cfg_test

import (
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var (
	pathToCompiled string
	err            error
)

type StringWriter struct {
	sb strings.Builder
}

func (sw *StringWriter) Write(bytes []byte) (int, error) {
	return sw.sb.Write(bytes)
}
func (sw *StringWriter) String() string {
	return sw.sb.String()
}

var _ = Describe("cfg", func() {
	BeforeEach(func() {
		pathToCompiled, err = gexec.Build("github.com/egoholic/cfg/test")
	})
	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})

	It("configures executable with flags", func() {
		var (
			outSb strings.Builder
			errSb strings.Builder
		)
		command := exec.Command(pathToCompiled, "-bool", "1", "-string", `test\ string`, "-strings", `testString1 text\ string\ 2`, "-int", "9", "-ints", "1", "2", "3", "5", "8")
		session, err := gexec.Start(command, &outSb, &errSb)
		Expect(err).ShouldNot(HaveOccurred())
		session.Wait()
		Expect(outSb.String()).To(Equal("BoolArg: true\nStringArg: test\\ string\nIntArg: 9\nDefaultStringArg: Default Value\nStringArrayArg: testString1 text\\ string\\ 2\nIntArrayArg: [1 2 3 5 8]\n\n"))
		Expect(errSb.String()).To(Equal(""))
		outSb.Reset()
		errSb.Reset()
		command = exec.Command(pathToCompiled, "-bool", "1", "-string", `test\ string`, "-strings", `testString1 text\ string\ 2`, "-int", "9", "-ints", "1", "2", "3", "5", "8", "-default", "changed")
		session, err = gexec.Start(command, &outSb, &errSb)
		Expect(err).ShouldNot(HaveOccurred())
		session.Wait()
		Expect(outSb.String()).To(Equal("BoolArg: true\nStringArg: test\\ string\nIntArg: 9\nDefaultStringArg: changed\nStringArrayArg: testString1 text\\ string\\ 2\nIntArrayArg: [1 2 3 5 8]\n\n"))
		Expect(errSb.String()).To(Equal(""))
	})

	It("configures executable with env vars", func() {
		var (
			outSb strings.Builder
			errSb strings.Builder
		)
		os.Setenv("BOOL", "1")
		os.Setenv("STRING", "test string")
		os.Setenv("STRINGS", `testString1:"text string 2"`)
		os.Setenv("INT", "9")
		os.Setenv("INTS", "1:2:3:5:8")
		command := exec.Command(pathToCompiled)
		session, err := gexec.Start(command, &outSb, &errSb)
		Expect(err).ShouldNot(HaveOccurred())
		session.Wait()
		Expect(outSb.String()).To(Equal("BoolArg: true\nStringArg: test string\nIntArg: 9\nDefaultStringArg: Default Value\nStringArrayArg: testString1,\"text string 2\"\nIntArrayArg: [1 2 3 5 8]\n\n"))
		Expect(errSb.String()).To(Equal(""))
		outSb.Reset()
		errSb.Reset()
		os.Setenv("DEFAULT", "changed")
		command = exec.Command(pathToCompiled)
		session, err = gexec.Start(command, &outSb, &errSb)
		Expect(err).ShouldNot(HaveOccurred())
		session.Wait()
		Expect(outSb.String()).To(Equal("BoolArg: true\nStringArg: test string\nIntArg: 9\nDefaultStringArg: changed\nStringArrayArg: testString1,\"text string 2\"\nIntArrayArg: [1 2 3 5 8]\n\n"))
		Expect(errSb.String()).To(Equal(""))
	})
})
