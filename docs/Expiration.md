# Expiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysToExpiration** | Pointer to **int32** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**ExpirationType** | Pointer to [**ExpirationType**](ExpirationType.md) |  | [optional] 
**Standard** | Pointer to **bool** |  | [optional] 
**SettlementType** | Pointer to [**SettlementType**](SettlementType.md) |  | [optional] 
**OptionRoots** | Pointer to **string** |  | [optional] 

## Methods

### NewExpiration

`func NewExpiration() *Expiration`

NewExpiration instantiates a new Expiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationWithDefaults

`func NewExpirationWithDefaults() *Expiration`

NewExpirationWithDefaults instantiates a new Expiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysToExpiration

`func (o *Expiration) GetDaysToExpiration() int32`

GetDaysToExpiration returns the DaysToExpiration field if non-nil, zero value otherwise.

### GetDaysToExpirationOk

`func (o *Expiration) GetDaysToExpirationOk() (*int32, bool)`

GetDaysToExpirationOk returns a tuple with the DaysToExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExpiration

`func (o *Expiration) SetDaysToExpiration(v int32)`

SetDaysToExpiration sets DaysToExpiration field to given value.

### HasDaysToExpiration

`func (o *Expiration) HasDaysToExpiration() bool`

HasDaysToExpiration returns a boolean if a field has been set.

### GetExpiration

`func (o *Expiration) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Expiration) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Expiration) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Expiration) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetExpirationType

`func (o *Expiration) GetExpirationType() ExpirationType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *Expiration) GetExpirationTypeOk() (*ExpirationType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *Expiration) SetExpirationType(v ExpirationType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *Expiration) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetStandard

`func (o *Expiration) GetStandard() bool`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *Expiration) GetStandardOk() (*bool, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *Expiration) SetStandard(v bool)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *Expiration) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetSettlementType

`func (o *Expiration) GetSettlementType() SettlementType`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *Expiration) GetSettlementTypeOk() (*SettlementType, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *Expiration) SetSettlementType(v SettlementType)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *Expiration) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetOptionRoots

`func (o *Expiration) GetOptionRoots() string`

GetOptionRoots returns the OptionRoots field if non-nil, zero value otherwise.

### GetOptionRootsOk

`func (o *Expiration) GetOptionRootsOk() (*string, bool)`

GetOptionRootsOk returns a tuple with the OptionRoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionRoots

`func (o *Expiration) SetOptionRoots(v string)`

SetOptionRoots sets OptionRoots field to given value.

### HasOptionRoots

`func (o *Expiration) HasOptionRoots() bool`

HasOptionRoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


