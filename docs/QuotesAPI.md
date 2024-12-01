# \QuotesAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuote**](QuotesAPI.md#GetQuote) | **Get** /{symbol_id}/quotes | Get Quote by single symbol.
[**GetQuotes**](QuotesAPI.md#GetQuotes) | **Get** /quotes | Get Quotes by list of symbols.



## GetQuote

> map[string]QuoteResponseObject GetQuote(ctx, symbolId).Fields(fields).Execute()

Get Quote by single symbol.

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
	symbolId := "TSLA" // string | Symbol of instrument
	fields := "quote,reference" // string | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending `quote, fundamental` in request will return quote and fundamental data in response. Dont send this attribute for full response. (optional) (default to "all")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuote(context.Background(), symbolId).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuote`: map[string]QuoteResponseObject
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbolId** | **string** | Symbol of instrument | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending &#x60;quote, fundamental&#x60; in request will return quote and fundamental data in response. Dont send this attribute for full response. | [default to &quot;all&quot;]

### Return type

[**map[string]QuoteResponseObject**](QuoteResponseObject.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuotes

> map[string]QuoteResponseObject GetQuotes(ctx).Symbols(symbols).Fields(fields).Indicative(indicative).Execute()

Get Quotes by list of symbols.

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
	symbols := "MRAD,EATOF,EBIZ,AAPL,BAC,AAAHX,AAAIX,$DJI,$SPX,MVEN,SOBS,TOITF,CNSWF,AMZN  230317C01360000,DJX   231215C00290000,/ESH23,./ADUF23C0.55,AUD/CAD" // string | Comma separated list of symbol(s) to look up a quote (optional)
	fields := "quote,reference" // string | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending `quote, fundamental` in request will return quote and fundamental data in response. Dont send this attribute for full response. (optional) (default to "all")
	indicative := false // bool | Include indicative symbol quotes for all ETF symbols in request. If ETF symbol ABC is in request and indicative=true API will return quotes for ABC and its corresponding indicative quote for $ABC.IV (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesAPI.GetQuotes(context.Background()).Symbols(symbols).Fields(fields).Indicative(indicative).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesAPI.GetQuotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuotes`: map[string]QuoteResponseObject
	fmt.Fprintf(os.Stdout, "Response from `QuotesAPI.GetQuotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | Comma separated list of symbol(s) to look up a quote | 
 **fields** | **string** | Request for subset of data by passing coma separated list of root nodes, possible root nodes are quote, fundamental, extended, reference, regular. Sending &#x60;quote, fundamental&#x60; in request will return quote and fundamental data in response. Dont send this attribute for full response. | [default to &quot;all&quot;]
 **indicative** | **bool** | Include indicative symbol quotes for all ETF symbols in request. If ETF symbol ABC is in request and indicative&#x3D;true API will return quotes for ABC and its corresponding indicative quote for $ABC.IV | 

### Return type

[**map[string]QuoteResponseObject**](QuoteResponseObject.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

