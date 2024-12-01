# QuoteIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WeekHigh** | Pointer to **float64** | Higest price traded in the past 12 months, or 52 weeks | [optional] 
**Var52WeekLow** | Pointer to **float64** | Lowest price traded in the past 12 months, or 52 weeks | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 

## Methods

### NewQuoteIndex

`func NewQuoteIndex() *QuoteIndex`

NewQuoteIndex instantiates a new QuoteIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteIndexWithDefaults

`func NewQuoteIndexWithDefaults() *QuoteIndex`

NewQuoteIndexWithDefaults instantiates a new QuoteIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WeekHigh

`func (o *QuoteIndex) GetVar52WeekHigh() float64`

GetVar52WeekHigh returns the Var52WeekHigh field if non-nil, zero value otherwise.

### GetVar52WeekHighOk

`func (o *QuoteIndex) GetVar52WeekHighOk() (*float64, bool)`

GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekHigh

`func (o *QuoteIndex) SetVar52WeekHigh(v float64)`

SetVar52WeekHigh sets Var52WeekHigh field to given value.

### HasVar52WeekHigh

`func (o *QuoteIndex) HasVar52WeekHigh() bool`

HasVar52WeekHigh returns a boolean if a field has been set.

### GetVar52WeekLow

`func (o *QuoteIndex) GetVar52WeekLow() float64`

GetVar52WeekLow returns the Var52WeekLow field if non-nil, zero value otherwise.

### GetVar52WeekLowOk

`func (o *QuoteIndex) GetVar52WeekLowOk() (*float64, bool)`

GetVar52WeekLowOk returns a tuple with the Var52WeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekLow

`func (o *QuoteIndex) SetVar52WeekLow(v float64)`

SetVar52WeekLow sets Var52WeekLow field to given value.

### HasVar52WeekLow

`func (o *QuoteIndex) HasVar52WeekLow() bool`

HasVar52WeekLow returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteIndex) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteIndex) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteIndex) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteIndex) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteIndex) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteIndex) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteIndex) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteIndex) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteIndex) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteIndex) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteIndex) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteIndex) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteIndex) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteIndex) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteIndex) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteIndex) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteIndex) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteIndex) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteIndex) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteIndex) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteIndex) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteIndex) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteIndex) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteIndex) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteIndex) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteIndex) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteIndex) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteIndex) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteIndex) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteIndex) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteIndex) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteIndex) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteIndex) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteIndex) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteIndex) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteIndex) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteIndex) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteIndex) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteIndex) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteIndex) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


