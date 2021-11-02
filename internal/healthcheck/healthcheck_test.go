package healthcheck_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/Rau9/library/internal/healthcheck"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Healthcheck", func() {
	var e *echo.Echo
	var req *http.Request
	var rec *httptest.ResponseRecorder
	var c echo.Context
	// BeforeEach test create the test webserver
	BeforeEach(func() {
		e = echo.New()
		req = httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)
		Expect(c).NotTo(BeNil())
	})
	Context("Handler", func() {
		When("A successful healthcheck validation", func() {
			It("should return no errors and the state of the dependencies", func() {
				h := healthcheck.Handler(dbClient, lgr)
				err := h(c)
				Expect(err).NotTo(HaveOccurred())
				Expect(http.StatusOK).To(Equal(rec.Code))

				jsonResponse := healthcheck.NewResponse()
				// parse the json response via the recorder and map it to a new
				// healthcheck.Response object
				err = json.Unmarshal(rec.Body.Bytes(), jsonResponse)

				Expect(err).NotTo(HaveOccurred())
				Expect(jsonResponse.DB.Status).To(Equal(true))
				Expect(jsonResponse.DB.Error).To(Equal(""))
				Expect(jsonResponse.DB.Time).To(BeNumerically(">", 0))
				Expect(jsonResponse.DB.Time).To(BeNumerically("<", 100000000))
			})
		})

		When("A failed healthcheck validation", func() {
			It("should return an error", func() {
				h := healthcheck.Handler(faultyDbClient, lgr)
				_ = h(c)
				Expect(http.StatusInternalServerError).To(Equal(rec.Code))

				jsonResponse := healthcheck.NewResponse()
				// parse the json response via the recorder and map it to a new
				// healthcheck.Response object
				err := json.Unmarshal(rec.Body.Bytes(), jsonResponse)

				Expect(err).NotTo(HaveOccurred())
				Expect(jsonResponse.DB.Status).To(Equal(false))
				Expect(jsonResponse.DB.Error).To(MatchRegexp("password authentication failed"))
				Expect(jsonResponse.DB.Time).To(BeNumerically(">", 0))
				Expect(jsonResponse.DB.Time).To(BeNumerically("<", 100000000))
			})
		})
	})
})
