# \CatsAPI

All URIs are relative to *https://cataas.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatGet**](CatsAPI.md#CatGet) | **Get** /cat/{id} | 
[**CatGetText**](CatsAPI.md#CatGetText) | **Get** /cat/{id}/says/{text} | 
[**CatRandom**](CatsAPI.md#CatRandom) | **Get** /cat | 
[**CatRandomTag**](CatsAPI.md#CatRandomTag) | **Get** /cat/{tag} | 
[**CatRandomTagText**](CatsAPI.md#CatRandomTagText) | **Get** /cat/{tag}/says/{text} | 
[**CatRandomText**](CatsAPI.md#CatRandomText) | **Get** /cat/says/{text} | 



## CatGet

> *os.File CatGet(ctx, id).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatGet(context.Background(), id).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatGetText

> *os.File CatGetText(ctx, id, text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	text := "text_example" // string | 
	font := "font_example" // string |  (optional) (default to "Impact")
	fontSize := int32(56) // int32 |  (optional) (default to 30)
	fontColor := "fontColor_example" // string |  (optional) (default to "#000")
	fontBackground := "fontBackground_example" // string |  (optional) (default to "none")
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatGetText(context.Background(), id, text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatGetText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatGetText`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatGetText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**text** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatGetTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **font** | **string** |  | [default to &quot;Impact&quot;]
 **fontSize** | **int32** |  | [default to 30]
 **fontColor** | **string** |  | [default to &quot;#000&quot;]
 **fontBackground** | **string** |  | [default to &quot;none&quot;]
 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatRandom

> *os.File CatRandom(ctx).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatRandom(context.Background()).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatRandom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatRandom`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatRandom`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatRandomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatRandomTag

> *os.File CatRandomTag(ctx, tag).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	tag := "tag_example" // string | 
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatRandomTag(context.Background(), tag).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatRandomTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatRandomTag`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatRandomTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatRandomTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatRandomTagText

> *os.File CatRandomTagText(ctx, tag, text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	tag := "tag_example" // string | 
	text := "text_example" // string | 
	font := "font_example" // string |  (optional) (default to "Impact")
	fontSize := int32(56) // int32 |  (optional) (default to 30)
	fontColor := "fontColor_example" // string |  (optional) (default to "#000")
	fontBackground := "fontBackground_example" // string |  (optional) (default to "none")
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatRandomTagText(context.Background(), tag, text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatRandomTagText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatRandomTagText`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatRandomTagText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 
**text** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatRandomTagTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **font** | **string** |  | [default to &quot;Impact&quot;]
 **fontSize** | **int32** |  | [default to 30]
 **fontColor** | **string** |  | [default to &quot;#000&quot;]
 **fontBackground** | **string** |  | [default to &quot;none&quot;]
 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CatRandomText

> *os.File CatRandomText(ctx, text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()





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
	text := "text_example" // string | 
	font := "font_example" // string |  (optional) (default to "Impact")
	fontSize := int32(56) // int32 |  (optional) (default to 30)
	fontColor := "fontColor_example" // string |  (optional) (default to "#000")
	fontBackground := "fontBackground_example" // string |  (optional) (default to "none")
	type_ := "type__example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	fit := "fit_example" // string |  (optional)
	position := "position_example" // string |  (optional) (default to "center")
	width := int32(56) // int32 |  (optional)
	height := int32(56) // int32 |  (optional)
	blur := int32(56) // int32 |  (optional)
	r := int32(56) // int32 | Red (optional)
	g := int32(56) // int32 | Green (optional)
	b := int32(56) // int32 | Blue (optional)
	brightness := float32(3.4) // float32 | Brightness multiplier (optional)
	saturation := float32(3.4) // float32 | Saturation multiplier (optional)
	hue := int32(56) // int32 | Hue rotation in degrees (optional)
	lightness := int32(56) // int32 | Lightness addend (optional)
	html := true // bool |  (optional)
	json := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatsAPI.CatRandomText(context.Background(), text).Font(font).FontSize(fontSize).FontColor(fontColor).FontBackground(fontBackground).Type_(type_).Filter(filter).Fit(fit).Position(position).Width(width).Height(height).Blur(blur).R(r).G(g).B(b).Brightness(brightness).Saturation(saturation).Hue(hue).Lightness(lightness).Html(html).Json(json).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatsAPI.CatRandomText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CatRandomText`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CatsAPI.CatRandomText`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**text** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCatRandomTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **font** | **string** |  | [default to &quot;Impact&quot;]
 **fontSize** | **int32** |  | [default to 30]
 **fontColor** | **string** |  | [default to &quot;#000&quot;]
 **fontBackground** | **string** |  | [default to &quot;none&quot;]
 **type_** | **string** |  | 
 **filter** | **string** |  | 
 **fit** | **string** |  | 
 **position** | **string** |  | [default to &quot;center&quot;]
 **width** | **int32** |  | 
 **height** | **int32** |  | 
 **blur** | **int32** |  | 
 **r** | **int32** | Red | 
 **g** | **int32** | Green | 
 **b** | **int32** | Blue | 
 **brightness** | **float32** | Brightness multiplier | 
 **saturation** | **float32** | Saturation multiplier | 
 **hue** | **int32** | Hue rotation in degrees | 
 **lightness** | **int32** | Lightness addend | 
 **html** | **bool** |  | 
 **json** | **bool** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

