# \SecurityAPI

All URIs are relative to *https://cataas.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiSecuritySignIn**](SecurityAPI.md#ApiSecuritySignIn) | **Post** /security/sign-in | 
[**ApiSecuritySignUp**](SecurityAPI.md#ApiSecuritySignUp) | **Post** /security/sign-up | 



## ApiSecuritySignIn

> ApiSecuritySignIn200Response ApiSecuritySignIn(ctx).ApiSecuritySignInRequest(apiSecuritySignInRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mccallry/cataas-go-sdk"
)

func main() {
	apiSecuritySignInRequest := *openapiclient.NewApiSecuritySignInRequest("Password_example") // ApiSecuritySignInRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAPI.ApiSecuritySignIn(context.Background()).ApiSecuritySignInRequest(apiSecuritySignInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiSecuritySignIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiSecuritySignIn`: ApiSecuritySignIn200Response
	fmt.Fprintf(os.Stdout, "Response from `SecurityAPI.ApiSecuritySignIn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSecuritySignInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiSecuritySignInRequest** | [**ApiSecuritySignInRequest**](ApiSecuritySignInRequest.md) |  | 

### Return type

[**ApiSecuritySignIn200Response**](ApiSecuritySignIn200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)


## ApiSecuritySignUp

> ApiSecuritySignUp(ctx).ApiSecuritySignUpRequest(apiSecuritySignUpRequest).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mccallry/cataas-go-sdk"
)

func main() {
	apiSecuritySignUpRequest := *openapiclient.NewApiSecuritySignUpRequest("Password_example", "Secret_example") // ApiSecuritySignUpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityAPI.ApiSecuritySignUp(context.Background()).ApiSecuritySignUpRequest(apiSecuritySignUpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAPI.ApiSecuritySignUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiSecuritySignUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiSecuritySignUpRequest** | [**ApiSecuritySignUpRequest**](ApiSecuritySignUpRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)

