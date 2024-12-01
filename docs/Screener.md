# Screener

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Change** | Pointer to **float64** | percent or value changed, by default its percent changed | [optional] 
**Description** | Pointer to **string** | Name of security | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Last** | Pointer to **float64** | what was last quoted price | [optional] 
**Symbol** | Pointer to **string** | schwab security symbol | [optional] 
**TotalVolume** | Pointer to **int64** |  | [optional] 

## Methods

### NewScreener

`func NewScreener() *Screener`

NewScreener instantiates a new Screener object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenerWithDefaults

`func NewScreenerWithDefaults() *Screener`

NewScreenerWithDefaults instantiates a new Screener object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChange

`func (o *Screener) GetChange() float64`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *Screener) GetChangeOk() (*float64, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *Screener) SetChange(v float64)`

SetChange sets Change field to given value.

### HasChange

`func (o *Screener) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetDescription

`func (o *Screener) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Screener) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Screener) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Screener) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *Screener) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Screener) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Screener) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Screener) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetLast

`func (o *Screener) GetLast() float64`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *Screener) GetLastOk() (*float64, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *Screener) SetLast(v float64)`

SetLast sets Last field to given value.

### HasLast

`func (o *Screener) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetSymbol

`func (o *Screener) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Screener) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Screener) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Screener) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTotalVolume

`func (o *Screener) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *Screener) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *Screener) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *Screener) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


