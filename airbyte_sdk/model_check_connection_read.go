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

// CheckConnectionRead struct for CheckConnectionRead
type CheckConnectionRead struct {
	Status  string             `json:"status"`
	Message *string            `json:"message,omitempty"`
	JobInfo SynchronousJobRead `json:"jobInfo"`
}

// NewCheckConnectionRead instantiates a new CheckConnectionRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckConnectionRead(status string, jobInfo SynchronousJobRead) *CheckConnectionRead {
	this := CheckConnectionRead{}
	this.Status = status
	this.JobInfo = jobInfo
	return &this
}

// NewCheckConnectionReadWithDefaults instantiates a new CheckConnectionRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckConnectionReadWithDefaults() *CheckConnectionRead {
	this := CheckConnectionRead{}
	return &this
}

// GetStatus returns the Status field value
func (o *CheckConnectionRead) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CheckConnectionRead) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CheckConnectionRead) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CheckConnectionRead) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckConnectionRead) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CheckConnectionRead) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CheckConnectionRead) SetMessage(v string) {
	o.Message = &v
}

// GetJobInfo returns the JobInfo field value
func (o *CheckConnectionRead) GetJobInfo() SynchronousJobRead {
	if o == nil {
		var ret SynchronousJobRead
		return ret
	}

	return o.JobInfo
}

// GetJobInfoOk returns a tuple with the JobInfo field value
// and a boolean to check if the value has been set.
func (o *CheckConnectionRead) GetJobInfoOk() (*SynchronousJobRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobInfo, true
}

// SetJobInfo sets field value
func (o *CheckConnectionRead) SetJobInfo(v SynchronousJobRead) {
	o.JobInfo = v
}

func (o CheckConnectionRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["jobInfo"] = o.JobInfo
	}
	return json.Marshal(toSerialize)
}

type NullableCheckConnectionRead struct {
	value *CheckConnectionRead
	isSet bool
}

func (v NullableCheckConnectionRead) Get() *CheckConnectionRead {
	return v.value
}

func (v *NullableCheckConnectionRead) Set(val *CheckConnectionRead) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckConnectionRead) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckConnectionRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckConnectionRead(val *CheckConnectionRead) *NullableCheckConnectionRead {
	return &NullableCheckConnectionRead{value: val, isSet: true}
}

func (v NullableCheckConnectionRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckConnectionRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}