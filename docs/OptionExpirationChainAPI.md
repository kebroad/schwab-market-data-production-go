# \OptionExpirationChainAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExpirationChain**](OptionExpirationChainAPI.md#GetExpirationChain) | **Get** /expirationchain | Get option expiration chain for an optionable symbol



## GetExpirationChain

> ExpirationChain GetExpirationChain(ctx).Symbol(symbol).Execute()

Get option expiration chain for an optionable symbol



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
	symbol := "AAPL" // string | Enter one symbol

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionExpirationChainAPI.GetExpirationChain(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionExpirationChainAPI.GetExpirationChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExpirationChain`: ExpirationChain
	fmt.Fprintf(os.Stdout, "Response from `OptionExpirationChainAPI.GetExpirationChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpirationChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Enter one symbol | 

### Return type

[**ExpirationChain**](ExpirationChain.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

