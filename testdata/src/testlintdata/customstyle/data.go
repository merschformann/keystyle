package customstyle

// CustomData uses a custom regex that requires keys to start with "key_" followed by lowercase letters.
type CustomData map[string]any

func literalValid() {
	_ = CustomData{
		"key_alpha": "v1",
		"key_beta":  "v2",
	}
}

func literalInvalid() {
	_ = CustomData{
		"key_alpha": "v1",
		"badkey":    "v2", // want `Key 'badkey' style should be custom`
	}
}

const (
	validCustomConst   = "key_gamma"
	invalidCustomConst = "gamma"
)

func constKeyValid() {
	_ = CustomData{validCustomConst: "v"}
}

func constKeyInvalid() {
	_ = CustomData{invalidCustomConst: "v"} // want `Key 'gamma' \(from identifier 'invalidCustomConst'\) style should be custom`
}
