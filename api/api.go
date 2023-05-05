package api

import (
	"encoding/json"
	_ "github.com/husanmusa/mh_api_gateway/api/docs"
	"github.com/husanmusa/mh_api_gateway/api/handlers"
	"github.com/husanmusa/mh_api_gateway/config"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	logger2 "github.com/saidamir98/udevs_pkg/logger"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://udevs.io
func SetUpRouter(h handlers.Handler, cfg config.Config, log logger2.LoggerI) *fiber.App {
	router := fiber.New(fiber.Config{JSONEncoder: json.Marshal, BodyLimit: 100 * 1024 * 1024})
	router.Use(logger.New(), cors.New())
	// router.Use(h.MaxAllowed(12000))
	router.Use(func(c *fiber.Ctx) error {
		log.Info("request", logger2.String("method", c.Method()), logger2.String("path", c.Path()),
			logger2.String("ip", c.Get("X-Forwarded-For")), logger2.String("ip", c.Get("X-Real-IP")))
		return c.Next()
	})
	router.Server().MaxConnsPerIP = 20
	router.Server().DisableKeepalive = true
	router.Server().IdleTimeout = time.Hour
	router.Server().LogAllErrors = true

	router.Get("/api/swagger/*", swagger.HandlerDefault)
	r := router.Group("/api")
	r.Get("/ping", h.Ping)
	r.Get("/config", h.GetConfig)

	// APPOINTMENT
	r.Post("/appointment", h.CreateAppointment)
	r.Get("/appointment/:appointment_id", h.GetAppointment)
	r.Delete("/appointment/:appointment_id", h.CancelAppointment)

	// DOCTOR
	r.Post("/doctor", h.CreateDoctor)
	r.Get("/doctor", h.GetListDoctors)
	r.Get("/doctor/:doctor_id", h.GetDoctor)
	r.Put("/doctor", h.UpdateDoctor)
	r.Delete("/doctor/:doctor_id", h.DeleteDoctor)

	// PATIENT
	r.Post("/patient", h.CreatePatient)
	r.Get("/patient", h.GetListPatients)
	r.Get("/patient/:patient_id", h.GetPatient)
	r.Put("/patient", h.UpdatePatient)
	r.Delete("/patient/:patient_id", h.DeletePatient)

	return router
}
