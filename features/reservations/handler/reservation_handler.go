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
		var input reservations.Reservation
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler : binding process error ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		res, err := rh.s.AddReservation(input)
		if err != nil {
			c.Logger().Error("handler : cannot add reservation", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}
