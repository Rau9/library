package storage_test

import (
	"github.com/Rau9/library/internal/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("Storage", func() {
	Context("RunMigrations", func() {
		When("Migrations failed", func() {

			dsn := "host=localhost user=postgr password=secure_pass_here dbname=catalog port=5432 sslmode=disable"
			dbclient, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

			It("should return error", func() {
				err := storage.RunMigrations(dbclient)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(MatchRegexp("failed to connect to"))
			})
		})

		When("Migrations succeeded", func() {

			dsn := "host=localhost user=postgres password=secure_pass_here dbname=catalog port=5432 sslmode=disable"
			dbclient, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

			It("should return an nil", func() {
				err := storage.RunMigrations(dbclient)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
