# ApiSecuritySignInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewApiSecuritySignInRequest

`func NewApiSecuritySignInRequest(password string, ) *ApiSecuritySignInRequest`

NewApiSecuritySignInRequest instantiates a new ApiSecuritySignInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecuritySignInRequestWithDefaults

`func NewApiSecuritySignInRequestWithDefaults() *ApiSecuritySignInRequest`

NewApiSecuritySignInRequestWithDefaults instantiates a new ApiSecuritySignInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ApiSecuritySignInRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiSecuritySignInRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiSecuritySignInRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiSecuritySignInRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ApiSecuritySignInRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiSecuritySignInRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiSecuritySignInRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.GEN.md#documentation-for-models) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints) [[Back to README]](../README.GEN.md)


