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

// checks if the Hours type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hours{}

// Hours struct for Hours
type Hours struct {
	Date *string `json:"date,omitempty"`
	MarketType *string `json:"marketType,omitempty"`
	Exchange *string `json:"exchange,omitempty"`
	Category *string `json:"category,omitempty"`
	Product *string `json:"product,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	IsOpen *bool `json:"isOpen,omitempty"`
	SessionHours *map[string][]Interval `json:"sessionHours,omitempty"`
}

// NewHours instantiates a new Hours object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHours() *Hours {
	this := Hours{}
	return &this
}

// NewHoursWithDefaults instantiates a new Hours object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHoursWithDefaults() *Hours {
	this := Hours{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Hours) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Hours) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Hours) SetDate(v string) {
	o.Date = &v
}

// GetMarketType returns the MarketType field value if set, zero value otherwise.
func (o *Hours) GetMarketType() string {
	if o == nil || IsNil(o.MarketType) {
		var ret string
		return ret
	}
	return *o.MarketType
}

// GetMarketTypeOk returns a tuple with the MarketType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetMarketTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MarketType) {
		return nil, false
	}
	return o.MarketType, true
}

// HasMarketType returns a boolean if a field has been set.
func (o *Hours) HasMarketType() bool {
	if o != nil && !IsNil(o.MarketType) {
		return true
	}

	return false
}

// SetMarketType gets a reference to the given string and assigns it to the MarketType field.
func (o *Hours) SetMarketType(v string) {
	o.MarketType = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Hours) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Hours) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Hours) SetExchange(v string) {
	o.Exchange = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Hours) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Hours) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Hours) SetCategory(v string) {
	o.Category = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *Hours) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *Hours) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *Hours) SetProduct(v string) {
	o.Product = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *Hours) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *Hours) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *Hours) SetProductName(v string) {
	o.ProductName = &v
}

// GetIsOpen returns the IsOpen field value if set, zero value otherwise.
func (o *Hours) GetIsOpen() bool {
	if o == nil || IsNil(o.IsOpen) {
		var ret bool
		return ret
	}
	return *o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetIsOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpen) {
		return nil, false
	}
	return o.IsOpen, true
}

// HasIsOpen returns a boolean if a field has been set.
func (o *Hours) HasIsOpen() bool {
	if o != nil && !IsNil(o.IsOpen) {
		return true
	}

	return false
}

// SetIsOpen gets a reference to the given bool and assigns it to the IsOpen field.
func (o *Hours) SetIsOpen(v bool) {
	o.IsOpen = &v
}

// GetSessionHours returns the SessionHours field value if set, zero value otherwise.
func (o *Hours) GetSessionHours() map[string][]Interval {
	if o == nil || IsNil(o.SessionHours) {
		var ret map[string][]Interval
		return ret
	}
	return *o.SessionHours
}

// GetSessionHoursOk returns a tuple with the SessionHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hours) GetSessionHoursOk() (*map[string][]Interval, bool) {
	if o == nil || IsNil(o.SessionHours) {
		return nil, false
	}
	return o.SessionHours, true
}

// HasSessionHours returns a boolean if a field has been set.
func (o *Hours) HasSessionHours() bool {
	if o != nil && !IsNil(o.SessionHours) {
		return true
	}

	return false
}

// SetSessionHours gets a reference to the given map[string][]Interval and assigns it to the SessionHours field.
func (o *Hours) SetSessionHours(v map[string][]Interval) {
	o.SessionHours = &v
}

func (o Hours) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hours) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.MarketType) {
		toSerialize["marketType"] = o.MarketType
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ProductName) {
		toSerialize["productName"] = o.ProductName
	}
	if !IsNil(o.IsOpen) {
		toSerialize["isOpen"] = o.IsOpen
	}
	if !IsNil(o.SessionHours) {
		toSerialize["sessionHours"] = o.SessionHours
	}
	return toSerialize, nil
}

type NullableHours struct {
	value *Hours
	isSet bool
}

func (v NullableHours) Get() *Hours {
	return v.value
}

func (v *NullableHours) Set(val *Hours) {
	v.value = val
	v.isSet = true
}

func (v NullableHours) IsSet() bool {
	return v.isSet
}

func (v *NullableHours) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHours(val *Hours) *NullableHours {
	return &NullableHours{value: val, isSet: true}
}

func (v NullableHours) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHours) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


