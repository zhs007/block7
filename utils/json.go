package block7utils

import (
	"bytes"

	"github.com/buger/jsonparser"
)

// GetJsonString - number to string
func GetJsonString(data []byte, keys ...string) (val string, err error) {
	v, t, _, e := jsonparser.Get(data, keys...)

	if e != nil {
		return "", e
	}

	if t == jsonparser.Number {
		return string(v), nil
		// if strings.Contains(string(v), ".") {
		// 	nf, err := jsonparser.ParseFloat(v)
		// 	if err != nil {
		// 		return "", err
		// 	}

		// 	str := strconv.FormatFloat(nf, 'E', -1, 64)

		// 	return str, nil
		// }

		// iv, err := jsonparser.ParseInt(v)
		// if err != nil {
		// 	return "", err
		// }

		// str := strconv.FormatInt(iv, 10)

		// return str, nil
	}

	if t != jsonparser.String {
		return "", ErrInvalidJsonString
	}

	// If no escapes return raw content
	if bytes.IndexByte(v, '\\') == -1 {
		return string(v), nil
	}

	return jsonparser.ParseString(v)
}

func GetJsonInt(data []byte, keys ...string) (val int64, err error) {
	v, t, _, e := jsonparser.Get(data, keys...)

	if e != nil {
		return 0, e
	}

	if t == jsonparser.String {
		// If no escapes return raw content
		if bytes.IndexByte(v, '\\') == -1 {
			n, err := String2Int64(string(v))
			if err != nil {
				return 0, err
			}

			return n, nil
		}

		s, err := jsonparser.ParseString(v)
		if err != nil {
			return 0, err
		}

		n, err := String2Int64(s)
		if err != nil {
			return 0, err
		}

		return n, nil
	}

	if t != jsonparser.Number {
		return 0, ErrInvalidJsonInt
	}

	return String2Int64(string(v))
}
