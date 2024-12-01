# Fundamental

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avg10DaysVolume** | Pointer to **float64** | Average 10 day volume | [optional] 
**Avg1YearVolume** | Pointer to **float64** | Average 1 day volume | [optional] 
**DeclarationDate** | Pointer to **time.Time** | Declaration date in yyyy-mm-ddThh:mm:ssZ | [optional] 
**DivAmount** | Pointer to **float64** | Dividend Amount | [optional] 
**DivExDate** | Pointer to **string** | Dividend date in yyyy-mm-ddThh:mm:ssZ | [optional] 
**DivFreq** | Pointer to [**NullableDivFreq**](DivFreq.md) |  | [optional] 
**DivPayAmount** | Pointer to **float64** | Dividend Pay Amount | [optional] 
**DivPayDate** | Pointer to **time.Time** | Dividend pay date in yyyy-mm-ddThh:mm:ssZ | [optional] 
**DivYield** | Pointer to **float64** | Dividend yield | [optional] 
**Eps** | Pointer to **float64** | Earnings per Share | [optional] 
**FundLeverageFactor** | Pointer to **float64** | Fund Leverage Factor + &gt; 0 &lt;- | [optional] 
**FundStrategy** | Pointer to [**NullableFundStrategy**](FundStrategy.md) |  | [optional] 
**NextDivExDate** | Pointer to **time.Time** | Next Dividend date | [optional] 
**NextDivPayDate** | Pointer to **time.Time** | Next Dividend pay date | [optional] 
**PeRatio** | Pointer to **float64** | P/E Ratio | [optional] 

## Methods

### NewFundamental

`func NewFundamental() *Fundamental`

NewFundamental instantiates a new Fundamental object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalWithDefaults

`func NewFundamentalWithDefaults() *Fundamental`

NewFundamentalWithDefaults instantiates a new Fundamental object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvg10DaysVolume

`func (o *Fundamental) GetAvg10DaysVolume() float64`

GetAvg10DaysVolume returns the Avg10DaysVolume field if non-nil, zero value otherwise.

### GetAvg10DaysVolumeOk

`func (o *Fundamental) GetAvg10DaysVolumeOk() (*float64, bool)`

GetAvg10DaysVolumeOk returns a tuple with the Avg10DaysVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg10DaysVolume

`func (o *Fundamental) SetAvg10DaysVolume(v float64)`

SetAvg10DaysVolume sets Avg10DaysVolume field to given value.

### HasAvg10DaysVolume

`func (o *Fundamental) HasAvg10DaysVolume() bool`

HasAvg10DaysVolume returns a boolean if a field has been set.

### GetAvg1YearVolume

`func (o *Fundamental) GetAvg1YearVolume() float64`

GetAvg1YearVolume returns the Avg1YearVolume field if non-nil, zero value otherwise.

### GetAvg1YearVolumeOk

`func (o *Fundamental) GetAvg1YearVolumeOk() (*float64, bool)`

GetAvg1YearVolumeOk returns a tuple with the Avg1YearVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg1YearVolume

`func (o *Fundamental) SetAvg1YearVolume(v float64)`

SetAvg1YearVolume sets Avg1YearVolume field to given value.

### HasAvg1YearVolume

`func (o *Fundamental) HasAvg1YearVolume() bool`

HasAvg1YearVolume returns a boolean if a field has been set.

### GetDeclarationDate

`func (o *Fundamental) GetDeclarationDate() time.Time`

GetDeclarationDate returns the DeclarationDate field if non-nil, zero value otherwise.

### GetDeclarationDateOk

`func (o *Fundamental) GetDeclarationDateOk() (*time.Time, bool)`

GetDeclarationDateOk returns a tuple with the DeclarationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationDate

`func (o *Fundamental) SetDeclarationDate(v time.Time)`

SetDeclarationDate sets DeclarationDate field to given value.

### HasDeclarationDate

