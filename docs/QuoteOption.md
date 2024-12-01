# QuoteOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var52WeekHigh** | Pointer to **float64** | Higest price traded in the past 12 months, or 52 weeks | [optional] 
**Var52WeekLow** | Pointer to **float64** | Lowest price traded in the past 12 months, or 52 weeks | [optional] 
**AskPrice** | Pointer to **float64** | Current Best Ask Price | [optional] 
**AskSize** | Pointer to **int32** | Number of shares for ask | [optional] 
**BidPrice** | Pointer to **float64** | Current Best Bid Price | [optional] 
**BidSize** | Pointer to **int32** | Number of shares for bid | [optional] 
**ClosePrice** | Pointer to **float64** | Previous day&#39;s closing price | [optional] 
**Delta** | Pointer to **float64** | Delta Value | [optional] 
**Gamma** | Pointer to **float64** | Gamma Value | [optional] 
**HighPrice** | Pointer to **float64** | Day&#39;s high trade price | [optional] 
**IndAskPrice** | Pointer to **float64** | Indicative Ask Price applicable only for Indicative Option Symbols | [optional] 
**IndBidPrice** | Pointer to **float64** | Indicative Bid Price applicable only for Indicative Option Symbols | [optional] 
**IndQuoteTime** | Pointer to **int64** | Indicative Quote Time in milliseconds since Epoch applicable only for Indicative Option Symbols | [optional] 
**ImpliedYield** | Pointer to **float64** | Implied Yield | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**LastSize** | Pointer to **int32** | Number of shares traded with last trade | [optional] 
**LowPrice** | Pointer to **float64** | Day&#39;s low trade price | [optional] 
**Mark** | Pointer to **float64** | Mark price | [optional] 
**MarkChange** | Pointer to **float64** | Mark Price change | [optional] 
**MarkPercentChange** | Pointer to **float64** | Mark Price percent change | [optional] 
**MoneyIntrinsicValue** | Pointer to **float64** | Money Intrinsic Value | [optional] 
**NetChange** | Pointer to **float64** | Current Last-Prev Close | [optional] 
**NetPercentChange** | Pointer to **float64** | Net Percentage Change | [optional] 
**OpenInterest** | Pointer to **float64** | Open Interest | [optional] 
**OpenPrice** | Pointer to **float64** | Price at market open | [optional] 
**QuoteTime** | Pointer to **int64** | Last quote time in milliseconds since Epoch | [optional] 
**Rho** | Pointer to **float64** | Rho Value | [optional] 
**SecurityStatus** | Pointer to **string** | Status of security | [optional] 
**TheoreticalOptionValue** | Pointer to **float64** | Theoretical option Value | [optional] 
**Theta** | Pointer to **float64** | Theta Value | [optional] 
**TimeValue** | Pointer to **float64** | Time Value | [optional] 
**TotalVolume** | Pointer to **int64** | Aggregated shares traded throughout the day, including pre/post market hours. | [optional] 
**TradeTime** | Pointer to **int64** | Last trade time in milliseconds since Epoch | [optional] 
**UnderlyingPrice** | Pointer to **float64** | Underlying Price | [optional] 
**Vega** | Pointer to **float64** | Vega Value | [optional] 
**Volatility** | Pointer to **float64** | Option Risk/Volatility Measurement | [optional] 

## Methods

### NewQuoteOption

`func NewQuoteOption() *QuoteOption`

NewQuoteOption instantiates a new QuoteOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteOptionWithDefaults

`func NewQuoteOptionWithDefaults() *QuoteOption`

NewQuoteOptionWithDefaults instantiates a new QuoteOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar52WeekHigh

`func (o *QuoteOption) GetVar52WeekHigh() float64`

GetVar52WeekHigh returns the Var52WeekHigh field if non-nil, zero value otherwise.

### GetVar52WeekHighOk

`func (o *QuoteOption) GetVar52WeekHighOk() (*float64, bool)`

GetVar52WeekHighOk returns a tuple with the Var52WeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekHigh

