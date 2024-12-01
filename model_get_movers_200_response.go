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

// checks if the GetMovers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMovers200Response{}

// GetMovers200Response struct for GetMovers200Response
type GetMovers200Response struct {
	Screeners []Screener `json:"screeners,omitempty"`
}

// NewGetMovers200Response instantiates a new GetMovers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMovers200Response() *GetMovers200Response {
	this := GetMovers200Response{}
	return &this
}

// NewGetMovers200ResponseWithDefaults instantiates a new GetMovers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMovers200ResponseWithDefaults() *GetMovers200Response {
	this := GetMovers200Response{}
	return &this
}

// GetScreeners returns the Screeners field value if set, zero value otherwise.
func (o *GetMovers200Response) GetScreeners() []Screener {
	if o == nil || IsNil(o.Screeners) {
		var ret []Screener
		return ret
	}
	return o.Screeners
}

// GetScreenersOk returns a tuple with the Screeners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMovers200Response) GetScreenersOk() ([]Screener, bool) {
	if o == nil || IsNil(o.Screeners) {
		return nil, false
	}
	return o.Screeners, true
}

// HasScreeners returns a boolean if a field has been set.
func (o *GetMovers200Response) HasScreeners() bool {
	if o != nil && !IsNil(o.Screeners) {
		return true
	}

	return false
}

// SetScreeners gets a reference to the given []Screener and assigns it to the Screeners field.
func (o *GetMovers200Response) SetScreeners(v []Screener) {
	o.Screeners = v
}

func (o GetMovers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMovers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Screeners) {
		toSerialize["screeners"] = o.Screeners
	}
	return toSerialize, nil
}

type NullableGetMovers200Response struct {
	value *GetMovers200Response
	isSet bool
}

func (v NullableGetMovers200Response) Get() *GetMovers200Response {
	return v.value
}

func (v *NullableGetMovers200Response) Set(val *GetMovers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMovers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMovers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMovers200Response(val *GetMovers200Response) *NullableGetMovers200Response {
	return &NullableGetMovers200Response{value: val, isSet: true}
}

func (v NullableGetMovers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMovers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

