# ReferenceEquity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cusip** | Pointer to **string** | CUSIP of Instrument | [optional] 
**Description** | Pointer to **string** | Description of Instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange Code | [optional] 
**ExchangeName** | Pointer to **string** | Exchange Name | [optional] 
**FsiDesc** | Pointer to **string** | FSI Desc | [optional] 
**HtbQuantity** | Pointer to **int32** | Hard to borrow quantity. | [optional] 
**HtbRate** | Pointer to **float64** | Hard to borrow rate. | [optional] 
**IsHardToBorrow** | Pointer to **bool** | is Hard to borrow security. | [optional] 
**IsShortable** | Pointer to **bool** | is shortable security. | [optional] 
**OtcMarketTier** | Pointer to **string** | OTC Market Tier | [optional] 

## Methods

### NewReferenceEquity

`func NewReferenceEquity() *ReferenceEquity`

NewReferenceEquity instantiates a new ReferenceEquity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceEquityWithDefaults

`func NewReferenceEquityWithDefaults() *ReferenceEquity`

NewReferenceEquityWithDefaults instantiates a new ReferenceEquity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCusip

`func (o *ReferenceEquity) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ReferenceEquity) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ReferenceEquity) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ReferenceEquity) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *ReferenceEquity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReferenceEquity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReferenceEquity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReferenceEquity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *ReferenceEquity) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReferenceEquity) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReferenceEquity) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReferenceEquity) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ReferenceEquity) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ReferenceEquity) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ReferenceEquity) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ReferenceEquity) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetFsiDesc

`func (o *ReferenceEquity) GetFsiDesc() string`

GetFsiDesc returns the FsiDesc field if non-nil, zero value otherwise.

### GetFsiDescOk

`func (o *ReferenceEquity) GetFsiDescOk() (*string, bool)`

GetFsiDescOk returns a tuple with the FsiDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsiDesc

`func (o *ReferenceEquity) SetFsiDesc(v string)`

SetFsiDesc sets FsiDesc field to given value.

### HasFsiDesc

`func (o *ReferenceEquity) HasFsiDesc() bool`

HasFsiDesc returns a boolean if a field has been set.

### GetHtbQuantity

`func (o *ReferenceEquity) GetHtbQuantity() int32`

GetHtbQuantity returns the HtbQuantity field if non-nil, zero value otherwise.

### GetHtbQuantityOk

`func (o *ReferenceEquity) GetHtbQuantityOk() (*int32, bool)`

GetHtbQuantityOk returns a tuple with the HtbQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtbQuantity

`func (o *ReferenceEquity) SetHtbQuantity(v int32)`

SetHtbQuantity sets HtbQuantity field to given value.

### HasHtbQuantity

`func (o *ReferenceEquity) HasHtbQuantity() bool`

HasHtbQuantity returns a boolean if a field has been set.

### GetHtbRate

`func (o *ReferenceEquity) GetHtbRate() float64`

GetHtbRate returns the HtbRate field if non-nil, zero value otherwise.

### GetHtbRateOk

`func (o *ReferenceEquity) GetHtbRateOk() (*float64, bool)`

GetHtbRateOk returns a tuple with the HtbRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtbRate

`func (o *ReferenceEquity) SetHtbRate(v float64)`

SetHtbRate sets HtbRate field to given value.

### HasHtbRate

`func (o *ReferenceEquity) HasHtbRate() bool`

HasHtbRate returns a boolean if a field has been set.

### GetIsHardToBorrow

`func (o *ReferenceEquity) GetIsHardToBorrow() bool`

GetIsHardToBorrow returns the IsHardToBorrow field if non-nil, zero value otherwise.

### GetIsHardToBorrowOk

`func (o *ReferenceEquity) GetIsHardToBorrowOk() (*bool, bool)`

GetIsHardToBorrowOk returns a tuple with the IsHardToBorrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardToBorrow

`func (o *ReferenceEquity) SetIsHardToBorrow(v bool)`

SetIsHardToBorrow sets IsHardToBorrow field to given value.

### HasIsHardToBorrow

`func (o *ReferenceEquity) HasIsHardToBorrow() bool`

HasIsHardToBorrow returns a boolean if a field has been set.

### GetIsShortable

`func (o *ReferenceEquity) GetIsShortable() bool`

GetIsShortable returns the IsShortable field if non-nil, zero value otherwise.

### GetIsShortableOk

`func (o *ReferenceEquity) GetIsShortableOk() (*bool, bool)`

GetIsShortableOk returns a tuple with the IsShortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShortable

`func (o *ReferenceEquity) SetIsShortable(v bool)`

SetIsShortable sets IsShortable field to given value.

### HasIsShortable

`func (o *ReferenceEquity) HasIsShortable() bool`

HasIsShortable returns a boolean if a field has been set.

### GetOtcMarketTier

`func (o *ReferenceEquity) GetOtcMarketTier() string`

GetOtcMarketTier returns the OtcMarketTier field if non-nil, zero value otherwise.

### GetOtcMarketTierOk

`func (o *ReferenceEquity) GetOtcMarketTierOk() (*string, bool)`

GetOtcMarketTierOk returns a tuple with the OtcMarketTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtcMarketTier

`func (o *ReferenceEquity) SetOtcMarketTier(v string)`

SetOtcMarketTier sets OtcMarketTier field to given value.

### HasOtcMarketTier

`func (o *ReferenceEquity) HasOtcMarketTier() bool`

HasOtcMarketTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


