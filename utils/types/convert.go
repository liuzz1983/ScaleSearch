package types

import "errors"

func ToBytes(value interface{}) ([]byte, error) {
	switch _value := value.(type) {
	case string:
		return []byte(_value), nil
	case []byte:
		return _value, nil
	default:
		return nil, errors.New("cant convert the value")
	}
}

func ToFloat(value interface{}) (float64, error) {

	switch _value := value.(type) {

	case uint8:
		return float64(_value), nil
	case uint16:
		return float64(_value), nil
	case uint32:
		return float64(_value), nil
	case uint64:
		return float64(_value), nil
	case uint:
		return float64(_value), nil
	case int8:
		return float64(_value), nil
	case int16:
		return float64(_value), nil
	case int32:
		return float64(_value), nil
	case int64:
		return float64(_value), nil
	case int:
		return float64(_value), nil
	case float32:
		return float64(_value), nil
	case float64:
		return float64(_value), nil
	default:
		return 0.0, errors.New("wrong type")
	}
}
