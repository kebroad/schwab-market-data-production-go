# QuoteMutualFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WeekHigh** | Pointer to **float64** | Higest price traded in the past 12 months, or 52 weeks | [optional] 
**Var52WeekLow** | Pointer to **float64** | Lowest price traded in the past 12 months, or 52 weeks | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**NAV** | Pointer to **float64** | Net Asset Value | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 

## Methods

### NewQuoteMutualFund

`func NewQuoteMutualFund() *QuoteMutualFund`

NewQuoteMutualFund instantiates a new QuoteMutualFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteMutualFundWithDefaults

`func NewQuoteMutualFundWithDefaults() *QuoteMutualFund`

NewQuoteMutualFundWithDefaults instantiates a new QuoteMutualFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WeekHigh

`func (o *QuoteMutualFund) GetVar52WeekHigh() float64`

GetVar52WeekHigh returns the Var52WeekHigh field if non-nil, zero value otherwise.

### GetVar52WeekHighOk

`func (o *QuoteMutualFund) GetVar52WeekHighOk() (*float64, bool)`

GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekHigh

`func (o *QuoteMutualFund) SetVar52WeekHigh(v float64)`

SetVar52WeekHigh sets Var52WeekHigh field to given value.

### HasVar52WeekHigh

`func (o *QuoteMutualFund) HasVar52WeekHigh() bool`

HasVar52WeekHigh returns a boolean if a field has been set.

### GetVar52WeekLow

`func (o *QuoteMutualFund) GetVar52WeekLow() float64`

GetVar52WeekLow returns the Var52WeekLow field if non-nil, zero value otherwise.

### GetVar52WeekLowOk

`func (o *QuoteMutualFund) GetVar52WeekLowOk() (*float64, bool)`

GetVar52WeekLowOk returns a tuple with the Var52WeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekLow

`func (o *QuoteMutualFund) SetVar52WeekLow(v float64)`

SetVar52WeekLow sets Var52WeekLow field to given value.

### HasVar52WeekLow

`func (o *QuoteMutualFund) HasVar52WeekLow() bool`

HasVar52WeekLow returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteMutualFund) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteMutualFund) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteMutualFund) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteMutualFund) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetNAV

`func (o *QuoteMutualFund) GetNAV() float64`

GetNAV returns the NAV field if non-nil, zero value otherwise.

### GetNAVOk

`func (o *QuoteMutualFund) GetNAVOk() (*float64, bool)`

GetNAVOk returns a tuple with the NAV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAV

`func (o *QuoteMutualFund) SetNAV(v float64)`

SetNAV sets NAV field to given value.

### HasNAV

`func (o *QuoteMutualFund) HasNAV() bool`

HasNAV returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteMutualFund) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteMutualFund) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteMutualFund) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteMutualFund) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteMutualFund) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteMutualFund) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteMutualFund) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteMutualFund) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteMutualFund) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteMutualFund) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteMutualFund) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteMutualFund) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteMutualFund) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteMutualFund) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteMutualFund) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteMutualFund) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteMutualFund) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteMutualFund) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteMutualFund) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteMutualFund) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


