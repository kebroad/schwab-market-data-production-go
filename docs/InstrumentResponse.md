# InstrumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**AssetType** | Pointer to **string** |  | [optional] 
**BondFactor** | Pointer to **string** |  | [optional] 
**BondMultiplier** | Pointer to **string** |  | [optional] 
**BondPrice** | Pointer to **float32** |  | [optional] 
**Fundamental** | Pointer to [**FundamentalInst**](FundamentalInst.md) |  | [optional] 
**InstrumentInfo** | Pointer to [**Instrument**](Instrument.md) |  | [optional] 
**BondInstrumentInfo** | Pointer to [**Bond**](Bond.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInstrumentResponse

`func NewInstrumentResponse() *InstrumentResponse`

NewInstrumentResponse instantiates a new InstrumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstrumentResponseWithDefaults

`func NewInstrumentResponseWithDefaults() *InstrumentResponse`

NewInstrumentResponseWithDefaults instantiates a new InstrumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCusip

`func (o *InstrumentResponse) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *InstrumentResponse) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *InstrumentResponse) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *InstrumentResponse) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *InstrumentResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InstrumentResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InstrumentResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *InstrumentResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *InstrumentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstrumentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstrumentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstrumentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *InstrumentResponse) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InstrumentResponse) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InstrumentResponse) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *InstrumentResponse) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetAssetType

`func (o *InstrumentResponse) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *InstrumentResponse) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *InstrumentResponse) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *InstrumentResponse) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetBondFactor

`func (o *InstrumentResponse) GetBondFactor() string`

GetBondFactor returns the BondFactor field if non-nil, zero value otherwise.

### GetBondFactorOk

`func (o *InstrumentResponse) GetBondFactorOk() (*string, bool)`

GetBondFactorOk returns a tuple with the BondFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondFactor

`func (o *InstrumentResponse) SetBondFactor(v string)`

SetBondFactor sets BondFactor field to given value.

### HasBondFactor

`func (o *InstrumentResponse) HasBondFactor() bool`

HasBondFactor returns a boolean if a field has been set.

### GetBondMultiplier

`func (o *InstrumentResponse) GetBondMultiplier() string`

GetBondMultiplier returns the BondMultiplier field if non-nil, zero value otherwise.

### GetBondMultiplierOk

`func (o *InstrumentResponse) GetBondMultiplierOk() (*string, bool)`

GetBondMultiplierOk returns a tuple with the BondMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondMultiplier

`func (o *InstrumentResponse) SetBondMultiplier(v string)`

SetBondMultiplier sets BondMultiplier field to given value.

### HasBondMultiplier

`func (o *InstrumentResponse) HasBondMultiplier() bool`

HasBondMultiplier returns a boolean if a field has been set.

### GetBondPrice

`func (o *InstrumentResponse) GetBondPrice() float32`

GetBondPrice returns the BondPrice field if non-nil, zero value otherwise.

### GetBondPriceOk

`func (o *InstrumentResponse) GetBondPriceOk() (*float32, bool)`

GetBondPriceOk returns a tuple with the BondPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondPrice

`func (o *InstrumentResponse) SetBondPrice(v float32)`

SetBondPrice sets BondPrice field to given value.

### HasBondPrice

`func (o *InstrumentResponse) HasBondPrice() bool`

HasBondPrice returns a boolean if a field has been set.

### GetFundamental

`func (o *InstrumentResponse) GetFundamental() FundamentalInst`

GetFundamental returns the Fundamental field if non-nil, zero value otherwise.

### GetFundamentalOk

`func (o *InstrumentResponse) GetFundamentalOk() (*FundamentalInst, bool)`

GetFundamentalOk returns a tuple with the Fundamental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundamental

`func (o *InstrumentResponse) SetFundamental(v FundamentalInst)`

SetFundamental sets Fundamental field to given value.

### HasFundamental

`func (o *InstrumentResponse) HasFundamental() bool`

HasFundamental returns a boolean if a field has been set.

### GetInstrumentInfo

`func (o *InstrumentResponse) GetInstrumentInfo() Instrument`

GetInstrumentInfo returns the InstrumentInfo field if non-nil, zero value otherwise.

### GetInstrumentInfoOk

`func (o *InstrumentResponse) GetInstrumentInfoOk() (*Instrument, bool)`

GetInstrumentInfoOk returns a tuple with the InstrumentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentInfo

`func (o *InstrumentResponse) SetInstrumentInfo(v Instrument)`

SetInstrumentInfo sets InstrumentInfo field to given value.

### HasInstrumentInfo

`func (o *InstrumentResponse) HasInstrumentInfo() bool`

HasInstrumentInfo returns a boolean if a field has been set.

### GetBondInstrumentInfo

`func (o *InstrumentResponse) GetBondInstrumentInfo() Bond`

GetBondInstrumentInfo returns the BondInstrumentInfo field if non-nil, zero value otherwise.

### GetBondInstrumentInfoOk

`func (o *InstrumentResponse) GetBondInstrumentInfoOk() (*Bond, bool)`

GetBondInstrumentInfoOk returns a tuple with the BondInstrumentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondInstrumentInfo

`func (o *InstrumentResponse) SetBondInstrumentInfo(v Bond)`

SetBondInstrumentInfo sets BondInstrumentInfo field to given value.

### HasBondInstrumentInfo

`func (o *InstrumentResponse) HasBondInstrumentInfo() bool`

HasBondInstrumentInfo returns a boolean if a field has been set.

### GetType

`func (o *InstrumentResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstrumentResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstrumentResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InstrumentResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


