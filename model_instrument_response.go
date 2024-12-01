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

// checks if the InstrumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstrumentResponse{}

// InstrumentResponse struct for InstrumentResponse
type InstrumentResponse struct {
	Cusip *string `json:"cusip,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
	AssetType *string `json:"assetType,omitempty"`
	BondFactor *string `json:"bondFactor,omitempty"`
	BondMultiplier *string `json:"bondMultiplier,omitempty"`
	BondPrice *float32 `json:"bondPrice,omitempty"`
	Fundamental *FundamentalInst `json:"fundamental,omitempty"`
	InstrumentInfo *Instrument `json:"instrumentInfo,omitempty"`
	BondInstrumentInfo *Bond `json:"bondInstrumentInfo,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewInstrumentResponse instantiates a new InstrumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstrumentResponse() *InstrumentResponse {
	this := InstrumentResponse{}
	return &this
}

// NewInstrumentResponseWithDefaults instantiates a new InstrumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstrumentResponseWithDefaults() *InstrumentResponse {
	this := InstrumentResponse{}
	return &this
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *InstrumentResponse) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *InstrumentResponse) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *InstrumentResponse) SetCusip(v string) {
	o.Cusip = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *InstrumentResponse) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *InstrumentResponse) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *InstrumentResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstrumentResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstrumentResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstrumentResponse) SetDescription(v string) {
	o.Description = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *InstrumentResponse) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *InstrumentResponse) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *InstrumentResponse) SetExchange(v string) {
	o.Exchange = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *InstrumentResponse) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *InstrumentResponse) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *InstrumentResponse) SetAssetType(v string) {
	o.AssetType = &v
}

// GetBondFactor returns the BondFactor field value if set, zero value otherwise.
func (o *InstrumentResponse) GetBondFactor() string {
	if o == nil || IsNil(o.BondFactor) {
		var ret string
		return ret
	}
	return *o.BondFactor
}

// GetBondFactorOk returns a tuple with the BondFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetBondFactorOk() (*string, bool) {
	if o == nil || IsNil(o.BondFactor) {
		return nil, false
	}
	return o.BondFactor, true
}

// HasBondFactor returns a boolean if a field has been set.
func (o *InstrumentResponse) HasBondFactor() bool {
	if o != nil && !IsNil(o.BondFactor) {
		return true
	}

	return false
}

// SetBondFactor gets a reference to the given string and assigns it to the BondFactor field.
func (o *InstrumentResponse) SetBondFactor(v string) {
	o.BondFactor = &v
}

// GetBondMultiplier returns the BondMultiplier field value if set, zero value otherwise.
func (o *InstrumentResponse) GetBondMultiplier() string {
	if o == nil || IsNil(o.BondMultiplier) {
		var ret string
		return ret
	}
	return *o.BondMultiplier
}

// GetBondMultiplierOk returns a tuple with the BondMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetBondMultiplierOk() (*string, bool) {
	if o == nil || IsNil(o.BondMultiplier) {
		return nil, false
	}
	return o.BondMultiplier, true
}

// HasBondMultiplier returns a boolean if a field has been set.
func (o *InstrumentResponse) HasBondMultiplier() bool {
	if o != nil && !IsNil(o.BondMultiplier) {
		return true
	}

	return false
}

// SetBondMultiplier gets a reference to the given string and assigns it to the BondMultiplier field.
func (o *InstrumentResponse) SetBondMultiplier(v string) {
	o.BondMultiplier = &v
}

// GetBondPrice returns the BondPrice field value if set, zero value otherwise.
func (o *InstrumentResponse) GetBondPrice() float32 {
	if o == nil || IsNil(o.BondPrice) {
		var ret float32
		return ret
	}
	return *o.BondPrice
}

// GetBondPriceOk returns a tuple with the BondPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetBondPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.BondPrice) {
		return nil, false
	}
	return o.BondPrice, true
}

// HasBondPrice returns a boolean if a field has been set.
func (o *InstrumentResponse) HasBondPrice() bool {
	if o != nil && !IsNil(o.BondPrice) {
		return true
	}

	return false
}

// SetBondPrice gets a reference to the given float32 and assigns it to the BondPrice field.
func (o *InstrumentResponse) SetBondPrice(v float32) {
	o.BondPrice = &v
}

