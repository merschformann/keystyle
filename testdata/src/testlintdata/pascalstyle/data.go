package pascalstyle

type EventData map[string]any

func literalValid() {
	_ = EventData{
		"UserName":  "alice",
		"AccountId": "123",
		"CreatedAt": "2024-01-01",
	}
}

func literalInvalid() {
	_ = EventData{
		"UserName":  "alice",
		"accountId": "123", // want `Key 'accountId' style should be PascalCase`
	}
}

var pointerData = &EventData{
	"ValidKey":    "v",
	"invalid_key": "v", // want `Key 'invalid_key' style should be PascalCase`
}

const (
	validPascalConst   = "OrderId"
	invalidPascalConst = "orderId"
)

func constKeyValid() {
	_ = EventData{validPascalConst: "v"}
}

func constKeyInvalid() {
	_ = EventData{invalidPascalConst: "v"} // want `Key 'orderId' \(from identifier 'invalidPascalConst'\) style should be PascalCase`
}
