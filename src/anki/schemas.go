package anki

type RequestSchema[T string | []uint64] struct {
	Action  string       `json:"action"`
	Version int          `json:"version"`
	Params  map[string]T `json:"params"`
}

type ResponseSchema[T []uint64 | []NoteSchema] struct {
	Result T       `json:"result"`
	Error  *string `json:"error"`
}

type NoteSchema struct {
	Fields map[string]FieldSchema `json:"fields"`
}

type FieldSchema struct {
	Value string `json:"value"`
}
