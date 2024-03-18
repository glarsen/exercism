package flatten

func Flatten(nested interface{}) []interface{} {
	flattened := make([]interface{}, 0)

	switch list := nested.(type) {
	case []interface{}:
		for _, element := range list {
			flattened = append(flattened, Flatten(element)...)
		}
	case interface{}:
		flattened = append(flattened, nested)
	}

	return flattened
}
