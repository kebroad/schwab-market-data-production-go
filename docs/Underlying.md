# Underlying

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ask** | Pointer to **float64** |  | [optional] 
**AskSize** | Pointer to **int32** |  | [optional] 
**Bid** | Pointer to **float64** |  | [optional] 
**BidSize** | Pointer to **int32** |  | [optional] 
**Change** | Pointer to **float64** |  | [optional] 
**Close** | Pointer to **float64** |  | [optional] 
**Delayed** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekHigh** | Pointer to **float64** |  | [optional] 
**FiftyTwoWeekLow** | Pointer to **float64** |  | [optional] 
**HighPrice** | Pointer to **float64** |  | [optional] 
**Last** | Pointer to **float64** |  | [optional] 
**LowPrice** | Pointer to **float64** |  | [optional] 
**Mark** | Pointer to **float64** |  | [optional] 
**MarkChange** | Pointer to **float64** |  | [optional] 
**MarkPercentChange** | Pointer to **float64** |  | [optional] 
**OpenPrice** | Pointer to **float64** |  | [optional] 
**PercentChange** | Pointer to **float64** |  | [optional] 
**QuoteTime** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TotalVolume** | Pointer to **int64** |  | [optional] 
**TradeTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewUnderlying

`func NewUnderlying() *Underlying`

NewUnderlying instantiates a new Underlying object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderlyingWithDefaults

`func NewUnderlyingWithDefaults() *Underlying`

NewUnderlyingWithDefaults instantiates a new Underlying object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsk

`func (o *Underlying) GetAsk() float64`

GetAsk returns the Ask field if non-nil, zero value otherwise.

### GetAskOk

`func (o *Underlying) GetAskOk() (*float64, bool)`

GetAskOk returns a tuple with the Ask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsk

`func (o *Underlying) SetAsk(v float64)`

SetAsk sets Ask field to given value.

### HasAsk

`func (o *Underlying) HasAsk() bool`

HasAsk returns a boolean if a field has been set.

### GetAskSize

`func (o *Underlying) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *Underlying) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *Underlying) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *Underlying) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBid

`func (o *Underlying) GetBid() float64`

GetBid returns the Bid field if non-nil, zero value otherwise.

### GetBidOk

`func (o *Underlying) GetBidOk() (*float64, bool)`

GetBidOk returns a tuple with the Bid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBid

`func (o *Underlying) SetBid(v float64)`

SetBid sets Bid field to given value.

### HasBid

`func (o *Underlying) HasBid() bool`

HasBid returns a boolean if a field has been set.

### GetBidSize

`func (o *Underlying) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *Underlying) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *Underlying) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *Underlying) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetChange

`func (o *Underlying) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Underlying) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Underlying) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *Underlying) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetClose

`func (o *Underlying) GetClose() float64`

GetClose returns the Close field if non-nil, zero value otherwise.

### GetCloseOk

`func (o *Underlying) GetCloseOk() (*float64, bool)`

GetCloseOk returns a tuple with the Close field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClose

`func (o *Underlying) SetClose(v float64)`

SetClose sets Close field to given value.

### HasClose

`func (o *Underlying) HasClose() bool`

HasClose returns a boolean if a field has been set.

### GetDelayed

`func (o *Underlying) GetDelayed() bool`

GetDelayed returns the Delayed field if non-nil, zero value otherwise.

### GetDelayedOk

`func (o *Underlying) GetDelayedOk() (*bool, bool)`

GetDelayedOk returns a tuple with the Delayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayed

`func (o *Underlying) SetDelayed(v bool)`

SetDelayed sets Delayed field to given value.

### HasDelayed

`func (o *Underlying) HasDelayed() bool`

HasDelayed returns a boolean if a field has been set.

### GetDescription

`func (o *Underlying) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Underlying) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Underlying) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Underlying) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeName

`func (o *Underlying) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *Underlying) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *Underlying) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *Underlying) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFiftyTwoWeekHigh

`func (o *Underlying) GetFiftyTwoWeekHigh() float64`

GetFiftyTwoWeekHigh returns the FiftyTwoWeekHigh field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighOk

`func (o *Underlying) GetFiftyTwoWeekHighOk() (*float64, bool)`

GetFiftyTwoWeekHighOk returns a tuple with the FiftyTwoWeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHigh

`func (o *Underlying) SetFiftyTwoWeekHigh(v float64)`

SetFiftyTwoWeekHigh sets FiftyTwoWeekHigh field to given value.

### HasFiftyTwoWeekHigh

`func (o *Underlying) HasFiftyTwoWeekHigh() bool`

HasFiftyTwoWeekHigh returns a boolean if a field has been set.

### GetFiftyTwoWeekLow

`func (o *Underlying) GetFiftyTwoWeekLow() float64`

GetFiftyTwoWeekLow returns the FiftyTwoWeekLow field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowOk

`func (o *Underlying) GetFiftyTwoWeekLowOk() (*float64, bool)`

GetFiftyTwoWeekLowOk returns a tuple with the FiftyTwoWeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLow

`func (o *Underlying) SetFiftyTwoWeekLow(v float64)`

SetFiftyTwoWeekLow sets FiftyTwoWeekLow field to given value.

### HasFiftyTwoWeekLow

`func (o *Underlying) HasFiftyTwoWeekLow() bool`

HasFiftyTwoWeekLow returns a boolean if a field has been set.

### GetHighPrice

`func (o *Underlying) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *Underlying) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *Underlying) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *Underlying) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLast

`func (o *Underlying) GetLast() float64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Underlying) GetLastOk() (*float64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Underlying) SetLast(v float64)`

SetLast sets Last field to given value.

### HasLast

`func (o *Underlying) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLowPrice

`func (o *Underlying) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *Underlying) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *Underlying) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *Underlying) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *Underlying) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *Underlying) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *Underlying) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *Underlying) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *Underlying) GetMarkChange() float64`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *Underlying) GetMarkChangeOk() (*float64, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *Underlying) SetMarkChange(v float64)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *Underlying) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *Underlying) GetMarkPercentChange() float64`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *Underlying) GetMarkPercentChangeOk() (*float64, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *Underlying) SetMarkPercentChange(v float64)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *Underlying) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetOpenPrice

`func (o *Underlying) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *Underlying) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *Underlying) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *Underlying) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetPercentChange

`func (o *Underlying) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *Underlying) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *Underlying) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *Underlying) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetQuoteTime

`func (o *Underlying) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *Underlying) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *Underlying) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *Underlying) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetSymbol

`func (o *Underlying) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Underlying) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Underlying) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Underlying) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalVolume

`func (o *Underlying) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *Underlying) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *Underlying) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *Underlying) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *Underlying) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *Underlying) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *Underlying) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *Underlying) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


