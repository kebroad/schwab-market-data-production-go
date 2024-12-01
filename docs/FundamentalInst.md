# FundamentalInst

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**High52** | Pointer to **float64** |  | [optional] 
**Low52** | Pointer to **float64** |  | [optional] 
**DividendAmount** | Pointer to **float64** |  | [optional] 
**DividendYield** | Pointer to **float64** |  | [optional] 
**DividendDate** | Pointer to **string** |  | [optional] 
**PeRatio** | Pointer to **float64** |  | [optional] 
**PegRatio** | Pointer to **float64** |  | [optional] 
**PbRatio** | Pointer to **float64** |  | [optional] 
**PrRatio** | Pointer to **float64** |  | [optional] 
**PcfRatio** | Pointer to **float64** |  | [optional] 
**GrossMarginTTM** | Pointer to **float64** |  | [optional] 
**GrossMarginMRQ** | Pointer to **float64** |  | [optional] 
**NetProfitMarginTTM** | Pointer to **float64** |  | [optional] 
**NetProfitMarginMRQ** | Pointer to **float64** |  | [optional] 
**OperatingMarginTTM** | Pointer to **float64** |  | [optional] 
**OperatingMarginMRQ** | Pointer to **float64** |  | [optional] 
**ReturnOnEquity** | Pointer to **float64** |  | [optional] 
**ReturnOnAssets** | Pointer to **float64** |  | [optional] 
**ReturnOnInvestment** | Pointer to **float64** |  | [optional] 
**QuickRatio** | Pointer to **float64** |  | [optional] 
**CurrentRatio** | Pointer to **float64** |  | [optional] 
**InterestCoverage** | Pointer to **float64** |  | [optional] 
**TotalDebtToCapital** | Pointer to **float64** |  | [optional] 
**LtDebtToEquity** | Pointer to **float64** |  | [optional] 
**TotalDebtToEquity** | Pointer to **float64** |  | [optional] 
**EpsTTM** | Pointer to **float64** |  | [optional] 
**EpsChangePercentTTM** | Pointer to **float64** |  | [optional] 
**EpsChangeYear** | Pointer to **float64** |  | [optional] 
**EpsChange** | Pointer to **float64** |  | [optional] 
**RevChangeYear** | Pointer to **float64** |  | [optional] 
**RevChangeTTM** | Pointer to **float64** |  | [optional] 
**RevChangeIn** | Pointer to **float64** |  | [optional] 
**SharesOutstanding** | Pointer to **float64** |  | [optional] 
**MarketCapFloat** | Pointer to **float64** |  | [optional] 
**MarketCap** | Pointer to **float64** |  | [optional] 
**BookValuePerShare** | Pointer to **float64** |  | [optional] 
**ShortIntToFloat** | Pointer to **float64** |  | [optional] 
**ShortIntDayToCover** | Pointer to **float64** |  | [optional] 
**DivGrowthRate3Year** | Pointer to **float64** |  | [optional] 
**DividendPayAmount** | Pointer to **float64** |  | [optional] 
**DividendPayDate** | Pointer to **string** |  | [optional] 
**Beta** | Pointer to **float64** |  | [optional] 
**Vol1DayAvg** | Pointer to **float64** |  | [optional] 
**Vol10DayAvg** | Pointer to **float64** |  | [optional] 
**Vol3MonthAvg** | Pointer to **float64** |  | [optional] 
**Avg10DaysVolume** | Pointer to **int64** |  | [optional] 
**Avg1DayVolume** | Pointer to **int64** |  | [optional] 
**Avg3MonthVolume** | Pointer to **int64** |  | [optional] 
**DeclarationDate** | Pointer to **string** |  | [optional] 
**DividendFreq** | Pointer to **int32** |  | [optional] 
**Eps** | Pointer to **float64** |  | [optional] 
**CorpactionDate** | Pointer to **string** |  | [optional] 
**DtnVolume** | Pointer to **int64** |  | [optional] 
**NextDividendPayDate** | Pointer to **string** |  | [optional] 
**NextDividendDate** | Pointer to **string** |  | [optional] 
**FundLeverageFactor** | Pointer to **float64** |  | [optional] 
**FundStrategy** | Pointer to **string** |  | [optional] 

## Methods

### NewFundamentalInst

`func NewFundamentalInst() *FundamentalInst`

NewFundamentalInst instantiates a new FundamentalInst object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundamentalInstWithDefaults

`func NewFundamentalInstWithDefaults() *FundamentalInst`

NewFundamentalInstWithDefaults instantiates a new FundamentalInst object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FundamentalInst) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FundamentalInst) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FundamentalInst) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FundamentalInst) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetHigh52

`func (o *FundamentalInst) GetHigh52() float64`

GetHigh52 returns the High52 field if non-nil, zero value otherwise.

### GetHigh52Ok

`func (o *FundamentalInst) GetHigh52Ok() (*float64, bool)`

GetHigh52Ok returns a tuple with the High52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh52

`func (o *FundamentalInst) SetHigh52(v float64)`

SetHigh52 sets High52 field to given value.

### HasHigh52

`func (o *FundamentalInst) HasHigh52() bool`

HasHigh52 returns a boolean if a field has been set.

### GetLow52

`func (o *FundamentalInst) GetLow52() float64`

GetLow52 returns the Low52 field if non-nil, zero value otherwise.

### GetLow52Ok

`func (o *FundamentalInst) GetLow52Ok() (*float64, bool)`

