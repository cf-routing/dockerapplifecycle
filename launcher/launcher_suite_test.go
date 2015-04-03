package main_test

import (
	"testing"

	. "github.com/cloudfoundry-incubator/docker_app_lifecycle/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/cloudfoundry-incubator/docker_app_lifecycle/Godeps/_workspace/src/github.com/onsi/gomega"
	"github.com/cloudfoundry-incubator/docker_app_lifecycle/Godeps/_workspace/src/github.com/onsi/gomega/gexec"
)

var launcher string

func TestDockerLifecycleLauncher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docker-App-Lifecycle-Launcher Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	launcherPath, err := gexec.Build("github.com/cloudfoundry-incubator/docker_app_lifecycle/launcher", "-race")
	Ω(err).ShouldNot(HaveOccurred())
	return []byte(launcherPath)
}, func(launcherPath []byte) {
	launcher = string(launcherPath)
})

var _ = SynchronizedAfterSuite(func() {
	//noop
}, func() {
	gexec.CleanupBuildArtifacts()
})
