# Cat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mimetype** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**Tags** | **[]string** |  | [default to []]

## Methods

### NewCat

`func NewCat(tags []string, ) *Cat`

NewCat instantiates a new Cat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatWithDefaults

`func NewCatWithDefaults() *Cat`

NewCatWithDefaults instantiates a new Cat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMimetype

`func (o *Cat) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *Cat) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *Cat) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *Cat) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetSize

`func (o *Cat) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Cat) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Cat) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Cat) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTags

`func (o *Cat) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Cat) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Cat) SetTags(v []string)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.GEN.md#documentation-for-models) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints) [[Back to README]](../README.GEN.md)


