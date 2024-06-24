/*
Cat as a service (CATAAS)

Cat as a service (CATAAS) is a REST API to spread peace and love (or not) thanks to cats.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the EditCat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EditCat{}

// EditCat struct for EditCat
type EditCat struct {
	Tags                 []string `json:"tags"`
	AdditionalProperties map[string]interface{}
}

type _EditCat EditCat

// NewEditCat instantiates a new EditCat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditCat(tags []string) *EditCat {
	this := EditCat{}
	this.Tags = tags
	return &this
}

// NewEditCatWithDefaults instantiates a new EditCat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditCatWithDefaults() *EditCat {
	this := EditCat{}
	return &this
}

// GetTags returns the Tags field value
func (o *EditCat) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *EditCat) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *EditCat) SetTags(v []string) {
	o.Tags = v
}

func (o EditCat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EditCat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EditCat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
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

	varEditCat := _EditCat{}

	err = json.Unmarshal(data, &varEditCat)

	if err != nil {
		return err
	}

	*o = EditCat(varEditCat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEditCat struct {
	value *EditCat
	isSet bool
}

func (v NullableEditCat) Get() *EditCat {
	return v.value
}

func (v *NullableEditCat) Set(val *EditCat) {
	v.value = val
	v.isSet = true
}

func (v NullableEditCat) IsSet() bool {
	return v.isSet
}

func (v *NullableEditCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditCat(val *EditCat) *NullableEditCat {
	return &NullableEditCat{value: val, isSet: true}
}

func (v NullableEditCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
