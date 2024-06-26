# SignUp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 

## Methods

### NewSignUp

`func NewSignUp(password string, secret string, ) *SignUp`

NewSignUp instantiates a new SignUp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignUpWithDefaults

`func NewSignUpWithDefaults() *SignUp`

NewSignUpWithDefaults instantiates a new SignUp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SignUp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SignUp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SignUp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SignUp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *SignUp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SignUp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SignUp) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *SignUp) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignUp) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignUp) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignUp) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecret

`func (o *SignUp) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SignUp) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SignUp) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.GEN.md#documentation-for-models) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints) [[Back to README]](../README.GEN.md)