`func (o *Fundamental) HasDeclarationDate() bool`

HasDeclarationDate returns a boolean if a field has been set.

### GetDivAmount

`func (o *Fundamental) GetDivAmount() float64`

GetDivAmount returns the DivAmount field if non-nil, zero value otherwise.

### GetDivAmountOk

`func (o *Fundamental) GetDivAmountOk() (*float64, bool)`

GetDivAmountOk returns a tuple with the DivAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivAmount

`func (o *Fundamental) SetDivAmount(v float64)`

SetDivAmount sets DivAmount field to given value.

### HasDivAmount

`func (o *Fundamental) HasDivAmount() bool`

HasDivAmount returns a boolean if a field has been set.

### GetDivExDate

`func (o *Fundamental) GetDivExDate() string`

GetDivExDate returns the DivExDate field if non-nil, zero value otherwise.

### GetDivExDateOk

`func (o *Fundamental) GetDivExDateOk() (*string, bool)`

GetDivExDateOk returns a tuple with the DivExDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivExDate

`func (o *Fundamental) SetDivExDate(v string)`

SetDivExDate sets DivExDate field to given value.

### HasDivExDate

`func (o *Fundamental) HasDivExDate() bool`

HasDivExDate returns a boolean if a field has been set.

### GetDivFreq

`func (o *Fundamental) GetDivFreq() DivFreq`

GetDivFreq returns the DivFreq field if non-nil, zero value otherwise.

### GetDivFreqOk

`func (o *Fundamental) GetDivFreqOk() (*DivFreq, bool)`

GetDivFreqOk returns a tuple with the DivFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivFreq

`func (o *Fundamental) SetDivFreq(v DivFreq)`

SetDivFreq sets DivFreq field to given value.

### HasDivFreq

`func (o *Fundamental) HasDivFreq() bool`

HasDivFreq returns a boolean if a field has been set.

### SetDivFreqNil

`func (o *Fundamental) SetDivFreqNil(b bool)`

 SetDivFreqNil sets the value for DivFreq to be an explicit nil

### UnsetDivFreq
`func (o *Fundamental) UnsetDivFreq()`

UnsetDivFreq ensures that no value is present for DivFreq, not even an explicit nil
### GetDivPayAmount

`func (o *Fundamental) GetDivPayAmount() float64`

GetDivPayAmount returns the DivPayAmount field if non-nil, zero value otherwise.

### GetDivPayAmountOk

`func (o *Fundamental) GetDivPayAmountOk() (*float64, bool)`

GetDivPayAmountOk returns a tuple with the DivPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivPayAmount

`func (o *Fundamental) SetDivPayAmount(v float64)`

SetDivPayAmount sets DivPayAmount field to given value.

### HasDivPayAmount

`func (o *Fundamental) HasDivPayAmount() bool`

HasDivPayAmount returns a boolean if a field has been set.

### GetDivPayDate

`func (o *Fundamental) GetDivPayDate() time.Time`

GetDivPayDate returns the DivPayDate field if non-nil, zero value otherwise.

### GetDivPayDateOk

`func (o *Fundamental) GetDivPayDateOk() (*time.Time, bool)`

GetDivPayDateOk returns a tuple with the DivPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivPayDate

`func (o *Fundamental) SetDivPayDate(v time.Time)`

SetDivPayDate sets DivPayDate field to given value.

### HasDivPayDate

`func (o *Fundamental) HasDivPayDate() bool`

HasDivPayDate returns a boolean if a field has been set.

### GetDivYield

`func (o *Fundamental) GetDivYield() float64`

GetDivYield returns the DivYield field if non-nil, zero value otherwise.

### GetDivYieldOk

`func (o *Fundamental) GetDivYieldOk() (*float64, bool)`

GetDivYieldOk returns a tuple with the DivYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivYield

`func (o *Fundamental) SetDivYield(v float64)`

SetDivYield sets DivYield field to given value.

### HasDivYield

