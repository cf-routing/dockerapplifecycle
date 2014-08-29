package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var spy string

func TestDockerCircusSpy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docker-Circus-Spy Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	spyPath, err := gexec.Build("github.com/cloudfoundry-incubator/docker-circus/spy")
	Ω(err).ShouldNot(HaveOccurred())
	return []byte(spyPath)
}, func(spyPath []byte) {
	spy = string(spyPath)
})

var _ = SynchronizedAfterSuite(func() {
	//noop
}, func() {
	gexec.CleanupBuildArtifacts()
})
