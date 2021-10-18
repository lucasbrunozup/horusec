package version_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Run horusec CLI with version argument", func() {
	var outBuffer *gbytes.Buffer
	var session *gexec.Session
	var err error

	BeforeEach(func() {
		outBuffer = gbytes.NewBuffer()
		session, err = gexec.Start(exec.Command("horusec", "version"), outBuffer, outBuffer)
	})

	It("execute command without error", func() {
		Expect(err).ShouldNot(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))
	})

	It("displays current version", func() {
		Eventually(outBuffer).Should(gbytes.Say("Actual version installed of the horusec is:"))
	})
})
