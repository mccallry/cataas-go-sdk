/*
Cat as a service (CATAAS)

Cat as a service (CATAAS) is a REST API to spread peace and love (or not) thanks to cats.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cataas

import (
	"encoding/json"
	"fmt"
)

// checks if the ApiSecuritySignInRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSecuritySignInRequest{}

// ApiSecuritySignInRequest struct for ApiSecuritySignInRequest
type ApiSecuritySignInRequest struct {
	Username             *string `json:"username,omitempty"`
	Password             string  `json:"password"`
	AdditionalProperties map[string]interface{}
}

type _ApiSecuritySignInRequest ApiSecuritySignInRequest

// NewApiSecuritySignInRequest instantiates a new ApiSecuritySignInRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecuritySignInRequest(password string) *ApiSecuritySignInRequest {
	this := ApiSecuritySignInRequest{}
	this.Password = password
	return &this
}

// NewApiSecuritySignInRequestWithDefaults instantiates a new ApiSecuritySignInRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecuritySignInRequestWithDefaults() *ApiSecuritySignInRequest {
	this := ApiSecuritySignInRequest{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApiSecuritySignInRequest) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecuritySignInRequest) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApiSecuritySignInRequest) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApiSecuritySignInRequest) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value
func (o *ApiSecuritySignInRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ApiSecuritySignInRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ApiSecuritySignInRequest) SetPassword(v string) {
	o.Password = v
}

func (o ApiSecuritySignInRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSecuritySignInRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	toSerialize["password"] = o.Password

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiSecuritySignInRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiSecuritySignInRequest := _ApiSecuritySignInRequest{}

	err = json.Unmarshal(data, &varApiSecuritySignInRequest)

	if err != nil {
		return err
	}

	*o = ApiSecuritySignInRequest(varApiSecuritySignInRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiSecuritySignInRequest struct {
	value *ApiSecuritySignInRequest
	isSet bool
}

func (v NullableApiSecuritySignInRequest) Get() *ApiSecuritySignInRequest {
	return v.value
}

func (v *NullableApiSecuritySignInRequest) Set(val *ApiSecuritySignInRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecuritySignInRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecuritySignInRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecuritySignInRequest(val *ApiSecuritySignInRequest) *NullableApiSecuritySignInRequest {
	return &NullableApiSecuritySignInRequest{value: val, isSet: true}
}

func (v NullableApiSecuritySignInRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecuritySignInRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
