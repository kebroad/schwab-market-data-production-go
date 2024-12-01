# OptionContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PutCall** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **float64** |  | [optional] 
**AskPrice** | Pointer to **float64** |  | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**MarkPrice** | Pointer to **float64** |  | [optional] 
**BidSize** | Pointer to **int32** |  | [optional] 
**AskSize** | Pointer to **int32** |  | [optional] 
**LastSize** | Pointer to **int32** |  | [optional] 
**HighPrice** | Pointer to **float64** |  | [optional] 
**LowPrice** | Pointer to **float64** |  | [optional] 
**OpenPrice** | Pointer to **float64** |  | [optional] 
**ClosePrice** | Pointer to **float64** |  | [optional] 
**TotalVolume** | Pointer to **int32** |  | [optional] 
**TradeDate** | Pointer to **float32** |  | [optional] 
**QuoteTimeInLong** | Pointer to **int32** |  | [optional] 
**TradeTimeInLong** | Pointer to **int32** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 
**Volatility** | Pointer to **float64** |  | [optional] 
**Delta** | Pointer to **float64** |  | [optional] 
**Gamma** | Pointer to **float64** |  | [optional] 
**Theta** | Pointer to **float64** |  | [optional] 
**Vega** | Pointer to **float64** |  | [optional] 
**Rho** | Pointer to **float64** |  | [optional] 
**TimeValue** | Pointer to **float64** |  | [optional] 
**OpenInterest** | Pointer to **float64** |  | [optional] 
**IsInTheMoney** | Pointer to **bool** |  | [optional] 
**TheoreticalOptionValue** | Pointer to **float64** |  | [optional] 
**TheoreticalVolatility** | Pointer to **float64** |  | [optional] 
**IsMini** | Pointer to **bool** |  | [optional] 
**IsNonStandard** | Pointer to **bool** |  | [optional] 
**OptionDeliverablesList** | Pointer to [**[]OptionDeliverables**](OptionDeliverables.md) |  | [optional] 
**StrikePrice** | Pointer to **float64** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**DaysToExpiration** | Pointer to **float32** |  | [optional] 
**ExpirationType** | Pointer to [**ExpirationType**](ExpirationType.md) |  | [optional] 
**LastTradingDay** | Pointer to **float32** |  | [optional] 
**Multiplier** | Pointer to **float64** |  | [optional] 
**SettlementType** | Pointer to [**SettlementType**](SettlementType.md) |  | [optional] 
**DeliverableNote** | Pointer to **string** |  | [optional] 
**IsIndexOption** | Pointer to **bool** |  | [optional] 
**PercentChange** | Pointer to **float64** |  | [optional] 
**MarkChange** | Pointer to **float64** |  | [optional] 
**MarkPercentChange** | Pointer to **float64** |  | [optional] 
**IsPennyPilot** | Pointer to **bool** |  | [optional] 
**IntrinsicValue** | Pointer to **float64** |  | [optional] 
**OptionRoot** | Pointer to **string** |  | [optional] 

## Methods

### NewOptionContract

`func NewOptionContract() *OptionContract`

NewOptionContract instantiates a new OptionContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionContractWithDefaults

`func NewOptionContractWithDefaults() *OptionContract`

NewOptionContractWithDefaults instantiates a new OptionContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPutCall

`func (o *OptionContract) GetPutCall() string`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *OptionContract) GetPutCallOk() (*string, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *OptionContract) SetPutCall(v string)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *OptionContract) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetSymbol

`func (o *OptionContract) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OptionContract) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OptionContract) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OptionContract) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *OptionContract) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionContract) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionContract) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionContract) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeName

`func (o *OptionContract) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *OptionContract) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *OptionContract) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *OptionContract) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetBidPrice

`func (o *OptionContract) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *OptionContract) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *OptionContract) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *OptionContract) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetAskPrice

`func (o *OptionContract) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *OptionContract) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *OptionContract) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *OptionContract) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *OptionContract) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *OptionContract) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *OptionContract) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *OptionContract) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetMarkPrice

`func (o *OptionContract) GetMarkPrice() float64`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *OptionContract) GetMarkPriceOk() (*float64, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *OptionContract) SetMarkPrice(v float64)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *OptionContract) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *OptionContract) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *OptionContract) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *OptionContract) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *OptionContract) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetAskSize

