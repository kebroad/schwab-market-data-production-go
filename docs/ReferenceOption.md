# ReferenceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | Pointer to [**ContractType**](ContractType.md) |  | [optional] 
**Cusip** | Pointer to **string** | CUSIP of Instrument | [optional] 
**DaysToExpiration** | Pointer to **int32** | Days to Expiration | [optional] 
**Deliverables** | Pointer to **string** | Unit of trade | [optional] 
**Description** | Pointer to **string** | Description of Instrument | [optional] 
**Exchange** | Pointer to **string** | Exchange Code | [optional] [default to "o"]
**ExchangeName** | Pointer to **string** | Exchange Name | [optional] 
**ExerciseType** | Pointer to [**ExerciseType**](ExerciseType.md) |  | [optional] 
**ExpirationDay** | Pointer to **int32** | Expiration Day | [optional] 
**ExpirationMonth** | Pointer to **int32** | Expiration Month | [optional] 
**ExpirationType** | Pointer to [**ExpirationType**](ExpirationType.md) |  | [optional] 
**ExpirationYear** | Pointer to **int32** | Expiration Year | [optional] 
**IsPennyPilot** | Pointer to **bool** | Is this contract part of the Penny Pilot program | [optional] 
**LastTradingDay** | Pointer to **int64** | milliseconds since epoch | [optional] 
**Multiplier** | Pointer to **float64** | Option multiplier | [optional] 
**SettlementType** | Pointer to [**SettlementType**](SettlementType.md) |  | [optional] 
**StrikePrice** | Pointer to **float64** | Strike Price | [optional] 
**Underlying** | Pointer to **string** | A company, index or fund name | [optional] 

## Methods

### NewReferenceOption

`func NewReferenceOption() *ReferenceOption`

NewReferenceOption instantiates a new ReferenceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceOptionWithDefaults

`func NewReferenceOptionWithDefaults() *ReferenceOption`

NewReferenceOptionWithDefaults instantiates a new ReferenceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *ReferenceOption) GetContractType() ContractType`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ReferenceOption) GetContractTypeOk() (*ContractType, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ReferenceOption) SetContractType(v ContractType)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ReferenceOption) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCusip

`func (o *ReferenceOption) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *ReferenceOption) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *ReferenceOption) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *ReferenceOption) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetDaysToExpiration

`func (o *ReferenceOption) GetDaysToExpiration() int32`

GetDaysToExpiration returns the DaysToExpiration field if non-nil, zero value otherwise.

### GetDaysToExpirationOk

`func (o *ReferenceOption) GetDaysToExpirationOk() (*int32, bool)`

GetDaysToExpirationOk returns a tuple with the DaysToExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExpiration

`func (o *ReferenceOption) SetDaysToExpiration(v int32)`

SetDaysToExpiration sets DaysToExpiration field to given value.

### HasDaysToExpiration

`func (o *ReferenceOption) HasDaysToExpiration() bool`

HasDaysToExpiration returns a boolean if a field has been set.

### GetDeliverables

`func (o *ReferenceOption) GetDeliverables() string`

GetDeliverables returns the Deliverables field if non-nil, zero value otherwise.

### GetDeliverablesOk

`func (o *ReferenceOption) GetDeliverablesOk() (*string, bool)`

GetDeliverablesOk returns a tuple with the Deliverables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverables

`func (o *ReferenceOption) SetDeliverables(v string)`

SetDeliverables sets Deliverables field to given value.

### HasDeliverables

`func (o *ReferenceOption) HasDeliverables() bool`

HasDeliverables returns a boolean if a field has been set.

### GetDescription

`func (o *ReferenceOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReferenceOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReferenceOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReferenceOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchange

`func (o *ReferenceOption) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *ReferenceOption) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *ReferenceOption) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *ReferenceOption) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExchangeName

