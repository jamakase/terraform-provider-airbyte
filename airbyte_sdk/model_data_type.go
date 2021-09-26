/*
Airbyte Configuration API

Airbyte Configuration API [https://airbyte.io](https://airbyte.io).  This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.  Here are some conventions that this API follows: * All endpoints are http POST methods. * All endpoints accept data via `application/json` request bodies. The API does not accept any data via query params. * The naming convention for endpoints is: localhost:8000/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. `localhost:8000/v1/connections/create`. * For all `update` methods, the whole object must be passed in, even the fields that did not change.  Change Management: * The major version of the API endpoint can be determined / specified in the URL `localhost:8080/v1/connections/create` * Minor version bumps will be invisible to the end user. The user cannot specify minor versions in requests. * All backwards incompatible changes will happen in major version bumps. We will not make backwards incompatible changes in minor version bumps. Examples of non-breaking changes (includes but not limited to...):   * Adding fields to request or response bodies.   * Adding new HTTP endpoints.

API version: 1.0.0
Contact: contact@airbyte.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airbyte_sdk

import (
	"encoding/json"
	"fmt"
)

// DataType the model 'DataType'
type DataType string

// List of DataType
const (
	DATATYPE_STRING  DataType = "string"
	DATATYPE_NUMBER  DataType = "number"
	DATATYPE_BOOLEAN DataType = "boolean"
	DATATYPE_OBJECT  DataType = "object"
	DATATYPE_ARRAY   DataType = "array"
)

var allowedDataTypeEnumValues = []DataType{
	"string",
	"number",
	"boolean",
	"object",
	"array",
}

func (v *DataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataType(value)
	for _, existing := range allowedDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataType", value)
}

// NewDataTypeFromValue returns a pointer to a valid DataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataTypeFromValue(v string) (*DataType, error) {
	ev := DataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataType: valid values are %v", v, allowedDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataType) IsValid() bool {
	for _, existing := range allowedDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataType value
func (v DataType) Ptr() *DataType {
	return &v
}

type NullableDataType struct {
	value *DataType
	isSet bool
}

func (v NullableDataType) Get() *DataType {
	return v.value
}

func (v *NullableDataType) Set(val *DataType) {
	v.value = val
	v.isSet = true
}

func (v NullableDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataType(val *DataType) *NullableDataType {
	return &NullableDataType{value: val, isSet: true}
}

func (v NullableDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
