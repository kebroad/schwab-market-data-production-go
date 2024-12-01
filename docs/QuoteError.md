# QuoteError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidCusips** | Pointer to **[]string** | list of invalid cusips from request | [optional] 
**InvalidSSIDs** | Pointer to **[]int64** | list of invalid SSIDs from request | [optional] 
**InvalidSymbols** | Pointer to **[]string** | list of invalid symbols from request | [optional] 

## Methods

### NewQuoteError

`func NewQuoteError() *QuoteError`

NewQuoteError instantiates a new QuoteError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteErrorWithDefaults

`func NewQuoteErrorWithDefaults() *QuoteError`

NewQuoteErrorWithDefaults instantiates a new QuoteError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidCusips

`func (o *QuoteError) GetInvalidCusips() []string`

GetInvalidCusips returns the InvalidCusips field if non-nil, zero value otherwise.

### GetInvalidCusipsOk

`func (o *QuoteError) GetInvalidCusipsOk() (*[]string, bool)`

GetInvalidCusipsOk returns a tuple with the InvalidCusips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCusips

`func (o *QuoteError) SetInvalidCusips(v []string)`

SetInvalidCusips sets InvalidCusips field to given value.

### HasInvalidCusips

`func (o *QuoteError) HasInvalidCusips() bool`

HasInvalidCusips returns a boolean if a field has been set.

### GetInvalidSSIDs

`func (o *QuoteError) GetInvalidSSIDs() []int64`

GetInvalidSSIDs returns the InvalidSSIDs field if non-nil, zero value otherwise.

### GetInvalidSSIDsOk

`func (o *QuoteError) GetInvalidSSIDsOk() (*[]int64, bool)`

GetInvalidSSIDsOk returns a tuple with the InvalidSSIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSSIDs

`func (o *QuoteError) SetInvalidSSIDs(v []int64)`

SetInvalidSSIDs sets InvalidSSIDs field to given value.

### HasInvalidSSIDs

`func (o *QuoteError) HasInvalidSSIDs() bool`

HasInvalidSSIDs returns a boolean if a field has been set.

### GetInvalidSymbols

`func (o *QuoteError) GetInvalidSymbols() []string`

GetInvalidSymbols returns the InvalidSymbols field if non-nil, zero value otherwise.

### GetInvalidSymbolsOk

`func (o *QuoteError) GetInvalidSymbolsOk() (*[]string, bool)`

GetInvalidSymbolsOk returns a tuple with the InvalidSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSymbols

`func (o *QuoteError) SetInvalidSymbols(v []string)`

SetInvalidSymbols sets InvalidSymbols field to given value.

### HasInvalidSymbols

`func (o *QuoteError) HasInvalidSymbols() bool`

HasInvalidSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


