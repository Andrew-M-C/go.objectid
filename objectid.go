package objectid

import (
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/valyala/fastrand"
)

// ObjectID is a series of bytes with length of 12 or 16
type ObjectID []byte

// Len is equivalent with len(id)
func (id ObjectID) Len() int {
	if id == nil {
		return 0
	}
	return len(id)
}

// String returns object id in hex format
func (id ObjectID) String() string {
	if 0 == id.Len() {
		return "nil"
	}
	return hex.EncodeToString([]byte(id))
}

// Time returns the time information stored in object id
func (id ObjectID) Time() time.Time {
	switch id.Len() {
	default:
		return time.Time{}

	case 12:
		var objID primitive.ObjectID
		copy(objID[:], id[:12])
		return objID.Timestamp()

	case 16:
		var objID primitive.ObjectID
		copy(objID[:], id[:12])
		t := objID.Timestamp()
		nano := convBytesToNanosec(id[12:])
		if nano > 0 && nano <= 999999999 {
			t = t.Add(time.Duration(nano))
		}
		return t
	}
}

// New12 generates a standard Mongo Object ID
func New12(t ...time.Time) ObjectID {
	var id primitive.ObjectID

	if len(t) == 0 {
		id = primitive.NewObjectID()
	} else {
		id = primitive.NewObjectIDFromTimestamp(t[0])
	}

	return id[:]
}

// New16 generates a extended Mongo Object ID, with tailing 4 bytes of nanoseconds
func New16(t ...time.Time) ObjectID {
	var tm time.Time
	if 0 == len(t) {
		tm = time.Now()
	} else {
		tm = t[0]
	}

	id := primitive.NewObjectIDFromTimestamp(tm)
	var tailing []byte

	if nano := tm.Nanosecond(); nano == 0 {
		tailing = rand4Bytes()
	} else {
		tailing = convNanosecToBytes(nano)
	}

	return append(id[:], tailing...)
}

func rand4Bytes() []byte {
	u := fastrand.Uint32()
	b := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(b, u)
	b[0] |= 0x80
	return b
}

func convNanosecToBytes(nano int) []byte {
	b := []byte{0, 0, 0, 0}
	binary.BigEndian.PutUint32(b, uint32(nano))
	return b
}

func convBytesToNanosec(b []byte) int {
	return int(binary.BigEndian.Uint32(b))
}
