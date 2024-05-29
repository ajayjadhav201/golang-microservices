package common

import "time"

// LocalTimeToUnix converts local time to Unix time (seconds since the epoch)
func LocalTimeToUnix(localTime time.Time) int64 {
	return localTime.Unix()
}

// UnixToLocalTime converts Unix time to local time
func UnixToLocalTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}

// useful time formats
//
// RFC3339Nano	2024-05-28T16:05:20.982793+05:30
// RFC850		Tuesday, 28-May-24 16:05:20 IST
// RFC1123		Tue, 28 May 2024 16:05:20 IST
// UTC			2024-05-28 10:35:20.9837966 +0000 UTC
// Unix			1716892520
// UnixMilli	1716 8925 2098 4
