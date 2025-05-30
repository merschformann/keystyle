package todo

// LogData is the map of which we want to lint the keys.
type LogData map[string]any

// TODO: ajsdlasj
func SomeFunc1() {
	data := LogData{
		"keyCamel1":       "value1",
		"incorrect_key_1": "value2",
	}
	_ = data["key1"]
	_ = someData
	_ = somePointerData
}

var (
	someData = LogData{
		"camelCase":     "value1",
		"IncorrectKey2": "value2",
	}
	somePointerData = &LogData{
		"camelCase":       "value3",
		"incorrect-key-3": "value4",
	}
)
