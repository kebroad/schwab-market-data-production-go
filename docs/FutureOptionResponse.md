# FutureOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMainType** | Pointer to [**AssetMainType**](AssetMainType.md) |  | [optional] 
**Ssid** | Pointer to **int64** | SSID of instrument | [optional] 
**Symbol** | Pointer to **string** | Symbol of instrument | [optional] 
**Realtime** | Pointer to **bool** | is quote realtime | [optional] 
**Quote** | Pointer to [**QuoteFutureOption**](QuoteFutureOption.md) |  | [optional] 
**Reference** | Pointer to [**ReferenceFutureOption**](ReferenceFutureOption.md) |  | [optional] 

## Methods

### NewFutureOptionResponse

`func NewFutureOptionResponse() *FutureOptionResponse`

NewFutureOptionResponse instantiates a new FutureOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFutureOptionResponseWithDefaults

`func NewFutureOptionResponseWithDefaults() *FutureOptionResponse`

NewFutureOptionResponseWithDefaults instantiates a new FutureOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMainType

`func (o *FutureOptionResponse) GetAssetMainType() AssetMainType`

GetAssetMainType returns the AssetMainType field if non-nil, zero value otherwise.

### GetAssetMainTypeOk

`func (o *FutureOptionResponse) GetAssetMainTypeOk() (*AssetMainType, bool)`

GetAssetMainTypeOk returns a tuple with the AssetMainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMainType

`func (o *FutureOptionResponse) SetAssetMainType(v AssetMainType)`

SetAssetMainType sets AssetMainType field to given value.

### HasAssetMainType

`func (o *FutureOptionResponse) HasAssetMainType() bool`

HasAssetMainType returns a boolean if a field has been set.

### GetSsid

`func (o *FutureOptionResponse) GetSsid() int64`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *FutureOptionResponse) GetSsidOk() (*int64, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *FutureOptionResponse) SetSsid(v int64)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *FutureOptionResponse) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSymbol

`func (o *FutureOptionResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FutureOptionResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FutureOptionResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FutureOptionResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRealtime

`func (o *FutureOptionResponse) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *FutureOptionResponse) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *FutureOptionResponse) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *FutureOptionResponse) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetQuote

`func (o *FutureOptionResponse) GetQuote() QuoteFutureOption`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *FutureOptionResponse) GetQuoteOk() (*QuoteFutureOption, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *FutureOptionResponse) SetQuote(v QuoteFutureOption)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *FutureOptionResponse) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReference

`func (o *FutureOptionResponse) GetReference() ReferenceFutureOption`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *FutureOptionResponse) GetReferenceOk() (*ReferenceFutureOption, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *FutureOptionResponse) SetReference(v ReferenceFutureOption)`

SetReference sets Reference field to given value.

### HasReference

`func (o *FutureOptionResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


