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

// SourceDefinitionSpecificationRead struct for SourceDefinitionSpecificationRead
type SourceDefinitionSpecificationRead struct {
	SourceDefinitionId string  `json:"sourceDefinitionId"`
	DocumentationUrl   *string `json:"documentationUrl,omitempty"`
	// The specification for what values are required to configure the sourceDefinition.
	ConnectionSpecification *map[string]interface{} `json:"connectionSpecification,omitempty"`
	AuthSpecification       *AuthSpecification      `json:"authSpecification,omitempty"`
	JobInfo                 SynchronousJobRead      `json:"jobInfo"`
}

// NewSourceDefinitionSpecificationRead instantiates a new SourceDefinitionSpecificationRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceDefinitionSpecificationRead(sourceDefinitionId string, jobInfo SynchronousJobRead) *SourceDefinitionSpecificationRead {
	this := SourceDefinitionSpecificationRead{}
	this.SourceDefinitionId = sourceDefinitionId
	this.JobInfo = jobInfo
	return &this
}

// NewSourceDefinitionSpecificationReadWithDefaults instantiates a new SourceDefinitionSpecificationRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceDefinitionSpecificationReadWithDefaults() *SourceDefinitionSpecificationRead {
	this := SourceDefinitionSpecificationRead{}
	return &this
}

// GetSourceDefinitionId returns the SourceDefinitionId field value
func (o *SourceDefinitionSpecificationRead) GetSourceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceDefinitionId
}

// GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionSpecificationRead) GetSourceDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceDefinitionId, true
}

// SetSourceDefinitionId sets field value
func (o *SourceDefinitionSpecificationRead) SetSourceDefinitionId(v string) {
	o.SourceDefinitionId = v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *SourceDefinitionSpecificationRead) GetDocumentationUrl() string {
	if o == nil || o.DocumentationUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceDefinitionSpecificationRead) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || o.DocumentationUrl == nil {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *SourceDefinitionSpecificationRead) HasDocumentationUrl() bool {
	if o != nil && o.DocumentationUrl != nil {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *SourceDefinitionSpecificationRead) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetConnectionSpecification returns the ConnectionSpecification field value if set, zero value otherwise.
func (o *SourceDefinitionSpecificationRead) GetConnectionSpecification() map[string]interface{} {
	if o == nil || o.ConnectionSpecification == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ConnectionSpecification
}

// GetConnectionSpecificationOk returns a tuple with the ConnectionSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceDefinitionSpecificationRead) GetConnectionSpecificationOk() (*map[string]interface{}, bool) {
	if o == nil || o.ConnectionSpecification == nil {
		return nil, false
	}
	return o.ConnectionSpecification, true
}

// HasConnectionSpecification returns a boolean if a field has been set.
func (o *SourceDefinitionSpecificationRead) HasConnectionSpecification() bool {
	if o != nil && o.ConnectionSpecification != nil {
		return true
	}

	return false
}

// SetConnectionSpecification gets a reference to the given map[string]interface{} and assigns it to the ConnectionSpecification field.
func (o *SourceDefinitionSpecificationRead) SetConnectionSpecification(v map[string]interface{}) {
	o.ConnectionSpecification = &v
}

// GetAuthSpecification returns the AuthSpecification field value if set, zero value otherwise.
func (o *SourceDefinitionSpecificationRead) GetAuthSpecification() AuthSpecification {
	if o == nil || o.AuthSpecification == nil {
		var ret AuthSpecification
		return ret
	}
	return *o.AuthSpecification
}

// GetAuthSpecificationOk returns a tuple with the AuthSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceDefinitionSpecificationRead) GetAuthSpecificationOk() (*AuthSpecification, bool) {
	if o == nil || o.AuthSpecification == nil {
		return nil, false
	}
	return o.AuthSpecification, true
}

// HasAuthSpecification returns a boolean if a field has been set.
func (o *SourceDefinitionSpecificationRead) HasAuthSpecification() bool {
	if o != nil && o.AuthSpecification != nil {
		return true
	}

	return false
}

// SetAuthSpecification gets a reference to the given AuthSpecification and assigns it to the AuthSpecification field.
func (o *SourceDefinitionSpecificationRead) SetAuthSpecification(v AuthSpecification) {
	o.AuthSpecification = &v
}

// GetJobInfo returns the JobInfo field value
func (o *SourceDefinitionSpecificationRead) GetJobInfo() SynchronousJobRead {
	if o == nil {
		var ret SynchronousJobRead
		return ret
	}

	return o.JobInfo
}

// GetJobInfoOk returns a tuple with the JobInfo field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionSpecificationRead) GetJobInfoOk() (*SynchronousJobRead, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobInfo, true
}

// SetJobInfo sets field value
func (o *SourceDefinitionSpecificationRead) SetJobInfo(v SynchronousJobRead) {
	o.JobInfo = v
}

func (o SourceDefinitionSpecificationRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceDefinitionId"] = o.SourceDefinitionId
	}
	if o.DocumentationUrl != nil {
		toSerialize["documentationUrl"] = o.DocumentationUrl
	}
	if o.ConnectionSpecification != nil {
		toSerialize["connectionSpecification"] = o.ConnectionSpecification
	}
	if o.AuthSpecification != nil {
		toSerialize["authSpecification"] = o.AuthSpecification
	}
	if true {
		toSerialize["jobInfo"] = o.JobInfo
	}
	return json.Marshal(toSerialize)
}

type NullableSourceDefinitionSpecificationRead struct {
	value *SourceDefinitionSpecificationRead
	isSet bool
}

func (v NullableSourceDefinitionSpecificationRead) Get() *SourceDefinitionSpecificationRead {
	return v.value
}

func (v *NullableSourceDefinitionSpecificationRead) Set(val *SourceDefinitionSpecificationRead) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceDefinitionSpecificationRead) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceDefinitionSpecificationRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceDefinitionSpecificationRead(val *SourceDefinitionSpecificationRead) *NullableSourceDefinitionSpecificationRead {
	return &NullableSourceDefinitionSpecificationRead{value: val, isSet: true}
}

func (v NullableSourceDefinitionSpecificationRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceDefinitionSpecificationRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}