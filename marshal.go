package main

import "fmt"

// MarshalDynamoValue make flatten dynamodb.AttributeValue
func MarshalDynamoValue(item interface{}) (interface{}, error) {
	var result interface{}
	for k, v := range item.(map[string]interface{}) {
		switch k {
		case "N", "BOOL", "S", "BS", "NULL", "NS", "SS":
			result = v
			break
		case "L":
			l := v.([]map[string]interface{})
			r := make([]interface{}, 0, len(l))
			for i := range l {
				v, err := MarshalDynamoValue(l[i])
				if err != nil {
					return nil, err
				}
				r = append(r, v)
			}
			result = r
			break
		case "M":
			r, err := MarshalDynamo(v)
			if err != nil {
				return nil, err
			}
			result = r
			break
		default:
			return nil, fmt.Errorf("unsupported format")
		}
	}

	return result, nil
}

// MarshalDynamo make flatten dynamodb.Item
func MarshalDynamo(item interface{}) (interface{}, error) {
	converted, ok := item.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data to marshal. item=%v", item)
	}

	result := make(map[string]interface{}, len(converted))
	for k1, v1 := range converted {
		v, err := MarshalDynamoValue(v1)
		if err != nil {
			return nil, err
		}

		result[k1] = v
	}

	return result, nil
}
