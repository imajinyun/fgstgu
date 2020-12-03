package timer

import "time"

// GetNowTime func.
func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(loc)
}

// GetCalculateTime func.
func GetCalculateTime(t time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return t.Add(duration), nil
}
