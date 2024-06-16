package forms

import (
	"errors"
	"strings"

	"models"
)

type ConfigForm models.ConfigJson

var (
	ErrInvalidAbsenceQuota = errors.New("invalid absensce quota")
	ErrInvalidDailyWorkHours = errors.New("invalid daily work hours")
	ErrInvalidWeeklyWorkHours = errors.New("invalid weekly work hours")
	ErrInvalidCheckInTime = errors.New("check in time should less than check out time")
	ErrMismatchWeeklyDailyWorkHours = errors.New("daily work hours should less than weekly work hours")
)

func (config *ConfigForm) Sanitize() {
	config.CheckInTime = strings.TrimSpace(config.CheckInTime)
	config.CheckOutTime = strings.TrimSpace(config.CheckOutTime)
}

func (config ConfigForm) Validate() (bool, error) {
	if config.AbsenceQuota <= 0 {
		return false, ErrInvalidAbsenceQuota
	}

	if config.DailyWorkHours <= 0 || config.WeeklyWorkHours > 24 {
		return false, ErrInvalidDailyWorkHours
	}

	if config.WeeklyWorkHours <= 0 || config.WeeklyWorkHours > 168 {
		return false, ErrInvalidWeeklyWorkHours
	}

	if config.CheckInTime >= config.CheckOutTime {
		return false, ErrInvalidCheckInTime
	}

	if config.WeeklyWorkHours < config.DailyWorkHours {
		return false, ErrMismatchWeeklyDailyWorkHours
	}

	return true, nil
}