`func (o *ReferenceOption) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ReferenceOption) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ReferenceOption) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ReferenceOption) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetExerciseType

`func (o *ReferenceOption) GetExerciseType() ExerciseType`

GetExerciseType returns the ExerciseType field if non-nil, zero value otherwise.

### GetExerciseTypeOk

`func (o *ReferenceOption) GetExerciseTypeOk() (*ExerciseType, bool)`

GetExerciseTypeOk returns a tuple with the ExerciseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExerciseType

`func (o *ReferenceOption) SetExerciseType(v ExerciseType)`

SetExerciseType sets ExerciseType field to given value.

### HasExerciseType

`func (o *ReferenceOption) HasExerciseType() bool`

HasExerciseType returns a boolean if a field has been set.

### GetExpirationDay

`func (o *ReferenceOption) GetExpirationDay() int32`

GetExpirationDay returns the ExpirationDay field if non-nil, zero value otherwise.

### GetExpirationDayOk

`func (o *ReferenceOption) GetExpirationDayOk() (*int32, bool)`

GetExpirationDayOk returns a tuple with the ExpirationDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDay

`func (o *ReferenceOption) SetExpirationDay(v int32)`

SetExpirationDay sets ExpirationDay field to given value.

### HasExpirationDay

`func (o *ReferenceOption) HasExpirationDay() bool`

HasExpirationDay returns a boolean if a field has been set.

### GetExpirationMonth

`func (o *ReferenceOption) GetExpirationMonth() int32`

GetExpirationMonth returns the ExpirationMonth field if non-nil, zero value otherwise.

### GetExpirationMonthOk

`func (o *ReferenceOption) GetExpirationMonthOk() (*int32, bool)`

GetExpirationMonthOk returns a tuple with the ExpirationMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationMonth

`func (o *ReferenceOption) SetExpirationMonth(v int32)`

SetExpirationMonth sets ExpirationMonth field to given value.

### HasExpirationMonth

`func (o *ReferenceOption) HasExpirationMonth() bool`

HasExpirationMonth returns a boolean if a field has been set.

### GetExpirationType

`func (o *ReferenceOption) GetExpirationType() ExpirationType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *ReferenceOption) GetExpirationTypeOk() (*ExpirationType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *ReferenceOption) SetExpirationType(v ExpirationType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *ReferenceOption) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetExpirationYear

`func (o *ReferenceOption) GetExpirationYear() int32`

GetExpirationYear returns the ExpirationYear field if non-nil, zero value otherwise.

### GetExpirationYearOk

`func (o *ReferenceOption) GetExpirationYearOk() (*int32, bool)`

GetExpirationYearOk returns a tuple with the ExpirationYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationYear

`func (o *ReferenceOption) SetExpirationYear(v int32)`

SetExpirationYear sets ExpirationYear field to given value.

### HasExpirationYear

`func (o *ReferenceOption) HasExpirationYear() bool`

HasExpirationYear returns a boolean if a field has been set.

### GetIsPennyPilot

`func (o *ReferenceOption) GetIsPennyPilot() bool`

GetIsPennyPilot returns the IsPennyPilot field if non-nil, zero value otherwise.

### GetIsPennyPilotOk

`func (o *ReferenceOption) GetIsPennyPilotOk() (*bool, bool)`

GetIsPennyPilotOk returns a tuple with the IsPennyPilot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPennyPilot

`func (o *ReferenceOption) SetIsPennyPilot(v bool)`

SetIsPennyPilot sets IsPennyPilot field to given value.

### HasIsPennyPilot

`func (o *ReferenceOption) HasIsPennyPilot() bool`

HasIsPennyPilot returns a boolean if a field has been set.

### GetLastTradingDay

`func (o *ReferenceOption) GetLastTradingDay() int64`

GetLastTradingDay returns the LastTradingDay field if non-nil, zero value otherwise.

### GetLastTradingDayOk

`func (o *ReferenceOption) GetLastTradingDayOk() (*int64, bool)`

GetLastTradingDayOk returns a tuple with the LastTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradingDay

`func (o *ReferenceOption) SetLastTradingDay(v int64)`

SetLastTradingDay sets LastTradingDay field to given value.

### HasLastTradingDay

`func (o *ReferenceOption) HasLastTradingDay() bool`

HasLastTradingDay returns a boolean if a field has been set.

### GetMultiplier

`func (o *ReferenceOption) GetMultiplier() float64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *ReferenceOption) GetMultiplierOk() (*float64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *ReferenceOption) SetMultiplier(v float64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *ReferenceOption) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetSettlementType

`func (o *ReferenceOption) GetSettlementType() SettlementType`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *ReferenceOption) GetSettlementTypeOk() (*SettlementType, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *ReferenceOption) SetSettlementType(v SettlementType)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *ReferenceOption) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetStrikePrice

`func (o *ReferenceOption) GetStrikePrice() float64`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *ReferenceOption) GetStrikePriceOk() (*float64, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *ReferenceOption) SetStrikePrice(v float64)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *ReferenceOption) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetUnderlying

`func (o *ReferenceOption) GetUnderlying() string`

GetUnderlying returns the Underlying field if non-nil, zero value otherwise.

### GetUnderlyingOk

`func (o *ReferenceOption) GetUnderlyingOk() (*string, bool)`

GetUnderlyingOk returns a tuple with the Underlying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlying

`func (o *ReferenceOption) SetUnderlying(v string)`

SetUnderlying sets Underlying field to given value.

### HasUnderlying

`func (o *ReferenceOption) HasUnderlying() bool`

HasUnderlying returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


