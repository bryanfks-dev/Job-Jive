package forms

import (
	"errors"
	"strings"
	"time"

	"models"
)

type ConfigForm models.ConfigJson

var (
	ErrInvalidAbsenceQuota          = errors.New("invalid absensce quota")
	ErrInvalidDailyWorkHours        = errors.New("invalid daily work hours")
	ErrInvalidWeeklyWorkHours       = errors.New("invalid weekly work hours")
	ErrInvalidCheckInTime           = errors.New("invalid check in time")
	ErrInvalidCheckOutTime          = errors.New("invalid check out time")
	ErrCheckInGTCheckOut            = errors.New("check in time should less than check out time")
	ErrInvalidMinCheckInMinutes     = errors.New("invalid minimum check in")
	ErrInvalidMaxCheckOutMinutes    = errors.New("invalid maximum check out")
	ErrMismatchWeeklyDailyWorkHours = errors.New("daily work hours should less than weekly work hours")
	ErrInvalidDeductionAmmounts     = errors.New("invalid deduction ammounts")
)

func (config *ConfigForm) Sanitize() {
	config.CheckInTime = strings.TrimSpace(config.CheckInTime)
	config.CheckOutTime = strings.TrimSpace(config.CheckOutTime)
}

func (config ConfigForm) Validate() (bool, error) {
	if config.AbsenceQuota < 0 {
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

	_, err := time.Parse(time.TimeOnly, config.CheckInTime+":00")

	if err != nil {
		return false, ErrInvalidCheckInTime
	}

	_, err = time.Parse(time.TimeOnly, config.CheckOutTime+":00")

	if err != nil {
		return false, ErrInvalidCheckOutTime
	}

	if config.MinCheckInMinutes <= 0 || config.MinCheckInMinutes > 1440 {
		return false, ErrCheckInGTCheckOut
	}

	if config.MaxCheckOutMinutes <= 0 || config.MaxCheckOutMinutes > 1440 {
		return false, ErrInvalidMaxCheckOutMinutes
	}

	if config.WeeklyWorkHours < config.DailyWorkHours {
		return false, ErrMismatchWeeklyDailyWorkHours
	}

	if config.DeductionAmounts < 0 {
		return false, ErrInvalidDeductionAmmounts
	}

	return true, nil
}