GetLow52Ok returns a tuple with the Low52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow52

`func (o *FundamentalInst) SetLow52(v float64)`

SetLow52 sets Low52 field to given value.

### HasLow52

`func (o *FundamentalInst) HasLow52() bool`

HasLow52 returns a boolean if a field has been set.

### GetDividendAmount

`func (o *FundamentalInst) GetDividendAmount() float64`

GetDividendAmount returns the DividendAmount field if non-nil, zero value otherwise.

### GetDividendAmountOk

`func (o *FundamentalInst) GetDividendAmountOk() (*float64, bool)`

GetDividendAmountOk returns a tuple with the DividendAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendAmount

`func (o *FundamentalInst) SetDividendAmount(v float64)`

SetDividendAmount sets DividendAmount field to given value.

### HasDividendAmount

`func (o *FundamentalInst) HasDividendAmount() bool`

HasDividendAmount returns a boolean if a field has been set.

### GetDividendYield

`func (o *FundamentalInst) GetDividendYield() float64`

GetDividendYield returns the DividendYield field if non-nil, zero value otherwise.

### GetDividendYieldOk

`func (o *FundamentalInst) GetDividendYieldOk() (*float64, bool)`

GetDividendYieldOk returns a tuple with the DividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendYield

`func (o *FundamentalInst) SetDividendYield(v float64)`

SetDividendYield sets DividendYield field to given value.

### HasDividendYield

`func (o *FundamentalInst) HasDividendYield() bool`

HasDividendYield returns a boolean if a field has been set.

### GetDividendDate

`func (o *FundamentalInst) GetDividendDate() string`

GetDividendDate returns the DividendDate field if non-nil, zero value otherwise.

### GetDividendDateOk

`func (o *FundamentalInst) GetDividendDateOk() (*string, bool)`

GetDividendDateOk returns a tuple with the DividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendDate

`func (o *FundamentalInst) SetDividendDate(v string)`

SetDividendDate sets DividendDate field to given value.

### HasDividendDate

`func (o *FundamentalInst) HasDividendDate() bool`

HasDividendDate returns a boolean if a field has been set.

### GetPeRatio

`func (o *FundamentalInst) GetPeRatio() float64`

GetPeRatio returns the PeRatio field if non-nil, zero value otherwise.

### GetPeRatioOk

`func (o *FundamentalInst) GetPeRatioOk() (*float64, bool)`

GetPeRatioOk returns a tuple with the PeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeRatio

`func (o *FundamentalInst) SetPeRatio(v float64)`

SetPeRatio sets PeRatio field to given value.

### HasPeRatio

`func (o *FundamentalInst) HasPeRatio() bool`

HasPeRatio returns a boolean if a field has been set.

### GetPegRatio

`func (o *FundamentalInst) GetPegRatio() float64`

GetPegRatio returns the PegRatio field if non-nil, zero value otherwise.

### GetPegRatioOk

`func (o *FundamentalInst) GetPegRatioOk() (*float64, bool)`

GetPegRatioOk returns a tuple with the PegRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegRatio

`func (o *FundamentalInst) SetPegRatio(v float64)`

SetPegRatio sets PegRatio field to given value.

### HasPegRatio

`func (o *FundamentalInst) HasPegRatio() bool`

HasPegRatio returns a boolean if a field has been set.

### GetPbRatio

`func (o *FundamentalInst) GetPbRatio() float64`

GetPbRatio returns the PbRatio field if non-nil, zero value otherwise.

### GetPbRatioOk

`func (o *FundamentalInst) GetPbRatioOk() (*float64, bool)`

GetPbRatioOk returns a tuple with the PbRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPbRatio

`func (o *FundamentalInst) SetPbRatio(v float64)`

SetPbRatio sets PbRatio field to given value.

### HasPbRatio

`func (o *FundamentalInst) HasPbRatio() bool`

HasPbRatio returns a boolean if a field has been set.

### GetPrRatio

`func (o *FundamentalInst) GetPrRatio() float64`

GetPrRatio returns the PrRatio field if non-nil, zero value otherwise.

### GetPrRatioOk

`func (o *FundamentalInst) GetPrRatioOk() (*float64, bool)`

GetPrRatioOk returns a tuple with the PrRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrRatio

`func (o *FundamentalInst) SetPrRatio(v float64)`

SetPrRatio sets PrRatio field to given value.

### HasPrRatio

`func (o *FundamentalInst) HasPrRatio() bool`

HasPrRatio returns a boolean if a field has been set.

### GetPcfRatio

`func (o *FundamentalInst) GetPcfRatio() float64`

GetPcfRatio returns the PcfRatio field if non-nil, zero value otherwise.

### GetPcfRatioOk

`func (o *FundamentalInst) GetPcfRatioOk() (*float64, bool)`

GetPcfRatioOk returns a tuple with the PcfRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfRatio

`func (o *FundamentalInst) SetPcfRatio(v float64)`

SetPcfRatio sets PcfRatio field to given value.

### HasPcfRatio

`func (o *FundamentalInst) HasPcfRatio() bool`

HasPcfRatio returns a boolean if a field has been set.

### GetGrossMarginTTM

`func (o *FundamentalInst) GetGrossMarginTTM() float64`

GetGrossMarginTTM returns the GrossMarginTTM field if non-nil, zero value otherwise.

### GetGrossMarginTTMOk

