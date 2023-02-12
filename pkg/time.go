package pkg

import (
	"fmt"
	"time"
)

type TimeString string

func Duration(format string, ts string) (string, error) {
	target, err := time.Parse(format, ts)
	if err != nil {
		return "", err
	}
	now := time.Now()
	duration := now.Sub(target)

	hours0 := int(duration.Hours())
	days := hours0 / 24
	months := days / 30
	years := months / 12
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60

	if years > 0 {
		return fmt.Sprintf("%d年前", years), nil
	}

	if months > 0 {
		return fmt.Sprintf("%dヶ月前", months), nil
	}

	if days > 0 {
		return fmt.Sprintf("%d日前", days), nil
	}

	if hours > 0 {
		return fmt.Sprintf("%d時間前", hours), nil
	}

	if mins > 0 {
		return fmt.Sprintf("%d分前", mins), nil
	}

	return fmt.Sprintf("%d秒前", secs), nil
}