`func (o *QuoteOption) SetVar52WeekHigh(v float64)`

SetVar52WeekHigh sets Var52WeekHigh field to given value.

### HasVar52WeekHigh

`func (o *QuoteOption) HasVar52WeekHigh() bool`

HasVar52WeekHigh returns a boolean if a field has been set.

### GetVar52WeekLow

`func (o *QuoteOption) GetVar52WeekLow() float64`

GetVar52WeekLow returns the Var52WeekLow field if non-nil, zero value otherwise.

### GetVar52WeekLowOk

`func (o *QuoteOption) GetVar52WeekLowOk() (*float64, bool)`

GetVar52WeekLowOk returns a tuple with the Var52WeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar52WeekLow

`func (o *QuoteOption) SetVar52WeekLow(v float64)`

SetVar52WeekLow sets Var52WeekLow field to given value.

### HasVar52WeekLow

`func (o *QuoteOption) HasVar52WeekLow() bool`

HasVar52WeekLow returns a boolean if a field has been set.

### GetAskPrice

`func (o *QuoteOption) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *QuoteOption) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *QuoteOption) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *QuoteOption) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskSize

`func (o *QuoteOption) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *QuoteOption) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *QuoteOption) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *QuoteOption) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.

### GetBidPrice

`func (o *QuoteOption) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *QuoteOption) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *QuoteOption) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *QuoteOption) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *QuoteOption) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *QuoteOption) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *QuoteOption) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *QuoteOption) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetClosePrice

`func (o *QuoteOption) GetClosePrice() float64`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *QuoteOption) GetClosePriceOk() (*float64, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *QuoteOption) SetClosePrice(v float64)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *QuoteOption) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### GetDelta

`func (o *QuoteOption) GetDelta() float64`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *QuoteOption) GetDeltaOk() (*float64, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *QuoteOption) SetDelta(v float64)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *QuoteOption) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### GetGamma

`func (o *QuoteOption) GetGamma() float64`

GetGamma returns the Gamma field if non-nil, zero value otherwise.

### GetGammaOk

`func (o *QuoteOption) GetGammaOk() (*float64, bool)`

GetGammaOk returns a tuple with the Gamma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGamma

`func (o *QuoteOption) SetGamma(v float64)`

SetGamma sets Gamma field to given value.

### HasGamma

`func (o *QuoteOption) HasGamma() bool`

HasGamma returns a boolean if a field has been set.

### GetHighPrice

`func (o *QuoteOption) GetHighPrice() float64`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *QuoteOption) GetHighPriceOk() (*float64, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *QuoteOption) SetHighPrice(v float64)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *QuoteOption) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetIndAskPrice

`func (o *QuoteOption) GetIndAskPrice() float64`

GetIndAskPrice returns the IndAskPrice field if non-nil, zero value otherwise.

### GetIndAskPriceOk

`func (o *QuoteOption) GetIndAskPriceOk() (*float64, bool)`

GetIndAskPriceOk returns a tuple with the IndAskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndAskPrice

`func (o *QuoteOption) SetIndAskPrice(v float64)`

SetIndAskPrice sets IndAskPrice field to given value.

### HasIndAskPrice

`func (o *QuoteOption) HasIndAskPrice() bool`

HasIndAskPrice returns a boolean if a field has been set.

### GetIndBidPrice

`func (o *QuoteOption) GetIndBidPrice() float64`

GetIndBidPrice returns the IndBidPrice field if non-nil, zero value otherwise.

### GetIndBidPriceOk

`func (o *QuoteOption) GetIndBidPriceOk() (*float64, bool)`

GetIndBidPriceOk returns a tuple with the IndBidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndBidPrice

`func (o *QuoteOption) SetIndBidPrice(v float64)`

SetIndBidPrice sets IndBidPrice field to given value.

### HasIndBidPrice

`func (o *QuoteOption) HasIndBidPrice() bool`

HasIndBidPrice returns a boolean if a field has been set.

### GetIndQuoteTime

`func (o *QuoteOption) GetIndQuoteTime() int64`

GetIndQuoteTime returns the IndQuoteTime field if non-nil, zero value otherwise.

### GetIndQuoteTimeOk

`func (o *QuoteOption) GetIndQuoteTimeOk() (*int64, bool)`

GetIndQuoteTimeOk returns a tuple with the IndQuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndQuoteTime

`func (o *QuoteOption) SetIndQuoteTime(v int64)`

SetIndQuoteTime sets IndQuoteTime field to given value.

### HasIndQuoteTime

`func (o *QuoteOption) HasIndQuoteTime() bool`

HasIndQuoteTime returns a boolean if a field has been set.

### GetImpliedYield

`func (o *QuoteOption) GetImpliedYield() float64`

GetImpliedYield returns the ImpliedYield field if non-nil, zero value otherwise.

### GetImpliedYieldOk

`func (o *QuoteOption) GetImpliedYieldOk() (*float64, bool)`

GetImpliedYieldOk returns a tuple with the ImpliedYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpliedYield

`func (o *QuoteOption) SetImpliedYield(v float64)`

SetImpliedYield sets ImpliedYield field to given value.

### HasImpliedYield

`func (o *QuoteOption) HasImpliedYield() bool`

HasImpliedYield returns a boolean if a field has been set.

### GetLastPrice

`func (o *QuoteOption) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *QuoteOption) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *QuoteOption) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *QuoteOption) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastSize

