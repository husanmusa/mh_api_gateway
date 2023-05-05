package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/mh_api_gateway/api/http"
	"github.com/husanmusa/mh_api_gateway/genproto/appointment_service"
)

// CreatePatient godoc
// @Security ApiKeyAuth
// @Summary Create a new patient
// @Description This API for creating a new patient
// ID create_patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param patient body appointment_service.Patient true "PatientCreateRequest"
// Success 201 {object} http.Response{data=string} "Patient data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/patient/ [POST]
func (h *Handler) CreatePatient(c *fiber.Ctx) error {
	var patient appointment_service.Patient

	err := c.BodyParser(&patient)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.PatientService().Create(
		c.Context(),
		&patient,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetListPatients godoc
// @Security ApiKeyAuth
// @Summary Get all patients
// @ID get_patient_list
// @Tags Patient
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} http.Response{data=appointment_service.GetListPatientsRequest} "GetPatientListResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/patient [GET]
func (h *Handler) GetListPatients(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.PatientService().GetList(
		c.Context(),
		&appointment_service.GetListPatientsRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// GetPatient godoc
// @Security ApiKeyAuth
// @Summary Get patient by patient_id
// @Param platform-id header string true "uuid" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// ID get_patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param patient_id path string true "patient_id"
// @Success 200 {object} http.Response{data=appointment_service.PatientId} "GetCountryResponseBody"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/patient/{patient_id} [GET]
func (h *Handler) GetPatient(c *fiber.Ctx) error {
	id := c.Params("patient_id")
	resp, err := h.services.PatientService().Get(
		c.Context(),
		&appointment_service.PatientId{
			PatientId: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdatePatient godoc
// @Security ApiKeyAuth
// @Summary Update patient
// ID update_patient
// @Tags Patient
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Patient data"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/patient/ [PUT]
func (h *Handler) UpdatePatient(c *fiber.Ctx) error {
	var patient appointment_service.Patient

	_, err := h.services.PatientService().Update(
		c.Context(),
		&patient,
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}
	return h.handleResponse(c, http.OK, "Updated Successfully")
}

// DeletePatient godoc
// @Security ApiKeyAuth
// @Summary Delete patient by_id
// @ID delete_patient
// @Tags Patient
// @Accept json
// @Produce json
// @Param patient_id path string false "patient_id"
// @Success 200 {object} http.Response{data=string} "deleted"
// @Failure 400 {object} http.Response{data=string} "Bad request"
// @Failure 500 {object} http.Response{data=string} "Internal server error"
// @Router /api/patient/{patient_id} [DELETE]
func (h *Handler) DeletePatient(c *fiber.Ctx) error {
	id := c.Params("patient_id")

	_, err := h.services.PatientService().Delete(
		c.Context(),
		&appointment_service.PatientId{
			PatientId: id,
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, "SUCCESS")
}
