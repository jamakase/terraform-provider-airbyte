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

// JobRead struct for JobRead
type JobRead struct {
	Id         int64         `json:"id"`
	ConfigType JobConfigType `json:"configType"`
	ConfigId   string        `json:"configId"`
	CreatedAt  int64         `json:"createdAt"`
	UpdatedAt  int64         `json:"updatedAt"`
	Status     JobStatus     `json:"status"`
}

// NewJobRead instantiates a new JobRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobRead(id int64, configType JobConfigType, configId string, createdAt int64, updatedAt int64, status JobStatus) *JobRead {
	this := JobRead{}
	this.Id = id
	this.ConfigType = configType
	this.ConfigId = configId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	return &this
}

// NewJobReadWithDefaults instantiates a new JobRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobReadWithDefaults() *JobRead {
	this := JobRead{}
	return &this
}

// GetId returns the Id field value
func (o *JobRead) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobRead) SetId(v int64) {
	o.Id = v
}

// GetConfigType returns the ConfigType field value
func (o *JobRead) GetConfigType() JobConfigType {
	if o == nil {
		var ret JobConfigType
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetConfigTypeOk() (*JobConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *JobRead) SetConfigType(v JobConfigType) {
	o.ConfigType = v
}

// GetConfigId returns the ConfigId field value
func (o *JobRead) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value
func (o *JobRead) SetConfigId(v string) {
	o.ConfigId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *JobRead) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *JobRead) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *JobRead) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *JobRead) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *JobRead) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobRead) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobRead) SetStatus(v JobStatus) {
	o.Status = v
}

func (o JobRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["configType"] = o.ConfigType
	}
	if true {
		toSerialize["configId"] = o.ConfigId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableJobRead struct {
	value *JobRead
	isSet bool
}

func (v NullableJobRead) Get() *JobRead {
	return v.value
}

func (v *NullableJobRead) Set(val *JobRead) {
	v.value = val
	v.isSet = true
}

func (v NullableJobRead) IsSet() bool {
	return v.isSet
}

func (v *NullableJobRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobRead(val *JobRead) *NullableJobRead {
	return &NullableJobRead{value: val, isSet: true}
}

func (v NullableJobRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
