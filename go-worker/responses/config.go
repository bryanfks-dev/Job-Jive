package responses

import (
	"models"
)

type ConfigResponse models.ConfigJson

func (config_response *ConfigResponse) Create(config_json models.ConfigJson) {
	config_response.CheckInTime = config_json.CheckInTime
	config_response.CheckOutTime = config_json.CheckOutTime
	config_response.MinCheckInMinutes = config_json.MinCheckInMinutes
	config_response.MaxCheckOutMinutes = config_json.MaxCheckOutMinutes
	config_response.AbsenceQuota = config_json.AbsenceQuota
	config_response.DailyWorkHours = config_json.DailyWorkHours
	config_response.WeeklyWorkHours = config_json.WeeklyWorkHours
	config_response.DeductionAmounts = config_json.DeductionAmounts
}
