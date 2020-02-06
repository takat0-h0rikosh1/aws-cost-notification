package util

import "time"

const location = "Asia/Tokyo"

var FixedTZLocation = genFixedTimeZoneLocation()

func genFixedTimeZoneLocation() *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	return loc
}
