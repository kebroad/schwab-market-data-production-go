# OptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMainType** | Pointer to [**AssetMainType**](AssetMainType.md) |  | [optional] 
**Ssid** | Pointer to **int64** | SSID of instrument | [optional] 
**Symbol** | Pointer to **string** | Symbol of instrument | [optional] 
**Realtime** | Pointer to **bool** | is quote realtime | [optional] 
**Quote** | Pointer to [**QuoteOption**](QuoteOption.md) |  | [optional] 
**Reference** | Pointer to [**ReferenceOption**](ReferenceOption.md) |  | [optional] 

## Methods

### NewOptionResponse

`func NewOptionResponse() *OptionResponse`

NewOptionResponse instantiates a new OptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionResponseWithDefaults

`func NewOptionResponseWithDefaults() *OptionResponse`

NewOptionResponseWithDefaults instantiates a new OptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMainType

`func (o *OptionResponse) GetAssetMainType() AssetMainType`

GetAssetMainType returns the AssetMainType field if non-nil, zero value otherwise.

### GetAssetMainTypeOk

`func (o *OptionResponse) GetAssetMainTypeOk() (*AssetMainType, bool)`

GetAssetMainTypeOk returns a tuple with the AssetMainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMainType

`func (o *OptionResponse) SetAssetMainType(v AssetMainType)`

SetAssetMainType sets AssetMainType field to given value.

### HasAssetMainType

`func (o *OptionResponse) HasAssetMainType() bool`

HasAssetMainType returns a boolean if a field has been set.

### GetSsid

`func (o *OptionResponse) GetSsid() int64`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *OptionResponse) GetSsidOk() (*int64, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *OptionResponse) SetSsid(v int64)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *OptionResponse) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRealtime

`func (o *OptionResponse) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *OptionResponse) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *OptionResponse) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *OptionResponse) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetQuote

`func (o *OptionResponse) GetQuote() QuoteOption`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *OptionResponse) GetQuoteOk() (*QuoteOption, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *OptionResponse) SetQuote(v QuoteOption)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *OptionResponse) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReference

`func (o *OptionResponse) GetReference() ReferenceOption`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *OptionResponse) GetReferenceOk() (*ReferenceOption, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *OptionResponse) SetReference(v ReferenceOption)`

SetReference sets Reference field to given value.

### HasReference

`func (o *OptionResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


