package pkg

import (
	"fmt"
	"time"
)

func AddPeriodToDate(startDate time.Time, period string) (time.Time, error) {
	switch period {
	case "weekly":
		return startDate.AddDate(0, 0, 7), nil
	case "monthly":
		return startDate.AddDate(0, 0, 30), nil
	case "yearly":
		return startDate.AddDate(0, 0, 365), nil
	default:
		return time.Time{}, fmt.Errorf("неподдерживаемый период: %s. Доступные значения: weekly, monthly, yearly", period)
	}
}
