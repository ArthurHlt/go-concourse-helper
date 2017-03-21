package go_concourse_helper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoConcourseHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoConcourseHelper Suite")
}