`func (o *Fundamental) HasDivYield() bool`

HasDivYield returns a boolean if a field has been set.

### GetEps

`func (o *Fundamental) GetEps() float64`

GetEps returns the Eps field if non-nil, zero value otherwise.

### GetEpsOk

`func (o *Fundamental) GetEpsOk() (*float64, bool)`

GetEpsOk returns a tuple with the Eps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEps

`func (o *Fundamental) SetEps(v float64)`

SetEps sets Eps field to given value.

### HasEps

`func (o *Fundamental) HasEps() bool`

HasEps returns a boolean if a field has been set.

### GetFundLeverageFactor

`func (o *Fundamental) GetFundLeverageFactor() float64`

GetFundLeverageFactor returns the FundLeverageFactor field if non-nil, zero value otherwise.

### GetFundLeverageFactorOk

`func (o *Fundamental) GetFundLeverageFactorOk() (*float64, bool)`

GetFundLeverageFactorOk returns a tuple with the FundLeverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundLeverageFactor

`func (o *Fundamental) SetFundLeverageFactor(v float64)`

SetFundLeverageFactor sets FundLeverageFactor field to given value.

### HasFundLeverageFactor

`func (o *Fundamental) HasFundLeverageFactor() bool`

HasFundLeverageFactor returns a boolean if a field has been set.

### GetFundStrategy

`func (o *Fundamental) GetFundStrategy() FundStrategy`

GetFundStrategy returns the FundStrategy field if non-nil, zero value otherwise.

### GetFundStrategyOk

`func (o *Fundamental) GetFundStrategyOk() (*FundStrategy, bool)`

GetFundStrategyOk returns a tuple with the FundStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundStrategy

`func (o *Fundamental) SetFundStrategy(v FundStrategy)`

SetFundStrategy sets FundStrategy field to given value.

### HasFundStrategy

`func (o *Fundamental) HasFundStrategy() bool`

HasFundStrategy returns a boolean if a field has been set.

### SetFundStrategyNil

`func (o *Fundamental) SetFundStrategyNil(b bool)`

 SetFundStrategyNil sets the value for FundStrategy to be an explicit nil

### UnsetFundStrategy
`func (o *Fundamental) UnsetFundStrategy()`

UnsetFundStrategy ensures that no value is present for FundStrategy, not even an explicit nil
### GetNextDivExDate

`func (o *Fundamental) GetNextDivExDate() time.Time`

GetNextDivExDate returns the NextDivExDate field if non-nil, zero value otherwise.

### GetNextDivExDateOk

`func (o *Fundamental) GetNextDivExDateOk() (*time.Time, bool)`

GetNextDivExDateOk returns a tuple with the NextDivExDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDivExDate

`func (o *Fundamental) SetNextDivExDate(v time.Time)`

SetNextDivExDate sets NextDivExDate field to given value.

### HasNextDivExDate

`func (o *Fundamental) HasNextDivExDate() bool`

HasNextDivExDate returns a boolean if a field has been set.

### GetNextDivPayDate

`func (o *Fundamental) GetNextDivPayDate() time.Time`

GetNextDivPayDate returns the NextDivPayDate field if non-nil, zero value otherwise.

### GetNextDivPayDateOk

`func (o *Fundamental) GetNextDivPayDateOk() (*time.Time, bool)`

GetNextDivPayDateOk returns a tuple with the NextDivPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDivPayDate

`func (o *Fundamental) SetNextDivPayDate(v time.Time)`

SetNextDivPayDate sets NextDivPayDate field to given value.

### HasNextDivPayDate

`func (o *Fundamental) HasNextDivPayDate() bool`

HasNextDivPayDate returns a boolean if a field has been set.

### GetPeRatio

`func (o *Fundamental) GetPeRatio() float64`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *Fundamental) GetPeRatioOk() (*float64, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *Fundamental) SetPeRatio(v float64)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *Fundamental) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


