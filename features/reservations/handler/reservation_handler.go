package handler

import (
	"net/http"
	"project/features/reservations"
	"project/helper"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ReservationHandler struct {
	s reservations.ReservationService
}

func NewReservationHandler(service reservations.ReservationService) *ReservationHandler {
	return &ReservationHandler{
		s: service,
	}
}

func (rh *ReservationHandler) GetAllReservations() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := rh.s.GetAllReservations(c.Get("user").(*jwt.Token))
		if err != nil {
			if strings.Contains(err.Error(), "Unauthorized user") {
				c.Logger().Error("Handler : Unauthorized user")
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("Unauthorized user", nil, http.StatusUnauthorized))
			}
			c.Logger().Error("Hanlder : cannot get reservations", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail to get reservations", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (rh *ReservationHandler) GetReservationsByUsername() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := rh.s.GetReservationsByUsername(c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error("Handler: cannot get reservation by username", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (rh *ReservationHandler) AddReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ReservationRequest
		var reservationData reservations.Reservation
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler : binding process error ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		if err := helper.CompareDate(input.Date); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(err.Error(), nil, http.StatusBadRequest))
		}

		reservationData.Date = input.Date
		reservationData.PaymentStatus = input.PaymentStatus
		reservationData.RoomID = input.RoomID

		res, err := rh.s.AddReservation(reservationData, c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error("handler : cannot add reservation ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail "+err.Error(), nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}

func (rh *ReservationHandler) UpdateReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input reservations.Reservation
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler : binding process error ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, binding process error", nil, http.StatusBadRequest))
		}
		if dateError := helper.CompareDate(input.Date); dateError != nil {
			c.Logger().Error("handler : error comparing data ", dateError.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+dateError.Error(), nil, http.StatusBadRequest))
		}

		res, err := rh.s.UpdateReservation(input, c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error("handler : error updating data ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (rh *ReservationHandler) DeleteReservation() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res, err := rh.s.GetReservationByID(id, c.Get("user").(*jwt.Token))
		if err != nil {
			if strings.Contains(err.Error(), "Unauthorized user") {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusUnauthorized))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		if dateErr := helper.VerifyCancelDate(res.Date); dateErr != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+dateErr.Error(), nil, http.StatusBadRequest))
		}

		if err := rh.s.DeleteReservation(id, c.Get("user").(*jwt.Token)); err != nil {
			if strings.Contains(err.Error(), "Unauthorized user") {
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusUnauthorized))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusNoContent, nil)
	}
}
