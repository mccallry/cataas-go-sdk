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

// checks if the SignUp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignUp{}

// SignUp struct for SignUp
type SignUp struct {
	Username             *string `json:"username,omitempty"`
	Password             string  `json:"password"`
	Email                *string `json:"email,omitempty"`
	Secret               string  `json:"secret"`
	AdditionalProperties map[string]interface{}
}

type _SignUp SignUp

// NewSignUp instantiates a new SignUp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignUp(password string, secret string) *SignUp {
	this := SignUp{}
	this.Password = password
	this.Secret = secret
	return &this
}

// NewSignUpWithDefaults instantiates a new SignUp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignUpWithDefaults() *SignUp {
	this := SignUp{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SignUp) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignUp) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SignUp) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SignUp) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value
func (o *SignUp) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SignUp) SetPassword(v string) {
	o.Password = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SignUp) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignUp) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SignUp) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SignUp) SetEmail(v string) {
	o.Email = &v
}

// GetSecret returns the Secret field value
func (o *SignUp) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *SignUp) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *SignUp) SetSecret(v string) {
	o.Secret = v
}

func (o SignUp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignUp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	toSerialize["password"] = o.Password
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["secret"] = o.Secret

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignUp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"secret",
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

	varSignUp := _SignUp{}

	err = json.Unmarshal(data, &varSignUp)

	if err != nil {
		return err
	}

	*o = SignUp(varSignUp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "email")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignUp struct {
	value *SignUp
	isSet bool
}

func (v NullableSignUp) Get() *SignUp {
	return v.value
}

func (v *NullableSignUp) Set(val *SignUp) {
	v.value = val
	v.isSet = true
}

func (v NullableSignUp) IsSet() bool {
	return v.isSet
}

func (v *NullableSignUp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignUp(val *SignUp) *NullableSignUp {
	return &NullableSignUp{value: val, isSet: true}
}

func (v NullableSignUp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignUp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
