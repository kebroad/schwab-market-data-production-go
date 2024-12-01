# ReferenceFutureOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**Description** | Pointer to **string** | Description of Instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange Code | [optional] 
**ExchangeName** | Pointer to **string** | Exchange Name | [optional] 
**Multiplier** | Pointer to **float64** | Option multiplier | [optional] 
**ExpirationDate** | Pointer to **int64** | date of expiration in long | [optional] 
**ExpirationStyle** | Pointer to **string** | Style of expiration | [optional] 
**StrikePrice** | Pointer to **float64** | Strike Price | [optional] 
**Underlying** | Pointer to **string** | A company, index or fund name | [optional] 

## Methods

### NewReferenceFutureOption

`func NewReferenceFutureOption() *ReferenceFutureOption`

NewReferenceFutureOption instantiates a new ReferenceFutureOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceFutureOptionWithDefaults

`func NewReferenceFutureOptionWithDefaults() *ReferenceFutureOption`

NewReferenceFutureOptionWithDefaults instantiates a new ReferenceFutureOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *ReferenceFutureOption) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ReferenceFutureOption) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ReferenceFutureOption) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ReferenceFutureOption) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetDescription

`func (o *ReferenceFutureOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReferenceFutureOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReferenceFutureOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReferenceFutureOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *ReferenceFutureOption) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReferenceFutureOption) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReferenceFutureOption) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReferenceFutureOption) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ReferenceFutureOption) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ReferenceFutureOption) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ReferenceFutureOption) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ReferenceFutureOption) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetMultiplier

`func (o *ReferenceFutureOption) GetMultiplier() float64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *ReferenceFutureOption) GetMultiplierOk() (*float64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *ReferenceFutureOption) SetMultiplier(v float64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *ReferenceFutureOption) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ReferenceFutureOption) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ReferenceFutureOption) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ReferenceFutureOption) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ReferenceFutureOption) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetExpirationStyle

`func (o *ReferenceFutureOption) GetExpirationStyle() string`

GetExpirationStyle returns the ExpirationStyle field if non-nil, zero value otherwise.

### GetExpirationStyleOk

`func (o *ReferenceFutureOption) GetExpirationStyleOk() (*string, bool)`

GetExpirationStyleOk returns a tuple with the ExpirationStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStyle

`func (o *ReferenceFutureOption) SetExpirationStyle(v string)`

SetExpirationStyle sets ExpirationStyle field to given value.

### HasExpirationStyle

`func (o *ReferenceFutureOption) HasExpirationStyle() bool`

HasExpirationStyle returns a boolean if a field has been set.

### GetStrikePrice

`func (o *ReferenceFutureOption) GetStrikePrice() float64`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *ReferenceFutureOption) GetStrikePriceOk() (*float64, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *ReferenceFutureOption) SetStrikePrice(v float64)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *ReferenceFutureOption) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetUnderlying

`func (o *ReferenceFutureOption) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *ReferenceFutureOption) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *ReferenceFutureOption) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *ReferenceFutureOption) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