`func (o *FundamentalInst) GetGrossMarginTTMOk() (*float64, bool)`

GetGrossMarginTTMOk returns a tuple with the GrossMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginTTM

`func (o *FundamentalInst) SetGrossMarginTTM(v float64)`

SetGrossMarginTTM sets GrossMarginTTM field to given value.

### HasGrossMarginTTM

`func (o *FundamentalInst) HasGrossMarginTTM() bool`

HasGrossMarginTTM returns a boolean if a field has been set.

### GetGrossMarginMRQ

`func (o *FundamentalInst) GetGrossMarginMRQ() float64`

GetGrossMarginMRQ returns the GrossMarginMRQ field if non-nil, zero value otherwise.

### GetGrossMarginMRQOk

`func (o *FundamentalInst) GetGrossMarginMRQOk() (*float64, bool)`

GetGrossMarginMRQOk returns a tuple with the GrossMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossMarginMRQ

`func (o *FundamentalInst) SetGrossMarginMRQ(v float64)`

SetGrossMarginMRQ sets GrossMarginMRQ field to given value.

### HasGrossMarginMRQ

`func (o *FundamentalInst) HasGrossMarginMRQ() bool`

HasGrossMarginMRQ returns a boolean if a field has been set.

### GetNetProfitMarginTTM

`func (o *FundamentalInst) GetNetProfitMarginTTM() float64`

GetNetProfitMarginTTM returns the NetProfitMarginTTM field if non-nil, zero value otherwise.

### GetNetProfitMarginTTMOk

`func (o *FundamentalInst) GetNetProfitMarginTTMOk() (*float64, bool)`

GetNetProfitMarginTTMOk returns a tuple with the NetProfitMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginTTM

`func (o *FundamentalInst) SetNetProfitMarginTTM(v float64)`

SetNetProfitMarginTTM sets NetProfitMarginTTM field to given value.

### HasNetProfitMarginTTM

`func (o *FundamentalInst) HasNetProfitMarginTTM() bool`

HasNetProfitMarginTTM returns a boolean if a field has been set.

### GetNetProfitMarginMRQ

`func (o *FundamentalInst) GetNetProfitMarginMRQ() float64`

GetNetProfitMarginMRQ returns the NetProfitMarginMRQ field if non-nil, zero value otherwise.

### GetNetProfitMarginMRQOk

`func (o *FundamentalInst) GetNetProfitMarginMRQOk() (*float64, bool)`

GetNetProfitMarginMRQOk returns a tuple with the NetProfitMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfitMarginMRQ

`func (o *FundamentalInst) SetNetProfitMarginMRQ(v float64)`

SetNetProfitMarginMRQ sets NetProfitMarginMRQ field to given value.

### HasNetProfitMarginMRQ

`func (o *FundamentalInst) HasNetProfitMarginMRQ() bool`

HasNetProfitMarginMRQ returns a boolean if a field has been set.

### GetOperatingMarginTTM

`func (o *FundamentalInst) GetOperatingMarginTTM() float64`

GetOperatingMarginTTM returns the OperatingMarginTTM field if non-nil, zero value otherwise.

### GetOperatingMarginTTMOk

`func (o *FundamentalInst) GetOperatingMarginTTMOk() (*float64, bool)`

GetOperatingMarginTTMOk returns a tuple with the OperatingMarginTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginTTM

`func (o *FundamentalInst) SetOperatingMarginTTM(v float64)`

SetOperatingMarginTTM sets OperatingMarginTTM field to given value.

### HasOperatingMarginTTM

`func (o *FundamentalInst) HasOperatingMarginTTM() bool`

HasOperatingMarginTTM returns a boolean if a field has been set.

### GetOperatingMarginMRQ

`func (o *FundamentalInst) GetOperatingMarginMRQ() float64`

GetOperatingMarginMRQ returns the OperatingMarginMRQ field if non-nil, zero value otherwise.

### GetOperatingMarginMRQOk

`func (o *FundamentalInst) GetOperatingMarginMRQOk() (*float64, bool)`

GetOperatingMarginMRQOk returns a tuple with the OperatingMarginMRQ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMarginMRQ

`func (o *FundamentalInst) SetOperatingMarginMRQ(v float64)`

SetOperatingMarginMRQ sets OperatingMarginMRQ field to given value.

### HasOperatingMarginMRQ

`func (o *FundamentalInst) HasOperatingMarginMRQ() bool`

HasOperatingMarginMRQ returns a boolean if a field has been set.

### GetReturnOnEquity

`func (o *FundamentalInst) GetReturnOnEquity() float64`

GetReturnOnEquity returns the ReturnOnEquity field if non-nil, zero value otherwise.

### GetReturnOnEquityOk

`func (o *FundamentalInst) GetReturnOnEquityOk() (*float64, bool)`

GetReturnOnEquityOk returns a tuple with the ReturnOnEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnEquity

`func (o *FundamentalInst) SetReturnOnEquity(v float64)`

SetReturnOnEquity sets ReturnOnEquity field to given value.

### HasReturnOnEquity

`func (o *FundamentalInst) HasReturnOnEquity() bool`

HasReturnOnEquity returns a boolean if a field has been set.

### GetReturnOnAssets

`func (o *FundamentalInst) GetReturnOnAssets() float64`

GetReturnOnAssets returns the ReturnOnAssets field if non-nil, zero value otherwise.

### GetReturnOnAssetsOk

