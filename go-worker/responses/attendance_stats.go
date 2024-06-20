package responses

import "models"

type AttendanceStatsReponse struct {
	UserId            int `json:"user_id"`
	CurrentWeekHours  int `json:"current_week_hours"`
	CurrentMonthHours int `json:"current_month_hours"`
	AnnualLeaves      int `json:"annual_leaves"`
}

func (attendance_stats_response *AttendanceStatsReponse) Create(attendance_stats models.AttendanceStats) {
	attendance_stats_response.UserId = attendance_stats.UserID
	attendance_stats_response.CurrentWeekHours = attendance_stats.CurrentWeekHours
	attendance_stats_response.CurrentMonthHours = attendance_stats.CurrentMonthHours
	attendance_stats_response.AnnualLeaves = attendance_stats.AnnualLeaves
}
