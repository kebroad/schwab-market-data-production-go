# ReferenceMutualFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cusip** | Pointer to **string** | CUSIP of Instrument | [optional] 
**Description** | Pointer to **string** | Description of Instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange Code | [optional] [default to "m"]
**ExchangeName** | Pointer to **string** | Exchange Name | [optional] [default to "MUTUAL_FUND"]

## Methods

### NewReferenceMutualFund

`func NewReferenceMutualFund() *ReferenceMutualFund`

NewReferenceMutualFund instantiates a new ReferenceMutualFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceMutualFundWithDefaults

`func NewReferenceMutualFundWithDefaults() *ReferenceMutualFund`

NewReferenceMutualFundWithDefaults instantiates a new ReferenceMutualFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCusip

`func (o *ReferenceMutualFund) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ReferenceMutualFund) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ReferenceMutualFund) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ReferenceMutualFund) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDescription

`func (o *ReferenceMutualFund) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReferenceMutualFund) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReferenceMutualFund) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReferenceMutualFund) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *ReferenceMutualFund) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReferenceMutualFund) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReferenceMutualFund) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReferenceMutualFund) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ReferenceMutualFund) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ReferenceMutualFund) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ReferenceMutualFund) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ReferenceMutualFund) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


