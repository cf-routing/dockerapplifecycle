package testingtsupport_test

import (
	. "github.com/cloudfoundry-incubator/docker-circus/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