`func (o *OptionContract) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *OptionContract) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *OptionContract) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *OptionContract) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetLastSize

`func (o *OptionContract) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *OptionContract) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *OptionContract) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *OptionContract) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetHighPrice

`func (o *OptionContract) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *OptionContract) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *OptionContract) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *OptionContract) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *OptionContract) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *OptionContract) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *OptionContract) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *OptionContract) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetOpenPrice

`func (o *OptionContract) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *OptionContract) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *OptionContract) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *OptionContract) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetClosePrice

`func (o *OptionContract) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *OptionContract) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *OptionContract) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *OptionContract) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetTotalVolume

`func (o *OptionContract) GetTotalVolume() int32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *OptionContract) GetTotalVolumeOk() (*int32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *OptionContract) SetTotalVolume(v int32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *OptionContract) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeDate

`func (o *OptionContract) GetTradeDate() float32`

GetTradeDate returns the TradeDate field if non-nil, zero value otherwise.

### GetTradeDateOk

`func (o *OptionContract) GetTradeDateOk() (*float32, bool)`

GetTradeDateOk returns a tuple with the TradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDate

`func (o *OptionContract) SetTradeDate(v float32)`

SetTradeDate sets TradeDate field to given value.

### HasTradeDate

`func (o *OptionContract) HasTradeDate() bool`

HasTradeDate returns a boolean if a field has been set.

### GetQuoteTimeInLong

`func (o *OptionContract) GetQuoteTimeInLong() int32`

GetQuoteTimeInLong returns the QuoteTimeInLong field if non-nil, zero value otherwise.

### GetQuoteTimeInLongOk

`func (o *OptionContract) GetQuoteTimeInLongOk() (*int32, bool)`

GetQuoteTimeInLongOk returns a tuple with the QuoteTimeInLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTimeInLong

`func (o *OptionContract) SetQuoteTimeInLong(v int32)`

SetQuoteTimeInLong sets QuoteTimeInLong field to given value.

### HasQuoteTimeInLong

`func (o *OptionContract) HasQuoteTimeInLong() bool`

HasQuoteTimeInLong returns a boolean if a field has been set.

### GetTradeTimeInLong

`func (o *OptionContract) GetTradeTimeInLong() int32`

GetTradeTimeInLong returns the TradeTimeInLong field if non-nil, zero value otherwise.

### GetTradeTimeInLongOk

`func (o *OptionContract) GetTradeTimeInLongOk() (*int32, bool)`

GetTradeTimeInLongOk returns a tuple with the TradeTimeInLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTimeInLong

`func (o *OptionContract) SetTradeTimeInLong(v int32)`

SetTradeTimeInLong sets TradeTimeInLong field to given value.

### HasTradeTimeInLong

`func (o *OptionContract) HasTradeTimeInLong() bool`

HasTradeTimeInLong returns a boolean if a field has been set.

### GetNetChange

