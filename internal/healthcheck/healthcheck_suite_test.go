package healthcheck_test

import (
	"testing"

	logger "github.com/Rau9/library/internal/logger"
	"github.com/Rau9/library/internal/storage"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHealthcheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Healthcheck Suite")
}

var dbClient, faultyDbClient *gorm.DB
var lgr *zap.Logger

/*
	Before running the test suite:
	- Create DB client
	- Run migration
	- Create faulty DB client
	- Create logger
*/
var _ = BeforeSuite(func() {
	var err error
	dsn := "host=localhost user=postgres password=secure_pass_here dbname=catalog port=5432 sslmode=disable"
	dbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())
	Expect(storage.RunMigrations(dbClient)).NotTo(HaveOccurred())

	dsnE := "host=localhost user=postgres password=invalid dbname=catalog port=5432 sslmode=disable"
	faultyDbClient, err = gorm.Open(postgres.Open(dsnE), &gorm.Config{})
	Expect(err).To(HaveOccurred())

	lgr, err = logger.NewProduction()
	Expect(err).NotTo(HaveOccurred())
})
