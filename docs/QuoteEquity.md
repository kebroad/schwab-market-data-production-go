# QuoteEquity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WeekHigh** | Pointer to **float64** | Higest price traded in the past 12 months, or 52 weeks | [optional] 
**Var52WeekLow** | Pointer to **float64** | Lowest price traded in the past 12 months, or 52 weeks | [optional] 
**AskMICId** | Pointer to **string** | ask MIC code | [optional] 
**AskPrice** | Pointer to **float64** | Current Best Ask Price | [optional] 
**AskSize** | Pointer to **int32** | Number of shares for ask | [optional] 
**AskTime** | Pointer to **int64** | Last ask time in milliseconds since Epoch | [optional] 
**BidMICId** | Pointer to **string** | bid MIC code | [optional] 
**BidPrice** | Pointer to **float64** | Current Best Bid Price | [optional] 
**BidSize** | Pointer to **int32** | Number of shares for bid | [optional] 
**BidTime** | Pointer to **int64** | Last bid time in milliseconds since Epoch | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**LastMICId** | Pointer to **string** | Last MIC Code | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastSize** | Pointer to **int32** | Number of shares traded with last trade | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**Mark** | Pointer to **float64** | Mark price | [optional] 
**MarkChange** | Pointer to **float64** | Mark Price change | [optional] 
**MarkPercentChange** | Pointer to **float64** | Mark Price percent change | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**QuoteTime** | Pointer to **int64** | Last quote time in milliseconds since Epoch | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 
**Volatility** | Pointer to **float64** | Option Risk/Volatility Measurement | [optional] 

## Methods

### NewQuoteEquity

`func NewQuoteEquity() *QuoteEquity`

NewQuoteEquity instantiates a new QuoteEquity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteEquityWithDefaults

`func NewQuoteEquityWithDefaults() *QuoteEquity`

NewQuoteEquityWithDefaults instantiates a new QuoteEquity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WeekHigh

`func (o *QuoteEquity) GetVar52WeekHigh() float64`

GetVar52WeekHigh returns the Var52WeekHigh field if non-nil, zero value otherwise.

### GetVar52WeekHighOk

`func (o *QuoteEquity) GetVar52WeekHighOk() (*float64, bool)`

GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekHigh

`func (o *QuoteEquity) SetVar52WeekHigh(v float64)`

SetVar52WeekHigh sets Var52WeekHigh field to given value.

### HasVar52WeekHigh

`func (o *QuoteEquity) HasVar52WeekHigh() bool`

HasVar52WeekHigh returns a boolean if a field has been set.

### GetVar52WeekLow

`func (o *QuoteEquity) GetVar52WeekLow() float64`

GetVar52WeekLow returns the Var52WeekLow field if non-nil, zero value otherwise.

### GetVar52WeekLowOk

`func (o *QuoteEquity) GetVar52WeekLowOk() (*float64, bool)`

GetVar52WeekLowOk returns a tuple with the Var52WeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekLow

`func (o *QuoteEquity) SetVar52WeekLow(v float64)`

SetVar52WeekLow sets Var52WeekLow field to given value.

### HasVar52WeekLow

`func (o *QuoteEquity) HasVar52WeekLow() bool`

HasVar52WeekLow returns a boolean if a field has been set.

### GetAskMICId

`func (o *QuoteEquity) GetAskMICId() string`

GetAskMICId returns the AskMICId field if non-nil, zero value otherwise.

### GetAskMICIdOk

`func (o *QuoteEquity) GetAskMICIdOk() (*string, bool)`

GetAskMICIdOk returns a tuple with the AskMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMICId

`func (o *QuoteEquity) SetAskMICId(v string)`

SetAskMICId sets AskMICId field to given value.

### HasAskMICId

`func (o *QuoteEquity) HasAskMICId() bool`

HasAskMICId returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteEquity) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteEquity) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteEquity) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteEquity) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteEquity) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteEquity) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteEquity) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteEquity) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetAskTime

