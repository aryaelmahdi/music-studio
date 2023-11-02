package helper

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type HelperDate interface {
	CompareDate(date string) error
	VerifyCancelDate(date string) error
}

func CompareDate(date string) error {
	dateParts := strings.Split(date, "-")
	if len(dateParts) != 3 {
		return errors.New("invalid date format")
	}

	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])

	if month > 12 || month < 0 {
		return errors.New("invalid month")
	}

	reservationDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	today := time.Now()

	if reservationDate.Before(today) {
		return errors.New("date has passed")
	}

	if reservationDate == today {
		return errors.New("date cannot be the same")
	}
	return nil
}

func VerifyCancelDate(date string) error {
	dateParts := strings.Split(date, "-")
	if len(dateParts) != 3 {
		return errors.New("invalid date format")
	}

	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])

	reservationDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	today := time.Now()

	daysDiff := int(reservationDate.Sub(today).Hours() / 24)

	if daysDiff <= 2 {
		return errors.New("cancelation is allowed up to 2 days before the reservation date")
	}

	return nil
}
