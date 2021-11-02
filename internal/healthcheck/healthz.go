package healthcheck

import (
	"net/http"
	"time"

	"github.com/Rau9/library/internal/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Response is the type that handles the serialised response of the
// healthcheck endpoint
type Response struct {
	DB health `json:"database"`
}

type health struct {
	Status bool          `json:"status"`
	Time   time.Duration `json:"time_ns"`
	Error  string        `json:"error"`
}

// timeTrack is a function that modifies a time.Duration pointer
// since the start time it was executed
func timeTrack(start time.Time, timer *time.Duration) {
	*timer = time.Since(start)
}

// checkDB validates the healthiness of the database
func (res *Response) checkDB(dbclient *gorm.DB, logger *zap.Logger) error {
	// track and measure the time it takes this function to execute
	defer timeTrack(time.Now(), &res.DB.Time)
	var hc models.Healthcheck
	tx := dbclient.FirstOrCreate(&hc, models.Healthcheck{Status: true})
	if tx.Error != nil {
		logger.Error("error checking db health",
			zap.String("check", "db"),
			zap.Error(tx.Error),
		)
		res.DB.Error = tx.Error.Error()
		res.DB.Status = false
		return tx.Error
	}
	res.DB.Status = true
	return nil
}

// NewResponse create a new healthcheck Response object
func NewResponse() *Response {
	return &Response{}
}

// Handler deals with /healthz requests
func Handler(dbclient *gorm.DB, logger *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		res := NewResponse()
		if err := res.checkDB(dbclient, logger); err != nil {
			return c.JSON(http.StatusInternalServerError, res)
		}
		return c.JSON(http.StatusOK, res)
	}
}
