# ExpirationChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ExpirationList** | Pointer to [**[]Expiration**](Expiration.md) |  | [optional] 

## Methods

### NewExpirationChain

`func NewExpirationChain() *ExpirationChain`

NewExpirationChain instantiates a new ExpirationChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpirationChainWithDefaults

`func NewExpirationChainWithDefaults() *ExpirationChain`

NewExpirationChainWithDefaults instantiates a new ExpirationChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ExpirationChain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpirationChain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpirationChain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExpirationChain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpirationList

`func (o *ExpirationChain) GetExpirationList() []Expiration`

GetExpirationList returns the ExpirationList field if non-nil, zero value otherwise.

### GetExpirationListOk

`func (o *ExpirationChain) GetExpirationListOk() (*[]Expiration, bool)`

GetExpirationListOk returns a tuple with the ExpirationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationList

`func (o *ExpirationChain) SetExpirationList(v []Expiration)`

SetExpirationList sets ExpirationList field to given value.

### HasExpirationList

`func (o *ExpirationChain) HasExpirationList() bool`

HasExpirationList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