`func (o *QuoteOption) GetLastSize() int32`

GetLastSize returns the LastSize field if non-nil, zero value otherwise.

### GetLastSizeOk

`func (o *QuoteOption) GetLastSizeOk() (*int32, bool)`

GetLastSizeOk returns a tuple with the LastSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSize

`func (o *QuoteOption) SetLastSize(v int32)`

SetLastSize sets LastSize field to given value.

### HasLastSize

`func (o *QuoteOption) HasLastSize() bool`

HasLastSize returns a boolean if a field has been set.

### GetLowPrice

`func (o *QuoteOption) GetLowPrice() float64`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *QuoteOption) GetLowPriceOk() (*float64, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *QuoteOption) SetLowPrice(v float64)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *QuoteOption) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetMark

`func (o *QuoteOption) GetMark() float64`

GetMark returns the Mark field if non-nil, zero value otherwise.

### GetMarkOk

`func (o *QuoteOption) GetMarkOk() (*float64, bool)`

GetMarkOk returns a tuple with the Mark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMark

`func (o *QuoteOption) SetMark(v float64)`

SetMark sets Mark field to given value.

### HasMark

`func (o *QuoteOption) HasMark() bool`

HasMark returns a boolean if a field has been set.

### GetMarkChange

`func (o *QuoteOption) GetMarkChange() float64`

GetMarkChange returns the MarkChange field if non-nil, zero value otherwise.

### GetMarkChangeOk

`func (o *QuoteOption) GetMarkChangeOk() (*float64, bool)`

GetMarkChangeOk returns a tuple with the MarkChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkChange

`func (o *QuoteOption) SetMarkChange(v float64)`

SetMarkChange sets MarkChange field to given value.

### HasMarkChange

`func (o *QuoteOption) HasMarkChange() bool`

HasMarkChange returns a boolean if a field has been set.

### GetMarkPercentChange

`func (o *QuoteOption) GetMarkPercentChange() float64`

GetMarkPercentChange returns the MarkPercentChange field if non-nil, zero value otherwise.

### GetMarkPercentChangeOk

`func (o *QuoteOption) GetMarkPercentChangeOk() (*float64, bool)`

GetMarkPercentChangeOk returns a tuple with the MarkPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPercentChange

`func (o *QuoteOption) SetMarkPercentChange(v float64)`

SetMarkPercentChange sets MarkPercentChange field to given value.

### HasMarkPercentChange

`func (o *QuoteOption) HasMarkPercentChange() bool`

HasMarkPercentChange returns a boolean if a field has been set.

