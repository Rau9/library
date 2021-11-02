package config_test

import (
	"os"

	"github.com/Rau9/library/internal/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("Init", func() {
		When("A POSTGRES_PASSWORD environment variable is informed", func() {

			os.Setenv("POSTGRES_PASSWORD", "secure_pass_here")

			It("should return no errors", func() {
				err := config.Init()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		When("A POSTGRES_PASSWORD environment variable is not informed", func() {

			It("should return an error", func() {
				os.Unsetenv("POSTGRES_PASSWORD")
				err := config.Init()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
