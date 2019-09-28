package extvals

// Object can hold any valid rpc values, stores its actual type in .T field.
type Object struct {
	T string
	V interface{}
}
