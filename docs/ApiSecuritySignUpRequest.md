# ApiSecuritySignUpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewApiSecuritySignUpRequest

`func NewApiSecuritySignUpRequest(password string, secret string, ) *ApiSecuritySignUpRequest`

NewApiSecuritySignUpRequest instantiates a new ApiSecuritySignUpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecuritySignUpRequestWithDefaults

`func NewApiSecuritySignUpRequestWithDefaults() *ApiSecuritySignUpRequest`

NewApiSecuritySignUpRequestWithDefaults instantiates a new ApiSecuritySignUpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ApiSecuritySignUpRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiSecuritySignUpRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiSecuritySignUpRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiSecuritySignUpRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ApiSecuritySignUpRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiSecuritySignUpRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiSecuritySignUpRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *ApiSecuritySignUpRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiSecuritySignUpRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiSecuritySignUpRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApiSecuritySignUpRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecret

`func (o *ApiSecuritySignUpRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApiSecuritySignUpRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApiSecuritySignUpRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.GEN.md#documentation-for-models) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints) [[Back to README]](../README.GEN.md)


