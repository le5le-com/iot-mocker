package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/le5le-com/uuid"
)

type Json map[string]interface{}

func Int(v interface{}) int {
	switch reply := v.(type) {
	case int:
		return reply
	case int8:
		return int(reply)
	case int16:
		return int(reply)
	case int32:
		return int(reply)
	case int64:
		return int(reply)
	case uint:
		return int(reply)
	case uint8:
		return int(reply)
	case uint16:
		return int(reply)
	case uint32:
		return int(reply)
	case uint64:
		return int(reply)
	case []byte:
		return int(binary.BigEndian.Uint64(reply))
	case string:
		n, _ := strconv.ParseInt(reply, 10, 0)
		return int(n)
	case nil:
		return 0
	case float64:
		return int(reply)
	}

	return 0
}

func Int2(v interface{}) (int, error) {
	switch reply := v.(type) {
	case int:
		return reply, nil
	case int8:
		return int(reply), nil
	case int16:
		return int(reply), nil
	case int32:
		return int(reply), nil
	case int64:
		return int(reply), nil
	case uint:
		return int(reply), nil
	case uint8:
		return int(reply), nil
	case uint16:
		return int(reply), nil
	case uint32:
		return int(reply), nil
	case uint64:
		return int(reply), nil
	case []byte:
		return int(binary.BigEndian.Uint64(reply)), nil
	case string:
		n, err := strconv.ParseInt(reply, 10, 0)
		return int(n), err
	case nil:
		return 0, nil
	case float64:
		return int(reply), nil
	}

	return 0, errors.New("unknown")
}

func Int8(v interface{}) int8 {
	switch reply := v.(type) {
	case int:
		return int8(reply)
	case int8:
		return reply
	case int16:
		return int8(reply)
	case int32:
		return int8(reply)
	case int64:
		return int8(reply)
	case uint:
		return int8(reply)
	case uint8:
		return int8(reply)
	case uint16:
		return int8(reply)
	case uint32:
		return int8(reply)
	case uint64:
		return int8(reply)
	case string:
		n, _ := strconv.ParseInt(reply, 10, 8)
		return int8(n)
	case nil:
		return 0
	case float64:
		return int8(reply)
	}

	return 0
}

func Int16(v interface{}) int16 {
	switch reply := v.(type) {
	case int:
		return int16(reply)
	case int8:
		return int16(reply)
	case int16:
		return int16(reply)
	case int32:
		return int16(reply)
	case int64:
		return int16(reply)
	case uint:
		return int16(reply)
	case uint8:
		return int16(reply)
	case uint16:
		return int16(reply)
	case uint32:
		return int16(reply)
	case uint64:
		return int16(reply)
	case []byte:
		return int16(binary.BigEndian.Uint16(reply))
	case string:
		n, _ := strconv.ParseInt(reply, 10, 16)
		return int16(n)
	case nil:
		return 0
	case float64:
		return int16(reply)
	}

	return 0
}

func Uint16(v interface{}) uint16 {
	switch reply := v.(type) {
	case int:
		return uint16(reply)
	case int8:
		return uint16(reply)
	case int16:
		return uint16(reply)
	case int32:
		return uint16(reply)
	case int64:
		return uint16(reply)
	case uint:
		return uint16(reply)
	case uint8:
		return uint16(reply)
	case uint16:
		return uint16(reply)
	case uint32:
		return uint16(reply)
	case uint64:
		return uint16(reply)
	case []byte:
		return binary.BigEndian.Uint16(reply)
	case string:
		n, _ := strconv.ParseInt(reply, 10, 16)
		return uint16(n)
	case nil:
		return 0
	case float64:
		return uint16(reply)
	case bool:
		if reply {
			return 1
		}
		return 0
	}

	return 0
}

func Int32(v interface{}) int32 {
	switch reply := v.(type) {
	case int:
		return int32(reply)
	case int8:
		return int32(reply)
	case int16:
		return int32(reply)
	case int32:
		return int32(reply)
	case int64:
		return int32(reply)
	case uint:
		return int32(reply)
	case uint8:
		return int32(reply)
	case uint16:
		return int32(reply)
	case uint32:
		return int32(reply)
	case uint64:
		return int32(reply)
	case []byte:
		return int32(binary.BigEndian.Uint32(reply))
	case string:
		n, _ := strconv.ParseInt(reply, 10, 32)
		return int32(n)
	case nil:
		return 0
	case float64:
		return int32(reply)
	}

	return 0
}

