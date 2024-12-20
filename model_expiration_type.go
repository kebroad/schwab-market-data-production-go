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

// ExpirationType M for End Of Month Expiration Calendar Cycle. (To match the last business day of the month), Q for Quarterly expirations (last business day of the quarter month MAR/JUN/SEP/DEC), W for Weekly expiration (also called Friday Short Term Expirations) and S for Expires 3rd Friday of the month (also known as regular options).
type ExpirationType string

// List of ExpirationType
const (
	EXPIRATIONTYPE_M ExpirationType = "M"
	EXPIRATIONTYPE_Q ExpirationType = "Q"
	EXPIRATIONTYPE_S ExpirationType = "S"
	EXPIRATIONTYPE_W ExpirationType = "W"
)

// All allowed values of ExpirationType enum
var AllowedExpirationTypeEnumValues = []ExpirationType{
	"M",
	"Q",
	"S",
	"W",
}

func (v *ExpirationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpirationType(value)
	for _, existing := range AllowedExpirationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpirationType", value)
}

// NewExpirationTypeFromValue returns a pointer to a valid ExpirationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpirationTypeFromValue(v string) (*ExpirationType, error) {
	ev := ExpirationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpirationType: valid values are %v", v, AllowedExpirationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpirationType) IsValid() bool {
	for _, existing := range AllowedExpirationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpirationType value
func (v ExpirationType) Ptr() *ExpirationType {
	return &v
}

type NullableExpirationType struct {
	value *ExpirationType
	isSet bool
}

func (v NullableExpirationType) Get() *ExpirationType {
	return v.value
}

func (v *NullableExpirationType) Set(val *ExpirationType) {
	v.value = val
	v.isSet = true
}

func (v NullableExpirationType) IsSet() bool {
	return v.isSet
}

func (v *NullableExpirationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpirationType(val *ExpirationType) *NullableExpirationType {
	return &NullableExpirationType{value: val, isSet: true}
}

func (v NullableExpirationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpirationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

