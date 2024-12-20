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

// checks if the OptionChain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionChain{}

// OptionChain struct for OptionChain
type OptionChain struct {
	Symbol *string `json:"symbol,omitempty"`
	Status *string `json:"status,omitempty"`
	Underlying *Underlying `json:"underlying,omitempty"`
	Strategy *string `json:"strategy,omitempty"`
	Interval *float64 `json:"interval,omitempty"`
	IsDelayed *bool `json:"isDelayed,omitempty"`
	IsIndex *bool `json:"isIndex,omitempty"`
	DaysToExpiration *float64 `json:"daysToExpiration,omitempty"`
	InterestRate *float64 `json:"interestRate,omitempty"`
	UnderlyingPrice *float64 `json:"underlyingPrice,omitempty"`
	Volatility *float64 `json:"volatility,omitempty"`
	CallExpDateMap *map[string]map[string]OptionContract `json:"callExpDateMap,omitempty"`
	PutExpDateMap *map[string]map[string]OptionContract `json:"putExpDateMap,omitempty"`
}

// NewOptionChain instantiates a new OptionChain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionChain() *OptionChain {
	this := OptionChain{}
	return &this
}

// NewOptionChainWithDefaults instantiates a new OptionChain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionChainWithDefaults() *OptionChain {
	this := OptionChain{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OptionChain) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OptionChain) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OptionChain) SetSymbol(v string) {
	o.Symbol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OptionChain) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OptionChain) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OptionChain) SetStatus(v string) {
	o.Status = &v
}

// GetUnderlying returns the Underlying field value if set, zero value otherwise.
func (o *OptionChain) GetUnderlying() Underlying {
	if o == nil || IsNil(o.Underlying) {
		var ret Underlying
		return ret
	}
	return *o.Underlying
}

// GetUnderlyingOk returns a tuple with the Underlying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetUnderlyingOk() (*Underlying, bool) {
	if o == nil || IsNil(o.Underlying) {
		return nil, false
	}
	return o.Underlying, true
}

// HasUnderlying returns a boolean if a field has been set.
func (o *OptionChain) HasUnderlying() bool {
	if o != nil && !IsNil(o.Underlying) {
		return true
	}

	return false
}

// SetUnderlying gets a reference to the given Underlying and assigns it to the Underlying field.
func (o *OptionChain) SetUnderlying(v Underlying) {
	o.Underlying = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *OptionChain) GetStrategy() string {
	if o == nil || IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *OptionChain) HasStrategy() bool {
	if o != nil && !IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *OptionChain) SetStrategy(v string) {
	o.Strategy = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *OptionChain) GetInterval() float64 {
	if o == nil || IsNil(o.Interval) {
		var ret float64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetIntervalOk() (*float64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *OptionChain) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given float64 and assigns it to the Interval field.
func (o *OptionChain) SetInterval(v float64) {
	o.Interval = &v
}

// GetIsDelayed returns the IsDelayed field value if set, zero value otherwise.
func (o *OptionChain) GetIsDelayed() bool {
	if o == nil || IsNil(o.IsDelayed) {
		var ret bool
		return ret
	}
	return *o.IsDelayed
}

// GetIsDelayedOk returns a tuple with the IsDelayed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetIsDelayedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDelayed) {
		return nil, false
	}
	return o.IsDelayed, true
}

// HasIsDelayed returns a boolean if a field has been set.
func (o *OptionChain) HasIsDelayed() bool {
	if o != nil && !IsNil(o.IsDelayed) {
		return true
	}

	return false
}

// SetIsDelayed gets a reference to the given bool and assigns it to the IsDelayed field.
func (o *OptionChain) SetIsDelayed(v bool) {
	o.IsDelayed = &v
}

// GetIsIndex returns the IsIndex field value if set, zero value otherwise.
func (o *OptionChain) GetIsIndex() bool {
	if o == nil || IsNil(o.IsIndex) {
		var ret bool
		return ret
	}
	return *o.IsIndex
}

