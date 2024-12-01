# QuoteResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetMainType** | Pointer to [**AssetMainType**](AssetMainType.md) |  | [optional] 
**AssetSubType** | Pointer to [**NullableMutualFundAssetSubType**](MutualFundAssetSubType.md) |  | [optional] 
**Ssid** | Pointer to **int64** | SSID of instrument | [optional] 
**Symbol** | Pointer to **string** | Symbol of instrument | [optional] 
**Realtime** | Pointer to **bool** | is quote realtime | [optional] 
**QuoteType** | Pointer to [**NullableQuoteType**](QuoteType.md) |  | [optional] 
**Extended** | Pointer to [**ExtendedMarket**](ExtendedMarket.md) |  | [optional] 
**Fundamental** | Pointer to [**Fundamental**](Fundamental.md) |  | [optional] 
**Quote** | Pointer to [**QuoteMutualFund**](QuoteMutualFund.md) |  | [optional] 
**Reference** | Pointer to [**ReferenceMutualFund**](ReferenceMutualFund.md) |  | [optional] 
**Regular** | Pointer to [**RegularMarket**](RegularMarket.md) |  | [optional] 
**InvalidCusips** | Pointer to **[]string** | list of invalid cusips from request | [optional] 
**InvalidSSIDs** | Pointer to **[]int64** | list of invalid SSIDs from request | [optional] 
**InvalidSymbols** | Pointer to **[]string** | list of invalid symbols from request | [optional] 

## Methods

### NewQuoteResponseObject

`func NewQuoteResponseObject() *QuoteResponseObject`

NewQuoteResponseObject instantiates a new QuoteResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseObjectWithDefaults

`func NewQuoteResponseObjectWithDefaults() *QuoteResponseObject`

NewQuoteResponseObjectWithDefaults instantiates a new QuoteResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetMainType

`func (o *QuoteResponseObject) GetAssetMainType() AssetMainType`

GetAssetMainType returns the AssetMainType field if non-nil, zero value otherwise.

### GetAssetMainTypeOk

`func (o *QuoteResponseObject) GetAssetMainTypeOk() (*AssetMainType, bool)`

GetAssetMainTypeOk returns a tuple with the AssetMainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetMainType

`func (o *QuoteResponseObject) SetAssetMainType(v AssetMainType)`

SetAssetMainType sets AssetMainType field to given value.

### HasAssetMainType

`func (o *QuoteResponseObject) HasAssetMainType() bool`

HasAssetMainType returns a boolean if a field has been set.

### GetAssetSubType

`func (o *QuoteResponseObject) GetAssetSubType() MutualFundAssetSubType`

GetAssetSubType returns the AssetSubType field if non-nil, zero value otherwise.

### GetAssetSubTypeOk

`func (o *QuoteResponseObject) GetAssetSubTypeOk() (*MutualFundAssetSubType, bool)`

GetAssetSubTypeOk returns a tuple with the AssetSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSubType

`func (o *QuoteResponseObject) SetAssetSubType(v MutualFundAssetSubType)`

SetAssetSubType sets AssetSubType field to given value.

### HasAssetSubType

`func (o *QuoteResponseObject) HasAssetSubType() bool`

HasAssetSubType returns a boolean if a field has been set.

### SetAssetSubTypeNil

`func (o *QuoteResponseObject) SetAssetSubTypeNil(b bool)`

 SetAssetSubTypeNil sets the value for AssetSubType to be an explicit nil

### UnsetAssetSubType
`func (o *QuoteResponseObject) UnsetAssetSubType()`

UnsetAssetSubType ensures that no value is present for AssetSubType, not even an explicit nil
### GetSsid

`func (o *QuoteResponseObject) GetSsid() int64`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *QuoteResponseObject) GetSsidOk() (*int64, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *QuoteResponseObject) SetSsid(v int64)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *QuoteResponseObject) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSymbol

`func (o *QuoteResponseObject) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QuoteResponseObject) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QuoteResponseObject) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QuoteResponseObject) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetRealtime

`func (o *QuoteResponseObject) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *QuoteResponseObject) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *QuoteResponseObject) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *QuoteResponseObject) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetQuoteType

`func (o *QuoteResponseObject) GetQuoteType() QuoteType`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *QuoteResponseObject) GetQuoteTypeOk() (*QuoteType, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *QuoteResponseObject) SetQuoteType(v QuoteType)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *QuoteResponseObject) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### SetQuoteTypeNil

`func (o *QuoteResponseObject) SetQuoteTypeNil(b bool)`

 SetQuoteTypeNil sets the value for QuoteType to be an explicit nil

