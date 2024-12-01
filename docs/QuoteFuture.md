# QuoteFuture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskMICId** | Pointer to **string** | ask MIC code | [optional] 
**AskPrice** | Pointer to **float64** | Current Best Ask Price | [optional] 
**AskSize** | Pointer to **int32** | Number of shares for ask | [optional] 
**AskTime** | Pointer to **int64** | Last ask time in milliseconds since Epoch | [optional] 
**BidMICId** | Pointer to **string** | bid MIC code | [optional] 
**BidPrice** | Pointer to **float64** | Current Best Bid Price | [optional] 
**BidSize** | Pointer to **int32** | Number of shares for bid | [optional] 
**BidTime** | Pointer to **int64** | Last bid time in milliseconds since Epoch | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**FuturePercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**LastMICId** | Pointer to **string** | Last MIC Code | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastSize** | Pointer to **int32** | Number of shares traded with last trade | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**Mark** | Pointer to **float64** | Mark price | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**OpenInterest** | Pointer to **int32** | Open interest | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**QuoteTime** | Pointer to **int64** | Last quote time in milliseconds since Epoch | [optional] 
**QuotedInSession** | Pointer to **bool** | quoted during trading session | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**SettleTime** | Pointer to **int64** | settlement time in milliseconds since Epoch | [optional] 
**Tick** | Pointer to **float64** | Tick Price | [optional] 
**TickAmount** | Pointer to **float64** | Tick Amount | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 

## Methods

### NewQuoteFuture

`func NewQuoteFuture() *QuoteFuture`

NewQuoteFuture instantiates a new QuoteFuture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteFutureWithDefaults

`func NewQuoteFutureWithDefaults() *QuoteFuture`

NewQuoteFutureWithDefaults instantiates a new QuoteFuture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskMICId

`func (o *QuoteFuture) GetAskMICId() string`

GetAskMICId returns the AskMICId field if non-nil, zero value otherwise.

### GetAskMICIdOk

`func (o *QuoteFuture) GetAskMICIdOk() (*string, bool)`

GetAskMICIdOk returns a tuple with the AskMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMICId

`func (o *QuoteFuture) SetAskMICId(v string)`

SetAskMICId sets AskMICId field to given value.

### HasAskMICId

`func (o *QuoteFuture) HasAskMICId() bool`

HasAskMICId returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteFuture) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteFuture) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteFuture) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteFuture) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteFuture) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteFuture) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteFuture) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteFuture) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetAskTime

`func (o *QuoteFuture) GetAskTime() int64`

GetAskTime returns the AskTime field if non-nil, zero value otherwise.

### GetAskTimeOk

`func (o *QuoteFuture) GetAskTimeOk() (*int64, bool)`

GetAskTimeOk returns a tuple with the AskTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskTime

`func (o *QuoteFuture) SetAskTime(v int64)`

SetAskTime sets AskTime field to given value.

### HasAskTime

`func (o *QuoteFuture) HasAskTime() bool`

HasAskTime returns a boolean if a field has been set.

### GetBidMICId

`func (o *QuoteFuture) GetBidMICId() string`

GetBidMICId returns the BidMICId field if non-nil, zero value otherwise.

### GetBidMICIdOk

`func (o *QuoteFuture) GetBidMICIdOk() (*string, bool)`

GetBidMICIdOk returns a tuple with the BidMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMICId

`func (o *QuoteFuture) SetBidMICId(v string)`

SetBidMICId sets BidMICId field to given value.

### HasBidMICId

`func (o *QuoteFuture) HasBidMICId() bool`

HasBidMICId returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteFuture) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteFuture) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteFuture) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteFuture) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteFuture) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteFuture) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteFuture) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteFuture) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetBidTime

`func (o *QuoteFuture) GetBidTime() int64`

GetBidTime returns the BidTime field if non-nil, zero value otherwise.

### GetBidTimeOk

`func (o *QuoteFuture) GetBidTimeOk() (*int64, bool)`

GetBidTimeOk returns a tuple with the BidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidTime

`func (o *QuoteFuture) SetBidTime(v int64)`

SetBidTime sets BidTime field to given value.

### HasBidTime

`func (o *QuoteFuture) HasBidTime() bool`

HasBidTime returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteFuture) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteFuture) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteFuture) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteFuture) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetFuturePercentChange

`func (o *QuoteFuture) GetFuturePercentChange() float64`

GetFuturePercentChange returns the FuturePercentChange field if non-nil, zero value otherwise.

