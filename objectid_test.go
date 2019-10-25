package objectid

import (
	"encoding/base64"
	"strings"
	"testing"
	"time"
)

func Test_New12_New16(t *testing.T) {
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

func Test_ObjectidBase64(t *testing.T) {
	b64 := base64.StdEncoding
	id16 := New16()
	t.Logf("id16: %v", id16)

	s16 := b64.EncodeToString(id16)
	t.Logf("base64 for id16: %s, len %d", s16, len(s16))

	id12 := New12()
	t.Logf("id12: %v", id12)

	s12 := b64.EncodeToString(id12)
	t.Logf("base64 for id12: %s, len %d", s12, len(s12))
}

func Test_NewByHex(t *testing.T) {
	// Firstly, test a normal one
	o := New16()
	s := o.String()
	newO, err := NewByHex(s)
	if err != nil {
		t.Errorf("parsed objectid %s failed: %v", s, err)
		return
	}
	if newO.String() != s {
		t.Errorf("parse objectid %v, but %v got", s, newO)
		return
	}

	s = strings.ToUpper(s)
	newO, err = NewByHex(s)
	if err != nil {
		t.Errorf("parsed objectid %s failed: %v", s, err)
		return
	}
	if newO.String() != strings.ToLower(s) {
		t.Errorf("parse objectid %v, but %v got", s, newO)
		return
	}

	o = New12()
	s = o.String()
	newO, err = NewByHex(s)
	if err != nil {
		t.Errorf("parsed objectid %s failed: %v", s, err)
		return
	}
	if newO.String() != s {
		t.Errorf("parse objectid %v, but %v got", s, newO)
		return
	}

	// Test error one - string illegal
	s = s[:23] + "H" // generate a illegal one
	_, err = NewByHex(s)
	if err == nil {
		t.Errorf("error expected for %s, but no error got", s)
		return
	}
	t.Logf("expected return: %v", err)

	// Test error one - length illegal
	s = New16().String()
	s += "AB"
	t.Logf("length for s: %d", len(s))
	_, err = NewByHex(s)
	if err == nil {
		t.Errorf("error expected for %s, but no error got", s)
		return
	}
	t.Logf("expected return: %v", err)

	return
}

func Test_NewByBytes(t *testing.T) {
	now := time.Now()
	b12 := []byte(New12(now))
	o12, err := NewByBytes(b12)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if o12.Time().Unix() != now.Unix() {
		t.Errorf("time not equal for o12")
		return
	}

	b16 := []byte(New16(now))
	o16, err := NewByBytes(b16)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if false == o16.Time().Equal(now) {
		t.Errorf("time not equal for o16")
		return
	}

	// test error
	_, err = NewByBytes(nil)
	if err == nil {
		t.Errorf("error expected but not caught")
		return
	}

	_, err = NewByBytes([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if err == nil {
		t.Errorf("error expected but not caught")
		return
	}

	return
}