`func (o *OptionContract) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *OptionContract) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *OptionContract) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *OptionContract) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetVolatility

`func (o *OptionContract) GetVolatility() float64`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *OptionContract) GetVolatilityOk() (*float64, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *OptionContract) SetVolatility(v float64)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *OptionContract) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.

### GetDelta

`func (o *OptionContract) GetDelta() float64`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *OptionContract) GetDeltaOk() (*float64, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *OptionContract) SetDelta(v float64)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *OptionContract) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetGamma

`func (o *OptionContract) GetGamma() float64`

GetGamma returns the Gamma field if non-nil, zero value otherwise.

### GetGammaOk

`func (o *OptionContract) GetGammaOk() (*float64, bool)`

GetGammaOk returns a tuple with the Gamma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamma

`func (o *OptionContract) SetGamma(v float64)`

SetGamma sets Gamma field to given value.

### HasGamma

`func (o *OptionContract) HasGamma() bool`

HasGamma returns a boolean if a field has been set.

### GetTheta

`func (o *OptionContract) GetTheta() float64`

GetTheta returns the Theta field if non-nil, zero value otherwise.

### GetThetaOk

`func (o *OptionContract) GetThetaOk() (*float64, bool)`

GetThetaOk returns a tuple with the Theta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheta

`func (o *OptionContract) SetTheta(v float64)`

SetTheta sets Theta field to given value.

### HasTheta

`func (o *OptionContract) HasTheta() bool`

HasTheta returns a boolean if a field has been set.

### GetVega

`func (o *OptionContract) GetVega() float64`

GetVega returns the Vega field if non-nil, zero value otherwise.

### GetVegaOk

`func (o *OptionContract) GetVegaOk() (*float64, bool)`

GetVegaOk returns a tuple with the Vega field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVega

`func (o *OptionContract) SetVega(v float64)`

SetVega sets Vega field to given value.

### HasVega

`func (o *OptionContract) HasVega() bool`

HasVega returns a boolean if a field has been set.

### GetRho

`func (o *OptionContract) GetRho() float64`

GetRho returns the Rho field if non-nil, zero value otherwise.

### GetRhoOk

`func (o *OptionContract) GetRhoOk() (*float64, bool)`

GetRhoOk returns a tuple with the Rho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRho

`func (o *OptionContract) SetRho(v float64)`

SetRho sets Rho field to given value.

### HasRho

`func (o *OptionContract) HasRho() bool`

HasRho returns a boolean if a field has been set.

### GetTimeValue

`func (o *OptionContract) GetTimeValue() float64`

GetTimeValue returns the TimeValue field if non-nil, zero value otherwise.

### GetTimeValueOk

`func (o *OptionContract) GetTimeValueOk() (*float64, bool)`

GetTimeValueOk returns a tuple with the TimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValue

`func (o *OptionContract) SetTimeValue(v float64)`

SetTimeValue sets TimeValue field to given value.

### HasTimeValue

`func (o *OptionContract) HasTimeValue() bool`

HasTimeValue returns a boolean if a field has been set.

### GetOpenInterest

`func (o *OptionContract) GetOpenInterest() float64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *OptionContract) GetOpenInterestOk() (*float64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *OptionContract) SetOpenInterest(v float64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *OptionContract) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetIsInTheMoney

`func (o *OptionContract) GetIsInTheMoney() bool`

GetIsInTheMoney returns the IsInTheMoney field if non-nil, zero value otherwise.

### GetIsInTheMoneyOk

`func (o *OptionContract) GetIsInTheMoneyOk() (*bool, bool)`

GetIsInTheMoneyOk returns a tuple with the IsInTheMoney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInTheMoney

`func (o *OptionContract) SetIsInTheMoney(v bool)`

SetIsInTheMoney sets IsInTheMoney field to given value.

### HasIsInTheMoney

`func (o *OptionContract) HasIsInTheMoney() bool`

HasIsInTheMoney returns a boolean if a field has been set.

### GetTheoreticalOptionValue

`func (o *OptionContract) GetTheoreticalOptionValue() float64`

GetTheoreticalOptionValue returns the TheoreticalOptionValue field if non-nil, zero value otherwise.

### GetTheoreticalOptionValueOk

`func (o *OptionContract) GetTheoreticalOptionValueOk() (*float64, bool)`

GetTheoreticalOptionValueOk returns a tuple with the TheoreticalOptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheoreticalOptionValue

`func (o *OptionContract) SetTheoreticalOptionValue(v float64)`

SetTheoreticalOptionValue sets TheoreticalOptionValue field to given value.

### HasTheoreticalOptionValue

`func (o *OptionContract) HasTheoreticalOptionValue() bool`

HasTheoreticalOptionValue returns a boolean if a field has been set.

### GetTheoreticalVolatility

`func (o *OptionContract) GetTheoreticalVolatility() float64`

GetTheoreticalVolatility returns the TheoreticalVolatility field if non-nil, zero value otherwise.

### GetTheoreticalVolatilityOk

`func (o *OptionContract) GetTheoreticalVolatilityOk() (*float64, bool)`

GetTheoreticalVolatilityOk returns a tuple with the TheoreticalVolatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheoreticalVolatility

`func (o *OptionContract) SetTheoreticalVolatility(v float64)`

SetTheoreticalVolatility sets TheoreticalVolatility field to given value.

### HasTheoreticalVolatility

`func (o *OptionContract) HasTheoreticalVolatility() bool`

HasTheoreticalVolatility returns a boolean if a field has been set.

### GetIsMini

`func (o *OptionContract) GetIsMini() bool`

GetIsMini returns the IsMini field if non-nil, zero value otherwise.

### GetIsMiniOk

`func (o *OptionContract) GetIsMiniOk() (*bool, bool)`

GetIsMiniOk returns a tuple with the IsMini field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMini

`func (o *OptionContract) SetIsMini(v bool)`

SetIsMini sets IsMini field to given value.

### HasIsMini

`func (o *OptionContract) HasIsMini() bool`