`func (o *FundamentalInst) GetReturnOnAssetsOk() (*float64, bool)`

GetReturnOnAssetsOk returns a tuple with the ReturnOnAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnAssets

`func (o *FundamentalInst) SetReturnOnAssets(v float64)`

SetReturnOnAssets sets ReturnOnAssets field to given value.

### HasReturnOnAssets

`func (o *FundamentalInst) HasReturnOnAssets() bool`

HasReturnOnAssets returns a boolean if a field has been set.

### GetReturnOnInvestment

`func (o *FundamentalInst) GetReturnOnInvestment() float64`

GetReturnOnInvestment returns the ReturnOnInvestment field if non-nil, zero value otherwise.

### GetReturnOnInvestmentOk

`func (o *FundamentalInst) GetReturnOnInvestmentOk() (*float64, bool)`

GetReturnOnInvestmentOk returns a tuple with the ReturnOnInvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOnInvestment

`func (o *FundamentalInst) SetReturnOnInvestment(v float64)`

SetReturnOnInvestment sets ReturnOnInvestment field to given value.

### HasReturnOnInvestment

`func (o *FundamentalInst) HasReturnOnInvestment() bool`

HasReturnOnInvestment returns a boolean if a field has been set.

### GetQuickRatio

`func (o *FundamentalInst) GetQuickRatio() float64`

GetQuickRatio returns the QuickRatio field if non-nil, zero value otherwise.

### GetQuickRatioOk

`func (o *FundamentalInst) GetQuickRatioOk() (*float64, bool)`

GetQuickRatioOk returns a tuple with the QuickRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickRatio

`func (o *FundamentalInst) SetQuickRatio(v float64)`

SetQuickRatio sets QuickRatio field to given value.

### HasQuickRatio

`func (o *FundamentalInst) HasQuickRatio() bool`

HasQuickRatio returns a boolean if a field has been set.

### GetCurrentRatio

`func (o *FundamentalInst) GetCurrentRatio() float64`

GetCurrentRatio returns the CurrentRatio field if non-nil, zero value otherwise.

### GetCurrentRatioOk

`func (o *FundamentalInst) GetCurrentRatioOk() (*float64, bool)`

GetCurrentRatioOk returns a tuple with the CurrentRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRatio

`func (o *FundamentalInst) SetCurrentRatio(v float64)`

SetCurrentRatio sets CurrentRatio field to given value.

### HasCurrentRatio

`func (o *FundamentalInst) HasCurrentRatio() bool`

HasCurrentRatio returns a boolean if a field has been set.

### GetInterestCoverage

`func (o *FundamentalInst) GetInterestCoverage() float64`

GetInterestCoverage returns the InterestCoverage field if non-nil, zero value otherwise.

### GetInterestCoverageOk

`func (o *FundamentalInst) GetInterestCoverageOk() (*float64, bool)`

GetInterestCoverageOk returns a tuple with the InterestCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCoverage

`func (o *FundamentalInst) SetInterestCoverage(v float64)`

SetInterestCoverage sets InterestCoverage field to given value.

### HasInterestCoverage

`func (o *FundamentalInst) HasInterestCoverage() bool`

HasInterestCoverage returns a boolean if a field has been set.

### GetTotalDebtToCapital

`func (o *FundamentalInst) GetTotalDebtToCapital() float64`

GetTotalDebtToCapital returns the TotalDebtToCapital field if non-nil, zero value otherwise.

### GetTotalDebtToCapitalOk

`func (o *FundamentalInst) GetTotalDebtToCapitalOk() (*float64, bool)`

GetTotalDebtToCapitalOk returns a tuple with the TotalDebtToCapital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToCapital

`func (o *FundamentalInst) SetTotalDebtToCapital(v float64)`

SetTotalDebtToCapital sets TotalDebtToCapital field to given value.

### HasTotalDebtToCapital

`func (o *FundamentalInst) HasTotalDebtToCapital() bool`

HasTotalDebtToCapital returns a boolean if a field has been set.

### GetLtDebtToEquity

`func (o *FundamentalInst) GetLtDebtToEquity() float64`

GetLtDebtToEquity returns the LtDebtToEquity field if non-nil, zero value otherwise.

### GetLtDebtToEquityOk

`func (o *FundamentalInst) GetLtDebtToEquityOk() (*float64, bool)`

GetLtDebtToEquityOk returns a tuple with the LtDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtDebtToEquity

`func (o *FundamentalInst) SetLtDebtToEquity(v float64)`

SetLtDebtToEquity sets LtDebtToEquity field to given value.

### HasLtDebtToEquity

`func (o *FundamentalInst) HasLtDebtToEquity() bool`

HasLtDebtToEquity returns a boolean if a field has been set.

### GetTotalDebtToEquity

`func (o *FundamentalInst) GetTotalDebtToEquity() float64`

GetTotalDebtToEquity returns the TotalDebtToEquity field if non-nil, zero value otherwise.

### GetTotalDebtToEquityOk

`func (o *FundamentalInst) GetTotalDebtToEquityOk() (*float64, bool)`

GetTotalDebtToEquityOk returns a tuple with the TotalDebtToEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtToEquity

`func (o *FundamentalInst) SetTotalDebtToEquity(v float64)`

SetTotalDebtToEquity sets TotalDebtToEquity field to given value.

### HasTotalDebtToEquity

`func (o *FundamentalInst) HasTotalDebtToEquity() bool`