### GetFuturePercentChangeOk

`func (o *QuoteFuture) GetFuturePercentChangeOk() (*float64, bool)`

GetFuturePercentChangeOk returns a tuple with the FuturePercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePercentChange

`func (o *QuoteFuture) SetFuturePercentChange(v float64)`

SetFuturePercentChange sets FuturePercentChange field to given value.

### HasFuturePercentChange

`func (o *QuoteFuture) HasFuturePercentChange() bool`

HasFuturePercentChange returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteFuture) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteFuture) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteFuture) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteFuture) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastMICId

`func (o *QuoteFuture) GetLastMICId() string`

GetLastMICId returns the LastMICId field if non-nil, zero value otherwise.

### GetLastMICIdOk

`func (o *QuoteFuture) GetLastMICIdOk() (*string, bool)`

GetLastMICIdOk returns a tuple with the LastMICId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMICId

`func (o *QuoteFuture) SetLastMICId(v string)`

SetLastMICId sets LastMICId field to given value.

### HasLastMICId

`func (o *QuoteFuture) HasLastMICId() bool`

HasLastMICId returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteFuture) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteFuture) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteFuture) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteFuture) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *QuoteFuture) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteFuture) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteFuture) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *QuoteFuture) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteFuture) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteFuture) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteFuture) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteFuture) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *QuoteFuture) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *QuoteFuture) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *QuoteFuture) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *QuoteFuture) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteFuture) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteFuture) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteFuture) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteFuture) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetOpenInterest

`func (o *QuoteFuture) GetOpenInterest() int32`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *QuoteFuture) GetOpenInterestOk() (*int32, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *QuoteFuture) SetOpenInterest(v int32)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *QuoteFuture) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteFuture) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteFuture) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteFuture) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteFuture) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetQuoteTime

`func (o *QuoteFuture) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *QuoteFuture) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *QuoteFuture) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *QuoteFuture) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetQuotedInSession

`func (o *QuoteFuture) GetQuotedInSession() bool`

GetQuotedInSession returns the QuotedInSession field if non-nil, zero value otherwise.

### GetQuotedInSessionOk

`func (o *QuoteFuture) GetQuotedInSessionOk() (*bool, bool)`

GetQuotedInSessionOk returns a tuple with the QuotedInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotedInSession

`func (o *QuoteFuture) SetQuotedInSession(v bool)`

SetQuotedInSession sets QuotedInSession field to given value.

### HasQuotedInSession

`func (o *QuoteFuture) HasQuotedInSession() bool`

HasQuotedInSession returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteFuture) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteFuture) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteFuture) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteFuture) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetSettleTime

`func (o *QuoteFuture) GetSettleTime() int64`

GetSettleTime returns the SettleTime field if non-nil, zero value otherwise.

### GetSettleTimeOk

`func (o *QuoteFuture) GetSettleTimeOk() (*int64, bool)`

GetSettleTimeOk returns a tuple with the SettleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleTime

`func (o *QuoteFuture) SetSettleTime(v int64)`

SetSettleTime sets SettleTime field to given value.

### HasSettleTime

`func (o *QuoteFuture) HasSettleTime() bool`

HasSettleTime returns a boolean if a field has been set.

### GetTick

`func (o *QuoteFuture) GetTick() float64`

GetTick returns the Tick field if non-nil, zero value otherwise.

### GetTickOk

`func (o *QuoteFuture) GetTickOk() (*float64, bool)`

GetTickOk returns a tuple with the Tick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTick

`func (o *QuoteFuture) SetTick(v float64)`

SetTick sets Tick field to given value.

### HasTick

`func (o *QuoteFuture) HasTick() bool`

HasTick returns a boolean if a field has been set.

### GetTickAmount

`func (o *QuoteFuture) GetTickAmount() float64`

GetTickAmount returns the TickAmount field if non-nil, zero value otherwise.

### GetTickAmountOk

`func (o *QuoteFuture) GetTickAmountOk() (*float64, bool)`

GetTickAmountOk returns a tuple with the TickAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickAmount

`func (o *QuoteFuture) SetTickAmount(v float64)`

SetTickAmount sets TickAmount field to given value.

### HasTickAmount

`func (o *QuoteFuture) HasTickAmount() bool`

HasTickAmount returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteFuture) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteFuture) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteFuture) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteFuture) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteFuture) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteFuture) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteFuture) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteFuture) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


