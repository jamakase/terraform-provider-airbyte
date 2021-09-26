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

// UploadRead struct for UploadRead
type UploadRead struct {
	Status     string  `json:"status"`
	ResourceId *string `json:"resourceId,omitempty"`
}

// NewUploadRead instantiates a new UploadRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadRead(status string) *UploadRead {
	this := UploadRead{}
	this.Status = status
	return &this
}

// NewUploadReadWithDefaults instantiates a new UploadRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadReadWithDefaults() *UploadRead {
	this := UploadRead{}
	return &this
}

// GetStatus returns the Status field value
func (o *UploadRead) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UploadRead) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UploadRead) SetStatus(v string) {
	o.Status = v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *UploadRead) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadRead) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *UploadRead) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *UploadRead) SetResourceId(v string) {
	o.ResourceId = &v
}

func (o UploadRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	return json.Marshal(toSerialize)
}

type NullableUploadRead struct {
	value *UploadRead
	isSet bool
}

func (v NullableUploadRead) Get() *UploadRead {
	return v.value
}

func (v *NullableUploadRead) Set(val *UploadRead) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadRead) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadRead(val *UploadRead) *NullableUploadRead {
	return &NullableUploadRead{value: val, isSet: true}
}

func (v NullableUploadRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}