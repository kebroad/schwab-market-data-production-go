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

// EquityAssetSubType Asset Sub Type (only there if applicable)
type EquityAssetSubType string

// List of EquityAssetSubType
const (
	EQUITYASSETSUBTYPE_COE EquityAssetSubType = "COE"
	EQUITYASSETSUBTYPE_PRF EquityAssetSubType = "PRF"
	EQUITYASSETSUBTYPE_ADR EquityAssetSubType = "ADR"
	EQUITYASSETSUBTYPE_GDR EquityAssetSubType = "GDR"
	EQUITYASSETSUBTYPE_CEF EquityAssetSubType = "CEF"
	EQUITYASSETSUBTYPE_ETF EquityAssetSubType = "ETF"
	EQUITYASSETSUBTYPE_ETN EquityAssetSubType = "ETN"
	EQUITYASSETSUBTYPE_UIT EquityAssetSubType = "UIT"
	EQUITYASSETSUBTYPE_WAR EquityAssetSubType = "WAR"
	EQUITYASSETSUBTYPE_RGT EquityAssetSubType = "RGT"
)

// All allowed values of EquityAssetSubType enum
var AllowedEquityAssetSubTypeEnumValues = []EquityAssetSubType{
	"COE",
	"PRF",
	"ADR",
	"GDR",
	"CEF",
	"ETF",
	"ETN",
	"UIT",
	"WAR",
	"RGT",
}

func (v *EquityAssetSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EquityAssetSubType(value)
	for _, existing := range AllowedEquityAssetSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EquityAssetSubType", value)
}

// NewEquityAssetSubTypeFromValue returns a pointer to a valid EquityAssetSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEquityAssetSubTypeFromValue(v string) (*EquityAssetSubType, error) {
	ev := EquityAssetSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EquityAssetSubType: valid values are %v", v, AllowedEquityAssetSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EquityAssetSubType) IsValid() bool {
	for _, existing := range AllowedEquityAssetSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EquityAssetSubType value
func (v EquityAssetSubType) Ptr() *EquityAssetSubType {
	return &v
}

type NullableEquityAssetSubType struct {
	value *EquityAssetSubType
	isSet bool
}

func (v NullableEquityAssetSubType) Get() *EquityAssetSubType {
	return v.value
}

func (v *NullableEquityAssetSubType) Set(val *EquityAssetSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityAssetSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityAssetSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityAssetSubType(val *EquityAssetSubType) *NullableEquityAssetSubType {
	return &NullableEquityAssetSubType{value: val, isSet: true}
}

func (v NullableEquityAssetSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityAssetSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

