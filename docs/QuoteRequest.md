# QuoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cusips** | Pointer to **[]string** | List of cusip, max of 500 of symbols+cusip+ssids | [optional] 
**Fields** | Pointer to **string** | comma separated list of nodes in each quote&lt;br/&gt; possible values are quote,fundamental,reference,extended,regular. Dont send this attribute for full response. | [optional] 
**Ssids** | Pointer to **[]int64** | List of Schwab securityid[SSID], max of 500 of symbols+cusip+ssids | [optional] 
**Symbols** | Pointer to **[]string** | List of symbols, max of 500 of symbols+cusip+ssids | [optional] 
**Realtime** | Pointer to **bool** | Get realtime quotes and skip entitlement check | [optional] 
**Indicative** | Pointer to **bool** | Include indicative symbol quotes for all ETF symbols in request. If ETF symbol ABC is in request and indicative&#x3D;true API will return quotes for ABC and its corresponding indicative quote for $ABC.IV | [optional] 

## Methods

### NewQuoteRequest

`func NewQuoteRequest() *QuoteRequest`

NewQuoteRequest instantiates a new QuoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteRequestWithDefaults

`func NewQuoteRequestWithDefaults() *QuoteRequest`

NewQuoteRequestWithDefaults instantiates a new QuoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCusips

`func (o *QuoteRequest) GetCusips() []string`

GetCusips returns the Cusips field if non-nil, zero value otherwise.

### GetCusipsOk

`func (o *QuoteRequest) GetCusipsOk() (*[]string, bool)`

GetCusipsOk returns a tuple with the Cusips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusips

`func (o *QuoteRequest) SetCusips(v []string)`

SetCusips sets Cusips field to given value.

### HasCusips

`func (o *QuoteRequest) HasCusips() bool`

HasCusips returns a boolean if a field has been set.

### GetFields

`func (o *QuoteRequest) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *QuoteRequest) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *QuoteRequest) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *QuoteRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetSsids

`func (o *QuoteRequest) GetSsids() []int64`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *QuoteRequest) GetSsidsOk() (*[]int64, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *QuoteRequest) SetSsids(v []int64)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *QuoteRequest) HasSsids() bool`

HasSsids returns a boolean if a field has been set.

### GetSymbols

`func (o *QuoteRequest) GetSymbols() []string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *QuoteRequest) GetSymbolsOk() (*[]string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *QuoteRequest) SetSymbols(v []string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *QuoteRequest) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetRealtime

`func (o *QuoteRequest) GetRealtime() bool`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *QuoteRequest) GetRealtimeOk() (*bool, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *QuoteRequest) SetRealtime(v bool)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *QuoteRequest) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetIndicative

`func (o *QuoteRequest) GetIndicative() bool`

GetIndicative returns the Indicative field if non-nil, zero value otherwise.

### GetIndicativeOk

`func (o *QuoteRequest) GetIndicativeOk() (*bool, bool)`

GetIndicativeOk returns a tuple with the Indicative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicative

`func (o *QuoteRequest) SetIndicative(v bool)`

SetIndicative sets Indicative field to given value.

### HasIndicative

`func (o *QuoteRequest) HasIndicative() bool`

HasIndicative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