HasTotalDebtToEquity returns a boolean if a field has been set.

### GetEpsTTM

`func (o *FundamentalInst) GetEpsTTM() float64`

GetEpsTTM returns the EpsTTM field if non-nil, zero value otherwise.

### GetEpsTTMOk

`func (o *FundamentalInst) GetEpsTTMOk() (*float64, bool)`

GetEpsTTMOk returns a tuple with the EpsTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsTTM

`func (o *FundamentalInst) SetEpsTTM(v float64)`

SetEpsTTM sets EpsTTM field to given value.

### HasEpsTTM

`func (o *FundamentalInst) HasEpsTTM() bool`

HasEpsTTM returns a boolean if a field has been set.

### GetEpsChangePercentTTM

`func (o *FundamentalInst) GetEpsChangePercentTTM() float64`

GetEpsChangePercentTTM returns the EpsChangePercentTTM field if non-nil, zero value otherwise.

### GetEpsChangePercentTTMOk

`func (o *FundamentalInst) GetEpsChangePercentTTMOk() (*float64, bool)`

GetEpsChangePercentTTMOk returns a tuple with the EpsChangePercentTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangePercentTTM

`func (o *FundamentalInst) SetEpsChangePercentTTM(v float64)`

SetEpsChangePercentTTM sets EpsChangePercentTTM field to given value.

### HasEpsChangePercentTTM

`func (o *FundamentalInst) HasEpsChangePercentTTM() bool`

HasEpsChangePercentTTM returns a boolean if a field has been set.

### GetEpsChangeYear

`func (o *FundamentalInst) GetEpsChangeYear() float64`

GetEpsChangeYear returns the EpsChangeYear field if non-nil, zero value otherwise.

### GetEpsChangeYearOk

`func (o *FundamentalInst) GetEpsChangeYearOk() (*float64, bool)`

GetEpsChangeYearOk returns a tuple with the EpsChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChangeYear

`func (o *FundamentalInst) SetEpsChangeYear(v float64)`

SetEpsChangeYear sets EpsChangeYear field to given value.

### HasEpsChangeYear

`func (o *FundamentalInst) HasEpsChangeYear() bool`

HasEpsChangeYear returns a boolean if a field has been set.

### GetEpsChange

`func (o *FundamentalInst) GetEpsChange() float64`

GetEpsChange returns the EpsChange field if non-nil, zero value otherwise.

### GetEpsChangeOk

`func (o *FundamentalInst) GetEpsChangeOk() (*float64, bool)`

GetEpsChangeOk returns a tuple with the EpsChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsChange

`func (o *FundamentalInst) SetEpsChange(v float64)`

SetEpsChange sets EpsChange field to given value.

### HasEpsChange

`func (o *FundamentalInst) HasEpsChange() bool`

HasEpsChange returns a boolean if a field has been set.

### GetRevChangeYear

`func (o *FundamentalInst) GetRevChangeYear() float64`

GetRevChangeYear returns the RevChangeYear field if non-nil, zero value otherwise.

### GetRevChangeYearOk

`func (o *FundamentalInst) GetRevChangeYearOk() (*float64, bool)`

GetRevChangeYearOk returns a tuple with the RevChangeYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeYear

`func (o *FundamentalInst) SetRevChangeYear(v float64)`

SetRevChangeYear sets RevChangeYear field to given value.

### HasRevChangeYear

`func (o *FundamentalInst) HasRevChangeYear() bool`

HasRevChangeYear returns a boolean if a field has been set.

### GetRevChangeTTM

`func (o *FundamentalInst) GetRevChangeTTM() float64`

GetRevChangeTTM returns the RevChangeTTM field if non-nil, zero value otherwise.

### GetRevChangeTTMOk

`func (o *FundamentalInst) GetRevChangeTTMOk() (*float64, bool)`

GetRevChangeTTMOk returns a tuple with the RevChangeTTM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeTTM

`func (o *FundamentalInst) SetRevChangeTTM(v float64)`

SetRevChangeTTM sets RevChangeTTM field to given value.

### HasRevChangeTTM

`func (o *FundamentalInst) HasRevChangeTTM() bool`

HasRevChangeTTM returns a boolean if a field has been set.

### GetRevChangeIn

`func (o *FundamentalInst) GetRevChangeIn() float64`

GetRevChangeIn returns the RevChangeIn field if non-nil, zero value otherwise.

### GetRevChangeInOk

`func (o *FundamentalInst) GetRevChangeInOk() (*float64, bool)`

GetRevChangeInOk returns a tuple with the RevChangeIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevChangeIn

`func (o *FundamentalInst) SetRevChangeIn(v float64)`

SetRevChangeIn sets RevChangeIn field to given value.

### HasRevChangeIn

`func (o *FundamentalInst) HasRevChangeIn() bool`

HasRevChangeIn returns a boolean if a field has been set.

### GetSharesOutstanding

`func (o *FundamentalInst) GetSharesOutstanding() float64`

GetSharesOutstanding returns the SharesOutstanding field if non-nil, zero value otherwise.

### GetSharesOutstandingOk

`func (o *FundamentalInst) GetSharesOutstandingOk() (*float64, bool)`

GetSharesOutstandingOk returns a tuple with the SharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesOutstanding

`func (o *FundamentalInst) SetSharesOutstanding(v float64)`

SetSharesOutstanding sets SharesOutstanding field to given value.

### HasSharesOutstanding

