package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

func AddGigasecond(start time.Time) time.Time {
	gigasecond := time.Duration(time.Second * 1000000000)
	return start.Add(gigasecond)
}
