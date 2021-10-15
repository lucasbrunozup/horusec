package version_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	"os/exec"
)


var _ = Describe("Run horusec CLI with version argument", func() {
	It("displays current version", func() {
		command := exec.Command("horusec", "version")
		session, err := Start(command, GinkgoWriter, GinkgoWriter)

		Expect(err).ShouldNot(HaveOccurred())

		Eventually(session.Wait()).Should(gbytes.Say("Actual version installed of the horusec is:"))
	})
})
