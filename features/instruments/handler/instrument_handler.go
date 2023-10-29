package handler

import (
	"net/http"
	"project/features/instruments"
	"project/helper"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type InstrumentHandler struct {
	s instruments.InstrumentService
}

func NewInstrumentHandler(service instruments.InstrumentService) instruments.InstrumentHandler {
	return &InstrumentHandler{
		s: service,
	}
}

func (ih *InstrumentHandler) GetAllInstruments() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ih.s.GetAllInstruments()
		if err != nil {
			c.Logger().Error("handler : cannot get instruments data :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (ih *InstrumentHandler) GetInstrumentByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res, err := ih.s.GetInstrumentByID(id)
		if err != nil {
			c.Logger().Error("handler : cannot get instrument data :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (ih *InstrumentHandler) AddInstrument() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input instruments.Instruments
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler : binding process error :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		res, err := ih.s.AddInstrument(input, c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error("handler : cannot add instruments :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (ih *InstrumentHandler) DeleteInstrument() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := ih.s.DeleteInstrument(id, c.Get("user").(*jwt.Token)); err != nil {
			c.Logger().Error("handler : cannot delete instrument :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusNoContent, helper.FormatResponse("success", nil, http.StatusNoContent))
	}
}

func (ih *InstrumentHandler) UpdateInstrument() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var input instruments.Instruments
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler : binding process error :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}

		res, err := ih.s.UpdateInstrument(id, input, c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error("handler : cannot get instruments data :" + err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}