`func (o *FundamentalInst) HasSharesOutstanding() bool`

HasSharesOutstanding returns a boolean if a field has been set.

### GetMarketCapFloat

`func (o *FundamentalInst) GetMarketCapFloat() float64`

GetMarketCapFloat returns the MarketCapFloat field if non-nil, zero value otherwise.

### GetMarketCapFloatOk

`func (o *FundamentalInst) GetMarketCapFloatOk() (*float64, bool)`

GetMarketCapFloatOk returns a tuple with the MarketCapFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCapFloat

`func (o *FundamentalInst) SetMarketCapFloat(v float64)`

SetMarketCapFloat sets MarketCapFloat field to given value.

### HasMarketCapFloat

`func (o *FundamentalInst) HasMarketCapFloat() bool`

HasMarketCapFloat returns a boolean if a field has been set.

### GetMarketCap

`func (o *FundamentalInst) GetMarketCap() float64`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *FundamentalInst) GetMarketCapOk() (*float64, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *FundamentalInst) SetMarketCap(v float64)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *FundamentalInst) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetBookValuePerShare

`func (o *FundamentalInst) GetBookValuePerShare() float64`

GetBookValuePerShare returns the BookValuePerShare field if non-nil, zero value otherwise.

### GetBookValuePerShareOk

`func (o *FundamentalInst) GetBookValuePerShareOk() (*float64, bool)`

GetBookValuePerShareOk returns a tuple with the BookValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValuePerShare

`func (o *FundamentalInst) SetBookValuePerShare(v float64)`

SetBookValuePerShare sets BookValuePerShare field to given value.

### HasBookValuePerShare

`func (o *FundamentalInst) HasBookValuePerShare() bool`

HasBookValuePerShare returns a boolean if a field has been set.

### GetShortIntToFloat

`func (o *FundamentalInst) GetShortIntToFloat() float64`

GetShortIntToFloat returns the ShortIntToFloat field if non-nil, zero value otherwise.

### GetShortIntToFloatOk

`func (o *FundamentalInst) GetShortIntToFloatOk() (*float64, bool)`

GetShortIntToFloatOk returns a tuple with the ShortIntToFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntToFloat

`func (o *FundamentalInst) SetShortIntToFloat(v float64)`

SetShortIntToFloat sets ShortIntToFloat field to given value.

### HasShortIntToFloat

`func (o *FundamentalInst) HasShortIntToFloat() bool`

HasShortIntToFloat returns a boolean if a field has been set.

### GetShortIntDayToCover

`func (o *FundamentalInst) GetShortIntDayToCover() float64`

GetShortIntDayToCover returns the ShortIntDayToCover field if non-nil, zero value otherwise.

### GetShortIntDayToCoverOk

`func (o *FundamentalInst) GetShortIntDayToCoverOk() (*float64, bool)`

GetShortIntDayToCoverOk returns a tuple with the ShortIntDayToCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortIntDayToCover

`func (o *FundamentalInst) SetShortIntDayToCover(v float64)`

SetShortIntDayToCover sets ShortIntDayToCover field to given value.

### HasShortIntDayToCover

`func (o *FundamentalInst) HasShortIntDayToCover() bool`

HasShortIntDayToCover returns a boolean if a field has been set.

### GetDivGrowthRate3Year

`func (o *FundamentalInst) GetDivGrowthRate3Year() float64`

GetDivGrowthRate3Year returns the DivGrowthRate3Year field if non-nil, zero value otherwise.

### GetDivGrowthRate3YearOk

`func (o *FundamentalInst) GetDivGrowthRate3YearOk() (*float64, bool)`

GetDivGrowthRate3YearOk returns a tuple with the DivGrowthRate3Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivGrowthRate3Year

`func (o *FundamentalInst) SetDivGrowthRate3Year(v float64)`

SetDivGrowthRate3Year sets DivGrowthRate3Year field to given value.

### HasDivGrowthRate3Year

`func (o *FundamentalInst) HasDivGrowthRate3Year() bool`

HasDivGrowthRate3Year returns a boolean if a field has been set.

### GetDividendPayAmount

`func (o *FundamentalInst) GetDividendPayAmount() float64`

GetDividendPayAmount returns the DividendPayAmount field if non-nil, zero value otherwise.

### GetDividendPayAmountOk

`func (o *FundamentalInst) GetDividendPayAmountOk() (*float64, bool)`

GetDividendPayAmountOk returns a tuple with the DividendPayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayAmount

`func (o *FundamentalInst) SetDividendPayAmount(v float64)`

SetDividendPayAmount sets DividendPayAmount field to given value.

### HasDividendPayAmount

`func (o *FundamentalInst) HasDividendPayAmount() bool`

HasDividendPayAmount returns a boolean if a field has been set.

### GetDividendPayDate

`func (o *FundamentalInst) GetDividendPayDate() string`

GetDividendPayDate returns the DividendPayDate field if non-nil, zero value otherwise.

### GetDividendPayDateOk

`func (o *FundamentalInst) GetDividendPayDateOk() (*string, bool)`

GetDividendPayDateOk returns a tuple with the DividendPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendPayDate

`func (o *FundamentalInst) SetDividendPayDate(v string)`

SetDividendPayDate sets DividendPayDate field to given value.

### HasDividendPayDate

`func (o *FundamentalInst) HasDividendPayDate() bool`

HasDividendPayDate returns a boolean if a field has been set.

### GetBeta