HasIsMini returns a boolean if a field has been set.

### GetIsNonStandard

`func (o *OptionContract) GetIsNonStandard() bool`

GetIsNonStandard returns the IsNonStandard field if non-nil, zero value otherwise.

### GetIsNonStandardOk

`func (o *OptionContract) GetIsNonStandardOk() (*bool, bool)`

GetIsNonStandardOk returns a tuple with the IsNonStandard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonStandard

`func (o *OptionContract) SetIsNonStandard(v bool)`

SetIsNonStandard sets IsNonStandard field to given value.

### HasIsNonStandard

`func (o *OptionContract) HasIsNonStandard() bool`

HasIsNonStandard returns a boolean if a field has been set.

### GetOptionDeliverablesList

`func (o *OptionContract) GetOptionDeliverablesList() []OptionDeliverables`

GetOptionDeliverablesList returns the OptionDeliverablesList field if non-nil, zero value otherwise.

### GetOptionDeliverablesListOk

`func (o *OptionContract) GetOptionDeliverablesListOk() (*[]OptionDeliverables, bool)`

GetOptionDeliverablesListOk returns a tuple with the OptionDeliverablesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeliverablesList

`func (o *OptionContract) SetOptionDeliverablesList(v []OptionDeliverables)`

SetOptionDeliverablesList sets OptionDeliverablesList field to given value.

### HasOptionDeliverablesList

`func (o *OptionContract) HasOptionDeliverablesList() bool`

HasOptionDeliverablesList returns a boolean if a field has been set.

### GetStrikePrice

`func (o *OptionContract) GetStrikePrice() float64`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *OptionContract) GetStrikePriceOk() (*float64, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *OptionContract) SetStrikePrice(v float64)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *OptionContract) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetExpirationDate

`func (o *OptionContract) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *OptionContract) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *OptionContract) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *OptionContract) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDaysToExpiration

`func (o *OptionContract) GetDaysToExpiration() float32`

GetDaysToExpiration returns the DaysToExpiration field if non-nil, zero value otherwise.

### GetDaysToExpirationOk

`func (o *OptionContract) GetDaysToExpirationOk() (*float32, bool)`

GetDaysToExpirationOk returns a tuple with the DaysToExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToExpiration

`func (o *OptionContract) SetDaysToExpiration(v float32)`

SetDaysToExpiration sets DaysToExpiration field to given value.

### HasDaysToExpiration

`func (o *OptionContract) HasDaysToExpiration() bool`

HasDaysToExpiration returns a boolean if a field has been set.

### GetExpirationType

`func (o *OptionContract) GetExpirationType() ExpirationType`

GetExpirationType returns the ExpirationType field if non-nil, zero value otherwise.

### GetExpirationTypeOk

`func (o *OptionContract) GetExpirationTypeOk() (*ExpirationType, bool)`

GetExpirationTypeOk returns a tuple with the ExpirationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationType

`func (o *OptionContract) SetExpirationType(v ExpirationType)`

SetExpirationType sets ExpirationType field to given value.

### HasExpirationType

`func (o *OptionContract) HasExpirationType() bool`

HasExpirationType returns a boolean if a field has been set.

### GetLastTradingDay

`func (o *OptionContract) GetLastTradingDay() float32`

GetLastTradingDay returns the LastTradingDay field if non-nil, zero value otherwise.

### GetLastTradingDayOk

`func (o *OptionContract) GetLastTradingDayOk() (*float32, bool)`

GetLastTradingDayOk returns a tuple with the LastTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradingDay

`func (o *OptionContract) SetLastTradingDay(v float32)`

SetLastTradingDay sets LastTradingDay field to given value.

### HasLastTradingDay

`func (o *OptionContract) HasLastTradingDay() bool`

HasLastTradingDay returns a boolean if a field has been set.

### GetMultiplier

`func (o *OptionContract) GetMultiplier() float64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *OptionContract) GetMultiplierOk() (*float64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *OptionContract) SetMultiplier(v float64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *OptionContract) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetSettlementType

`func (o *OptionContract) GetSettlementType() SettlementType`

GetSettlementType returns the SettlementType field if non-nil, zero value otherwise.

### GetSettlementTypeOk

`func (o *OptionContract) GetSettlementTypeOk() (*SettlementType, bool)`

GetSettlementTypeOk returns a tuple with the SettlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementType

`func (o *OptionContract) SetSettlementType(v SettlementType)`