`func (o *QuoteEquity) GetAskTime() int64`

GetAskTime returns the AskTime field if non-nil, zero value otherwise.

### GetAskTimeOk

`func (o *QuoteEquity) GetAskTimeOk() (*int64, bool)`

GetAskTimeOk returns a tuple with the AskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskTime

`func (o *QuoteEquity) SetAskTime(v int64)`

SetAskTime sets AskTime field to given value.

### HasAskTime

`func (o *QuoteEquity) HasAskTime() bool`

HasAskTime returns a boolean if a field has been set.

### GetBidMICId

`func (o *QuoteEquity) GetBidMICId() string`

GetBidMICId returns the BidMICId field if non-nil, zero value otherwise.

### GetBidMICIdOk

`func (o *QuoteEquity) GetBidMICIdOk() (*string, bool)`

GetBidMICIdOk returns a tuple with the BidMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMICId

`func (o *QuoteEquity) SetBidMICId(v string)`

SetBidMICId sets BidMICId field to given value.

### HasBidMICId

`func (o *QuoteEquity) HasBidMICId() bool`

HasBidMICId returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteEquity) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteEquity) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteEquity) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteEquity) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteEquity) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteEquity) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteEquity) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteEquity) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetBidTime

`func (o *QuoteEquity) GetBidTime() int64`

GetBidTime returns the BidTime field if non-nil, zero value otherwise.

### GetBidTimeOk

`func (o *QuoteEquity) GetBidTimeOk() (*int64, bool)`

GetBidTimeOk returns a tuple with the BidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidTime

`func (o *QuoteEquity) SetBidTime(v int64)`

SetBidTime sets BidTime field to given value.

### HasBidTime

`func (o *QuoteEquity) HasBidTime() bool`

HasBidTime returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteEquity) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteEquity) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteEquity) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteEquity) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteEquity) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteEquity) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteEquity) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteEquity) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastMICId

`func (o *QuoteEquity) GetLastMICId() string`

GetLastMICId returns the LastMICId field if non-nil, zero value otherwise.

### GetLastMICIdOk

`func (o *QuoteEquity) GetLastMICIdOk() (*string, bool)`

GetLastMICIdOk returns a tuple with the LastMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMICId

`func (o *QuoteEquity) SetLastMICId(v string)`

SetLastMICId sets LastMICId field to given value.

### HasLastMICId

`func (o *QuoteEquity) HasLastMICId() bool`

HasLastMICId returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteEquity) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteEquity) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteEquity) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteEquity) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *QuoteEquity) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteEquity) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteEquity) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *QuoteEquity) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteEquity) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteEquity) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteEquity) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteEquity) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *QuoteEquity) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *QuoteEquity) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *QuoteEquity) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *QuoteEquity) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *QuoteEquity) GetMarkChange() float64`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *QuoteEquity) GetMarkChangeOk() (*float64, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *QuoteEquity) SetMarkChange(v float64)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *QuoteEquity) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *QuoteEquity) GetMarkPercentChange() float64`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *QuoteEquity) GetMarkPercentChangeOk() (*float64, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *QuoteEquity) SetMarkPercentChange(v float64)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *QuoteEquity) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteEquity) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteEquity) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteEquity) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteEquity) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteEquity) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteEquity) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteEquity) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteEquity) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteEquity) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteEquity) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteEquity) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteEquity) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetQuoteTime

`func (o *QuoteEquity) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *QuoteEquity) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *QuoteEquity) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *QuoteEquity) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteEquity) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteEquity) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteEquity) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteEquity) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteEquity) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteEquity) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteEquity) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteEquity) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteEquity) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteEquity) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteEquity) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteEquity) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.

### GetVolatility

`func (o *QuoteEquity) GetVolatility() float64`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *QuoteEquity) GetVolatilityOk() (*float64, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *QuoteEquity) SetVolatility(v float64)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *QuoteEquity) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