func Uint32(v interface{}) uint32 {
	switch reply := v.(type) {
	case int:
		return uint32(reply)
	case int8:
		return uint32(reply)
	case int16:
		return uint32(reply)
	case int32:
		return uint32(reply)
	case int64:
		return uint32(reply)
	case uint:
		return uint32(reply)
	case uint8:
		return uint32(reply)
	case uint16:
		return uint32(reply)
	case uint32:
		return uint32(reply)
	case uint64:
		return uint32(reply)
	case float64:
		return uint32(reply)
	case []byte:
		return binary.BigEndian.Uint32(reply)
	case string:
		n, _ := strconv.ParseInt(reply, 10, 32)
		return uint32(n)
	case nil:
		return 0
	}

	return 0
}

func Int64(v interface{}) int64 {
	switch reply := v.(type) {
	case int:
		return int64(reply)
	case int8:
		return int64(reply)
	case int16:
		return int64(reply)
	case int32:
		return int64(reply)
	case int64:
		return reply
	case uint:
		return int64(reply)
	case uint8:
		return int64(reply)
	case uint16:
		return int64(reply)
	case uint32:
		return int64(reply)
	case uint64:
		return int64(reply)
	case []byte:
		return int64(binary.BigEndian.Uint64(reply))
	case string:
		n, _ := strconv.ParseInt(reply, 10, 64)
		return n
	case float64:
		return int64(reply)

	case nil:
		return 0
	}

	return 0
}

func Uint64(v interface{}) uint64 {
	switch reply := v.(type) {
	case int:
		return uint64(reply)
	case int8:
		return uint64(reply)
	case int16:
		return uint64(reply)
	case int32:
		return uint64(reply)
	case int64:
		return uint64(reply)
	case uint:
		return uint64(reply)
	case uint8:
		return uint64(reply)
	case uint16:
		return uint64(reply)
	case uint32:
		return uint64(reply)
	case uint64:
		return uint64(reply)
	case []byte:
		return binary.BigEndian.Uint64(reply)
	case string:
		n, _ := strconv.ParseInt(reply, 10, 64)
		return uint64(n)
	case nil:
		return 0
	case float64:
		return uint64(reply)
	}

	return 0
}

func I64(v interface{}) (int64, error) {
	switch reply := v.(type) {
	case int:
		return int64(reply), nil
	case int8:
		return int64(reply), nil
	case int16:
		return int64(reply), nil
	case int32:
		return int64(reply), nil
	case int64:
		return reply, nil
	case uint:
		return int64(reply), nil
	case uint8:
		return int64(reply), nil
	case uint16:
		return int64(reply), nil
	case uint32:
		return int64(reply), nil
	case uint64:
		return int64(reply), nil
	case []byte:
		return int64(binary.BigEndian.Uint64(reply)), nil
	case string:
		return strconv.ParseInt(reply, 10, 64)
	case nil:
		return 0, errors.New("Nil")
	case float64:
		return int64(reply), nil
	}

	return 0, errors.New("Unknown")
}

func Float32(v interface{}) float32 {
	switch reply := v.(type) {
	case nil:
		return 0.0
	case string:
		f, _ := strconv.ParseFloat(reply, 32)
		return float32(f)
	case int8:
		return float32(reply)
	case int16:
		return float32(reply)
	case int32:
		return float32(reply)
	case int64:
		return float32(reply)
	case uint:
		return float32(reply)
	case uint8:
		return float32(reply)
	case uint16:
		return float32(reply)
	case uint32:
		return float32(reply)
	case uint64:
		return float32(reply)
	case float32:
		return reply
	case float64:
		return float32(reply)
	}

	return 0.0
}

func Float64(v interface{}) float64 {
	switch reply := v.(type) {
	case nil:
		return 0.0
	case string:
		f, _ := strconv.ParseFloat(reply, 64)
		return f
	case int8:
		return float64(reply)
	case int16:
		return float64(reply)
	case int32:
		return float64(reply)
	case int64:
		return float64(reply)
	case uint:
		return float64(reply)
	case uint8:
		return float64(reply)
	case uint16:
		return float64(reply)
	case uint32:
		return float64(reply)
	case uint64:
		return float64(reply)
	case float32:
		return float64(reply)
	case float64:
		return reply
	}

	return 0.0
}

func F64(v interface{}) (float64, error) {
	switch reply := v.(type) {
	case nil:
		return 0.0, errors.New("Nil")
	case string:
		return strconv.ParseFloat(reply, 64)
	case int8:
		return float64(reply), nil
	case int16:
		return float64(reply), nil
	case int32:
		return float64(reply), nil
	case int64:
		return float64(reply), nil
	case uint:
		return float64(reply), nil
	case uint8:
		return float64(reply), nil
	case uint16:
		return float64(reply), nil
	case uint32:
		return float64(reply), nil
	case uint64:
		return float64(reply), nil
	case float32:
		return float64(reply), nil
	case float64:
		return reply, nil
	}

	return 0.0, errors.New("Unknown")
}

