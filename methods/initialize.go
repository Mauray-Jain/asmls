package methods

import "encoding/json"

func Initialize(params *json.RawMessage) (any, error) {
	res := map[string]any{
		"capabilities": map[string]any{},
		"serverInfo": map[string]string{
			"name":    "asmls",
			"version": "0.0.1",
		},
	}
	return res, nil
}

func InitializedNotif(_ *json.RawMessage) (error) {
	return nil
}
