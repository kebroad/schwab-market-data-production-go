# QuoteForex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WeekHigh** | Pointer to **float64** | Higest price traded in the past 12 months, or 52 weeks | [optional] 
**Var52WeekLow** | Pointer to **float64** | Lowest price traded in the past 12 months, or 52 weeks | [optional] 
**AskPrice** | Pointer to **float64** | Current Best Ask Price | [optional] 
**AskSize** | Pointer to **int32** | Number of shares for ask | [optional] 
**BidPrice** | Pointer to **float64** | Current Best Bid Price | [optional] 
**BidSize** | Pointer to **int32** | Number of shares for bid | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastSize** | Pointer to **int32** | Number of shares traded with last trade | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**Mark** | Pointer to **float64** | Mark price | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**QuoteTime** | Pointer to **int64** | Last quote time in milliseconds since Epoch | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**Tick** | Pointer to **float64** | Tick Price | [optional] 
**TickAmount** | Pointer to **float64** | Tick Amount | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 

## Methods

### NewQuoteForex

`func NewQuoteForex() *QuoteForex`

NewQuoteForex instantiates a new QuoteForex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteForexWithDefaults

`func NewQuoteForexWithDefaults() *QuoteForex`

NewQuoteForexWithDefaults instantiates a new QuoteForex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WeekHigh

`func (o *QuoteForex) GetVar52WeekHigh() float64`

GetVar52WeekHigh returns the Var52WeekHigh field if non-nil, zero value otherwise.

### GetVar52WeekHighOk

`func (o *QuoteForex) GetVar52WeekHighOk() (*float64, bool)`

GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekHigh

`func (o *QuoteForex) SetVar52WeekHigh(v float64)`

SetVar52WeekHigh sets Var52WeekHigh field to given value.

### HasVar52WeekHigh

`func (o *QuoteForex) HasVar52WeekHigh() bool`

HasVar52WeekHigh returns a boolean if a field has been set.

### GetVar52WeekLow

`func (o *QuoteForex) GetVar52WeekLow() float64`

GetVar52WeekLow returns the Var52WeekLow field if non-nil, zero value otherwise.

### GetVar52WeekLowOk

`func (o *QuoteForex) GetVar52WeekLowOk() (*float64, bool)`

GetVar52WeekLowOk returns a tuple with the Var52WeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekLow

`func (o *QuoteForex) SetVar52WeekLow(v float64)`

SetVar52WeekLow sets Var52WeekLow field to given value.

### HasVar52WeekLow

`func (o *QuoteForex) HasVar52WeekLow() bool`

HasVar52WeekLow returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteForex) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteForex) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteForex) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteForex) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteForex) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteForex) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteForex) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteForex) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteForex) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteForex) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteForex) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteForex) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteForex) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteForex) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteForex) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteForex) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteForex) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteForex) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteForex) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteForex) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteForex) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteForex) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteForex) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteForex) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteForex) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteForex) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteForex) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteForex) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *QuoteForex) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteForex) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteForex) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *QuoteForex) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteForex) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteForex) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteForex) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteForex) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *QuoteForex) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *QuoteForex) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *QuoteForex) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *QuoteForex) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteForex) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteForex) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteForex) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteForex) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteForex) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteForex) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteForex) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteForex) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteForex) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteForex) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteForex) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteForex) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetQuoteTime

`func (o *QuoteForex) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *QuoteForex) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *QuoteForex) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *QuoteForex) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteForex) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteForex) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteForex) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteForex) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTick

`func (o *QuoteForex) GetTick() float64`

GetTick returns the Tick field if non-nil, zero value otherwise.

### GetTickOk

`func (o *QuoteForex) GetTickOk() (*float64, bool)`

GetTickOk returns a tuple with the Tick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTick

`func (o *QuoteForex) SetTick(v float64)`

SetTick sets Tick field to given value.

### HasTick

`func (o *QuoteForex) HasTick() bool`

HasTick returns a boolean if a field has been set.

### GetTickAmount

`func (o *QuoteForex) GetTickAmount() float64`

GetTickAmount returns the TickAmount field if non-nil, zero value otherwise.

### GetTickAmountOk

`func (o *QuoteForex) GetTickAmountOk() (*float64, bool)`

GetTickAmountOk returns a tuple with the TickAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickAmount

`func (o *QuoteForex) SetTickAmount(v float64)`

SetTickAmount sets TickAmount field to given value.

### HasTickAmount

`func (o *QuoteForex) HasTickAmount() bool`

HasTickAmount returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteForex) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteForex) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteForex) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteForex) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteForex) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteForex) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteForex) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteForex) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


