# AdminCatsBrowse200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Validated** | **bool** |  | 
**File** | **string** |  | 
**Mimetype** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 
**Tags** | **[]string** |  | [default to []]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdminCatsBrowse200ResponseInner

`func NewAdminCatsBrowse200ResponseInner(validated bool, file string, tags []string, ) *AdminCatsBrowse200ResponseInner`

NewAdminCatsBrowse200ResponseInner instantiates a new AdminCatsBrowse200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminCatsBrowse200ResponseInnerWithDefaults

`func NewAdminCatsBrowse200ResponseInnerWithDefaults() *AdminCatsBrowse200ResponseInner`

NewAdminCatsBrowse200ResponseInnerWithDefaults instantiates a new AdminCatsBrowse200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdminCatsBrowse200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminCatsBrowse200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminCatsBrowse200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminCatsBrowse200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValidated

`func (o *AdminCatsBrowse200ResponseInner) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *AdminCatsBrowse200ResponseInner) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *AdminCatsBrowse200ResponseInner) SetValidated(v bool)`

SetValidated sets Validated field to given value.


### GetFile

`func (o *AdminCatsBrowse200ResponseInner) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *AdminCatsBrowse200ResponseInner) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *AdminCatsBrowse200ResponseInner) SetFile(v string)`

SetFile sets File field to given value.


### GetMimetype

`func (o *AdminCatsBrowse200ResponseInner) GetMimetype() string`

GetMimetype returns the Mimetype field if non-nil, zero value otherwise.

### GetMimetypeOk

`func (o *AdminCatsBrowse200ResponseInner) GetMimetypeOk() (*string, bool)`

GetMimetypeOk returns a tuple with the Mimetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimetype

`func (o *AdminCatsBrowse200ResponseInner) SetMimetype(v string)`

SetMimetype sets Mimetype field to given value.

### HasMimetype

`func (o *AdminCatsBrowse200ResponseInner) HasMimetype() bool`

HasMimetype returns a boolean if a field has been set.

### GetSize

`func (o *AdminCatsBrowse200ResponseInner) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AdminCatsBrowse200ResponseInner) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AdminCatsBrowse200ResponseInner) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *AdminCatsBrowse200ResponseInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTags

`func (o *AdminCatsBrowse200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdminCatsBrowse200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdminCatsBrowse200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *AdminCatsBrowse200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminCatsBrowse200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminCatsBrowse200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminCatsBrowse200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdminCatsBrowse200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdminCatsBrowse200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdminCatsBrowse200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdminCatsBrowse200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.GEN.md#documentation-for-models) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints) [[Back to README]](../README.GEN.md)


