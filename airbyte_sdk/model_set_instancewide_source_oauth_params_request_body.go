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
)

// SetInstancewideSourceOauthParamsRequestBody struct for SetInstancewideSourceOauthParamsRequestBody
type SetInstancewideSourceOauthParamsRequestBody struct {
	SourceDefinitionId string                            `json:"sourceDefinitionId"`
	Params             map[string]map[string]interface{} `json:"params"`
}

// NewSetInstancewideSourceOauthParamsRequestBody instantiates a new SetInstancewideSourceOauthParamsRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetInstancewideSourceOauthParamsRequestBody(sourceDefinitionId string, params map[string]map[string]interface{}) *SetInstancewideSourceOauthParamsRequestBody {
	this := SetInstancewideSourceOauthParamsRequestBody{}
	this.SourceDefinitionId = sourceDefinitionId
	this.Params = params
	return &this
}

// NewSetInstancewideSourceOauthParamsRequestBodyWithDefaults instantiates a new SetInstancewideSourceOauthParamsRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetInstancewideSourceOauthParamsRequestBodyWithDefaults() *SetInstancewideSourceOauthParamsRequestBody {
	this := SetInstancewideSourceOauthParamsRequestBody{}
	return &this
}

// GetSourceDefinitionId returns the SourceDefinitionId field value
func (o *SetInstancewideSourceOauthParamsRequestBody) GetSourceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceDefinitionId
}

// GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *SetInstancewideSourceOauthParamsRequestBody) GetSourceDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDefinitionId, true
}

// SetSourceDefinitionId sets field value
func (o *SetInstancewideSourceOauthParamsRequestBody) SetSourceDefinitionId(v string) {
	o.SourceDefinitionId = v
}

// GetParams returns the Params field value
func (o *SetInstancewideSourceOauthParamsRequestBody) GetParams() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *SetInstancewideSourceOauthParamsRequestBody) GetParamsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *SetInstancewideSourceOauthParamsRequestBody) SetParams(v map[string]map[string]interface{}) {
	o.Params = v
}

func (o SetInstancewideSourceOauthParamsRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceDefinitionId"] = o.SourceDefinitionId
	}
	if true {
		toSerialize["params"] = o.Params
	}
	return json.Marshal(toSerialize)
}

type NullableSetInstancewideSourceOauthParamsRequestBody struct {
	value *SetInstancewideSourceOauthParamsRequestBody
	isSet bool
}

func (v NullableSetInstancewideSourceOauthParamsRequestBody) Get() *SetInstancewideSourceOauthParamsRequestBody {
	return v.value
}

func (v *NullableSetInstancewideSourceOauthParamsRequestBody) Set(val *SetInstancewideSourceOauthParamsRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetInstancewideSourceOauthParamsRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetInstancewideSourceOauthParamsRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetInstancewideSourceOauthParamsRequestBody(val *SetInstancewideSourceOauthParamsRequestBody) *NullableSetInstancewideSourceOauthParamsRequestBody {
	return &NullableSetInstancewideSourceOauthParamsRequestBody{value: val, isSet: true}
}

func (v NullableSetInstancewideSourceOauthParamsRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetInstancewideSourceOauthParamsRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}