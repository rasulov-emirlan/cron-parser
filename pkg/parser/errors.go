package parser

import "errors"

var (
	ErrInvalidMinutes    = errors.New("parser: invalid minutes")
	ErrInvalidHours      = errors.New("parser: invalid hours")
	ErrInvalidDayOfMonth = errors.New("parser: invalid day of month")
	ErrInvalidMonth      = errors.New("parser: invalid month")
	ErrInvalidDayOfWeek  = errors.New("parser: invalid day of week")

	ErrInvalidInput = errors.New("parser: invalid input")
)
