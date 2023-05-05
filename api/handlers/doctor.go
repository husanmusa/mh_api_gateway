package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/mh_api_gateway/api/http"
	"github.com/husanmusa/mh_api_gateway/genproto/appointment_service"
)

// CreateDoctor godoc
// @Security ApiKeyAuth
// @Summary Create a new doctor
// @Description This API for creating a new doctor
// ID create_doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param doctor body appointment_service.Doctor true "DoctorCreateRequest"
// Success 201 {object} http.Response{data=string} "Doctor data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/doctor/ [POST]
func (h *Handler) CreateDoctor(c *fiber.Ctx) error {
	var doctor appointment_service.Doctor

	err := c.BodyParser(&doctor)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.DoctorService().Create(
		c.Context(),
		&doctor,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetListDoctors godoc
// @Security ApiKeyAuth
// @Summary Get all doctors
// @ID get_doctor_list
// @Tags Doctor
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=appointment_service.GetListDoctorsRequest} "GetDoctorListResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/doctor [GET]
func (h *Handler) GetListDoctors(c *fiber.Ctx) error {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.DoctorService().GetList(
		c.Context(),
		&appointment_service.GetListDoctorsRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// GetDoctor godoc
// @Security ApiKeyAuth
// @Summary Get doctor by doctor_id
// @Param platform-id header string true "uuid" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// ID get_doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param doctor_id path string true "doctor_id"
// @Success 200 {object} http.Response{data=appointment_service.DoctorId} "GetCountryResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/doctor/{doctor_id} [GET]
func (h *Handler) GetDoctor(c *fiber.Ctx) error {

	id := c.Params("doctor_id")
	resp, err := h.services.DoctorService().Get(
		c.Context(),
		&appointment_service.DoctorId{
			DoctorId: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateDoctor godoc
// @Security ApiKeyAuth
// @Summary Update doctor
// ID update_doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Doctor data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/doctor/ [PUT]
func (h *Handler) UpdateDoctor(c *fiber.Ctx) error {
	var doctor appointment_service.Doctor

	_, err := h.services.DoctorService().Update(
		c.Context(),
		&doctor,
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}
	return h.handleResponse(c, http.OK, "Updated Successfully")
}

// DeleteDoctor godoc
// @Security ApiKeyAuth
// @Summary Delete doctor by_id
// @ID delete_doctor
// @Tags Doctor
// @Accept json
// @Produce json
// @Param doctor_id path string false "doctor_id"
// @Success 200 {object} http.Response{data=string} "deleted"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/doctor/{doctor_id} [DELETE]
func (h *Handler) DeleteDoctor(c *fiber.Ctx) error {

	id := c.Params("doctor_id")

	_, err := h.services.DoctorService().Delete(
		c.Context(),
		&appointment_service.DoctorId{
			DoctorId: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}
