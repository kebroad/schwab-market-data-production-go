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
	"gopkg.in/validator.v2"
	"fmt"
)

// QuoteResponseObject - struct for QuoteResponseObject
type QuoteResponseObject struct {
	EquityResponse *EquityResponse
	ForexResponse *ForexResponse
	FutureOptionResponse *FutureOptionResponse
	FutureResponse *FutureResponse
	IndexResponse *IndexResponse
	MutualFundResponse *MutualFundResponse
	OptionResponse *OptionResponse
	QuoteError *QuoteError
}

// EquityResponseAsQuoteResponseObject is a convenience function that returns EquityResponse wrapped in QuoteResponseObject
func EquityResponseAsQuoteResponseObject(v *EquityResponse) QuoteResponseObject {
	return QuoteResponseObject{
		EquityResponse: v,
	}
}

// ForexResponseAsQuoteResponseObject is a convenience function that returns ForexResponse wrapped in QuoteResponseObject
func ForexResponseAsQuoteResponseObject(v *ForexResponse) QuoteResponseObject {
	return QuoteResponseObject{
		ForexResponse: v,
	}
}

// FutureOptionResponseAsQuoteResponseObject is a convenience function that returns FutureOptionResponse wrapped in QuoteResponseObject
func FutureOptionResponseAsQuoteResponseObject(v *FutureOptionResponse) QuoteResponseObject {
	return QuoteResponseObject{
		FutureOptionResponse: v,
	}
}

// FutureResponseAsQuoteResponseObject is a convenience function that returns FutureResponse wrapped in QuoteResponseObject
func FutureResponseAsQuoteResponseObject(v *FutureResponse) QuoteResponseObject {
	return QuoteResponseObject{
		FutureResponse: v,
	}
}

// IndexResponseAsQuoteResponseObject is a convenience function that returns IndexResponse wrapped in QuoteResponseObject
func IndexResponseAsQuoteResponseObject(v *IndexResponse) QuoteResponseObject {
	return QuoteResponseObject{
		IndexResponse: v,
	}
}

// MutualFundResponseAsQuoteResponseObject is a convenience function that returns MutualFundResponse wrapped in QuoteResponseObject
func MutualFundResponseAsQuoteResponseObject(v *MutualFundResponse) QuoteResponseObject {
	return QuoteResponseObject{
		MutualFundResponse: v,
	}
}

// OptionResponseAsQuoteResponseObject is a convenience function that returns OptionResponse wrapped in QuoteResponseObject
func OptionResponseAsQuoteResponseObject(v *OptionResponse) QuoteResponseObject {
	return QuoteResponseObject{
		OptionResponse: v,
	}
}

