package handler

import (
	"net/http"
	"project/features/payments"
	"project/helper"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	s payments.PaymentService
}

func NewPaymentHandler(service payments.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		s: service,
	}
}

func (ph *PaymentHandler) CreatePayment() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res, orderID, err := ph.s.CreatePayment(id)
		registrationToken := c.Request().Header.Get("Authorization")
		if err != nil {
			c.Logger().Error("Handler : cannot create payment", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		if res != nil && res.Token != "" {
			errSendMessage := ph.s.SendEmail(registrationToken, orderID, res.Token)
			if errSendMessage != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("failed, "+errSendMessage.Error(), nil, http.StatusBadRequest))
			}
		} else {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("failed, missing payment token", nil, http.StatusBadRequest))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}
