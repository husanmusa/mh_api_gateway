package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/mh_api_gateway/api/http"
	"github.com/husanmusa/mh_api_gateway/genproto/appointment_service"
)

// CreateAppointment godoc
// @Security ApiKeyAuth
// @Summary Create a new appointment
// @Description This API for creating a new appointment
// ID create_appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param appointment body appointment_service.Appointment true "AppointmentCreateRequest"
// Success 201 {object} http.Response{data=string} "Appointment data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/appointment/ [POST]
func (h *Handler) CreateAppointment(c *fiber.Ctx) error {
	var appointment appointment_service.Appointment

	err := c.BodyParser(&appointment)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.AppointmentService().Create(
		c.Context(),
		&appointment,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetAppointment godoc
// @Security ApiKeyAuth
// @Summary Get appointment by appointment_id
// ID get_appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param appointment_id path string true "appointment_id"
// @Success 200 {object} http.Response{data=appointment_service.AppointmentId} "GetAppointment ResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/appointment/{appointment_id} [GET]
func (h *Handler) GetAppointment(c *fiber.Ctx) error {

	id := c.Params("appointment_id")
	resp, err := h.services.AppointmentService().Get(
		c.Context(),
		&appointment_service.AppointmentId{
			Id: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// CancelAppointment godoc
// @Security ApiKeyAuth
// @Summary Delete appointment by_id
// @ID delete_appointment
// @Tags Appointment
// @Accept json
// @Produce json
// @Param appointment_id path string false "appointment_id"
// @Success 200 {object} http.Response{data=string} "Success Cancel"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/appointment/{appointment_id} [DELETE]
func (h *Handler) CancelAppointment(c *fiber.Ctx) error {

	id := c.Params("appointment_id")

	_, err := h.services.AppointmentService().Cancel(
		c.Context(),
		&appointment_service.AppointmentId{
			Id: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}
