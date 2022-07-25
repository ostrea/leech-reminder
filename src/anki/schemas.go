package anki

type requestSchema[T string | []uint64] struct {
	Action  string       `json:"action"`
	Version int          `json:"version"`
	Params  map[string]T `json:"params"`
}

type responseSchema[T []uint64 | []noteSchema] struct {
	Result T       `json:"result"`
	Error  *string `json:"error"`
}

type noteSchema struct {
	Fields map[string]fieldSchema `json:"fields"`
}

type fieldSchema struct {
	Value string `json:"value"`
}