`func (o *FundamentalInst) GetBeta() float64`

GetBeta returns the Beta field if non-nil, zero value otherwise.

### GetBetaOk

`func (o *FundamentalInst) GetBetaOk() (*float64, bool)`

GetBetaOk returns a tuple with the Beta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeta

`func (o *FundamentalInst) SetBeta(v float64)`

SetBeta sets Beta field to given value.

### HasBeta

`func (o *FundamentalInst) HasBeta() bool`

HasBeta returns a boolean if a field has been set.

### GetVol1DayAvg

`func (o *FundamentalInst) GetVol1DayAvg() float64`

GetVol1DayAvg returns the Vol1DayAvg field if non-nil, zero value otherwise.

### GetVol1DayAvgOk

`func (o *FundamentalInst) GetVol1DayAvgOk() (*float64, bool)`

GetVol1DayAvgOk returns a tuple with the Vol1DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol1DayAvg

`func (o *FundamentalInst) SetVol1DayAvg(v float64)`

SetVol1DayAvg sets Vol1DayAvg field to given value.

### HasVol1DayAvg

`func (o *FundamentalInst) HasVol1DayAvg() bool`

HasVol1DayAvg returns a boolean if a field has been set.

### GetVol10DayAvg

`func (o *FundamentalInst) GetVol10DayAvg() float64`

GetVol10DayAvg returns the Vol10DayAvg field if non-nil, zero value otherwise.

### GetVol10DayAvgOk

`func (o *FundamentalInst) GetVol10DayAvgOk() (*float64, bool)`

GetVol10DayAvgOk returns a tuple with the Vol10DayAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol10DayAvg

`func (o *FundamentalInst) SetVol10DayAvg(v float64)`

SetVol10DayAvg sets Vol10DayAvg field to given value.

### HasVol10DayAvg

`func (o *FundamentalInst) HasVol10DayAvg() bool`

HasVol10DayAvg returns a boolean if a field has been set.

### GetVol3MonthAvg

`func (o *FundamentalInst) GetVol3MonthAvg() float64`

GetVol3MonthAvg returns the Vol3MonthAvg field if non-nil, zero value otherwise.

### GetVol3MonthAvgOk

`func (o *FundamentalInst) GetVol3MonthAvgOk() (*float64, bool)`

GetVol3MonthAvgOk returns a tuple with the Vol3MonthAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVol3MonthAvg

`func (o *FundamentalInst) SetVol3MonthAvg(v float64)`

SetVol3MonthAvg sets Vol3MonthAvg field to given value.

### HasVol3MonthAvg

`func (o *FundamentalInst) HasVol3MonthAvg() bool`

HasVol3MonthAvg returns a boolean if a field has been set.

### GetAvg10DaysVolume

`func (o *FundamentalInst) GetAvg10DaysVolume() int64`

GetAvg10DaysVolume returns the Avg10DaysVolume field if non-nil, zero value otherwise.

### GetAvg10DaysVolumeOk

`func (o *FundamentalInst) GetAvg10DaysVolumeOk() (*int64, bool)`

GetAvg10DaysVolumeOk returns a tuple with the Avg10DaysVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg10DaysVolume

`func (o *FundamentalInst) SetAvg10DaysVolume(v int64)`

SetAvg10DaysVolume sets Avg10DaysVolume field to given value.

### HasAvg10DaysVolume

`func (o *FundamentalInst) HasAvg10DaysVolume() bool`

HasAvg10DaysVolume returns a boolean if a field has been set.

### GetAvg1DayVolume

`func (o *FundamentalInst) GetAvg1DayVolume() int64`

GetAvg1DayVolume returns the Avg1DayVolume field if non-nil, zero value otherwise.

### GetAvg1DayVolumeOk

`func (o *FundamentalInst) GetAvg1DayVolumeOk() (*int64, bool)`

GetAvg1DayVolumeOk returns a tuple with the Avg1DayVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg1DayVolume

`func (o *FundamentalInst) SetAvg1DayVolume(v int64)`

SetAvg1DayVolume sets Avg1DayVolume field to given value.

### HasAvg1DayVolume

`func (o *FundamentalInst) HasAvg1DayVolume() bool`

HasAvg1DayVolume returns a boolean if a field has been set.

### GetAvg3MonthVolume

`func (o *FundamentalInst) GetAvg3MonthVolume() int64`

GetAvg3MonthVolume returns the Avg3MonthVolume field if non-nil, zero value otherwise.

### GetAvg3MonthVolumeOk

`func (o *FundamentalInst) GetAvg3MonthVolumeOk() (*int64, bool)`

GetAvg3MonthVolumeOk returns a tuple with the Avg3MonthVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg3MonthVolume

`func (o *FundamentalInst) SetAvg3MonthVolume(v int64)`

SetAvg3MonthVolume sets Avg3MonthVolume field to given value.

### HasAvg3MonthVolume

`func (o *FundamentalInst) HasAvg3MonthVolume() bool`

HasAvg3MonthVolume returns a boolean if a field has been set.

### GetDeclarationDate

`func (o *FundamentalInst) GetDeclarationDate() string`

GetDeclarationDate returns the DeclarationDate field if non-nil, zero value otherwise.

### GetDeclarationDateOk

`func (o *FundamentalInst) GetDeclarationDateOk() (*string, bool)`

GetDeclarationDateOk returns a tuple with the DeclarationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclarationDate

