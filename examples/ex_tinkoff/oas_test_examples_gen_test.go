// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/jx"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBrokerAccountType_EncodeDecode(t *testing.T) {
	var typ BrokerAccountType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 BrokerAccountType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCandle_EncodeDecode(t *testing.T) {
	var typ Candle
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Candle
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCandleResolution_EncodeDecode(t *testing.T) {
	var typ CandleResolution
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CandleResolution
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCandles_EncodeDecode(t *testing.T) {
	var typ Candles
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Candles
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCandlesResponse_EncodeDecode(t *testing.T) {
	var typ CandlesResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CandlesResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCurrencies_EncodeDecode(t *testing.T) {
	var typ Currencies
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Currencies
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCurrency_EncodeDecode(t *testing.T) {
	var typ Currency
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Currency
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCurrencyPosition_EncodeDecode(t *testing.T) {
	var typ CurrencyPosition
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CurrencyPosition
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestEmpty_EncodeDecode(t *testing.T) {
	var typ Empty
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Empty
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestEmptyPayload_EncodeDecode(t *testing.T) {
	var typ EmptyPayload
	typ = make(EmptyPayload)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 EmptyPayload
	typ2 = make(EmptyPayload)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestErrorPayload_EncodeDecode(t *testing.T) {
	var typ ErrorPayload
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ErrorPayload
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestInstrumentType_EncodeDecode(t *testing.T) {
	var typ InstrumentType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 InstrumentType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLimitOrderRequest_EncodeDecode(t *testing.T) {
	var typ LimitOrderRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LimitOrderRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestLimitOrderResponse_EncodeDecode(t *testing.T) {
	var typ LimitOrderResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 LimitOrderResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMarketInstrument_EncodeDecode(t *testing.T) {
	var typ MarketInstrument
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MarketInstrument
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMarketInstrumentList_EncodeDecode(t *testing.T) {
	var typ MarketInstrumentList
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MarketInstrumentList
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMarketInstrumentListResponse_EncodeDecode(t *testing.T) {
	var typ MarketInstrumentListResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MarketInstrumentListResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMarketOrderRequest_EncodeDecode(t *testing.T) {
	var typ MarketOrderRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MarketOrderRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMarketOrderResponse_EncodeDecode(t *testing.T) {
	var typ MarketOrderResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MarketOrderResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMoneyAmount_EncodeDecode(t *testing.T) {
	var typ MoneyAmount
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MoneyAmount
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperation_EncodeDecode(t *testing.T) {
	var typ Operation
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Operation
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperationStatus_EncodeDecode(t *testing.T) {
	var typ OperationStatus
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OperationStatus
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperationTrade_EncodeDecode(t *testing.T) {
	var typ OperationTrade
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OperationTrade
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperationType_EncodeDecode(t *testing.T) {
	var typ OperationType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OperationType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperationTypeWithCommission_EncodeDecode(t *testing.T) {
	var typ OperationTypeWithCommission
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OperationTypeWithCommission
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperations_EncodeDecode(t *testing.T) {
	var typ Operations
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Operations
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOperationsResponse_EncodeDecode(t *testing.T) {
	var typ OperationsResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OperationsResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrder_EncodeDecode(t *testing.T) {
	var typ Order
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Order
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrderResponse_EncodeDecode(t *testing.T) {
	var typ OrderResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OrderResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrderStatus_EncodeDecode(t *testing.T) {
	var typ OrderStatus
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OrderStatus
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrderType_EncodeDecode(t *testing.T) {
	var typ OrderType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OrderType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrderbook_EncodeDecode(t *testing.T) {
	var typ Orderbook
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Orderbook
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrderbookResponse_EncodeDecode(t *testing.T) {
	var typ OrderbookResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OrderbookResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOrdersResponse_EncodeDecode(t *testing.T) {
	var typ OrdersResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OrdersResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPlacedLimitOrder_EncodeDecode(t *testing.T) {
	var typ PlacedLimitOrder
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PlacedLimitOrder
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPlacedMarketOrder_EncodeDecode(t *testing.T) {
	var typ PlacedMarketOrder
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PlacedMarketOrder
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPortfolio_EncodeDecode(t *testing.T) {
	var typ Portfolio
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Portfolio
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPortfolioCurrenciesResponse_EncodeDecode(t *testing.T) {
	var typ PortfolioCurrenciesResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PortfolioCurrenciesResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPortfolioPosition_EncodeDecode(t *testing.T) {
	var typ PortfolioPosition
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PortfolioPosition
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPortfolioResponse_EncodeDecode(t *testing.T) {
	var typ PortfolioResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PortfolioResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxAccount_EncodeDecode(t *testing.T) {
	var typ SandboxAccount
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxAccount
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxCurrency_EncodeDecode(t *testing.T) {
	var typ SandboxCurrency
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxCurrency
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxRegisterRequest_EncodeDecode(t *testing.T) {
	var typ SandboxRegisterRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxRegisterRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxRegisterResponse_EncodeDecode(t *testing.T) {
	var typ SandboxRegisterResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxRegisterResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxSetCurrencyBalanceRequest_EncodeDecode(t *testing.T) {
	var typ SandboxSetCurrencyBalanceRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxSetCurrencyBalanceRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSandboxSetPositionBalanceRequest_EncodeDecode(t *testing.T) {
	var typ SandboxSetPositionBalanceRequest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SandboxSetPositionBalanceRequest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSearchMarketInstrument_EncodeDecode(t *testing.T) {
	var typ SearchMarketInstrument
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SearchMarketInstrument
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSearchMarketInstrumentResponse_EncodeDecode(t *testing.T) {
	var typ SearchMarketInstrumentResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SearchMarketInstrumentResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTradeStatus_EncodeDecode(t *testing.T) {
	var typ TradeStatus
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TradeStatus
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestUserAccount_EncodeDecode(t *testing.T) {
	var typ UserAccount
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 UserAccount
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestUserAccounts_EncodeDecode(t *testing.T) {
	var typ UserAccounts
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 UserAccounts
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestUserAccountsResponse_EncodeDecode(t *testing.T) {
	var typ UserAccountsResponse
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 UserAccountsResponse
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
