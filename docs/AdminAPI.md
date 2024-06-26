# \AdminAPI

All URIs are relative to *https://cataas.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCatsBrowse**](AdminAPI.md#AdminCatsBrowse) | **Get** /admin/cats | 
[**AdminCatsDelete**](AdminAPI.md#AdminCatsDelete) | **Delete** /admin/cats/{id} | 
[**AdminCatsEdit**](AdminAPI.md#AdminCatsEdit) | **Patch** /admin/cats/{id} | 
[**AdminCatsValidate**](AdminAPI.md#AdminCatsValidate) | **Put** /admin/cats/{id}/validate | 



## AdminCatsBrowse

> []AdminCatsBrowse200ResponseInner AdminCatsBrowse(ctx).Limit(limit).Skip(skip).Execute()





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
	limit := int32(56) // int32 |  (optional)
	skip := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminAPI.AdminCatsBrowse(context.Background()).Limit(limit).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCatsBrowse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminCatsBrowse`: []AdminCatsBrowse200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCatsBrowse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCatsBrowseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **skip** | **int32** |  | 

### Return type

[**[]AdminCatsBrowse200ResponseInner**](AdminCatsBrowse200ResponseInner.md)

### Authorization

[jwt](../README.GEN.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)


## AdminCatsDelete

> AdminCatsDelete(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminCatsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCatsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCatsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.GEN.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)


## AdminCatsEdit

> AdminCatsEdit(ctx, id).AdminCatsEditRequest(adminCatsEditRequest).Execute()





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
	id := "id_example" // string | 
	adminCatsEditRequest := *openapiclient.NewAdminCatsEditRequest() // AdminCatsEditRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminCatsEdit(context.Background(), id).AdminCatsEditRequest(adminCatsEditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCatsEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCatsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminCatsEditRequest** | [**AdminCatsEditRequest**](AdminCatsEditRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.GEN.md#jwt)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)


## AdminCatsValidate

> AdminCatsValidate(ctx, id).Execute()





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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminAPI.AdminCatsValidate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCatsValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminCatsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[jwt](../README.GEN.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.GEN.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.GEN.md#documentation-for-models)
[[Back to README]](../README.GEN.md)

