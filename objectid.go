/*
Package objectid provides MongoDB-objectId-like with 16-byte-length support (32 bytes in string). This could be replacement of UUID.

The simplest way to implement 16-bytes long object ID (32 bytes as hex string) is adding random binaries in the tailing. However, as nanoseconds are also quite randomized enough, thus I use nanoseconds instead.

	This objectId has many advantages comparing to UUID:
	1. An ObjectID has timestamp information, which can be very useful in many case. Such as database sharding.
	2. An ObjectID is lead by timestamp, therefore it can be ordered.
*/
package objectid

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/valyala/fastrand"
)

// ObjectID is a series of bytes with length of 12 or 16. The 16-bytes-lengthed ObjectID is compatable with standard MongoDB objectId.
type ObjectID []byte

// Len is equivalent with len(id)
func (id ObjectID) Len() int {
	if id == nil {
		return 0
	}
	return len(id)
}

// String returns object id in hex format.
func (id ObjectID) String() string {
	if 0 == id.Len() {
		return "nil"
	}
	return hex.EncodeToString([]byte(id))
}

// Time returns the time information stored in object ID. If the id is extended (16 bytes length), nanoseconds will also be parsed.
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

// New12 generates a standard MongoDB object ID. The time.Time parameter is optional. If no time given, it generates ObjectID with current time.
func New12(t ...time.Time) ObjectID {
	var id primitive.ObjectID

	if len(t) == 0 {
		id = primitive.NewObjectID()
	} else {
		id = primitive.NewObjectIDFromTimestamp(t[0])
	}

	return id[:]
}

// New16 generates a extended MongoDB object ID, with tailing 4 bytes of nanoseconds. The time.Time parameter is optional. If no time given, it generates ObjectID with current time.
func New16(t ...time.Time) ObjectID {
	var tm time.Time
	if 0 == len(t) {
		tm = time.Now()
	} else {
		tm = t[0]
	}

	id := primitive.NewObjectIDFromTimestamp(tm)
	tailing := convTimeToTailingBytes(&tm)
	return append(id[:], tailing...)
}

// NewByHex parse a objectid from given hex string
func NewByHex(s string) (ObjectID, error) {
	switch l := len(s); l {
	case 24, 32:
		// OK, continue
	default:
		return nil, fmt.Errorf("invalid hex length %d", l)
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return ObjectID(b), nil
}

// NewByBytes parse a objectid from given byte slice
func NewByBytes(b []byte) (ObjectID, error) {
	if nil == b {
		return nil, fmt.Errorf("nil bytes")
	}
	switch l := len(b); l {
	case 12, 16:
		return ObjectID(b), nil
	default:
		return nil, fmt.Errorf("invalid hex length %d", l)
	}
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

func convTimeToTailingBytes(t *time.Time) []byte {
	nano := t.Nanosecond()
	if 0 == nano {
		return rand4Bytes()
	}
	return convNanosecToBytes(nano)
}
