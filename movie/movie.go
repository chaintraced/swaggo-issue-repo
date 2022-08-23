package movie

import "encoding/json"

type CreateMovie struct {
	Name  string
	Actor Field[Actor]
}

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

type Actor struct {
	Name string
}
