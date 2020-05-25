package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"
)

const (
	// Each EPOCH is 30 seconds
	POODLE_EPOCH_MILLIS = 30 * 1000

	DEFAULT_DRIFT_MILLIS_LOW  = 300
	DEFAULT_DRIFT_MILLIS_HIGH = 500

	DEFAULT_ETC_DIR = "/etc/poodle"
	DEFAULT_LIB_DIR = "/var/lib/poodle"
	DEFAULT_LOG_DIR = "/var/log/poodle"

	DEFAULT_SECRET = "poodle"

	DEFAULT_PUDP_PORT = 31415
	DEFAULT_QUIC_PORT = 31416

	MAX_KEY_LENGTH    = 4 * 1024    // Maximum  4 KB Key Length
	MAX_VALUE_LENGTH  = 56 * 1024   // Maximum 56 KB Value Length
	MAX_SCHEME_LENGTH = 2 * 1024    // Maximum  2 KB Scheme Length
	MAX_ATTR_GROUPS   = 256         // maximum 256 Attribute Groups per Key
	MAX_DATA_LENGTH   = 64*1024 - 1 // Maximum 64 KB - 1 Data Length
	MAX_PACKET_LENGTH = 64*1024 - 1 // Maximum 64 KB - 1 Packet Length

	CLS_NODE       = 1
	CLS_CLUSTER    = 2
	CLS_UNIVERSE   = 3
	CLS_SERVICE    = 4
	CLS_FEDERATION = 5
)

////////////////////////////////////////////////////////////////////////////////
// utilities

func µ(a ...interface{}) []interface{} {
	return a
}

func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}

func EqByteArray(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func EqUint16Array(a, b []uint16) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func EqUint32Array(a, b []uint32) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func EqUint64Array(a, b []uint64) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Int64ToTime(nano int64) *time.Time {
	t := time.Unix(0, nano)
	return &t
}

func BytesToTime(buf []byte) (*time.Time, error) {
	if len(buf) < 8 {
		return nil, fmt.Errorf("BytesToTime - buf length less than 8 bytes [%x]", buf)
	}
	nano := binary.BigEndian.Uint64(buf[:8])
	return Int64ToTime(int64(nano)), nil
}

func TimeToBytes(t *time.Time) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(t.UnixNano()))
	return buf
}

func ByteArrayToBigInt(data []byte) *big.Int {
	result := new(big.Int)
	result.SetBytes(data)
	return result
}

func BigIntToByteArray(d *big.Int) []byte {
	input := d.Bytes()
	if len(input) == 32 {
		return input
	} else if len(input) > 32 {
		return input[len(input)-32:]
	} else {
		buf := make([]byte, 32-len(input))
		return append(buf, input[:]...)
	}
}

func Int64ToByteArray(input int64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, uint64(input))
	return result
}

func ByteArrayToInt64(buf []byte) int64 {
	var data uint64
	err := binary.Read(bytes.NewReader(buf), binary.BigEndian, &data)
	if err != nil {
		return 0
	}
	return int64(data)
}

func Int32ToByteArray(input int32) []byte {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(input))
	return result
}

func ByteArrayToInt32(buf []byte) int32 {
	var data uint32
	err := binary.Read(bytes.NewReader(buf), binary.BigEndian, &data)
	if err != nil {
		return 0
	}
	return int32(data)
}

func Uint32ToByteArray(input uint32) []byte {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(input))
	return result
}

func ByteArrayToUint32(buf []byte) uint32 {
	var data uint32
	err := binary.Read(bytes.NewReader(buf), binary.BigEndian, &data)
	if err != nil {
		return 0
	}
	return data
}
