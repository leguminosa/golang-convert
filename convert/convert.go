package convert

import (
	"strconv"
	"strings"
)

// ToInt converts anything to int
func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch v := v.(type) {
	case uint:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case int:
		return v
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		result, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return 0
		}
		return result
	case []byte:
		result, err := strconv.Atoi(strings.TrimSpace(string(v)))
		if err != nil {
			return 0
		}
		return result
	default:
		return 0
	}
}

// ToString converts anything to string
func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case int:
		return strconv.Itoa(v)
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return ""
	}
}
