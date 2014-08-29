package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Soldier", func() {
	var appDir string

	BeforeEach(func() {
		var err error

		appDir, err = ioutil.TempDir("", "app-dir")
		Ω(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		os.RemoveAll(appDir)
	})

	It("executes it with $HOME as the given dir", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "echo HOME set to $HOME"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("HOME set to " + appDir))
	})

	It("executes it with $TMPDIR as the given dir + /tmp", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "echo TMPDIR set to $TMPDIR"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("TMPDIR set to " + appDir + "/tmp"))
	})

	It("executes with the environment of the caller", func() {
		os.Setenv("CALLERENV", "some-value")

		session, err := gexec.Start(
			exec.Command(soldier, appDir, "echo CALLERENV set to $CALLERENV"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("CALLERENV set to some-value"))
	})

	It("changes to the app directory when running", func() {
		session, err := gexec.Start(
			exec.Command(soldier, appDir, "echo PWD is $(pwd)"),
			GinkgoWriter,
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gbytes.Say("PWD is .*" + filepath.Base(appDir)))
	})

	It("munges VCAP_APPLICATION appropriately", func() {
		outBuf := new(bytes.Buffer)

		cmd := exec.Command(soldier, appDir, "echo $VCAP_APPLICATION")
		cmd.Env = append(
			os.Environ(),
			"PORT=8080",
			"CF_INSTANCE_GUID=some-instance-guid",
			"CF_INSTANCE_INDEX=123",
			`VCAP_APPLICATION={"foo":1}`,
		)

		session, err := gexec.Start(
			cmd,
			io.MultiWriter(GinkgoWriter, outBuf),
			GinkgoWriter,
		)
		Ω(err).ShouldNot(HaveOccurred())

		Eventually(session).Should(gexec.Exit(0))

		vcapApplication := map[string]interface{}{}

		err = json.Unmarshal(outBuf.Bytes(), &vcapApplication)
		Ω(err).ShouldNot(HaveOccurred())

		Ω(vcapApplication["host"]).Should(Equal("0.0.0.0"))
		Ω(vcapApplication["port"]).Should(Equal(float64(8080)))
		Ω(vcapApplication["instance_index"]).Should(Equal(float64(123)))
		Ω(vcapApplication["instance_id"]).Should(Equal("some-instance-guid"))
		Ω(vcapApplication["foo"]).Should(Equal(float64(1)))
	})
})
