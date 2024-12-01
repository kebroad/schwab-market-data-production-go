# EquityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMainType** | Pointer to [**AssetMainType**](AssetMainType.md) |  | [optional] 
**AssetSubType** | Pointer to [**NullableEquityAssetSubType**](EquityAssetSubType.md) |  | [optional] 
**Ssid** | Pointer to **int64** | SSID of instrument | [optional] 
**Symbol** | Pointer to **string** | Symbol of instrument | [optional] 
**Realtime** | Pointer to **bool** | is quote realtime | [optional] 
**QuoteType** | Pointer to [**NullableQuoteType**](QuoteType.md) |  | [optional] 
**Extended** | Pointer to [**ExtendedMarket**](ExtendedMarket.md) |  | [optional] 
**Fundamental** | Pointer to [**Fundamental**](Fundamental.md) |  | [optional] 
**Quote** | Pointer to [**QuoteEquity**](QuoteEquity.md) |  | [optional] 
**Reference** | Pointer to [**ReferenceEquity**](ReferenceEquity.md) |  | [optional] 
**Regular** | Pointer to [**RegularMarket**](RegularMarket.md) |  | [optional] 

## Methods

### NewEquityResponse

`func NewEquityResponse() *EquityResponse`

NewEquityResponse instantiates a new EquityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquityResponseWithDefaults

`func NewEquityResponseWithDefaults() *EquityResponse`

NewEquityResponseWithDefaults instantiates a new EquityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMainType

`func (o *EquityResponse) GetAssetMainType() AssetMainType`

GetAssetMainType returns the AssetMainType field if non-nil, zero value otherwise.

### GetAssetMainTypeOk

`func (o *EquityResponse) GetAssetMainTypeOk() (*AssetMainType, bool)`

GetAssetMainTypeOk returns a tuple with the AssetMainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMainType

`func (o *EquityResponse) SetAssetMainType(v AssetMainType)`

SetAssetMainType sets AssetMainType field to given value.

### HasAssetMainType

`func (o *EquityResponse) HasAssetMainType() bool`

HasAssetMainType returns a boolean if a field has been set.

### GetAssetSubType

`func (o *EquityResponse) GetAssetSubType() EquityAssetSubType`

GetAssetSubType returns the AssetSubType field if non-nil, zero value otherwise.

### GetAssetSubTypeOk

`func (o *EquityResponse) GetAssetSubTypeOk() (*EquityAssetSubType, bool)`

GetAssetSubTypeOk returns a tuple with the AssetSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSubType

`func (o *EquityResponse) SetAssetSubType(v EquityAssetSubType)`

SetAssetSubType sets AssetSubType field to given value.

### HasAssetSubType

`func (o *EquityResponse) HasAssetSubType() bool`

HasAssetSubType returns a boolean if a field has been set.

### SetAssetSubTypeNil

`func (o *EquityResponse) SetAssetSubTypeNil(b bool)`

 SetAssetSubTypeNil sets the value for AssetSubType to be an explicit nil

### UnsetAssetSubType
`func (o *EquityResponse) UnsetAssetSubType()`

UnsetAssetSubType ensures that no value is present for AssetSubType, not even an explicit nil
### GetSsid

`func (o *EquityResponse) GetSsid() int64`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *EquityResponse) GetSsidOk() (*int64, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *EquityResponse) SetSsid(v int64)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *EquityResponse) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSymbol

`func (o *EquityResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *EquityResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *EquityResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *EquityResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRealtime

`func (o *EquityResponse) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *EquityResponse) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *EquityResponse) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *EquityResponse) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetQuoteType

`func (o *EquityResponse) GetQuoteType() QuoteType`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *EquityResponse) GetQuoteTypeOk() (*QuoteType, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *EquityResponse) SetQuoteType(v QuoteType)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *EquityResponse) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### SetQuoteTypeNil

`func (o *EquityResponse) SetQuoteTypeNil(b bool)`

 SetQuoteTypeNil sets the value for QuoteType to be an explicit nil

### UnsetQuoteType
`func (o *EquityResponse) UnsetQuoteType()`

UnsetQuoteType ensures that no value is present for QuoteType, not even an explicit nil
### GetExtended

`func (o *EquityResponse) GetExtended() ExtendedMarket`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *EquityResponse) GetExtendedOk() (*ExtendedMarket, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *EquityResponse) SetExtended(v ExtendedMarket)`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *EquityResponse) HasExtended() bool`

HasExtended returns a boolean if a field has been set.

### GetFundamental

`func (o *EquityResponse) GetFundamental() Fundamental`

GetFundamental returns the Fundamental field if non-nil, zero value otherwise.

### GetFundamentalOk

`func (o *EquityResponse) GetFundamentalOk() (*Fundamental, bool)`

GetFundamentalOk returns a tuple with the Fundamental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundamental

`func (o *EquityResponse) SetFundamental(v Fundamental)`

SetFundamental sets Fundamental field to given value.

### HasFundamental

`func (o *EquityResponse) HasFundamental() bool`

HasFundamental returns a boolean if a field has been set.

### GetQuote

`func (o *EquityResponse) GetQuote() QuoteEquity`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *EquityResponse) GetQuoteOk() (*QuoteEquity, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *EquityResponse) SetQuote(v QuoteEquity)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *EquityResponse) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReference

`func (o *EquityResponse) GetReference() ReferenceEquity`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *EquityResponse) GetReferenceOk() (*ReferenceEquity, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *EquityResponse) SetReference(v ReferenceEquity)`

SetReference sets Reference field to given value.

### HasReference

`func (o *EquityResponse) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetRegular

`func (o *EquityResponse) GetRegular() RegularMarket`

GetRegular returns the Regular field if non-nil, zero value otherwise.

### GetRegularOk

`func (o *EquityResponse) GetRegularOk() (*RegularMarket, bool)`

GetRegularOk returns a tuple with the Regular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegular

`func (o *EquityResponse) SetRegular(v RegularMarket)`

SetRegular sets Regular field to given value.

### HasRegular

`func (o *EquityResponse) HasRegular() bool`

HasRegular returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


