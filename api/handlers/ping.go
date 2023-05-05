package handlers

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/husanmusa/mh_api_gateway/api/http"
	"github.com/husanmusa/mh_api_gateway/config"
)

// Ping godoc
// @ID ping
// @Router /api/ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Response data"
// @Failure 500 {object} http.Response{}
func (h *Handler) Ping(c *fiber.Ctx) error {
	return h.handleResponse(c, http.OK, "pong")
}

// GetConfig godoc
// @ID get_config
// @Router /api/config [GET]
// @Summary get config data on the debug mode
// @Description show service config data when the service environment set to debug mode
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=config.Config} "Response data"
// @Failure 400 {object} http.Response{}
func (h *Handler) GetConfig(c *fiber.Ctx) error {
	switch h.cfg.Environment {
	case config.DebugMode:
		return h.handleResponse(c, http.OK, h.cfg)
	case config.TestMode:
		return h.handleResponse(c, http.OK, h.cfg.Environment)
	case config.ReleaseMode:
		return h.handleResponse(c, http.OK, "private data")
	}

	return h.handleResponse(c, http.BadEnvironment, "wrong environment value passed")
}
