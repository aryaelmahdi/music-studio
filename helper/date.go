package helper

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func CaompareDate(date string) error {
	dateParts := strings.Split(date, "-")
	if len(dateParts) != 3 {
		return errors.New("invalid date format")
	}

	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])

	reservationDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	today := time.Now()

	if reservationDate.Before(today) {
		return errors.New("date has passed")
	}
	return nil
}
