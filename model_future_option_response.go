/*
Market Data

Trader API - Market data

API version: 1.0.0
Contact: TraderAPI@Schwab.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the FutureOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FutureOptionResponse{}

// FutureOptionResponse Quote info of Future Option security
type FutureOptionResponse struct {
	AssetMainType *AssetMainType `json:"assetMainType,omitempty"`
	// SSID of instrument
	Ssid *int64 `json:"ssid,omitempty"`
	// Symbol of instrument
	Symbol *string `json:"symbol,omitempty"`
	// is quote realtime
	Realtime *bool `json:"realtime,omitempty"`
	Quote *QuoteFutureOption `json:"quote,omitempty"`
	Reference *ReferenceFutureOption `json:"reference,omitempty"`
}

// NewFutureOptionResponse instantiates a new FutureOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFutureOptionResponse() *FutureOptionResponse {
	this := FutureOptionResponse{}
	return &this
}

// NewFutureOptionResponseWithDefaults instantiates a new FutureOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFutureOptionResponseWithDefaults() *FutureOptionResponse {
	this := FutureOptionResponse{}
	return &this
}

// GetAssetMainType returns the AssetMainType field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetAssetMainType() AssetMainType {
	if o == nil || IsNil(o.AssetMainType) {
		var ret AssetMainType
		return ret
	}
	return *o.AssetMainType
}

// GetAssetMainTypeOk returns a tuple with the AssetMainType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetAssetMainTypeOk() (*AssetMainType, bool) {
	if o == nil || IsNil(o.AssetMainType) {
		return nil, false
	}
	return o.AssetMainType, true
}

// HasAssetMainType returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasAssetMainType() bool {
	if o != nil && !IsNil(o.AssetMainType) {
		return true
	}

	return false
}

// SetAssetMainType gets a reference to the given AssetMainType and assigns it to the AssetMainType field.
func (o *FutureOptionResponse) SetAssetMainType(v AssetMainType) {
	o.AssetMainType = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetSsid() int64 {
	if o == nil || IsNil(o.Ssid) {
		var ret int64
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetSsidOk() (*int64, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given int64 and assigns it to the Ssid field.
func (o *FutureOptionResponse) SetSsid(v int64) {
	o.Ssid = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FutureOptionResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRealtime returns the Realtime field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetRealtime() bool {
	if o == nil || IsNil(o.Realtime) {
		var ret bool
		return ret
	}
	return *o.Realtime
}

// GetRealtimeOk returns a tuple with the Realtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetRealtimeOk() (*bool, bool) {
	if o == nil || IsNil(o.Realtime) {
		return nil, false
	}
	return o.Realtime, true
}

// HasRealtime returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasRealtime() bool {
	if o != nil && !IsNil(o.Realtime) {
		return true
	}

	return false
}

// SetRealtime gets a reference to the given bool and assigns it to the Realtime field.
func (o *FutureOptionResponse) SetRealtime(v bool) {
	o.Realtime = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetQuote() QuoteFutureOption {
	if o == nil || IsNil(o.Quote) {
		var ret QuoteFutureOption
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetQuoteOk() (*QuoteFutureOption, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given QuoteFutureOption and assigns it to the Quote field.
func (o *FutureOptionResponse) SetQuote(v QuoteFutureOption) {
	o.Quote = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *FutureOptionResponse) GetReference() ReferenceFutureOption {
	if o == nil || IsNil(o.Reference) {
		var ret ReferenceFutureOption
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureOptionResponse) GetReferenceOk() (*ReferenceFutureOption, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *FutureOptionResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given ReferenceFutureOption and assigns it to the Reference field.
func (o *FutureOptionResponse) SetReference(v ReferenceFutureOption) {
	o.Reference = &v
}

func (o FutureOptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FutureOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetMainType) {
		toSerialize["assetMainType"] = o.AssetMainType
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Realtime) {
		toSerialize["realtime"] = o.Realtime
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	return toSerialize, nil
}

type NullableFutureOptionResponse struct {
	value *FutureOptionResponse
	isSet bool
}

func (v NullableFutureOptionResponse) Get() *FutureOptionResponse {
	return v.value
}

func (v *NullableFutureOptionResponse) Set(val *FutureOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFutureOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFutureOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFutureOptionResponse(val *FutureOptionResponse) *NullableFutureOptionResponse {
	return &NullableFutureOptionResponse{value: val, isSet: true}
}

func (v NullableFutureOptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFutureOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