// QuoteErrorAsQuoteResponseObject is a convenience function that returns QuoteError wrapped in QuoteResponseObject
func QuoteErrorAsQuoteResponseObject(v *QuoteError) QuoteResponseObject {
	return QuoteResponseObject{
		QuoteError: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QuoteResponseObject) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into EquityResponse
	err = newStrictDecoder(data).Decode(&dst.EquityResponse)
	if err == nil {
		jsonEquityResponse, _ := json.Marshal(dst.EquityResponse)
		if string(jsonEquityResponse) == "{}" { // empty struct
			dst.EquityResponse = nil
		} else {
			if err = validator.Validate(dst.EquityResponse); err != nil {
				dst.EquityResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.EquityResponse = nil
	}

	// try to unmarshal data into ForexResponse
	err = newStrictDecoder(data).Decode(&dst.ForexResponse)
	if err == nil {
		jsonForexResponse, _ := json.Marshal(dst.ForexResponse)
		if string(jsonForexResponse) == "{}" { // empty struct
			dst.ForexResponse = nil
		} else {
			if err = validator.Validate(dst.ForexResponse); err != nil {
				dst.ForexResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ForexResponse = nil
	}

	// try to unmarshal data into FutureOptionResponse
	err = newStrictDecoder(data).Decode(&dst.FutureOptionResponse)
	if err == nil {
		jsonFutureOptionResponse, _ := json.Marshal(dst.FutureOptionResponse)
		if string(jsonFutureOptionResponse) == "{}" { // empty struct
			dst.FutureOptionResponse = nil
		} else {
			if err = validator.Validate(dst.FutureOptionResponse); err != nil {
				dst.FutureOptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FutureOptionResponse = nil
	}

	// try to unmarshal data into FutureResponse
	err = newStrictDecoder(data).Decode(&dst.FutureResponse)
	if err == nil {
		jsonFutureResponse, _ := json.Marshal(dst.FutureResponse)
		if string(jsonFutureResponse) == "{}" { // empty struct
			dst.FutureResponse = nil
		} else {
			if err = validator.Validate(dst.FutureResponse); err != nil {
				dst.FutureResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.FutureResponse = nil
	}

	// try to unmarshal data into IndexResponse
	err = newStrictDecoder(data).Decode(&dst.IndexResponse)
	if err == nil {
		jsonIndexResponse, _ := json.Marshal(dst.IndexResponse)
		if string(jsonIndexResponse) == "{}" { // empty struct
			dst.IndexResponse = nil
		} else {
			if err = validator.Validate(dst.IndexResponse); err != nil {
				dst.IndexResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.IndexResponse = nil
	}

	// try to unmarshal data into MutualFundResponse
	err = newStrictDecoder(data).Decode(&dst.MutualFundResponse)
	if err == nil {
		jsonMutualFundResponse, _ := json.Marshal(dst.MutualFundResponse)
		if string(jsonMutualFundResponse) == "{}" { // empty struct
			dst.MutualFundResponse = nil
		} else {
			if err = validator.Validate(dst.MutualFundResponse); err != nil {
				dst.MutualFundResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.MutualFundResponse = nil
	}

	// try to unmarshal data into OptionResponse
	err = newStrictDecoder(data).Decode(&dst.OptionResponse)
	if err == nil {
		jsonOptionResponse, _ := json.Marshal(dst.OptionResponse)
		if string(jsonOptionResponse) == "{}" { // empty struct
			dst.OptionResponse = nil
		} else {
			if err = validator.Validate(dst.OptionResponse); err != nil {
				dst.OptionResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.OptionResponse = nil
	}

	// try to unmarshal data into QuoteError
	err = newStrictDecoder(data).Decode(&dst.QuoteError)
	if err == nil {
		jsonQuoteError, _ := json.Marshal(dst.QuoteError)
		if string(jsonQuoteError) == "{}" { // empty struct
			dst.QuoteError = nil
		} else {
			if err = validator.Validate(dst.QuoteError); err != nil {
				dst.QuoteError = nil
			} else {
				match++
			}
		}
	} else {
		dst.QuoteError = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.EquityResponse = nil
		dst.ForexResponse = nil
		dst.FutureOptionResponse = nil
		dst.FutureResponse = nil
		dst.IndexResponse = nil
		dst.MutualFundResponse = nil
		dst.OptionResponse = nil
		dst.QuoteError = nil

		return fmt.Errorf("data matches more than one schema in oneOf(QuoteResponseObject)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QuoteResponseObject)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QuoteResponseObject) MarshalJSON() ([]byte, error) {
	if src.EquityResponse != nil {
		return json.Marshal(&src.EquityResponse)
	}

	if src.ForexResponse != nil {
		return json.Marshal(&src.ForexResponse)
	}

	if src.FutureOptionResponse != nil {
		return json.Marshal(&src.FutureOptionResponse)
	}

	if src.FutureResponse != nil {
		return json.Marshal(&src.FutureResponse)
	}

	if src.IndexResponse != nil {
		return json.Marshal(&src.IndexResponse)
	}

	if src.MutualFundResponse != nil {
		return json.Marshal(&src.MutualFundResponse)
	}

	if src.OptionResponse != nil {
		return json.Marshal(&src.OptionResponse)
	}

	if src.QuoteError != nil {
		return json.Marshal(&src.QuoteError)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QuoteResponseObject) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EquityResponse != nil {
		return obj.EquityResponse
	}

	if obj.ForexResponse != nil {
		return obj.ForexResponse
	}

	if obj.FutureOptionResponse != nil {
		return obj.FutureOptionResponse
	}

	if obj.FutureResponse != nil {
		return obj.FutureResponse
	}

	if obj.IndexResponse != nil {
		return obj.IndexResponse
	}

	if obj.MutualFundResponse != nil {
		return obj.MutualFundResponse
	}

	if obj.OptionResponse != nil {
		return obj.OptionResponse
	}

	if obj.QuoteError != nil {
		return obj.QuoteError
	}

	// all schemas are nil
	return nil
}

type NullableQuoteResponseObject struct {
	value *QuoteResponseObject
	isSet bool
}

func (v NullableQuoteResponseObject) Get() *QuoteResponseObject {
	return v.value
}

func (v *NullableQuoteResponseObject) Set(val *QuoteResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResponseObject(val *QuoteResponseObject) *NullableQuoteResponseObject {
	return &NullableQuoteResponseObject{value: val, isSet: true}
}

func (v NullableQuoteResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


