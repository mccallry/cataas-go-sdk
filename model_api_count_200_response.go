/*
Cat as a service (CATAAS)

Cat as a service (CATAAS) is a REST API to spread peace and love (or not) thanks to cats.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiCount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCount200Response{}

// ApiCount200Response struct for ApiCount200Response
type ApiCount200Response struct {
	Count                *float32 `json:"count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiCount200Response ApiCount200Response

// NewApiCount200Response instantiates a new ApiCount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCount200Response() *ApiCount200Response {
	this := ApiCount200Response{}
	return &this
}

// NewApiCount200ResponseWithDefaults instantiates a new ApiCount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCount200ResponseWithDefaults() *ApiCount200Response {
	this := ApiCount200Response{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApiCount200Response) GetCount() float32 {
	if o == nil || IsNil(o.Count) {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCount200Response) GetCountOk() (*float32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApiCount200Response) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *ApiCount200Response) SetCount(v float32) {
	o.Count = &v
}

func (o ApiCount200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiCount200Response) UnmarshalJSON(data []byte) (err error) {
	varApiCount200Response := _ApiCount200Response{}

	err = json.Unmarshal(data, &varApiCount200Response)

	if err != nil {
		return err
	}

	*o = ApiCount200Response(varApiCount200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiCount200Response struct {
	value *ApiCount200Response
	isSet bool
}

func (v NullableApiCount200Response) Get() *ApiCount200Response {
	return v.value
}

func (v *NullableApiCount200Response) Set(val *ApiCount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCount200Response(val *ApiCount200Response) *NullableApiCount200Response {
	return &NullableApiCount200Response{value: val, isSet: true}
}

func (v NullableApiCount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
