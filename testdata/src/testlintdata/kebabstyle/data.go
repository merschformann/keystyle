package kebabstyle

type HeaderData map[string]any

func literalValid() {
	_ = HeaderData{
		"content-type": "application/json",
		"x-request-id": "abc123",
	}
}

func literalInvalid() {
	_ = HeaderData{
		"content-type": "application/json",
		"X-Request-Id": "abc123", // want `Key 'X-Request-Id' style should be kebab-case`
	}
}

const (
	validKebabConst   = "x-trace-id"
	invalidKebabConst = "X_Trace_Id"
)

func constKeyValid() {
	_ = HeaderData{validKebabConst: "v"}
}

func constKeyInvalid() {
	_ = HeaderData{invalidKebabConst: "v"} // want `Key 'X_Trace_Id' \(from identifier 'invalidKebabConst'\) style should be kebab-case`
}