// GetFundamental returns the Fundamental field value if set, zero value otherwise.
func (o *InstrumentResponse) GetFundamental() FundamentalInst {
	if o == nil || IsNil(o.Fundamental) {
		var ret FundamentalInst
		return ret
	}
	return *o.Fundamental
}

// GetFundamentalOk returns a tuple with the Fundamental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetFundamentalOk() (*FundamentalInst, bool) {
	if o == nil || IsNil(o.Fundamental) {
		return nil, false
	}
	return o.Fundamental, true
}

// HasFundamental returns a boolean if a field has been set.
func (o *InstrumentResponse) HasFundamental() bool {
	if o != nil && !IsNil(o.Fundamental) {
		return true
	}

	return false
}

// SetFundamental gets a reference to the given FundamentalInst and assigns it to the Fundamental field.
func (o *InstrumentResponse) SetFundamental(v FundamentalInst) {
	o.Fundamental = &v
}

// GetInstrumentInfo returns the InstrumentInfo field value if set, zero value otherwise.
func (o *InstrumentResponse) GetInstrumentInfo() Instrument {
	if o == nil || IsNil(o.InstrumentInfo) {
		var ret Instrument
		return ret
	}
	return *o.InstrumentInfo
}

// GetInstrumentInfoOk returns a tuple with the InstrumentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetInstrumentInfoOk() (*Instrument, bool) {
	if o == nil || IsNil(o.InstrumentInfo) {
		return nil, false
	}
	return o.InstrumentInfo, true
}

// HasInstrumentInfo returns a boolean if a field has been set.
func (o *InstrumentResponse) HasInstrumentInfo() bool {
	if o != nil && !IsNil(o.InstrumentInfo) {
		return true
	}

	return false
}

// SetInstrumentInfo gets a reference to the given Instrument and assigns it to the InstrumentInfo field.
func (o *InstrumentResponse) SetInstrumentInfo(v Instrument) {
	o.InstrumentInfo = &v
}

// GetBondInstrumentInfo returns the BondInstrumentInfo field value if set, zero value otherwise.
func (o *InstrumentResponse) GetBondInstrumentInfo() Bond {
	if o == nil || IsNil(o.BondInstrumentInfo) {
		var ret Bond
		return ret
	}
	return *o.BondInstrumentInfo
}

// GetBondInstrumentInfoOk returns a tuple with the BondInstrumentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetBondInstrumentInfoOk() (*Bond, bool) {
	if o == nil || IsNil(o.BondInstrumentInfo) {
		return nil, false
	}
	return o.BondInstrumentInfo, true
}

// HasBondInstrumentInfo returns a boolean if a field has been set.
func (o *InstrumentResponse) HasBondInstrumentInfo() bool {
	if o != nil && !IsNil(o.BondInstrumentInfo) {
		return true
	}

	return false
}

// SetBondInstrumentInfo gets a reference to the given Bond and assigns it to the BondInstrumentInfo field.
func (o *InstrumentResponse) SetBondInstrumentInfo(v Bond) {
	o.BondInstrumentInfo = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InstrumentResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstrumentResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InstrumentResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InstrumentResponse) SetType(v string) {
	o.Type = &v
}

func (o InstrumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstrumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cusip) {
		toSerialize["cusip"] = o.Cusip
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	if !IsNil(o.BondFactor) {
		toSerialize["bondFactor"] = o.BondFactor
	}
	if !IsNil(o.BondMultiplier) {
		toSerialize["bondMultiplier"] = o.BondMultiplier
	}
	if !IsNil(o.BondPrice) {
		toSerialize["bondPrice"] = o.BondPrice
	}
	if !IsNil(o.Fundamental) {
		toSerialize["fundamental"] = o.Fundamental
	}
	if !IsNil(o.InstrumentInfo) {
		toSerialize["instrumentInfo"] = o.InstrumentInfo
	}
	if !IsNil(o.BondInstrumentInfo) {
		toSerialize["bondInstrumentInfo"] = o.BondInstrumentInfo
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableInstrumentResponse struct {
	value *InstrumentResponse
	isSet bool
}

func (v NullableInstrumentResponse) Get() *InstrumentResponse {
	return v.value
}

func (v *NullableInstrumentResponse) Set(val *InstrumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInstrumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInstrumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstrumentResponse(val *InstrumentResponse) *NullableInstrumentResponse {
	return &NullableInstrumentResponse{value: val, isSet: true}
}

func (v NullableInstrumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstrumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

