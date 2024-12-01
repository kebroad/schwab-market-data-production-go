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

// checks if the ReferenceForex type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferenceForex{}

// ReferenceForex Reference data of Forex security
type ReferenceForex struct {
	// Description of Instrument
	Description *string `json:"description,omitempty"`
	// Exchange Code
	Exchange *string `json:"exchange,omitempty"`
	// Exchange Name
	ExchangeName *string `json:"exchangeName,omitempty"`
	// is FOREX tradable
	IsTradable *bool `json:"isTradable,omitempty"`
	// Market marker
	MarketMaker *string `json:"marketMaker,omitempty"`
	// Product name
	Product *string `json:"product,omitempty"`
	// Trading hours
	TradingHours *string `json:"tradingHours,omitempty"`
}

// NewReferenceForex instantiates a new ReferenceForex object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceForex() *ReferenceForex {
	this := ReferenceForex{}
	return &this
}

// NewReferenceForexWithDefaults instantiates a new ReferenceForex object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceForexWithDefaults() *ReferenceForex {
	this := ReferenceForex{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ReferenceForex) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ReferenceForex) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ReferenceForex) SetDescription(v string) {
	o.Description = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *ReferenceForex) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *ReferenceForex) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *ReferenceForex) SetExchange(v string) {
	o.Exchange = &v
}

// GetExchangeName returns the ExchangeName field value if set, zero value otherwise.
func (o *ReferenceForex) GetExchangeName() string {
	if o == nil || IsNil(o.ExchangeName) {
		var ret string
		return ret
	}
	return *o.ExchangeName
}

// GetExchangeNameOk returns a tuple with the ExchangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetExchangeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExchangeName) {
		return nil, false
	}
	return o.ExchangeName, true
}

// HasExchangeName returns a boolean if a field has been set.
func (o *ReferenceForex) HasExchangeName() bool {
	if o != nil && !IsNil(o.ExchangeName) {
		return true
	}

	return false
}

// SetExchangeName gets a reference to the given string and assigns it to the ExchangeName field.
func (o *ReferenceForex) SetExchangeName(v string) {
	o.ExchangeName = &v
}

// GetIsTradable returns the IsTradable field value if set, zero value otherwise.
func (o *ReferenceForex) GetIsTradable() bool {
	if o == nil || IsNil(o.IsTradable) {
		var ret bool
		return ret
	}
	return *o.IsTradable
}

// GetIsTradableOk returns a tuple with the IsTradable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetIsTradableOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTradable) {
		return nil, false
	}
	return o.IsTradable, true
}

// HasIsTradable returns a boolean if a field has been set.
func (o *ReferenceForex) HasIsTradable() bool {
	if o != nil && !IsNil(o.IsTradable) {
		return true
	}

	return false
}

// SetIsTradable gets a reference to the given bool and assigns it to the IsTradable field.
func (o *ReferenceForex) SetIsTradable(v bool) {
	o.IsTradable = &v
}

// GetMarketMaker returns the MarketMaker field value if set, zero value otherwise.
func (o *ReferenceForex) GetMarketMaker() string {
	if o == nil || IsNil(o.MarketMaker) {
		var ret string
		return ret
	}
	return *o.MarketMaker
}

// GetMarketMakerOk returns a tuple with the MarketMaker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetMarketMakerOk() (*string, bool) {
	if o == nil || IsNil(o.MarketMaker) {
		return nil, false
	}
	return o.MarketMaker, true
}

// HasMarketMaker returns a boolean if a field has been set.
func (o *ReferenceForex) HasMarketMaker() bool {
	if o != nil && !IsNil(o.MarketMaker) {
		return true
	}

	return false
}

// SetMarketMaker gets a reference to the given string and assigns it to the MarketMaker field.
func (o *ReferenceForex) SetMarketMaker(v string) {
	o.MarketMaker = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *ReferenceForex) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *ReferenceForex) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *ReferenceForex) SetProduct(v string) {
	o.Product = &v
}

// GetTradingHours returns the TradingHours field value if set, zero value otherwise.
func (o *ReferenceForex) GetTradingHours() string {
	if o == nil || IsNil(o.TradingHours) {
		var ret string
		return ret
	}
	return *o.TradingHours
}

// GetTradingHoursOk returns a tuple with the TradingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceForex) GetTradingHoursOk() (*string, bool) {
	if o == nil || IsNil(o.TradingHours) {
		return nil, false
	}
	return o.TradingHours, true
}

// HasTradingHours returns a boolean if a field has been set.
func (o *ReferenceForex) HasTradingHours() bool {
	if o != nil && !IsNil(o.TradingHours) {
		return true
	}

	return false
}

// SetTradingHours gets a reference to the given string and assigns it to the TradingHours field.
func (o *ReferenceForex) SetTradingHours(v string) {
	o.TradingHours = &v
}

func (o ReferenceForex) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferenceForex) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.ExchangeName) {
		toSerialize["exchangeName"] = o.ExchangeName
	}
	if !IsNil(o.IsTradable) {
		toSerialize["isTradable"] = o.IsTradable
	}
	if !IsNil(o.MarketMaker) {
		toSerialize["marketMaker"] = o.MarketMaker
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.TradingHours) {
		toSerialize["tradingHours"] = o.TradingHours
	}
	return toSerialize, nil
}

type NullableReferenceForex struct {
	value *ReferenceForex
	isSet bool
}

func (v NullableReferenceForex) Get() *ReferenceForex {
	return v.value
}

func (v *NullableReferenceForex) Set(val *ReferenceForex) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceForex) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceForex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceForex(val *ReferenceForex) *NullableReferenceForex {
	return &NullableReferenceForex{value: val, isSet: true}
}

func (v NullableReferenceForex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceForex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


