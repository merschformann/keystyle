package todo

// LogData is the map of which we want to lint the keys.
type LogData map[string]any

// -- String literal keys --

func literalValid() {
	_ = LogData{
		"keyCamel1": "value1",
		"keyTwo":    "value2",
	}
}

func literalInvalid() {
	_ = LogData{
		"keyCamel1":       "value1",
		"incorrect_key_1": "value2", // want `Key 'incorrect_key_1' style should be camelCase`
	}
}

// -- Pointer to LogData --

var somePointerData = &LogData{
	"camelCase":       "value3",
	"incorrect-key-3": "value4", // want `Key 'incorrect-key-3' style should be camelCase`
}

// -- Const keys: value is resolved and checked --

const (
	validConstKey   = "accountId"  // camelCase — no diagnostic
	invalidConstKey = "Account_Id" // not camelCase — expect diagnostic
)

func constKeyValid() {
	_ = LogData{validConstKey: "value"}
}

func constKeyInvalid() {
	_ = LogData{invalidConstKey: "value"} // want `Key 'Account_Id' \(from identifier 'invalidConstKey'\) style should be camelCase`
}

// -- Multiple consts in one literal; stops after first bad key --

const (
	anotherValidKey   = "userId"
	anotherInvalidKey = "User_Id"
)

func constKeyMixed() {
	_ = LogData{
		anotherValidKey:   "v1",
		anotherInvalidKey: "v2", // want `Key 'User_Id' \(from identifier 'anotherInvalidKey'\) style should be camelCase`
	}
}

// -- Var keys: value is not a compile-time constant, skip check --

var runtimeKey = "Not_CamelCase" // would fail style check if inspected

func varKeySkipped() {
	// runtimeKey is a var, not a const — its value cannot be known statically,
	// so the linter must skip it without reporting a diagnostic.
	_ = LogData{runtimeKey: "value"}
}
