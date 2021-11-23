package airbyte

import "encoding/json"

// unmarshalConnectorConfigJSON is a convenience func for unmarshalling
// `config_json` field.
func unmarshalConnectorConfigJSON(configJSON string) (map[string]interface{}, error) {
	dashboardJSON := map[string]interface{}{}
	err := json.Unmarshal([]byte(configJSON), &dashboardJSON)
	if err != nil {
		return nil, err
	}
	return dashboardJSON, nil
}

// validateConnectorConfigJSON is the ValidateFunc for `config_json`. It
// ensures its value is valid JSON.
func validateConnectorConfigJSON(config interface{}, k string) ([]string, []error) {
	configJSON := config.(string)
	configMap := map[string]interface{}{}
	err := json.Unmarshal([]byte(configJSON), &configMap)
	if err != nil {
		return nil, []error{err}
	}
	return nil, nil
}

// normalizeConnectorConfigJSON is the StateFunc for the `config_json` field.
//
func normalizeConnectorConfigJSON(config interface{}) string {
	var dashboardJSON map[string]interface{}
	switch c := config.(type) {
	case map[string]interface{}:
		dashboardJSON = c
	case string:
		var err error
		dashboardJSON, err = unmarshalConnectorConfigJSON(c)
		if err != nil {
			return c
		}
	}
	j, _ := json.Marshal(dashboardJSON)
	return string(j)
}