`func (o *FundamentalInst) SetDeclarationDate(v string)`

SetDeclarationDate sets DeclarationDate field to given value.

### HasDeclarationDate

`func (o *FundamentalInst) HasDeclarationDate() bool`

HasDeclarationDate returns a boolean if a field has been set.

### GetDividendFreq

`func (o *FundamentalInst) GetDividendFreq() int32`

GetDividendFreq returns the DividendFreq field if non-nil, zero value otherwise.

### GetDividendFreqOk

`func (o *FundamentalInst) GetDividendFreqOk() (*int32, bool)`

GetDividendFreqOk returns a tuple with the DividendFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendFreq

`func (o *FundamentalInst) SetDividendFreq(v int32)`

SetDividendFreq sets DividendFreq field to given value.

### HasDividendFreq

`func (o *FundamentalInst) HasDividendFreq() bool`

HasDividendFreq returns a boolean if a field has been set.

### GetEps

`func (o *FundamentalInst) GetEps() float64`

GetEps returns the Eps field if non-nil, zero value otherwise.

### GetEpsOk

`func (o *FundamentalInst) GetEpsOk() (*float64, bool)`

GetEpsOk returns a tuple with the Eps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEps

`func (o *FundamentalInst) SetEps(v float64)`

SetEps sets Eps field to given value.

### HasEps

`func (o *FundamentalInst) HasEps() bool`

HasEps returns a boolean if a field has been set.

### GetCorpactionDate

`func (o *FundamentalInst) GetCorpactionDate() string`

GetCorpactionDate returns the CorpactionDate field if non-nil, zero value otherwise.

### GetCorpactionDateOk

`func (o *FundamentalInst) GetCorpactionDateOk() (*string, bool)`

GetCorpactionDateOk returns a tuple with the CorpactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorpactionDate

`func (o *FundamentalInst) SetCorpactionDate(v string)`

SetCorpactionDate sets CorpactionDate field to given value.

### HasCorpactionDate

`func (o *FundamentalInst) HasCorpactionDate() bool`

HasCorpactionDate returns a boolean if a field has been set.

### GetDtnVolume

`func (o *FundamentalInst) GetDtnVolume() int64`

GetDtnVolume returns the DtnVolume field if non-nil, zero value otherwise.

### GetDtnVolumeOk

`func (o *FundamentalInst) GetDtnVolumeOk() (*int64, bool)`

GetDtnVolumeOk returns a tuple with the DtnVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtnVolume

`func (o *FundamentalInst) SetDtnVolume(v int64)`

SetDtnVolume sets DtnVolume field to given value.

### HasDtnVolume

`func (o *FundamentalInst) HasDtnVolume() bool`

HasDtnVolume returns a boolean if a field has been set.

### GetNextDividendPayDate

`func (o *FundamentalInst) GetNextDividendPayDate() string`

GetNextDividendPayDate returns the NextDividendPayDate field if non-nil, zero value otherwise.

### GetNextDividendPayDateOk

`func (o *FundamentalInst) GetNextDividendPayDateOk() (*string, bool)`

GetNextDividendPayDateOk returns a tuple with the NextDividendPayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDividendPayDate

`func (o *FundamentalInst) SetNextDividendPayDate(v string)`

SetNextDividendPayDate sets NextDividendPayDate field to given value.

### HasNextDividendPayDate

`func (o *FundamentalInst) HasNextDividendPayDate() bool`

HasNextDividendPayDate returns a boolean if a field has been set.

### GetNextDividendDate

`func (o *FundamentalInst) GetNextDividendDate() string`

GetNextDividendDate returns the NextDividendDate field if non-nil, zero value otherwise.

### GetNextDividendDateOk

`func (o *FundamentalInst) GetNextDividendDateOk() (*string, bool)`

GetNextDividendDateOk returns a tuple with the NextDividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDividendDate

`func (o *FundamentalInst) SetNextDividendDate(v string)`

SetNextDividendDate sets NextDividendDate field to given value.

### HasNextDividendDate

`func (o *FundamentalInst) HasNextDividendDate() bool`

HasNextDividendDate returns a boolean if a field has been set.

### GetFundLeverageFactor

`func (o *FundamentalInst) GetFundLeverageFactor() float64`

GetFundLeverageFactor returns the FundLeverageFactor field if non-nil, zero value otherwise.

### GetFundLeverageFactorOk

`func (o *FundamentalInst) GetFundLeverageFactorOk() (*float64, bool)`

GetFundLeverageFactorOk returns a tuple with the FundLeverageFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundLeverageFactor

`func (o *FundamentalInst) SetFundLeverageFactor(v float64)`

SetFundLeverageFactor sets FundLeverageFactor field to given value.

### HasFundLeverageFactor

`func (o *FundamentalInst) HasFundLeverageFactor() bool`

HasFundLeverageFactor returns a boolean if a field has been set.

### GetFundStrategy

`func (o *FundamentalInst) GetFundStrategy() string`

GetFundStrategy returns the FundStrategy field if non-nil, zero value otherwise.

### GetFundStrategyOk

`func (o *FundamentalInst) GetFundStrategyOk() (*string, bool)`

GetFundStrategyOk returns a tuple with the FundStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundStrategy

`func (o *FundamentalInst) SetFundStrategy(v string)`

SetFundStrategy sets FundStrategy field to given value.

### HasFundStrategy

`func (o *FundamentalInst) HasFundStrategy() bool`

HasFundStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


