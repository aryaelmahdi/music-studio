package handler

import (
	"fmt"
	"net/http"
	"project/features/rooms"
	"project/helper"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type RoomHandler struct {
	s rooms.RoomService
}

func NewRoomHandler(service rooms.RoomService) rooms.RoomHandler {
	return &RoomHandler{
		s: service,
	}
}

func (rh *RoomHandler) AddRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input rooms.Rooms
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		res, err := rh.s.AddRoom(input, c.Get("user").(*jwt.Token))
		if err != nil {
			if strings.Contains(err.Error(), "Unauthorized user") {
				c.Logger().Error("handler: Unauthorized user")
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusUnauthorized))
			}
			c.Logger().Error("handler: input process error :", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse("success", res, http.StatusCreated))
	}
}

func (rh *RoomHandler) DeleteRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.Param("id")
		res, err := rh.s.DeleteRoom(input, c.Get("user").(*jwt.Token))
		if err != nil {
			if strings.Contains(err.Error(), "Unauthorized user") {
				c.Logger().Error("handler: Unauthorized user")
				return c.JSON(http.StatusUnauthorized, helper.FormatResponse("fail, "+err.Error(), nil, http.StatusUnauthorized))
			}
			c.Logger().Error("handler: delete process error:", err.Error())
		}

		return c.JSON(http.StatusNoContent, helper.FormatResponse("room :"+fmt.Sprint(res)+" deleted", nil, http.StatusNoContent))
	}
}

func (rh *RoomHandler) GetAllRooms() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := rh.s.GetAllRooms()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse("fail", nil, http.StatusInternalServerError))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (rh *RoomHandler) GetRoomByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := c.Param("id")
		res, err := rh.s.GetRoomByID(input)
		if err != nil {
			c.Logger().Error("handler: get process error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}

func (rh *RoomHandler) UpdateRoom() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input rooms.Rooms
		roomID := c.Param("id")
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}

		input.RoomID = roomID
		res, err := rh.s.UpdateRoom(roomID, input)
		if err != nil {
			c.Logger().Error("handler: update process error:", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse("fail", nil, http.StatusBadRequest))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse("success", res, http.StatusOK))
	}
}
