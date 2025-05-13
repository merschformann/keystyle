package todo

// LogData is the map of which we want to lint the keys.
type LogData map[string]any

func SomeFunc1() {
	data := LogData{
		"key1": "value1",
	}
	_ = data["key1"]
}

var (
	someData = LogData{
		"camelCase":  "value1",
		"PascalCase": "value2",
	}
)
