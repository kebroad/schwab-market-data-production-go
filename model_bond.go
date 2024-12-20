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

// checks if the Bond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bond{}

// Bond struct for Bond
type Bond struct {
	Cusip *string `json:"cusip,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	Description *string `json:"description,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
	AssetType *string `json:"assetType,omitempty"`
	BondFactor *string `json:"bondFactor,omitempty"`
	BondMultiplier *string `json:"bondMultiplier,omitempty"`
	BondPrice *float32 `json:"bondPrice,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewBond instantiates a new Bond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBond() *Bond {
	this := Bond{}
	return &this
}

// NewBondWithDefaults instantiates a new Bond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBondWithDefaults() *Bond {
	this := Bond{}
	return &this
}

// GetCusip returns the Cusip field value if set, zero value otherwise.
func (o *Bond) GetCusip() string {
	if o == nil || IsNil(o.Cusip) {
		var ret string
		return ret
	}
	return *o.Cusip
}

// GetCusipOk returns a tuple with the Cusip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetCusipOk() (*string, bool) {
	if o == nil || IsNil(o.Cusip) {
		return nil, false
	}
	return o.Cusip, true
}

// HasCusip returns a boolean if a field has been set.
func (o *Bond) HasCusip() bool {
	if o != nil && !IsNil(o.Cusip) {
		return true
	}

	return false
}

// SetCusip gets a reference to the given string and assigns it to the Cusip field.
func (o *Bond) SetCusip(v string) {
	o.Cusip = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Bond) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Bond) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Bond) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Bond) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Bond) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Bond) SetDescription(v string) {
	o.Description = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Bond) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Bond) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Bond) SetExchange(v string) {
	o.Exchange = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *Bond) GetAssetType() string {
	if o == nil || IsNil(o.AssetType) {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetAssetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *Bond) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *Bond) SetAssetType(v string) {
	o.AssetType = &v
}

// GetBondFactor returns the BondFactor field value if set, zero value otherwise.
func (o *Bond) GetBondFactor() string {
	if o == nil || IsNil(o.BondFactor) {
		var ret string
		return ret
	}
	return *o.BondFactor
}

// GetBondFactorOk returns a tuple with the BondFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetBondFactorOk() (*string, bool) {
	if o == nil || IsNil(o.BondFactor) {
		return nil, false
	}
	return o.BondFactor, true
}

// HasBondFactor returns a boolean if a field has been set.
func (o *Bond) HasBondFactor() bool {
	if o != nil && !IsNil(o.BondFactor) {
		return true
	}

	return false
}

// SetBondFactor gets a reference to the given string and assigns it to the BondFactor field.
func (o *Bond) SetBondFactor(v string) {
	o.BondFactor = &v
}

// GetBondMultiplier returns the BondMultiplier field value if set, zero value otherwise.
func (o *Bond) GetBondMultiplier() string {
	if o == nil || IsNil(o.BondMultiplier) {
		var ret string
		return ret
	}
	return *o.BondMultiplier
}

// GetBondMultiplierOk returns a tuple with the BondMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetBondMultiplierOk() (*string, bool) {
	if o == nil || IsNil(o.BondMultiplier) {
		return nil, false
	}
	return o.BondMultiplier, true
}

// HasBondMultiplier returns a boolean if a field has been set.
func (o *Bond) HasBondMultiplier() bool {
	if o != nil && !IsNil(o.BondMultiplier) {
		return true
	}

	return false
}

// SetBondMultiplier gets a reference to the given string and assigns it to the BondMultiplier field.
func (o *Bond) SetBondMultiplier(v string) {
	o.BondMultiplier = &v
}

// GetBondPrice returns the BondPrice field value if set, zero value otherwise.
func (o *Bond) GetBondPrice() float32 {
	if o == nil || IsNil(o.BondPrice) {
		var ret float32
		return ret
	}
	return *o.BondPrice
}

// GetBondPriceOk returns a tuple with the BondPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetBondPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.BondPrice) {
		return nil, false
	}
	return o.BondPrice, true
}

// HasBondPrice returns a boolean if a field has been set.
func (o *Bond) HasBondPrice() bool {
	if o != nil && !IsNil(o.BondPrice) {
		return true
	}

	return false
}

// SetBondPrice gets a reference to the given float32 and assigns it to the BondPrice field.
func (o *Bond) SetBondPrice(v float32) {
	o.BondPrice = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Bond) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bond) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Bond) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Bond) SetType(v string) {
	o.Type = &v
}

func (o Bond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Bond) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableBond struct {
	value *Bond
	isSet bool
}

func (v NullableBond) Get() *Bond {
	return v.value
}

func (v *NullableBond) Set(val *Bond) {
	v.value = val
	v.isSet = true
}

func (v NullableBond) IsSet() bool {
	return v.isSet
}

func (v *NullableBond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBond(val *Bond) *NullableBond {
	return &NullableBond{value: val, isSet: true}
}

func (v NullableBond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


