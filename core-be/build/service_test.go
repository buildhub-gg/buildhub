package build_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/buildhub-gg/buildhub/build"
)

var _ = Describe("Service", func() {
	victim := build.NewService(build.NewInMemoryBuildRepository())

	Context("Create a new build", func() {
		When("the build has no name", func() {
			ID, err := victim.Create(&build.Build {})

			It("Fails to create a build", func() {
				Expect(ID).To(BeEmpty())
				Expect(err).NotTo(BeNil())
			})
		})
		When("the build has a name and no items", func() {
			ID, err := victim.Create(&build.Build {
				Name: "Test name",
			})

			It("Creates the build successfully", func() {
				Expect(ID).NotTo(BeEmpty())
				Expect(err).To(BeNil())
			})
		})
	})
})
