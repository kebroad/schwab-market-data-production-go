# \OptionChainsAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChain**](OptionChainsAPI.md#GetChain) | **Get** /chains | Get option chain for an optionable Symbol



## GetChain

> OptionChain GetChain(ctx).Symbol(symbol).ContractType(contractType).StrikeCount(strikeCount).IncludeUnderlyingQuote(includeUnderlyingQuote).Strategy(strategy).Interval(interval).Strike(strike).Range_(range_).FromDate(fromDate).ToDate(toDate).Volatility(volatility).UnderlyingPrice(underlyingPrice).InterestRate(interestRate).DaysToExpiration(daysToExpiration).ExpMonth(expMonth).OptionType(optionType).Entitlement(entitlement).Execute()

Get option chain for an optionable Symbol



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
	symbol := "AAPL" // string | Enter one symbol
	contractType := "contractType_example" // string | Contract Type (optional)
	strikeCount := int32(56) // int32 | The Number of strikes to return above or below the at-the-money price (optional)
	includeUnderlyingQuote := true // bool | Underlying quotes to be included (optional)
	strategy := "strategy_example" // string | OptionChain strategy. Default is SINGLE. ANALYTICAL allows the use of volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values. (optional)
	interval := float64(1.2) // float64 | Strike interval for spread strategy chains (see strategy param) (optional)
	strike := float64(1.2) // float64 | Strike Price (optional)
	range_ := "range__example" // string | Range(ITM/NTM/OTM etc.) (optional)
	fromDate := time.Now() // string | From date(pattern: yyyy-MM-dd) (optional)
	toDate := time.Now() // string | To date (pattern: yyyy-MM-dd) (optional)
	volatility := float64(1.2) // float64 | Volatility to use in calculations.  Applies only to ANALYTICAL strategy chains (see strategy param) (optional)
	underlyingPrice := float64(1.2) // float64 | Underlying price to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) (optional)
	interestRate := float64(1.2) // float64 | Interest rate to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) (optional)
	daysToExpiration := int32(56) // int32 | Days to expiration to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) (optional)
	expMonth := "expMonth_example" // string | Expiration month (optional)
	optionType := "optionType_example" // string | Option Type (optional)
	entitlement := "entitlement_example" // string | Applicable only if its retail token, entitlement of client PP-PayingPro, NP-NonPro and PN-NonPayingPro (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OptionChainsAPI.GetChain(context.Background()).Symbol(symbol).ContractType(contractType).StrikeCount(strikeCount).IncludeUnderlyingQuote(includeUnderlyingQuote).Strategy(strategy).Interval(interval).Strike(strike).Range_(range_).FromDate(fromDate).ToDate(toDate).Volatility(volatility).UnderlyingPrice(underlyingPrice).InterestRate(interestRate).DaysToExpiration(daysToExpiration).ExpMonth(expMonth).OptionType(optionType).Entitlement(entitlement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OptionChainsAPI.GetChain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChain`: OptionChain
	fmt.Fprintf(os.Stdout, "Response from `OptionChainsAPI.GetChain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | Enter one symbol | 
 **contractType** | **string** | Contract Type | 
 **strikeCount** | **int32** | The Number of strikes to return above or below the at-the-money price | 
 **includeUnderlyingQuote** | **bool** | Underlying quotes to be included | 
 **strategy** | **string** | OptionChain strategy. Default is SINGLE. ANALYTICAL allows the use of volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values. | 
 **interval** | **float64** | Strike interval for spread strategy chains (see strategy param) | 
 **strike** | **float64** | Strike Price | 
 **range_** | **string** | Range(ITM/NTM/OTM etc.) | 
 **fromDate** | **string** | From date(pattern: yyyy-MM-dd) | 
 **toDate** | **string** | To date (pattern: yyyy-MM-dd) | 
 **volatility** | **float64** | Volatility to use in calculations.  Applies only to ANALYTICAL strategy chains (see strategy param) | 
 **underlyingPrice** | **float64** | Underlying price to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) | 
 **interestRate** | **float64** | Interest rate to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) | 
 **daysToExpiration** | **int32** | Days to expiration to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param) | 
 **expMonth** | **string** | Expiration month | 
 **optionType** | **string** | Option Type | 
 **entitlement** | **string** | Applicable only if its retail token, entitlement of client PP-PayingPro, NP-NonPro and PN-NonPayingPro | 

### Return type

[**OptionChain**](OptionChain.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

