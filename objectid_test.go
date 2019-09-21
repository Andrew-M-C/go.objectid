package objectid

import (
	"testing"
	"time"
)

func Test_New12(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var tm time.Time

	// test standard generation
	id12 := New12()
	tm = id12.Time()
	t.Logf("Go id12: %v, time %v", id12, tm)
	if tm.IsZero() {
		t.Errorf("Got id12 zero time")
		return
	}

	id16 := New16()
	tm = id16.Time()
	t.Logf("Go id16: %v, time %v", id16, tm)
	if tm.IsZero() {
		t.Errorf("Got id16 zero time")
		return
	}

	// test a abnormal id
	idError := ObjectID([]byte{1, 2, 3, 4})
	tm = idError.Time()
	if false == tm.IsZero() {
		t.Errorf("invalid object id not detected")
		return
	}
	t.Logf("abnormal id checked")

	// test specified time
	tm = time.Now().In(loc) // try not using UTC
	perfectTm := tm.Add(-time.Duration(tm.Nanosecond()))
	if tm.Equal(perfectTm) {
		tm = tm.Add(time.Nanosecond) // add decimal part
	}
	t.Logf("perfect time: %v", perfectTm)
	t.Logf("actual time:  %v", tm)

	id12 = New12(tm)
	if false == id12.Time().Equal(perfectTm) {
		t.Errorf("Got id12 time not equal (%v <> %v)!", id12.Time(), perfectTm)
		return
	}
	t.Logf("id12 time checked")

	id16 = New16(tm)
	if false == id16.Time().Equal(tm) {
		t.Errorf("Got id16 time not equal (%v <> %v)!", id16.Time(), tm)
		return
	}
	t.Logf("id16 time checked")

	id16 = New16(perfectTm)
	if false == id16.Time().Equal(perfectTm) {
		t.Errorf("Got id16 time not equal (%v <> %v)!", id16.Time(), perfectTm)
		return
	}
	t.Logf("id16 with perfect time checked")

	return
}

func Test_MiscErrors(t *testing.T) {
	var nilBytes []byte
	var tm time.Time
	nilID := ObjectID(nilBytes)
	tm = nilID.Time()
	if false == tm.IsZero() {
		t.Errorf("abnormal time should be zero!")
		return
	}
	if "nil" != nilID.String() {
		t.Errorf("abnormal id should be nil!")
		return
	}
}
