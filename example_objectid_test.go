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

func ExampleNewByHex() {
	hex := "5DA5360F51AD44C91EB2C7291D946728"
	o, err := objectid.NewByHex(hex)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("hex:", o)
	// Output:
	// hex: 5da5360f51ad44c91eb2c7291d946728
}