SetSettlementType sets SettlementType field to given value.

### HasSettlementType

`func (o *OptionContract) HasSettlementType() bool`

HasSettlementType returns a boolean if a field has been set.

### GetDeliverableNote

`func (o *OptionContract) GetDeliverableNote() string`

GetDeliverableNote returns the DeliverableNote field if non-nil, zero value otherwise.

### GetDeliverableNoteOk

`func (o *OptionContract) GetDeliverableNoteOk() (*string, bool)`

GetDeliverableNoteOk returns a tuple with the DeliverableNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverableNote

`func (o *OptionContract) SetDeliverableNote(v string)`

SetDeliverableNote sets DeliverableNote field to given value.

### HasDeliverableNote

`func (o *OptionContract) HasDeliverableNote() bool`

HasDeliverableNote returns a boolean if a field has been set.

### GetIsIndexOption

`func (o *OptionContract) GetIsIndexOption() bool`

GetIsIndexOption returns the IsIndexOption field if non-nil, zero value otherwise.

### GetIsIndexOptionOk

`func (o *OptionContract) GetIsIndexOptionOk() (*bool, bool)`

GetIsIndexOptionOk returns a tuple with the IsIndexOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIndexOption

`func (o *OptionContract) SetIsIndexOption(v bool)`

SetIsIndexOption sets IsIndexOption field to given value.

### HasIsIndexOption

`func (o *OptionContract) HasIsIndexOption() bool`

HasIsIndexOption returns a boolean if a field has been set.

### GetPercentChange

`func (o *OptionContract) GetPercentChange() float64`

GetPercentChange returns the PercentChange field if non-nil, zero value otherwise.

### GetPercentChangeOk

`func (o *OptionContract) GetPercentChangeOk() (*float64, bool)`

GetPercentChangeOk returns a tuple with the PercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange

`func (o *OptionContract) SetPercentChange(v float64)`

SetPercentChange sets PercentChange field to given value.

### HasPercentChange

`func (o *OptionContract) HasPercentChange() bool`

HasPercentChange returns a boolean if a field has been set.

### GetMarkChange

`func (o *OptionContract) GetMarkChange() float64`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *OptionContract) GetMarkChangeOk() (*float64, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *OptionContract) SetMarkChange(v float64)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *OptionContract) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *OptionContract) GetMarkPercentChange() float64`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *OptionContract) GetMarkPercentChangeOk() (*float64, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *OptionContract) SetMarkPercentChange(v float64)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *OptionContract) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetIsPennyPilot

`func (o *OptionContract) GetIsPennyPilot() bool`

GetIsPennyPilot returns the IsPennyPilot field if non-nil, zero value otherwise.

### GetIsPennyPilotOk

`func (o *OptionContract) GetIsPennyPilotOk() (*bool, bool)`

GetIsPennyPilotOk returns a tuple with the IsPennyPilot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPennyPilot

`func (o *OptionContract) SetIsPennyPilot(v bool)`

SetIsPennyPilot sets IsPennyPilot field to given value.

### HasIsPennyPilot

`func (o *OptionContract) HasIsPennyPilot() bool`

HasIsPennyPilot returns a boolean if a field has been set.

### GetIntrinsicValue

`func (o *OptionContract) GetIntrinsicValue() float64`

GetIntrinsicValue returns the IntrinsicValue field if non-nil, zero value otherwise.

### GetIntrinsicValueOk

`func (o *OptionContract) GetIntrinsicValueOk() (*float64, bool)`

GetIntrinsicValueOk returns a tuple with the IntrinsicValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrinsicValue

`func (o *OptionContract) SetIntrinsicValue(v float64)`

SetIntrinsicValue sets IntrinsicValue field to given value.

### HasIntrinsicValue

`func (o *OptionContract) HasIntrinsicValue() bool`

HasIntrinsicValue returns a boolean if a field has been set.

### GetOptionRoot

`func (o *OptionContract) GetOptionRoot() string`

GetOptionRoot returns the OptionRoot field if non-nil, zero value otherwise.

### GetOptionRootOk

`func (o *OptionContract) GetOptionRootOk() (*string, bool)`

GetOptionRootOk returns a tuple with the OptionRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionRoot

`func (o *OptionContract) SetOptionRoot(v string)`

SetOptionRoot sets OptionRoot field to given value.

### HasOptionRoot

`func (o *OptionContract) HasOptionRoot() bool`

HasOptionRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