// GetIsIndexOk returns a tuple with the IsIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetIsIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIndex) {
		return nil, false
	}
	return o.IsIndex, true
}

// HasIsIndex returns a boolean if a field has been set.
func (o *OptionChain) HasIsIndex() bool {
	if o != nil && !IsNil(o.IsIndex) {
		return true
	}

	return false
}

// SetIsIndex gets a reference to the given bool and assigns it to the IsIndex field.
func (o *OptionChain) SetIsIndex(v bool) {
	o.IsIndex = &v
}

// GetDaysToExpiration returns the DaysToExpiration field value if set, zero value otherwise.
func (o *OptionChain) GetDaysToExpiration() float64 {
	if o == nil || IsNil(o.DaysToExpiration) {
		var ret float64
		return ret
	}
	return *o.DaysToExpiration
}

// GetDaysToExpirationOk returns a tuple with the DaysToExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetDaysToExpirationOk() (*float64, bool) {
	if o == nil || IsNil(o.DaysToExpiration) {
		return nil, false
	}
	return o.DaysToExpiration, true
}

// HasDaysToExpiration returns a boolean if a field has been set.
func (o *OptionChain) HasDaysToExpiration() bool {
	if o != nil && !IsNil(o.DaysToExpiration) {
		return true
	}

	return false
}

// SetDaysToExpiration gets a reference to the given float64 and assigns it to the DaysToExpiration field.
func (o *OptionChain) SetDaysToExpiration(v float64) {
	o.DaysToExpiration = &v
}

// GetInterestRate returns the InterestRate field value if set, zero value otherwise.
func (o *OptionChain) GetInterestRate() float64 {
	if o == nil || IsNil(o.InterestRate) {
		var ret float64
		return ret
	}
	return *o.InterestRate
}

// GetInterestRateOk returns a tuple with the InterestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetInterestRateOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestRate) {
		return nil, false
	}
	return o.InterestRate, true
}

// HasInterestRate returns a boolean if a field has been set.
func (o *OptionChain) HasInterestRate() bool {
	if o != nil && !IsNil(o.InterestRate) {
		return true
	}

	return false
}

// SetInterestRate gets a reference to the given float64 and assigns it to the InterestRate field.
func (o *OptionChain) SetInterestRate(v float64) {
	o.InterestRate = &v
}

// GetUnderlyingPrice returns the UnderlyingPrice field value if set, zero value otherwise.
func (o *OptionChain) GetUnderlyingPrice() float64 {
	if o == nil || IsNil(o.UnderlyingPrice) {
		var ret float64
		return ret
	}
	return *o.UnderlyingPrice
}

// GetUnderlyingPriceOk returns a tuple with the UnderlyingPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetUnderlyingPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.UnderlyingPrice) {
		return nil, false
	}
	return o.UnderlyingPrice, true
}

// HasUnderlyingPrice returns a boolean if a field has been set.
func (o *OptionChain) HasUnderlyingPrice() bool {
	if o != nil && !IsNil(o.UnderlyingPrice) {
		return true
	}

	return false
}

// SetUnderlyingPrice gets a reference to the given float64 and assigns it to the UnderlyingPrice field.
func (o *OptionChain) SetUnderlyingPrice(v float64) {
	o.UnderlyingPrice = &v
}

// GetVolatility returns the Volatility field value if set, zero value otherwise.
func (o *OptionChain) GetVolatility() float64 {
	if o == nil || IsNil(o.Volatility) {
		var ret float64
		return ret
	}
	return *o.Volatility
}

// GetVolatilityOk returns a tuple with the Volatility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetVolatilityOk() (*float64, bool) {
	if o == nil || IsNil(o.Volatility) {
		return nil, false
	}
	return o.Volatility, true
}

// HasVolatility returns a boolean if a field has been set.
func (o *OptionChain) HasVolatility() bool {
	if o != nil && !IsNil(o.Volatility) {
		return true
	}

	return false
}

