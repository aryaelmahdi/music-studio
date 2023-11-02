package handler

import (
	"encoding/json"
	"io"
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
		res, orderID, email, err := ph.s.CreatePayment(id)
		if err != nil {
			c.Logger().Error("Handler : cannot create payment", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		if res != nil && res.Token != "" {
			errSendMessage := ph.s.SendEmail(email, orderID, res.Token)
			if errSendMessage != nil {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse("failed, "+errSendMessage.Error(), nil, http.StatusBadRequest))
			}
		} else {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("failed, missing payment token", nil, http.StatusBadRequest))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}

func (ph *PaymentHandler) GetNotification() echo.HandlerFunc {
	return func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			c.Logger().Error("cannot read body")
		}
		defer c.Request().Body.Close()

		var notificationData map[string]interface{}
		if err := json.Unmarshal(body, &notificationData); err != nil {
			c.Logger().Error("cannot unmarhsall body")
		}
		orderID := notificationData["order_id"].(string)

		if err := ph.s.ConfirmedPayment(orderID); err != nil {
			c.Logger().Error("handler : cannot confirmed payment")
		}
		if err := ph.s.ConfirmedPaymentEmail(orderID); err != nil {
			c.Logger().Error("handler : cannot send email")
		}
		return c.String(http.StatusOK, "OK")
	}
}
