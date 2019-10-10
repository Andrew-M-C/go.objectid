package objectid_test

import (
	"fmt"
	"time"

	objectid "github.com/Andrew-M-C/go.objectid"
)

func ExampleObjectID_Time() {
	t := time.Date(2020, 1, 1, 12, 0, 0, 123456789, time.UTC)
	id := objectid.New16(t)
	fmt.Println("time: ", id.Time())
	// Output:
	// time:  2020-01-01 12:00:00.123456789 +0000 UTC
}
