package middleware_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/Rau9/library/internal/middleware"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

var _ = Describe("ZapLogger", func() {
	When("A request is processed by the middleware", func() {
		It("should produce a log entry with predefined zap fields", func() {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/myendpoint", nil)
			req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			obs, logs := observer.New(zap.DebugLevel)
			logger := zap.New(obs)
			// validate logger with a generic endpoint
			h := func(c echo.Context) error {
				return c.String(http.StatusOK, "")
			}

			err := middleware.ZapLogger(logger)(h)(c)
			logFields := logs.AllUntimed()[0].ContextMap()
			Expect(err).NotTo(HaveOccurred())
			Expect(http.StatusOK).To(Equal(rec.Code))
			Expect(logs.Len()).To(Equal(1))
			Expect(logFields["status"]).To(Equal(int64(http.StatusOK)))
			Expect(logFields["request"]).To(Equal("GET /myendpoint"))
			Expect(logFields["host"]).To(Equal("example.com"))
			Expect(logFields["size"]).To(Equal(int64(0)))
		})
	})
})
