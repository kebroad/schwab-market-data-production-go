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
	"fmt"
)

// FundStrategy FundStrategy \"A\" - Active \"L\" - Leveraged \"P\" - Passive \"Q\" - Quantitative \"S\" - Short
type FundStrategy string

// List of FundStrategy
const (
	FUNDSTRATEGY_A FundStrategy = "A"
	FUNDSTRATEGY_L FundStrategy = "L"
	FUNDSTRATEGY_P FundStrategy = "P"
	FUNDSTRATEGY_Q FundStrategy = "Q"
	FUNDSTRATEGY_S FundStrategy = "S"
)

// All allowed values of FundStrategy enum
var AllowedFundStrategyEnumValues = []FundStrategy{
	"A",
	"L",
	"P",
	"Q",
	"S",
}

func (v *FundStrategy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FundStrategy(value)
	for _, existing := range AllowedFundStrategyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FundStrategy", value)
}

// NewFundStrategyFromValue returns a pointer to a valid FundStrategy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFundStrategyFromValue(v string) (*FundStrategy, error) {
	ev := FundStrategy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FundStrategy: valid values are %v", v, AllowedFundStrategyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FundStrategy) IsValid() bool {
	for _, existing := range AllowedFundStrategyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FundStrategy value
func (v FundStrategy) Ptr() *FundStrategy {
	return &v
}

type NullableFundStrategy struct {
	value *FundStrategy
	isSet bool
}

func (v NullableFundStrategy) Get() *FundStrategy {
	return v.value
}

func (v *NullableFundStrategy) Set(val *FundStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableFundStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableFundStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundStrategy(val *FundStrategy) *NullableFundStrategy {
	return &NullableFundStrategy{value: val, isSet: true}
}

func (v NullableFundStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