func String(v interface{}) string {
	switch reply := v.(type) {
	case string:
		return reply
	case int:
		return strconv.Itoa(reply)
	case uint:
		return strconv.Itoa(int(reply))
	case int8:
		return strconv.Itoa(int(reply))
	case uint8:
		return strconv.Itoa(int(reply))
	case int16:
		return strconv.Itoa(int(reply))
	case uint16:
		return strconv.Itoa(int(reply))
	case int32:
		return strconv.Itoa(int(reply))
	case uint32:
		return strconv.Itoa(int(reply))
	case int64:
		return strconv.Itoa(int(reply))
	case uint64:
		return strconv.Itoa(int(reply))
	case float32:
		return strconv.FormatFloat(float64(reply), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(reply, 'f', -1, 64)
	case []byte:
		return *(*string)(unsafe.Pointer(&reply))
	case uuid.UUID:
		return reply.ShortString()
	case nil:
		return ""
	case bool:
		if reply {
			return "true"
		}
		return "false"
	case map[string]interface{}, Json:
		b, err := json.Marshal(reply)
		if err == nil {
			return string(b)
		}
		return err.Error()
	}

	return ""
}

func String2(v interface{}) (string, error) {
	switch reply := v.(type) {
	case string:
		return reply, nil
	case int:
		return strconv.Itoa(reply), nil
	case int8:
		return strconv.Itoa(int(reply)), nil
	case int16:
		return strconv.Itoa(int(reply)), nil
	case int32:
		return strconv.Itoa(int(reply)), nil
	case int64:
		return strconv.Itoa(int(reply)), nil
	case []byte:
		return *(*string)(unsafe.Pointer(&reply)), nil
	case nil:
		return "", nil
	case bool:
		if reply {
			return "true", nil
		}
		return "false", nil
	case map[string]interface{}:
		b, err := json.Marshal(reply)
		if err == nil {
			return string(b), err
		}
		return "", err
	}

	return "", errors.New("unknown")
}

func Strings(v interface{}) []string {
	var arr []string = make([]string, 0)
	switch reply := v.(type) {
	case []interface{}:
		for _, v := range reply {
			switch s := v.(type) {
			case string:
				arr = append(arr, s)
			}
		}
	case []string:
		return reply
	}

	return arr
}

func Bool(v interface{}) bool {
	switch reply := v.(type) {
	case bool:
		return reply
	case string:
		return reply == "true" || reply == "1"
	case int:
		return reply != 0
	case int32:
		return reply != 0
	case int64:
		return reply != 0
	}

	return false
}

func Bool2(v interface{}) (bool, error) {
	switch reply := v.(type) {
	case bool:
		return reply, nil
	case string:
		if reply == "true" || reply == "1" {
			return true, nil
		} else if reply == "false" || reply == "0" {
			return false, nil
		}
		return false, errors.New("Unknown")
	case int:
		return reply != 0, nil
	case int32:
		return reply != 0, nil
	case int64:
		return reply != 0, nil
	}

	return false, errors.New("Unknown")
}

func Byte(v interface{}) byte {
	switch reply := v.(type) {
	case byte:
		return reply
	case int8:
		return byte(reply)
	case int16:
		return byte(reply)
	case int32:
		return byte(reply)
	case int64:
		return byte(reply)
	case int:
		return byte(reply)
	case []byte:
		n, _ := strconv.ParseInt(string(reply), 16, 0)
		return byte(n)
	case string:
		n, _ := strconv.ParseInt(reply, 16, 0)
		return byte(n)
	case nil:
		return 0
	case float64:
		return byte(reply)
	}

	return 0
}

func Array(v interface{}) []interface{} {
	switch reply := v.(type) {
	case []interface{}:
		return reply
	case nil:
		return nil
	}

	return nil
}

func Bytes2Bits(data []byte) []uint8 {
	dst := make([]uint8, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint8(7 - i)
			dst = append(dst, uint8((v>>move)&1))
		}
	}
	return dst
}

func BinaryStringToBytes(b string) []byte {
	var out []byte
	var str string

	for i := len(b); i > 0; i -= 8 {
		if i-8 < 0 {
			str = string(b[0:i])
		} else {
			str = string(b[i-8 : i])
		}
		v, err := strconv.ParseUint(str, 2, 8)
		if err != nil {
			return nil
		}
		out = append([]byte{byte(v)}, out...)
	}
	return out
}

