package field

import "encoding/json"

// Field struct is used to understand if the key was set or not on JSON. This is important when parsing partial updates such as in a PATCH request where `nil` might actually mean `delete` or similar
type Field[T any] struct {
	value T
	IsSet bool // Set means the value has been set
}

// UnmarshalJSON is a custom function to unmarshal JSON into a Field
func (i *Field[T]) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.IsSet = true

	var temp T
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.value = temp
	return nil
}
