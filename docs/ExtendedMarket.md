# ExtendedMarket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float64** | Extended market ask price | [optional] 
**AskSize** | Pointer to **int32** | Extended market ask size | [optional] 
**BidPrice** | Pointer to **float64** | Extended market bid price | [optional] 
**BidSize** | Pointer to **int32** | Extended market bid size | [optional] 
**LastPrice** | Pointer to **float64** | Extended market last price | [optional] 
**LastSize** | Pointer to **int32** | Regular market last size | [optional] 
**Mark** | Pointer to **float64** | mark price | [optional] 
**QuoteTime** | Pointer to **int64** | Extended market quote time in milliseconds since Epoch | [optional] 
**TotalVolume** | Pointer to **float32** | Total volume | [optional] 
**TradeTime** | Pointer to **int64** | Extended market trade time in milliseconds since Epoch | [optional] 

## Methods

### NewExtendedMarket

`func NewExtendedMarket() *ExtendedMarket`

NewExtendedMarket instantiates a new ExtendedMarket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedMarketWithDefaults

`func NewExtendedMarketWithDefaults() *ExtendedMarket`

NewExtendedMarketWithDefaults instantiates a new ExtendedMarket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *ExtendedMarket) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *ExtendedMarket) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *ExtendedMarket) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *ExtendedMarket) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *ExtendedMarket) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *ExtendedMarket) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *ExtendedMarket) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *ExtendedMarket) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *ExtendedMarket) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *ExtendedMarket) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *ExtendedMarket) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *ExtendedMarket) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *ExtendedMarket) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *ExtendedMarket) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *ExtendedMarket) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *ExtendedMarket) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetLastPrice

`func (o *ExtendedMarket) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *ExtendedMarket) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *ExtendedMarket) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *ExtendedMarket) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *ExtendedMarket) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *ExtendedMarket) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *ExtendedMarket) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *ExtendedMarket) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetMark

`func (o *ExtendedMarket) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *ExtendedMarket) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *ExtendedMarket) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *ExtendedMarket) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetQuoteTime

`func (o *ExtendedMarket) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *ExtendedMarket) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *ExtendedMarket) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *ExtendedMarket) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetTotalVolume

`func (o *ExtendedMarket) GetTotalVolume() float32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *ExtendedMarket) GetTotalVolumeOk() (*float32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *ExtendedMarket) SetTotalVolume(v float32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *ExtendedMarket) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *ExtendedMarket) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *ExtendedMarket) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *ExtendedMarket) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *ExtendedMarket) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


