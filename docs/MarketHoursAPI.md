# \MarketHoursAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketHour**](MarketHoursAPI.md#GetMarketHour) | **Get** /markets/{market_id} | Get Market Hours for a single market.
[**GetMarketHours**](MarketHoursAPI.md#GetMarketHours) | **Get** /markets | Get Market Hours for different markets.



## GetMarketHour

> map[string]map[string]Hours GetMarketHour(ctx, marketId).Date(date).Execute()

Get Market Hours for a single market.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	marketId := "marketId_example" // string | market id
	date := time.Now() // string | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketHoursAPI.GetMarketHour(context.Background(), marketId).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketHoursAPI.GetMarketHour``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketHour`: map[string]map[string]Hours
	fmt.Fprintf(os.Stdout, "Response from `MarketHoursAPI.GetMarketHour`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | **string** | market id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHourRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD | 

### Return type

[**map[string]map[string]Hours**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketHours

> map[string]map[string]Hours GetMarketHours(ctx).Markets(markets).Date(date).Execute()

Get Market Hours for different markets.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	markets := []string{"Markets_example"} // []string | List of markets
	date := time.Now() // string | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketHoursAPI.GetMarketHours(context.Background()).Markets(markets).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketHoursAPI.GetMarketHours``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketHours`: map[string]map[string]Hours
	fmt.Fprintf(os.Stdout, "Response from `MarketHoursAPI.GetMarketHours`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketHoursRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markets** | **[]string** | List of markets | 
 **date** | **string** | Valid date range is from currentdate to 1 year from today. It will default to current day if not entered. Date format:YYYY-MM-DD | 

### Return type

[**map[string]map[string]Hours**](map.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