// SetVolatility gets a reference to the given float64 and assigns it to the Volatility field.
func (o *OptionChain) SetVolatility(v float64) {
	o.Volatility = &v
}

// GetCallExpDateMap returns the CallExpDateMap field value if set, zero value otherwise.
func (o *OptionChain) GetCallExpDateMap() map[string]map[string]OptionContract {
	if o == nil || IsNil(o.CallExpDateMap) {
		var ret map[string]map[string]OptionContract
		return ret
	}
	return *o.CallExpDateMap
}

// GetCallExpDateMapOk returns a tuple with the CallExpDateMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetCallExpDateMapOk() (*map[string]map[string]OptionContract, bool) {
	if o == nil || IsNil(o.CallExpDateMap) {
		return nil, false
	}
	return o.CallExpDateMap, true
}

// HasCallExpDateMap returns a boolean if a field has been set.
func (o *OptionChain) HasCallExpDateMap() bool {
	if o != nil && !IsNil(o.CallExpDateMap) {
		return true
	}

	return false
}

// SetCallExpDateMap gets a reference to the given map[string]map[string]OptionContract and assigns it to the CallExpDateMap field.
func (o *OptionChain) SetCallExpDateMap(v map[string]map[string]OptionContract) {
	o.CallExpDateMap = &v
}

// GetPutExpDateMap returns the PutExpDateMap field value if set, zero value otherwise.
func (o *OptionChain) GetPutExpDateMap() map[string]map[string]OptionContract {
	if o == nil || IsNil(o.PutExpDateMap) {
		var ret map[string]map[string]OptionContract
		return ret
	}
	return *o.PutExpDateMap
}

// GetPutExpDateMapOk returns a tuple with the PutExpDateMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionChain) GetPutExpDateMapOk() (*map[string]map[string]OptionContract, bool) {
	if o == nil || IsNil(o.PutExpDateMap) {
		return nil, false
	}
	return o.PutExpDateMap, true
}

// HasPutExpDateMap returns a boolean if a field has been set.
func (o *OptionChain) HasPutExpDateMap() bool {
	if o != nil && !IsNil(o.PutExpDateMap) {
		return true
	}

	return false
}

// SetPutExpDateMap gets a reference to the given map[string]map[string]OptionContract and assigns it to the PutExpDateMap field.
func (o *OptionChain) SetPutExpDateMap(v map[string]map[string]OptionContract) {
	o.PutExpDateMap = &v
}

func (o OptionChain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionChain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Underlying) {
		toSerialize["underlying"] = o.Underlying
	}
	if !IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.IsDelayed) {
		toSerialize["isDelayed"] = o.IsDelayed
	}
	if !IsNil(o.IsIndex) {
		toSerialize["isIndex"] = o.IsIndex
	}
	if !IsNil(o.DaysToExpiration) {
		toSerialize["daysToExpiration"] = o.DaysToExpiration
	}
	if !IsNil(o.InterestRate) {
		toSerialize["interestRate"] = o.InterestRate
	}
	if !IsNil(o.UnderlyingPrice) {
		toSerialize["underlyingPrice"] = o.UnderlyingPrice
	}
	if !IsNil(o.Volatility) {
		toSerialize["volatility"] = o.Volatility
	}
	if !IsNil(o.CallExpDateMap) {
		toSerialize["callExpDateMap"] = o.CallExpDateMap
	}
	if !IsNil(o.PutExpDateMap) {
		toSerialize["putExpDateMap"] = o.PutExpDateMap
	}
	return toSerialize, nil
}

type NullableOptionChain struct {
	value *OptionChain
	isSet bool
}

func (v NullableOptionChain) Get() *OptionChain {
	return v.value
}

func (v *NullableOptionChain) Set(val *OptionChain) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionChain) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionChain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionChain(val *OptionChain) *NullableOptionChain {
	return &NullableOptionChain{value: val, isSet: true}
}

func (v NullableOptionChain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionChain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


