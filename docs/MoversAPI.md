# \MoversAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovers**](MoversAPI.md#GetMovers) | **Get** /movers/{symbol_id} | Get Movers for a specific index.



## GetMovers

> GetMovers200Response GetMovers(ctx, symbolId).Sort(sort).Frequency(frequency).Execute()

Get Movers for a specific index.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-market-data-production-go"
)

func main() {
	symbolId := "$DJI" // string | Index Symbol
	sort := "VOLUME" // string | Sort by a particular attribute (optional)
	frequency := int32(56) // int32 | To return movers with the specified directions of up or down (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoversAPI.GetMovers(context.Background(), symbolId).Sort(sort).Frequency(frequency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoversAPI.GetMovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovers`: GetMovers200Response
	fmt.Fprintf(os.Stdout, "Response from `MoversAPI.GetMovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbolId** | **string** | Index Symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort by a particular attribute | 
 **frequency** | **int32** | To return movers with the specified directions of up or down | [default to 0]

### Return type

[**GetMovers200Response**](GetMovers200Response.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

