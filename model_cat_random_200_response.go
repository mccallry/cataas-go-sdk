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

// checks if the CatRandom200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatRandom200Response{}

// CatRandom200Response struct for CatRandom200Response
type CatRandom200Response struct {
	Id *string `json:"_id,omitempty"`
	Mimetype *string `json:"mimetype,omitempty"`
	Size *float32 `json:"size,omitempty"`
	Tags []string `json:"tags"`
	AdditionalProperties map[string]interface{}
}

type _CatRandom200Response CatRandom200Response

// NewCatRandom200Response instantiates a new CatRandom200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatRandom200Response(tags []string) *CatRandom200Response {
	this := CatRandom200Response{}
	this.Tags = tags
	return &this
}

// NewCatRandom200ResponseWithDefaults instantiates a new CatRandom200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatRandom200ResponseWithDefaults() *CatRandom200Response {
	this := CatRandom200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatRandom200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatRandom200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatRandom200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatRandom200Response) SetId(v string) {
	o.Id = &v
}

// GetMimetype returns the Mimetype field value if set, zero value otherwise.
func (o *CatRandom200Response) GetMimetype() string {
	if o == nil || IsNil(o.Mimetype) {
		var ret string
		return ret
	}
	return *o.Mimetype
}

// GetMimetypeOk returns a tuple with the Mimetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatRandom200Response) GetMimetypeOk() (*string, bool) {
	if o == nil || IsNil(o.Mimetype) {
		return nil, false
	}
	return o.Mimetype, true
}

// HasMimetype returns a boolean if a field has been set.
func (o *CatRandom200Response) HasMimetype() bool {
	if o != nil && !IsNil(o.Mimetype) {
		return true
	}

	return false
}

// SetMimetype gets a reference to the given string and assigns it to the Mimetype field.
func (o *CatRandom200Response) SetMimetype(v string) {
	o.Mimetype = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CatRandom200Response) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatRandom200Response) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CatRandom200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *CatRandom200Response) SetSize(v float32) {
	o.Size = &v
}

// GetTags returns the Tags field value
func (o *CatRandom200Response) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *CatRandom200Response) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *CatRandom200Response) SetTags(v []string) {
	o.Tags = v
}

func (o CatRandom200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatRandom200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Mimetype) {
		toSerialize["mimetype"] = o.Mimetype
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CatRandom200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCatRandom200Response := _CatRandom200Response{}

	err = json.Unmarshal(data, &varCatRandom200Response)

	if err != nil {
		return err
	}

	*o = CatRandom200Response(varCatRandom200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_id")
		delete(additionalProperties, "mimetype")
		delete(additionalProperties, "size")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCatRandom200Response struct {
	value *CatRandom200Response
	isSet bool
}

func (v NullableCatRandom200Response) Get() *CatRandom200Response {
	return v.value
}

func (v *NullableCatRandom200Response) Set(val *CatRandom200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCatRandom200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCatRandom200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatRandom200Response(val *CatRandom200Response) *NullableCatRandom200Response {
	return &NullableCatRandom200Response{value: val, isSet: true}
}

func (v NullableCatRandom200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatRandom200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

