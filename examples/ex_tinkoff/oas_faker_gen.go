// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

// SetFake set fake values.
func (s *BrokerAccountType) SetFake() {
	*s = BrokerAccountTypeTinkoff
}

// SetFake set fake values.
func (s *Candle) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Interval.SetFake()
		}
	}
	{

		{
			s.O = float64(0)
		}
	}
	{

		{
			s.C = float64(0)
		}
	}
	{

		{
			s.H = float64(0)
		}
	}
	{

		{
			s.L = float64(0)
		}
	}
	{

		{
			s.V = int32(1)
		}
	}
	{

		{
			s.Time = time.Now()
		}
	}
}

// SetFake set fake values.
func (s *CandleResolution) SetFake() {
	*s = CandleResolution1min
}

// SetFake set fake values.
func (s *Candles) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Interval.SetFake()
		}
	}
	{

		{
			s.Candles = nil
			for i := 0; i < 0; i++ {
				var elem Candle

				{
					elem.SetFake()
				}
				s.Candles = append(s.Candles, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *CandlesResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Currencies) SetFake() {
	{

		{
			s.Currencies = nil
			for i := 0; i < 0; i++ {
				var elem CurrencyPosition

				{
					elem.SetFake()
				}
				s.Currencies = append(s.Currencies, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *Currency) SetFake() {
	*s = CurrencyRUB
}

// SetFake set fake values.
func (s *CurrencyPosition) SetFake() {
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Balance = float64(0)
		}
	}
	{

		{
			s.Blocked.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Empty) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
	{

		{
			s.Status = ""
		}
	}
}

// SetFake set fake values.
func (s *EmptyPayload) SetFake() {
}

// SetFake set fake values.
func (s *Error) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *ErrorPayload) SetFake() {
	{

		{
			s.Message.SetFake()
		}
	}
	{

		{
			s.Code.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *InstrumentType) SetFake() {
	*s = InstrumentTypeStock
}

// SetFake set fake values.
func (s *LimitOrderRequest) SetFake() {
	{

		{
			s.Lots = int32(1)
		}
	}
	{

		{
			s.Operation.SetFake()
		}
	}
	{

		{
			s.Price = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *LimitOrderResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MarketInstrument) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Ticker = ""
		}
	}
	{

		{
			s.Isin.SetFake()
		}
	}
	{

		{
			s.MinPriceIncrement.SetFake()
		}
	}
	{

		{
			s.Lot = int32(1)
		}
	}
	{

		{
			s.MinQuantity.SetFake()
		}
	}
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Name = ""
		}
	}
	{

		{
			s.Type.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MarketInstrumentList) SetFake() {
	{

		{
			s.Total = int32(1)
		}
	}
	{

		{
			s.Instruments = nil
			for i := 0; i < 0; i++ {
				var elem MarketInstrument

				{
					elem.SetFake()
				}
				s.Instruments = append(s.Instruments, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *MarketInstrumentListResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MarketOrderRequest) SetFake() {
	{

		{
			s.Lots = int32(1)
		}
	}
	{

		{
			s.Operation.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MarketOrderResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *MoneyAmount) SetFake() {
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Value = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *Operation) SetFake() {
	{

		{
			s.ID = ""
		}
	}
	{

		{
			s.Status.SetFake()
		}
	}
	{

		{
			s.Trades = nil
			for i := 0; i < 0; i++ {
				var elem OperationTrade

				{
					elem.SetFake()
				}
				s.Trades = append(s.Trades, elem)
			}
		}
	}
	{

		{
			s.Commission.SetFake()
		}
	}
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Payment = float64(0)
		}
	}
	{

		{
			s.Price.SetFake()
		}
	}
	{

		{
			s.Quantity.SetFake()
		}
	}
	{

		{
			s.QuantityExecuted.SetFake()
		}
	}
	{

		{
			s.Figi.SetFake()
		}
	}
	{

		{
			s.InstrumentType.SetFake()
		}
	}
	{

		{
			s.IsMarginCall = true
		}
	}
	{

		{
			s.Date = time.Now()
		}
	}
	{

		{
			s.OperationType.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OperationStatus) SetFake() {
	*s = OperationStatusDone
}

// SetFake set fake values.
func (s *OperationTrade) SetFake() {
	{

		{
			s.TradeId = ""
		}
	}
	{

		{
			s.Date = time.Now()
		}
	}
	{

		{
			s.Price = float64(0)
		}
	}
	{

		{
			s.Quantity = int32(1)
		}
	}
}

// SetFake set fake values.
func (s *OperationType) SetFake() {
	*s = OperationTypeBuy
}

// SetFake set fake values.
func (s *OperationTypeWithCommission) SetFake() {
	*s = OperationTypeWithCommissionBuy
}

// SetFake set fake values.
func (s *Operations) SetFake() {
	{

		{
			s.Operations = nil
			for i := 0; i < 0; i++ {
				var elem Operation

				{
					elem.SetFake()
				}
				s.Operations = append(s.Operations, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *OperationsResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OptBrokerAccountType) SetFake() {
	var elem BrokerAccountType

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptCurrency) SetFake() {
	var elem Currency

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptFloat64) SetFake() {
	var elem float64

	{
		elem = float64(0)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInstrumentType) SetFake() {
	var elem InstrumentType

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptInt32) SetFake() {
	var elem int32

	{
		elem = int32(1)
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptMoneyAmount) SetFake() {
	var elem MoneyAmount

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptOperationTypeWithCommission) SetFake() {
	var elem OperationTypeWithCommission

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptSandboxRegisterRequest) SetFake() {
	var elem SandboxRegisterRequest

	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *OptString) SetFake() {
	var elem string

	{
		elem = ""
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *Order) SetFake() {
	{

		{
			s.OrderId = ""
		}
	}
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Operation.SetFake()
		}
	}
	{

		{
			s.Status.SetFake()
		}
	}
	{

		{
			s.RequestedLots = int32(1)
		}
	}
	{

		{
			s.ExecutedLots = int32(1)
		}
	}
	{

		{
			s.Type.SetFake()
		}
	}
	{

		{
			s.Price = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *OrderResponse) SetFake() {
	{

		{
			s.Price = float64(0)
		}
	}
	{

		{
			s.Quantity = int32(1)
		}
	}
}

// SetFake set fake values.
func (s *OrderStatus) SetFake() {
	*s = OrderStatusNew
}

// SetFake set fake values.
func (s *OrderType) SetFake() {
	*s = OrderTypeLimit
}

// SetFake set fake values.
func (s *Orderbook) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Depth = int32(1)
		}
	}
	{

		{
			s.Bids = nil
			for i := 0; i < 0; i++ {
				var elem OrderResponse

				{
					elem.SetFake()
				}
				s.Bids = append(s.Bids, elem)
			}
		}
	}
	{

		{
			s.Asks = nil
			for i := 0; i < 0; i++ {
				var elem OrderResponse

				{
					elem.SetFake()
				}
				s.Asks = append(s.Asks, elem)
			}
		}
	}
	{

		{
			s.TradeStatus.SetFake()
		}
	}
	{

		{
			s.MinPriceIncrement = float64(0)
		}
	}
	{

		{
			s.FaceValue.SetFake()
		}
	}
	{

		{
			s.LastPrice.SetFake()
		}
	}
	{

		{
			s.ClosePrice.SetFake()
		}
	}
	{

		{
			s.LimitUp.SetFake()
		}
	}
	{

		{
			s.LimitDown.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OrderbookResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *OrdersResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload = nil
			for i := 0; i < 0; i++ {
				var elem Order

				{
					elem.SetFake()
				}
				s.Payload = append(s.Payload, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *PlacedLimitOrder) SetFake() {
	{

		{
			s.OrderId = ""
		}
	}
	{

		{
			s.Operation.SetFake()
		}
	}
	{

		{
			s.Status.SetFake()
		}
	}
	{

		{
			s.RejectReason.SetFake()
		}
	}
	{

		{
			s.Message.SetFake()
		}
	}
	{

		{
			s.RequestedLots = int(1)
		}
	}
	{

		{
			s.ExecutedLots = int(1)
		}
	}
	{

		{
			s.Commission.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PlacedMarketOrder) SetFake() {
	{

		{
			s.OrderId = ""
		}
	}
	{

		{
			s.Operation.SetFake()
		}
	}
	{

		{
			s.Status.SetFake()
		}
	}
	{

		{
			s.RejectReason.SetFake()
		}
	}
	{

		{
			s.Message.SetFake()
		}
	}
	{

		{
			s.RequestedLots = int(1)
		}
	}
	{

		{
			s.ExecutedLots = int(1)
		}
	}
	{

		{
			s.Commission.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *Portfolio) SetFake() {
	{

		{
			s.Positions = nil
			for i := 0; i < 0; i++ {
				var elem PortfolioPosition

				{
					elem.SetFake()
				}
				s.Positions = append(s.Positions, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *PortfolioCurrenciesResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *PortfolioPosition) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Ticker.SetFake()
		}
	}
	{

		{
			s.Isin.SetFake()
		}
	}
	{

		{
			s.InstrumentType.SetFake()
		}
	}
	{

		{
			s.Balance = float64(0)
		}
	}
	{

		{
			s.Blocked.SetFake()
		}
	}
	{

		{
			s.ExpectedYield.SetFake()
		}
	}
	{

		{
			s.Lots = int32(1)
		}
	}
	{

		{
			s.AveragePositionPrice.SetFake()
		}
	}
	{

		{
			s.AveragePositionPriceNoNkd.SetFake()
		}
	}
	{

		{
			s.Name = ""
		}
	}
}

// SetFake set fake values.
func (s *PortfolioResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SandboxAccount) SetFake() {
	{

		{
			s.BrokerAccountType.SetFake()
		}
	}
	{

		{
			s.BrokerAccountId = ""
		}
	}
}

// SetFake set fake values.
func (s *SandboxCurrency) SetFake() {
	*s = SandboxCurrencyRUB
}

// SetFake set fake values.
func (s *SandboxRegisterRequest) SetFake() {
	{

		{
			s.BrokerAccountType.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SandboxRegisterResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SandboxSetCurrencyBalanceRequest) SetFake() {
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Balance = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *SandboxSetPositionBalanceRequest) SetFake() {
	{

		{
			s.Figi.SetFake()
		}
	}
	{

		{
			s.Balance = float64(0)
		}
	}
}

// SetFake set fake values.
func (s *SearchMarketInstrument) SetFake() {
	{

		{
			s.Figi = ""
		}
	}
	{

		{
			s.Ticker = ""
		}
	}
	{

		{
			s.Isin.SetFake()
		}
	}
	{

		{
			s.MinPriceIncrement.SetFake()
		}
	}
	{

		{
			s.Lot = int32(1)
		}
	}
	{

		{
			s.Currency.SetFake()
		}
	}
	{

		{
			s.Name = ""
		}
	}
	{

		{
			s.Type.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *SearchMarketInstrumentResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *TradeStatus) SetFake() {
	*s = TradeStatusNormalTrading
}

// SetFake set fake values.
func (s *UserAccount) SetFake() {
	{

		{
			s.BrokerAccountType.SetFake()
		}
	}
	{

		{
			s.BrokerAccountId = ""
		}
	}
}

// SetFake set fake values.
func (s *UserAccounts) SetFake() {
	{

		{
			s.Accounts = nil
			for i := 0; i < 0; i++ {
				var elem UserAccount

				{
					elem.SetFake()
				}
				s.Accounts = append(s.Accounts, elem)
			}
		}
	}
}

// SetFake set fake values.
func (s *UserAccountsResponse) SetFake() {
	{

		{
			s.TrackingId = ""
		}
	}
	{

		{
			s.Status = ""
		}
	}
	{

		{
			s.Payload.SetFake()
		}
	}
}
