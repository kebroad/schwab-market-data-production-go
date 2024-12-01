# CandleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candles** | Pointer to [**[]Candle**](Candle.md) |  | [optional] 
**Empty** | Pointer to **bool** |  | [optional] 
**PreviousClose** | Pointer to **float64** |  | [optional] 
**PreviousCloseDate** | Pointer to **int64** |  | [optional] 
**PreviousCloseDateISO8601** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewCandleList

`func NewCandleList() *CandleList`

NewCandleList instantiates a new CandleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandleListWithDefaults

`func NewCandleListWithDefaults() *CandleList`

NewCandleListWithDefaults instantiates a new CandleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandles

`func (o *CandleList) GetCandles() []Candle`

GetCandles returns the Candles field if non-nil, zero value otherwise.

### GetCandlesOk

`func (o *CandleList) GetCandlesOk() (*[]Candle, bool)`

GetCandlesOk returns a tuple with the Candles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandles

`func (o *CandleList) SetCandles(v []Candle)`

SetCandles sets Candles field to given value.

### HasCandles

`func (o *CandleList) HasCandles() bool`

HasCandles returns a boolean if a field has been set.

### GetEmpty

`func (o *CandleList) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *CandleList) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *CandleList) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *CandleList) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetPreviousClose

`func (o *CandleList) GetPreviousClose() float64`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *CandleList) GetPreviousCloseOk() (*float64, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *CandleList) SetPreviousClose(v float64)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *CandleList) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetPreviousCloseDate

`func (o *CandleList) GetPreviousCloseDate() int64`

GetPreviousCloseDate returns the PreviousCloseDate field if non-nil, zero value otherwise.

### GetPreviousCloseDateOk

`func (o *CandleList) GetPreviousCloseDateOk() (*int64, bool)`

GetPreviousCloseDateOk returns a tuple with the PreviousCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCloseDate

`func (o *CandleList) SetPreviousCloseDate(v int64)`

SetPreviousCloseDate sets PreviousCloseDate field to given value.

### HasPreviousCloseDate

`func (o *CandleList) HasPreviousCloseDate() bool`

HasPreviousCloseDate returns a boolean if a field has been set.

### GetPreviousCloseDateISO8601

`func (o *CandleList) GetPreviousCloseDateISO8601() string`

GetPreviousCloseDateISO8601 returns the PreviousCloseDateISO8601 field if non-nil, zero value otherwise.

### GetPreviousCloseDateISO8601Ok

`func (o *CandleList) GetPreviousCloseDateISO8601Ok() (*string, bool)`

GetPreviousCloseDateISO8601Ok returns a tuple with the PreviousCloseDateISO8601 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCloseDateISO8601

`func (o *CandleList) SetPreviousCloseDateISO8601(v string)`

SetPreviousCloseDateISO8601 sets PreviousCloseDateISO8601 field to given value.

### HasPreviousCloseDateISO8601

`func (o *CandleList) HasPreviousCloseDateISO8601() bool`

HasPreviousCloseDateISO8601 returns a boolean if a field has been set.

### GetSymbol

`func (o *CandleList) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CandleList) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CandleList) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CandleList) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