### GetMoneyIntrinsicValue

`func (o *QuoteOption) GetMoneyIntrinsicValue() float64`

GetMoneyIntrinsicValue returns the MoneyIntrinsicValue field if non-nil, zero value otherwise.

### GetMoneyIntrinsicValueOk

`func (o *QuoteOption) GetMoneyIntrinsicValueOk() (*float64, bool)`

GetMoneyIntrinsicValueOk returns a tuple with the MoneyIntrinsicValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyIntrinsicValue

`func (o *QuoteOption) SetMoneyIntrinsicValue(v float64)`

SetMoneyIntrinsicValue sets MoneyIntrinsicValue field to given value.

### HasMoneyIntrinsicValue

`func (o *QuoteOption) HasMoneyIntrinsicValue() bool`

HasMoneyIntrinsicValue returns a boolean if a field has been set.

### GetNetChange

`func (o *QuoteOption) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *QuoteOption) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *QuoteOption) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *QuoteOption) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.

### GetNetPercentChange

`func (o *QuoteOption) GetNetPercentChange() float64`

GetNetPercentChange returns the NetPercentChange field if non-nil, zero value otherwise.

### GetNetPercentChangeOk

`func (o *QuoteOption) GetNetPercentChangeOk() (*float64, bool)`

GetNetPercentChangeOk returns a tuple with the NetPercentChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetPercentChange

`func (o *QuoteOption) SetNetPercentChange(v float64)`

SetNetPercentChange sets NetPercentChange field to given value.

### HasNetPercentChange

`func (o *QuoteOption) HasNetPercentChange() bool`

HasNetPercentChange returns a boolean if a field has been set.

### GetOpenInterest

`func (o *QuoteOption) GetOpenInterest() float64`

GetOpenInterest returns the OpenInterest field if non-nil, zero value otherwise.

### GetOpenInterestOk

`func (o *QuoteOption) GetOpenInterestOk() (*float64, bool)`

GetOpenInterestOk returns a tuple with the OpenInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInterest

`func (o *QuoteOption) SetOpenInterest(v float64)`

SetOpenInterest sets OpenInterest field to given value.

### HasOpenInterest

`func (o *QuoteOption) HasOpenInterest() bool`

HasOpenInterest returns a boolean if a field has been set.

### GetOpenPrice

`func (o *QuoteOption) GetOpenPrice() float64`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *QuoteOption) GetOpenPriceOk() (*float64, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *QuoteOption) SetOpenPrice(v float64)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *QuoteOption) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetQuoteTime

`func (o *QuoteOption) GetQuoteTime() int64`

GetQuoteTime returns the QuoteTime field if non-nil, zero value otherwise.

### GetQuoteTimeOk

`func (o *QuoteOption) GetQuoteTimeOk() (*int64, bool)`

GetQuoteTimeOk returns a tuple with the QuoteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTime

`func (o *QuoteOption) SetQuoteTime(v int64)`

SetQuoteTime sets QuoteTime field to given value.

### HasQuoteTime

`func (o *QuoteOption) HasQuoteTime() bool`

HasQuoteTime returns a boolean if a field has been set.

### GetRho

`func (o *QuoteOption) GetRho() float64`

GetRho returns the Rho field if non-nil, zero value otherwise.

### GetRhoOk

`func (o *QuoteOption) GetRhoOk() (*float64, bool)`

GetRhoOk returns a tuple with the Rho field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRho

`func (o *QuoteOption) SetRho(v float64)`

SetRho sets Rho field to given value.

### HasRho

`func (o *QuoteOption) HasRho() bool`

HasRho returns a boolean if a field has been set.

### GetSecurityStatus

`func (o *QuoteOption) GetSecurityStatus() string`

GetSecurityStatus returns the SecurityStatus field if non-nil, zero value otherwise.

### GetSecurityStatusOk

`func (o *QuoteOption) GetSecurityStatusOk() (*string, bool)`

GetSecurityStatusOk returns a tuple with the SecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatus

`func (o *QuoteOption) SetSecurityStatus(v string)`

SetSecurityStatus sets SecurityStatus field to given value.

### HasSecurityStatus

`func (o *QuoteOption) HasSecurityStatus() bool`

HasSecurityStatus returns a boolean if a field has been set.

### GetTheoreticalOptionValue

`func (o *QuoteOption) GetTheoreticalOptionValue() float64`

GetTheoreticalOptionValue returns the TheoreticalOptionValue field if non-nil, zero value otherwise.

### GetTheoreticalOptionValueOk

`func (o *QuoteOption) GetTheoreticalOptionValueOk() (*float64, bool)`

GetTheoreticalOptionValueOk returns a tuple with the TheoreticalOptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheoreticalOptionValue

`func (o *QuoteOption) SetTheoreticalOptionValue(v float64)`

SetTheoreticalOptionValue sets TheoreticalOptionValue field to given value.

### HasTheoreticalOptionValue

`func (o *QuoteOption) HasTheoreticalOptionValue() bool`

HasTheoreticalOptionValue returns a boolean if a field has been set.

### GetTheta

`func (o *QuoteOption) GetTheta() float64`

GetTheta returns the Theta field if non-nil, zero value otherwise.

### GetThetaOk

`func (o *QuoteOption) GetThetaOk() (*float64, bool)`

GetThetaOk returns a tuple with the Theta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheta

`func (o *QuoteOption) SetTheta(v float64)`

SetTheta sets Theta field to given value.

### HasTheta

`func (o *QuoteOption) HasTheta() bool`

HasTheta returns a boolean if a field has been set.

### GetTimeValue

`func (o *QuoteOption) GetTimeValue() float64`

GetTimeValue returns the TimeValue field if non-nil, zero value otherwise.

### GetTimeValueOk

`func (o *QuoteOption) GetTimeValueOk() (*float64, bool)`

GetTimeValueOk returns a tuple with the TimeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValue

`func (o *QuoteOption) SetTimeValue(v float64)`

SetTimeValue sets TimeValue field to given value.

### HasTimeValue

`func (o *QuoteOption) HasTimeValue() bool`

HasTimeValue returns a boolean if a field has been set.

### GetTotalVolume

`func (o *QuoteOption) GetTotalVolume() int64`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *QuoteOption) GetTotalVolumeOk() (*int64, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *QuoteOption) SetTotalVolume(v int64)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *QuoteOption) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetTradeTime

`func (o *QuoteOption) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *QuoteOption) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *QuoteOption) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *QuoteOption) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.

