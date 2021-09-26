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

// DestinationOauthConsentRequest struct for DestinationOauthConsentRequest
type DestinationOauthConsentRequest struct {
	DestinationDefinitionId string `json:"destinationDefinitionId"`
	WorkspaceId             string `json:"workspaceId"`
	// The url to redirect to after getting the user consent
	RedirectUrl string `json:"redirectUrl"`
}

// NewDestinationOauthConsentRequest instantiates a new DestinationOauthConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationOauthConsentRequest(destinationDefinitionId string, workspaceId string, redirectUrl string) *DestinationOauthConsentRequest {
	this := DestinationOauthConsentRequest{}
	this.DestinationDefinitionId = destinationDefinitionId
	this.WorkspaceId = workspaceId
	this.RedirectUrl = redirectUrl
	return &this
}

// NewDestinationOauthConsentRequestWithDefaults instantiates a new DestinationOauthConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationOauthConsentRequestWithDefaults() *DestinationOauthConsentRequest {
	this := DestinationOauthConsentRequest{}
	return &this
}

// GetDestinationDefinitionId returns the DestinationDefinitionId field value
func (o *DestinationOauthConsentRequest) GetDestinationDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationDefinitionId
}

// GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field value
// and a boolean to check if the value has been set.
func (o *DestinationOauthConsentRequest) GetDestinationDefinitionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationDefinitionId, true
}

// SetDestinationDefinitionId sets field value
func (o *DestinationOauthConsentRequest) SetDestinationDefinitionId(v string) {
	o.DestinationDefinitionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *DestinationOauthConsentRequest) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *DestinationOauthConsentRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *DestinationOauthConsentRequest) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *DestinationOauthConsentRequest) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *DestinationOauthConsentRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *DestinationOauthConsentRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

func (o DestinationOauthConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinationDefinitionId"] = o.DestinationDefinitionId
	}
	if true {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if true {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationOauthConsentRequest struct {
	value *DestinationOauthConsentRequest
	isSet bool
}

func (v NullableDestinationOauthConsentRequest) Get() *DestinationOauthConsentRequest {
	return v.value
}

func (v *NullableDestinationOauthConsentRequest) Set(val *DestinationOauthConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationOauthConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationOauthConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationOauthConsentRequest(val *DestinationOauthConsentRequest) *NullableDestinationOauthConsentRequest {
	return &NullableDestinationOauthConsentRequest{value: val, isSet: true}
}

func (v NullableDestinationOauthConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationOauthConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
