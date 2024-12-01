# MutualFundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMainType** | Pointer to [**AssetMainType**](AssetMainType.md) |  | [optional] 
**AssetSubType** | Pointer to [**NullableMutualFundAssetSubType**](MutualFundAssetSubType.md) |  | [optional] 
**Ssid** | Pointer to **int64** | SSID of instrument | [optional] 
**Symbol** | Pointer to **string** | Symbol of instrument | [optional] 
**Realtime** | Pointer to **bool** | is quote realtime | [optional] 
**Fundamental** | Pointer to [**Fundamental**](Fundamental.md) |  | [optional] 
**Quote** | Pointer to [**QuoteMutualFund**](QuoteMutualFund.md) |  | [optional] 
**Reference** | Pointer to [**ReferenceMutualFund**](ReferenceMutualFund.md) |  | [optional] 

## Methods

### NewMutualFundResponse

`func NewMutualFundResponse() *MutualFundResponse`

NewMutualFundResponse instantiates a new MutualFundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundResponseWithDefaults

`func NewMutualFundResponseWithDefaults() *MutualFundResponse`

NewMutualFundResponseWithDefaults instantiates a new MutualFundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMainType

`func (o *MutualFundResponse) GetAssetMainType() AssetMainType`

GetAssetMainType returns the AssetMainType field if non-nil, zero value otherwise.

### GetAssetMainTypeOk

`func (o *MutualFundResponse) GetAssetMainTypeOk() (*AssetMainType, bool)`

GetAssetMainTypeOk returns a tuple with the AssetMainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMainType

`func (o *MutualFundResponse) SetAssetMainType(v AssetMainType)`

SetAssetMainType sets AssetMainType field to given value.

### HasAssetMainType

`func (o *MutualFundResponse) HasAssetMainType() bool`

HasAssetMainType returns a boolean if a field has been set.

### GetAssetSubType

`func (o *MutualFundResponse) GetAssetSubType() MutualFundAssetSubType`

GetAssetSubType returns the AssetSubType field if non-nil, zero value otherwise.

### GetAssetSubTypeOk

`func (o *MutualFundResponse) GetAssetSubTypeOk() (*MutualFundAssetSubType, bool)`

GetAssetSubTypeOk returns a tuple with the AssetSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSubType

`func (o *MutualFundResponse) SetAssetSubType(v MutualFundAssetSubType)`

SetAssetSubType sets AssetSubType field to given value.

### HasAssetSubType

`func (o *MutualFundResponse) HasAssetSubType() bool`

HasAssetSubType returns a boolean if a field has been set.

### SetAssetSubTypeNil

`func (o *MutualFundResponse) SetAssetSubTypeNil(b bool)`

 SetAssetSubTypeNil sets the value for AssetSubType to be an explicit nil

### UnsetAssetSubType
`func (o *MutualFundResponse) UnsetAssetSubType()`

UnsetAssetSubType ensures that no value is present for AssetSubType, not even an explicit nil
### GetSsid

`func (o *MutualFundResponse) GetSsid() int64`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *MutualFundResponse) GetSsidOk() (*int64, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *MutualFundResponse) SetSsid(v int64)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *MutualFundResponse) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSymbol

`func (o *MutualFundResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MutualFundResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRealtime

`func (o *MutualFundResponse) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *MutualFundResponse) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *MutualFundResponse) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *MutualFundResponse) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetFundamental

`func (o *MutualFundResponse) GetFundamental() Fundamental`

GetFundamental returns the Fundamental field if non-nil, zero value otherwise.

### GetFundamentalOk

`func (o *MutualFundResponse) GetFundamentalOk() (*Fundamental, bool)`

GetFundamentalOk returns a tuple with the Fundamental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundamental

`func (o *MutualFundResponse) SetFundamental(v Fundamental)`

SetFundamental sets Fundamental field to given value.

### HasFundamental

`func (o *MutualFundResponse) HasFundamental() bool`

HasFundamental returns a boolean if a field has been set.

### GetQuote

`func (o *MutualFundResponse) GetQuote() QuoteMutualFund`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *MutualFundResponse) GetQuoteOk() (*QuoteMutualFund, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *MutualFundResponse) SetQuote(v QuoteMutualFund)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *MutualFundResponse) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReference

`func (o *MutualFundResponse) GetReference() ReferenceMutualFund`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *MutualFundResponse) GetReferenceOk() (*ReferenceMutualFund, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *MutualFundResponse) SetReference(v ReferenceMutualFund)`

SetReference sets Reference field to given value.

### HasReference

`func (o *MutualFundResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


