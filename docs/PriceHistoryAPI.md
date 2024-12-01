# \PriceHistoryAPI

All URIs are relative to *https://api.schwabapi.com/marketdata/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriceHistory**](PriceHistoryAPI.md#GetPriceHistory) | **Get** /pricehistory | Get PriceHistory for a single symbol and date ranges.



## GetPriceHistory

> CandleList GetPriceHistory(ctx).Symbol(symbol).PeriodType(periodType).Period(period).FrequencyType(frequencyType).Frequency(frequency).StartDate(startDate).EndDate(endDate).NeedExtendedHoursData(needExtendedHoursData).NeedPreviousClose(needPreviousClose).Execute()

Get PriceHistory for a single symbol and date ranges.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	symbol := "AAPL" // string | The Equity symbol used to look up price history
	periodType := "periodType_example" // string | The chart period being requested. (optional)
	period := int32(56) // int32 | The number of chart period types.<br><br> If the periodType is <br> &#8226; <b>day</b> - valid values are 1, 2, 3, 4, 5, 10<br> &#8226; <b>month</b> - valid values are 1, 2, 3, 6<br> &#8226; <b>year</b> - valid values are 1, 2, 3, 5, 10, 15, 20<br> &#8226; <b>ytd</b> - valid values are 1<br><br> If   the period is not specified and the periodType is<br> &#8226; <b>day</b> - default period is 10.<br> &#8226; <b>month</b> - default period is 1.<br> &#8226; <b>year</b> - default period is 1.<br> &#8226; <b>ytd</b> - default period is 1.<br> (optional)
	frequencyType := "frequencyType_example" // string | The time frequencyType<br><br> If the periodType is <br> &#8226; <b>day</b> - valid value is minute<br> &#8226; <b>month</b> - valid values are daily, weekly<br> &#8226; <b>year</b> - valid values are daily, weekly, monthly<br> &#8226; <b>ytd</b> - valid values are daily, weekly<br><br> If frequencyType  is not specified, default value depends on the periodType<br> &#8226; <b>day</b> - defaulted to minute.<br> &#8226; <b>month</b> - defaulted to weekly.<br> &#8226; <b>year</b> - defaulted to monthly.<br> &#8226; <b>ytd</b> - defaulted to weekly.<br> (optional)
	frequency := int32(56) // int32 | The time frequency duration<br><br> If the frequencyType is <br> &#8226; <b>minute</b> - valid values are 1, 5, 10, 15, 30<br> &#8226; <b>daily</b> - valid value is 1<br> &#8226; <b>weekly</b> - valid value is 1<br> &#8226; <b>monthly</b> - valid value is 1<br><br> If frequency  is not specified, default value is <b>1</b><br> (optional)
	startDate := int64(789) // int64 | The start date, Time   in milliseconds since the UNIX epoch eg 1451624400000<br>If not   specified startDate will be (endDate - period) excluding weekends and holidays. (optional)
	endDate := int64(789) // int64 | The end date, Time   in milliseconds since the UNIX epoch eg 1451624400000<br> If not   specified, the endDate will default to the market close of previous business day. (optional)
	needExtendedHoursData := true // bool | Need extended hours data (optional)
	needPreviousClose := true // bool | Need previous close price/date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PriceHistoryAPI.GetPriceHistory(context.Background()).Symbol(symbol).PeriodType(periodType).Period(period).FrequencyType(frequencyType).Frequency(frequency).StartDate(startDate).EndDate(endDate).NeedExtendedHoursData(needExtendedHoursData).NeedPreviousClose(needPreviousClose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PriceHistoryAPI.GetPriceHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriceHistory`: CandleList
	fmt.Fprintf(os.Stdout, "Response from `PriceHistoryAPI.GetPriceHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPriceHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | **string** | The Equity symbol used to look up price history | 
 **periodType** | **string** | The chart period being requested. | 
 **period** | **int32** | The number of chart period types.&lt;br&gt;&lt;br&gt; If the periodType is &lt;br&gt; &amp;#8226; &lt;b&gt;day&lt;/b&gt; - valid values are 1, 2, 3, 4, 5, 10&lt;br&gt; &amp;#8226; &lt;b&gt;month&lt;/b&gt; - valid values are 1, 2, 3, 6&lt;br&gt; &amp;#8226; &lt;b&gt;year&lt;/b&gt; - valid values are 1, 2, 3, 5, 10, 15, 20&lt;br&gt; &amp;#8226; &lt;b&gt;ytd&lt;/b&gt; - valid values are 1&lt;br&gt;&lt;br&gt; If   the period is not specified and the periodType is&lt;br&gt; &amp;#8226; &lt;b&gt;day&lt;/b&gt; - default period is 10.&lt;br&gt; &amp;#8226; &lt;b&gt;month&lt;/b&gt; - default period is 1.&lt;br&gt; &amp;#8226; &lt;b&gt;year&lt;/b&gt; - default period is 1.&lt;br&gt; &amp;#8226; &lt;b&gt;ytd&lt;/b&gt; - default period is 1.&lt;br&gt; | 
 **frequencyType** | **string** | The time frequencyType&lt;br&gt;&lt;br&gt; If the periodType is &lt;br&gt; &amp;#8226; &lt;b&gt;day&lt;/b&gt; - valid value is minute&lt;br&gt; &amp;#8226; &lt;b&gt;month&lt;/b&gt; - valid values are daily, weekly&lt;br&gt; &amp;#8226; &lt;b&gt;year&lt;/b&gt; - valid values are daily, weekly, monthly&lt;br&gt; &amp;#8226; &lt;b&gt;ytd&lt;/b&gt; - valid values are daily, weekly&lt;br&gt;&lt;br&gt; If frequencyType  is not specified, default value depends on the periodType&lt;br&gt; &amp;#8226; &lt;b&gt;day&lt;/b&gt; - defaulted to minute.&lt;br&gt; &amp;#8226; &lt;b&gt;month&lt;/b&gt; - defaulted to weekly.&lt;br&gt; &amp;#8226; &lt;b&gt;year&lt;/b&gt; - defaulted to monthly.&lt;br&gt; &amp;#8226; &lt;b&gt;ytd&lt;/b&gt; - defaulted to weekly.&lt;br&gt; | 
 **frequency** | **int32** | The time frequency duration&lt;br&gt;&lt;br&gt; If the frequencyType is &lt;br&gt; &amp;#8226; &lt;b&gt;minute&lt;/b&gt; - valid values are 1, 5, 10, 15, 30&lt;br&gt; &amp;#8226; &lt;b&gt;daily&lt;/b&gt; - valid value is 1&lt;br&gt; &amp;#8226; &lt;b&gt;weekly&lt;/b&gt; - valid value is 1&lt;br&gt; &amp;#8226; &lt;b&gt;monthly&lt;/b&gt; - valid value is 1&lt;br&gt;&lt;br&gt; If frequency  is not specified, default value is &lt;b&gt;1&lt;/b&gt;&lt;br&gt; | 
 **startDate** | **int64** | The start date, Time   in milliseconds since the UNIX epoch eg 1451624400000&lt;br&gt;If not   specified startDate will be (endDate - period) excluding weekends and holidays. | 
 **endDate** | **int64** | The end date, Time   in milliseconds since the UNIX epoch eg 1451624400000&lt;br&gt; If not   specified, the endDate will default to the market close of previous business day. | 
 **needExtendedHoursData** | **bool** | Need extended hours data | 
 **needPreviousClose** | **bool** | Need previous close price/date | 

### Return type

[**CandleList**](CandleList.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

