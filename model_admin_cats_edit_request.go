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

// checks if the AdminCatsEditRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminCatsEditRequest{}

// AdminCatsEditRequest struct for AdminCatsEditRequest
type AdminCatsEditRequest struct {
	Tags                 *string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdminCatsEditRequest AdminCatsEditRequest

// NewAdminCatsEditRequest instantiates a new AdminCatsEditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCatsEditRequest() *AdminCatsEditRequest {
	this := AdminCatsEditRequest{}
	return &this
}

// NewAdminCatsEditRequestWithDefaults instantiates a new AdminCatsEditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCatsEditRequestWithDefaults() *AdminCatsEditRequest {
	this := AdminCatsEditRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdminCatsEditRequest) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCatsEditRequest) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdminCatsEditRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *AdminCatsEditRequest) SetTags(v string) {
	o.Tags = &v
}

func (o AdminCatsEditRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminCatsEditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdminCatsEditRequest) UnmarshalJSON(data []byte) (err error) {
	varAdminCatsEditRequest := _AdminCatsEditRequest{}

	err = json.Unmarshal(data, &varAdminCatsEditRequest)

	if err != nil {
		return err
	}

	*o = AdminCatsEditRequest(varAdminCatsEditRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdminCatsEditRequest struct {
	value *AdminCatsEditRequest
	isSet bool
}

func (v NullableAdminCatsEditRequest) Get() *AdminCatsEditRequest {
	return v.value
}

func (v *NullableAdminCatsEditRequest) Set(val *AdminCatsEditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCatsEditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCatsEditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCatsEditRequest(val *AdminCatsEditRequest) *NullableAdminCatsEditRequest {
	return &NullableAdminCatsEditRequest{value: val, isSet: true}
}

func (v NullableAdminCatsEditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCatsEditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