### GetUnderlyingPrice

`func (o *QuoteOption) GetUnderlyingPrice() float64`

GetUnderlyingPrice returns the UnderlyingPrice field if non-nil, zero value otherwise.

### GetUnderlyingPriceOk

`func (o *QuoteOption) GetUnderlyingPriceOk() (*float64, bool)`

GetUnderlyingPriceOk returns a tuple with the UnderlyingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingPrice

`func (o *QuoteOption) SetUnderlyingPrice(v float64)`

SetUnderlyingPrice sets UnderlyingPrice field to given value.

### HasUnderlyingPrice

`func (o *QuoteOption) HasUnderlyingPrice() bool`

HasUnderlyingPrice returns a boolean if a field has been set.

### GetVega

`func (o *QuoteOption) GetVega() float64`

GetVega returns the Vega field if non-nil, zero value otherwise.

### GetVegaOk

`func (o *QuoteOption) GetVegaOk() (*float64, bool)`

GetVegaOk returns a tuple with the Vega field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVega

`func (o *QuoteOption) SetVega(v float64)`

SetVega sets Vega field to given value.

### HasVega

`func (o *QuoteOption) HasVega() bool`

HasVega returns a boolean if a field has been set.

### GetVolatility

`func (o *QuoteOption) GetVolatility() float64`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *QuoteOption) GetVolatilityOk() (*float64, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *QuoteOption) SetVolatility(v float64)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *QuoteOption) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