func Time(v interface{}) time.Time {
	switch reply := v.(type) {
	case time.Time:
		return reply
	case int:
		return time.Unix(int64(reply), 0)
	case int64:
		return time.Unix(reply, 0)
	case uint:
		return time.Unix(int64(reply), 0)
	case uint64:
		return time.Unix(int64(reply), 0)
	case string:
		t, _ := time.Parse("2006-01-02T15:04:05.999999999Z07", reply)
		return t
	}

	return time.Unix(0, 0)
}

func Unix2(v interface{}) (int64, error) {
	switch reply := v.(type) {
	case time.Time:
		return reply.Unix(), nil
	case int:
		return int64(reply), nil
	case int64:
		return reply, nil
	case uint:
		return int64(reply), nil
	case uint64:
		return int64(reply), nil
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", reply)
		if err == nil {
			return t.Unix(), nil
		}
		return 0, err
	}

	return 0, errors.New("Unknown")
}

func UUIDs(v []string) []uuid.UUID {
	var arr []uuid.UUID = make([]uuid.UUID, 0)
	for _, s := range v {
		id, err := uuid.Parse(s)
		if err == nil {
			arr = append(arr, id)
		}
	}

	return arr
}

func Between(v any, expr string) bool {
	expr = strings.TrimSpace(expr)
	parts := strings.Split(expr, ",")
	if len(parts) != 2 {
		return false
	}

	start, end := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	startOpen, endOpen := strings.HasPrefix(start, "("), strings.HasSuffix(end, ")")

	valFloat, f := v.(float64)
	if f {
		startVal, err := strconv.ParseFloat(strings.Trim(start, "()[] "), 64)
		if err != nil {
			return false
		}

		endVal, err := strconv.ParseFloat(strings.Trim(end, "()[] "), 64)
		if err != nil {
			return false
		}

		switch {
		case startOpen && endOpen: // (start, end)
			return startVal < valFloat && valFloat < endVal
		case startOpen && !endOpen: // (start, end]
			return startVal < valFloat && valFloat <= endVal
		case !startOpen && endOpen: // [start, end)
			return startVal <= valFloat && valFloat < endVal
		default: // [start, end]
			return startVal <= valFloat && valFloat <= endVal
		}
	} else {
		a, err := I64(v)
		if err != nil {
			return false
		}

		startVal, err := strconv.ParseInt(strings.Trim(start, "()[] "), 10, 64)
		if err != nil {
			return false
		}

		endVal, err := strconv.ParseInt(strings.Trim(end, "()[] "), 10, 64)
		if err != nil {
			return false
		}

		switch {
		case startOpen && endOpen: // (start, end)
			return startVal < a && a < endVal
		case startOpen && !endOpen: // (start, end]
			return startVal < a && a <= endVal
		case !startOpen && endOpen: // [start, end)
			return startVal <= a && a < endVal
		default: // [start, end]
			return startVal <= a && a <= endVal
		}
	}
}

func Belong(val any, expr string) bool {
	expr = strings.Trim(expr, "[] ")
	rangesAndValues := strings.Split(expr, ",")
	inRange := false

	for i := range rangesAndValues {
		rangesAndValues[i] = strings.TrimSpace(rangesAndValues[i])
	}

	// 判断val的类型，进行相应的比较
	switch val := val.(type) {
	case float64:
		for _, part := range rangesAndValues {
			if strings.Contains(part, "..") {
				// 处理范围
				bounds := strings.Split(part, "..")
				if len(bounds) == 2 {
					start, err1 := strconv.ParseFloat(strings.TrimSpace(bounds[0]), 64)
					end, err2 := strconv.ParseFloat(strings.TrimSpace(bounds[1]), 64)
					if err1 == nil && err2 == nil && val >= start && val <= end {
						inRange = true
						break
					}
				}
			} else {
				// 处理单个值
				value, err := strconv.ParseFloat(part, 64)
				if err == nil && val == value {
					inRange = true
					break
				}
			}
		}
	case string:
		valStr := val
		for _, part := range rangesAndValues {
			if valStr == part {
				inRange = true
				break
			}
		}
	default:
		val64, err := I64(val)
		if err != nil {
			return false
		}

		for _, part := range rangesAndValues {
			if strings.Contains(part, "..") {
				bounds := strings.Split(part, "..")
				if len(bounds) == 2 {
					start, err1 := strconv.ParseInt(strings.TrimSpace(bounds[0]), 10, 64)
					end, err2 := strconv.ParseInt(strings.TrimSpace(bounds[1]), 10, 64)
					if err1 == nil && err2 == nil && val64 >= start && val64 <= end {
						inRange = true
						break
					}
				}
			} else {
				// 处理单个值
				value, err := strconv.ParseInt(part, 10, 64)
				if err == nil && val == value {
					inRange = true
					break
				}
			}
		}
	}

	return inRange
}
