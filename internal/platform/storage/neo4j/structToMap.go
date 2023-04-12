package neo4j

import "reflect"

func structToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	valueType := reflect.TypeOf(data)
	value := reflect.ValueOf(data)

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		if tag, ok := field.Tag.Lookup("json"); ok {
			var name string
			if tag == "" {
				name = field.Name
			} else {
				name = tag
			}
			result[name] = value.Field(i).Interface()
		}
	}

	return result
}