### UnsetQuoteType
`func (o *QuoteResponseObject) UnsetQuoteType()`

UnsetQuoteType ensures that no value is present for QuoteType, not even an explicit nil
### GetExtended

`func (o *QuoteResponseObject) GetExtended() ExtendedMarket`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *QuoteResponseObject) GetExtendedOk() (*ExtendedMarket, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *QuoteResponseObject) SetExtended(v ExtendedMarket)`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *QuoteResponseObject) HasExtended() bool`

HasExtended returns a boolean if a field has been set.

### GetFundamental

`func (o *QuoteResponseObject) GetFundamental() Fundamental`

GetFundamental returns the Fundamental field if non-nil, zero value otherwise.

### GetFundamentalOk

`func (o *QuoteResponseObject) GetFundamentalOk() (*Fundamental, bool)`

GetFundamentalOk returns a tuple with the Fundamental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundamental

`func (o *QuoteResponseObject) SetFundamental(v Fundamental)`

SetFundamental sets Fundamental field to given value.

### HasFundamental

`func (o *QuoteResponseObject) HasFundamental() bool`

HasFundamental returns a boolean if a field has been set.

### GetQuote

`func (o *QuoteResponseObject) GetQuote() QuoteMutualFund`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *QuoteResponseObject) GetQuoteOk() (*QuoteMutualFund, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *QuoteResponseObject) SetQuote(v QuoteMutualFund)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *QuoteResponseObject) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetReference

`func (o *QuoteResponseObject) GetReference() ReferenceMutualFund`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *QuoteResponseObject) GetReferenceOk() (*ReferenceMutualFund, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *QuoteResponseObject) SetReference(v ReferenceMutualFund)`

SetReference sets Reference field to given value.

### HasReference

`func (o *QuoteResponseObject) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetRegular

`func (o *QuoteResponseObject) GetRegular() RegularMarket`

GetRegular returns the Regular field if non-nil, zero value otherwise.

### GetRegularOk

`func (o *QuoteResponseObject) GetRegularOk() (*RegularMarket, bool)`

GetRegularOk returns a tuple with the Regular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegular

`func (o *QuoteResponseObject) SetRegular(v RegularMarket)`

SetRegular sets Regular field to given value.

### HasRegular

`func (o *QuoteResponseObject) HasRegular() bool`

HasRegular returns a boolean if a field has been set.

### GetInvalidCusips

`func (o *QuoteResponseObject) GetInvalidCusips() []string`

GetInvalidCusips returns the InvalidCusips field if non-nil, zero value otherwise.

### GetInvalidCusipsOk

`func (o *QuoteResponseObject) GetInvalidCusipsOk() (*[]string, bool)`

GetInvalidCusipsOk returns a tuple with the InvalidCusips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCusips

`func (o *QuoteResponseObject) SetInvalidCusips(v []string)`

SetInvalidCusips sets InvalidCusips field to given value.

### HasInvalidCusips

`func (o *QuoteResponseObject) HasInvalidCusips() bool`

HasInvalidCusips returns a boolean if a field has been set.

### GetInvalidSSIDs

`func (o *QuoteResponseObject) GetInvalidSSIDs() []int64`

GetInvalidSSIDs returns the InvalidSSIDs field if non-nil, zero value otherwise.

### GetInvalidSSIDsOk

`func (o *QuoteResponseObject) GetInvalidSSIDsOk() (*[]int64, bool)`

GetInvalidSSIDsOk returns a tuple with the InvalidSSIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSSIDs

`func (o *QuoteResponseObject) SetInvalidSSIDs(v []int64)`

SetInvalidSSIDs sets InvalidSSIDs field to given value.

### HasInvalidSSIDs

`func (o *QuoteResponseObject) HasInvalidSSIDs() bool`

HasInvalidSSIDs returns a boolean if a field has been set.

### GetInvalidSymbols

`func (o *QuoteResponseObject) GetInvalidSymbols() []string`

GetInvalidSymbols returns the InvalidSymbols field if non-nil, zero value otherwise.

### GetInvalidSymbolsOk

`func (o *QuoteResponseObject) GetInvalidSymbolsOk() (*[]string, bool)`

GetInvalidSymbolsOk returns a tuple with the InvalidSymbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidSymbols

`func (o *QuoteResponseObject) SetInvalidSymbols(v []string)`

SetInvalidSymbols sets InvalidSymbols field to given value.

### HasInvalidSymbols

`func (o *QuoteResponseObject) HasInvalidSymbols() bool`

HasInvalidSymbols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


