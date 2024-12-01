# QuoteFutureOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskMICId** | Pointer to **string** | ask MIC code | [optional] 
**AskPrice** | Pointer to **float64** | Current Best Ask Price | [optional] 
**AskSize** | Pointer to **int32** | Number of shares for ask | [optional] 
**BidMICId** | Pointer to **string** | bid MIC code | [optional] 
**BidPrice** | Pointer to **float64** | Current Best Bid Price | [optional] 
**BidSize** | Pointer to **int32** | Number of shares for bid | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**LastMICId** | Pointer to **string** | Last MIC Code | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastSize** | Pointer to **int32** | Number of shares traded with last trade | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**Mark** | Pointer to **float64** | Mark price | [optional] 
**MarkChange** | Pointer to **float64** | Mark Price change | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**OpenInterest** | Pointer to **int32** | Open Interest | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**QuoteTime** | Pointer to **int64** | Last quote time in milliseconds since Epoch | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**SettlemetPrice** | Pointer to **float64** | Price at market open | [optional] 
**Tick** | Pointer to **float64** | Tick Price | [optional] 
**TickAmount** | Pointer to **float64** | Tick Amount | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 

## Methods

### NewQuoteFutureOption

`func NewQuoteFutureOption() *QuoteFutureOption`

NewQuoteFutureOption instantiates a new QuoteFutureOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteFutureOptionWithDefaults

`func NewQuoteFutureOptionWithDefaults() *QuoteFutureOption`

NewQuoteFutureOptionWithDefaults instantiates a new QuoteFutureOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskMICId

`func (o *QuoteFutureOption) GetAskMICId() string`

GetAskMICId returns the AskMICId field if non-nil, zero value otherwise.

### GetAskMICIdOk

`func (o *QuoteFutureOption) GetAskMICIdOk() (*string, bool)`

GetAskMICIdOk returns a tuple with the AskMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMICId

`func (o *QuoteFutureOption) SetAskMICId(v string)`

SetAskMICId sets AskMICId field to given value.

### HasAskMICId

`func (o *QuoteFutureOption) HasAskMICId() bool`

HasAskMICId returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteFutureOption) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteFutureOption) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteFutureOption) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteFutureOption) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteFutureOption) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteFutureOption) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteFutureOption) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteFutureOption) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidMICId

`func (o *QuoteFutureOption) GetBidMICId() string`

GetBidMICId returns the BidMICId field if non-nil, zero value otherwise.

### GetBidMICIdOk

`func (o *QuoteFutureOption) GetBidMICIdOk() (*string, bool)`

GetBidMICIdOk returns a tuple with the BidMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMICId

`func (o *QuoteFutureOption) SetBidMICId(v string)`

SetBidMICId sets BidMICId field to given value.

### HasBidMICId

`func (o *QuoteFutureOption) HasBidMICId() bool`

HasBidMICId returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteFutureOption) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteFutureOption) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteFutureOption) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteFutureOption) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteFutureOption) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteFutureOption) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteFutureOption) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteFutureOption) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteFutureOption) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteFutureOption) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteFutureOption) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteFutureOption) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteFutureOption) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteFutureOption) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteFutureOption) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteFutureOption) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastMICId

`func (o *QuoteFutureOption) GetLastMICId() string`

GetLastMICId returns the LastMICId field if non-nil, zero value otherwise.

### GetLastMICIdOk

`func (o *QuoteFutureOption) GetLastMICIdOk() (*string, bool)`

GetLastMICIdOk returns a tuple with the LastMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMICId

`func (o *QuoteFutureOption) SetLastMICId(v string)`

SetLastMICId sets LastMICId field to given value.

### HasLastMICId

`func (o *QuoteFutureOption) HasLastMICId() bool`

HasLastMICId returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteFutureOption) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteFutureOption) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteFutureOption) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteFutureOption) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *QuoteFutureOption) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteFutureOption) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteFutureOption) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *QuoteFutureOption) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteFutureOption) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteFutureOption) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteFutureOption) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteFutureOption) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *QuoteFutureOption) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *QuoteFutureOption) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *QuoteFutureOption) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *QuoteFutureOption) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *QuoteFutureOption) GetMarkChange() float64`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *QuoteFutureOption) GetMarkChangeOk() (*float64, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *QuoteFutureOption) SetMarkChange(v float64)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *QuoteFutureOption) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteFutureOption) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteFutureOption) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteFutureOption) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteFutureOption) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteFutureOption) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteFutureOption) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteFutureOption) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteFutureOption) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetOpenInterest

`func (o *QuoteFutureOption) GetOpenInterest() int32`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *QuoteFutureOption) GetOpenInterestOk() (*int32, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *QuoteFutureOption) SetOpenInterest(v int32)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *QuoteFutureOption) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteFutureOption) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteFutureOption) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteFutureOption) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteFutureOption) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetQuoteTime

`func (o *QuoteFutureOption) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *QuoteFutureOption) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *QuoteFutureOption) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *QuoteFutureOption) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteFutureOption) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteFutureOption) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteFutureOption) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteFutureOption) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSettlemetPrice

`func (o *QuoteFutureOption) GetSettlemetPrice() float64`

GetSettlemetPrice returns the SettlemetPrice field if non-nil, zero value otherwise.

### GetSettlemetPriceOk

`func (o *QuoteFutureOption) GetSettlemetPriceOk() (*float64, bool)`

GetSettlemetPriceOk returns a tuple with the SettlemetPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlemetPrice

`func (o *QuoteFutureOption) SetSettlemetPrice(v float64)`

SetSettlemetPrice sets SettlemetPrice field to given value.

### HasSettlemetPrice

`func (o *QuoteFutureOption) HasSettlemetPrice() bool`

HasSettlemetPrice returns a boolean if a field has been set.

### GetTick

`func (o *QuoteFutureOption) GetTick() float64`

GetTick returns the Tick field if non-nil, zero value otherwise.

### GetTickOk

`func (o *QuoteFutureOption) GetTickOk() (*float64, bool)`

GetTickOk returns a tuple with the Tick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTick

`func (o *QuoteFutureOption) SetTick(v float64)`

SetTick sets Tick field to given value.

### HasTick

`func (o *QuoteFutureOption) HasTick() bool`

HasTick returns a boolean if a field has been set.

### GetTickAmount

`func (o *QuoteFutureOption) GetTickAmount() float64`

GetTickAmount returns the TickAmount field if non-nil, zero value otherwise.

### GetTickAmountOk

`func (o *QuoteFutureOption) GetTickAmountOk() (*float64, bool)`

GetTickAmountOk returns a tuple with the TickAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickAmount

`func (o *QuoteFutureOption) SetTickAmount(v float64)`

SetTickAmount sets TickAmount field to given value.

### HasTickAmount

`func (o *QuoteFutureOption) HasTickAmount() bool`

HasTickAmount returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteFutureOption) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteFutureOption) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteFutureOption) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteFutureOption) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteFutureOption) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteFutureOption) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteFutureOption) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteFutureOption) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


