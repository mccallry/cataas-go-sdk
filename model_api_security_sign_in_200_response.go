/*
Cat as a service (CATAAS)

Cat as a service (CATAAS) is a REST API to spread peace and love (or not) thanks to cats.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cataas

import (
	"encoding/json"
)

// checks if the ApiSecuritySignIn200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSecuritySignIn200Response{}

// ApiSecuritySignIn200Response struct for ApiSecuritySignIn200Response
type ApiSecuritySignIn200Response struct {
	Token                *string `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiSecuritySignIn200Response ApiSecuritySignIn200Response

// NewApiSecuritySignIn200Response instantiates a new ApiSecuritySignIn200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecuritySignIn200Response() *ApiSecuritySignIn200Response {
	this := ApiSecuritySignIn200Response{}
	return &this
}

// NewApiSecuritySignIn200ResponseWithDefaults instantiates a new ApiSecuritySignIn200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecuritySignIn200ResponseWithDefaults() *ApiSecuritySignIn200Response {
	this := ApiSecuritySignIn200Response{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ApiSecuritySignIn200Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecuritySignIn200Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ApiSecuritySignIn200Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ApiSecuritySignIn200Response) SetToken(v string) {
	o.Token = &v
}

func (o ApiSecuritySignIn200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSecuritySignIn200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiSecuritySignIn200Response) UnmarshalJSON(data []byte) (err error) {
	varApiSecuritySignIn200Response := _ApiSecuritySignIn200Response{}

	err = json.Unmarshal(data, &varApiSecuritySignIn200Response)

	if err != nil {
		return err
	}

	*o = ApiSecuritySignIn200Response(varApiSecuritySignIn200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiSecuritySignIn200Response struct {
	value *ApiSecuritySignIn200Response
	isSet bool
}

func (v NullableApiSecuritySignIn200Response) Get() *ApiSecuritySignIn200Response {
	return v.value
}

func (v *NullableApiSecuritySignIn200Response) Set(val *ApiSecuritySignIn200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecuritySignIn200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecuritySignIn200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecuritySignIn200Response(val *ApiSecuritySignIn200Response) *NullableApiSecuritySignIn200Response {
	return &NullableApiSecuritySignIn200Response{value: val, isSet: true}
}

func (v NullableApiSecuritySignIn200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecuritySignIn200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
