# \InstrumentsAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstruments**](InstrumentsAPI.md#GetInstruments) | **Get** /instruments | Get Instruments by symbols and projections.
[**GetInstrumentsByCusip**](InstrumentsAPI.md#GetInstrumentsByCusip) | **Get** /instruments/{cusip_id} | Get Instrument by specific cusip



## GetInstruments

> GetInstruments200Response GetInstruments(ctx).Symbol(symbol).Projection(projection).Execute()

Get Instruments by symbols and projections.



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
	symbol := "symbol_example" // string | symbol of a security
	projection := "projection_example" // string | search by

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstrumentsAPI.GetInstruments(context.Background()).Symbol(symbol).Projection(projection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsAPI.GetInstruments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstruments`: GetInstruments200Response
	fmt.Fprintf(os.Stdout, "Response from `InstrumentsAPI.GetInstruments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | symbol of a security | 
 **projection** | **string** | search by | 

### Return type

[**GetInstruments200Response**](GetInstruments200Response.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstrumentsByCusip

> InstrumentResponse GetInstrumentsByCusip(ctx, cusipId).Execute()

Get Instrument by specific cusip



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
	cusipId := "cusipId_example" // string | cusip of a security

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstrumentsAPI.GetInstrumentsByCusip(context.Background(), cusipId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstrumentsAPI.GetInstrumentsByCusip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstrumentsByCusip`: InstrumentResponse
	fmt.Fprintf(os.Stdout, "Response from `InstrumentsAPI.GetInstrumentsByCusip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cusipId** | **string** | cusip of a security | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstrumentsByCusipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InstrumentResponse**](InstrumentResponse.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

