# ReferenceFuture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of Instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange Code | [optional] 
**ExchangeName** | Pointer to **string** | Exchange Name | [optional] 
**FutureActiveSymbol** | Pointer to **string** | Active symbol | [optional] 
**FutureExpirationDate** | Pointer to **float32** | Future expiration date in milliseconds since epoch | [optional] 
**FutureIsActive** | Pointer to **bool** | Future is active | [optional] 
**FutureMultiplier** | Pointer to **float64** | Future multiplier | [optional] 
**FuturePriceFormat** | Pointer to **string** | Price format | [optional] 
**FutureSettlementPrice** | Pointer to **float64** | Future Settlement Price | [optional] 
**FutureTradingHours** | Pointer to **string** | Trading Hours | [optional] 
**Product** | Pointer to **string** | Futures product symbol | [optional] 

## Methods

### NewReferenceFuture

`func NewReferenceFuture() *ReferenceFuture`

NewReferenceFuture instantiates a new ReferenceFuture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceFutureWithDefaults

`func NewReferenceFutureWithDefaults() *ReferenceFuture`

NewReferenceFutureWithDefaults instantiates a new ReferenceFuture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ReferenceFuture) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReferenceFuture) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReferenceFuture) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReferenceFuture) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *ReferenceFuture) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReferenceFuture) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReferenceFuture) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReferenceFuture) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ReferenceFuture) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ReferenceFuture) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ReferenceFuture) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ReferenceFuture) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFutureActiveSymbol

`func (o *ReferenceFuture) GetFutureActiveSymbol() string`

GetFutureActiveSymbol returns the FutureActiveSymbol field if non-nil, zero value otherwise.

### GetFutureActiveSymbolOk

`func (o *ReferenceFuture) GetFutureActiveSymbolOk() (*string, bool)`

GetFutureActiveSymbolOk returns a tuple with the FutureActiveSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureActiveSymbol

`func (o *ReferenceFuture) SetFutureActiveSymbol(v string)`

SetFutureActiveSymbol sets FutureActiveSymbol field to given value.

### HasFutureActiveSymbol

`func (o *ReferenceFuture) HasFutureActiveSymbol() bool`

HasFutureActiveSymbol returns a boolean if a field has been set.

### GetFutureExpirationDate

`func (o *ReferenceFuture) GetFutureExpirationDate() float32`

GetFutureExpirationDate returns the FutureExpirationDate field if non-nil, zero value otherwise.

### GetFutureExpirationDateOk

`func (o *ReferenceFuture) GetFutureExpirationDateOk() (*float32, bool)`

GetFutureExpirationDateOk returns a tuple with the FutureExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureExpirationDate

`func (o *ReferenceFuture) SetFutureExpirationDate(v float32)`

SetFutureExpirationDate sets FutureExpirationDate field to given value.

### HasFutureExpirationDate

`func (o *ReferenceFuture) HasFutureExpirationDate() bool`

HasFutureExpirationDate returns a boolean if a field has been set.

### GetFutureIsActive

`func (o *ReferenceFuture) GetFutureIsActive() bool`

GetFutureIsActive returns the FutureIsActive field if non-nil, zero value otherwise.

### GetFutureIsActiveOk

`func (o *ReferenceFuture) GetFutureIsActiveOk() (*bool, bool)`

GetFutureIsActiveOk returns a tuple with the FutureIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureIsActive

`func (o *ReferenceFuture) SetFutureIsActive(v bool)`

SetFutureIsActive sets FutureIsActive field to given value.

### HasFutureIsActive

`func (o *ReferenceFuture) HasFutureIsActive() bool`

HasFutureIsActive returns a boolean if a field has been set.

### GetFutureMultiplier

`func (o *ReferenceFuture) GetFutureMultiplier() float64`

GetFutureMultiplier returns the FutureMultiplier field if non-nil, zero value otherwise.

### GetFutureMultiplierOk

`func (o *ReferenceFuture) GetFutureMultiplierOk() (*float64, bool)`

GetFutureMultiplierOk returns a tuple with the FutureMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureMultiplier

`func (o *ReferenceFuture) SetFutureMultiplier(v float64)`

SetFutureMultiplier sets FutureMultiplier field to given value.

### HasFutureMultiplier

`func (o *ReferenceFuture) HasFutureMultiplier() bool`

HasFutureMultiplier returns a boolean if a field has been set.

### GetFuturePriceFormat

`func (o *ReferenceFuture) GetFuturePriceFormat() string`

GetFuturePriceFormat returns the FuturePriceFormat field if non-nil, zero value otherwise.

### GetFuturePriceFormatOk

`func (o *ReferenceFuture) GetFuturePriceFormatOk() (*string, bool)`

GetFuturePriceFormatOk returns a tuple with the FuturePriceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePriceFormat

`func (o *ReferenceFuture) SetFuturePriceFormat(v string)`

SetFuturePriceFormat sets FuturePriceFormat field to given value.

### HasFuturePriceFormat

`func (o *ReferenceFuture) HasFuturePriceFormat() bool`

HasFuturePriceFormat returns a boolean if a field has been set.

### GetFutureSettlementPrice

`func (o *ReferenceFuture) GetFutureSettlementPrice() float64`

GetFutureSettlementPrice returns the FutureSettlementPrice field if non-nil, zero value otherwise.

### GetFutureSettlementPriceOk

`func (o *ReferenceFuture) GetFutureSettlementPriceOk() (*float64, bool)`

GetFutureSettlementPriceOk returns a tuple with the FutureSettlementPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureSettlementPrice

`func (o *ReferenceFuture) SetFutureSettlementPrice(v float64)`

SetFutureSettlementPrice sets FutureSettlementPrice field to given value.

### HasFutureSettlementPrice

`func (o *ReferenceFuture) HasFutureSettlementPrice() bool`

HasFutureSettlementPrice returns a boolean if a field has been set.

### GetFutureTradingHours

`func (o *ReferenceFuture) GetFutureTradingHours() string`

GetFutureTradingHours returns the FutureTradingHours field if non-nil, zero value otherwise.

### GetFutureTradingHoursOk

`func (o *ReferenceFuture) GetFutureTradingHoursOk() (*string, bool)`

GetFutureTradingHoursOk returns a tuple with the FutureTradingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureTradingHours

`func (o *ReferenceFuture) SetFutureTradingHours(v string)`

SetFutureTradingHours sets FutureTradingHours field to given value.

### HasFutureTradingHours

`func (o *ReferenceFuture) HasFutureTradingHours() bool`

HasFutureTradingHours returns a boolean if a field has been set.

### GetProduct

`func (o *ReferenceFuture) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ReferenceFuture) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ReferenceFuture) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ReferenceFuture) HasProduct() bool`

HasProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


