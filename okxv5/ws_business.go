package okxv5

import (
	"github.com/msw-x/moon/ulog"
	"github.com/msw-x/moon/uws"
	"strconv"
	"time"
)

type WsBusiness struct {
	c             *WsClient
	onConnected   func()
	subscriptions *Subscriptions
}

func NewWsBusiness() *WsBusiness {
	o := new(WsBusiness)
	o.c = NewWsClient()
	o.subscriptions = NewSubscriptions(o)
	o.c.WithPath("v5/business")
	return o
}

func (o *WsBusiness) Close() {
	o.c.Close()
}

func (o *WsBusiness) Transport() *uws.Options {
	return o.c.Transport()
}

func (o *WsBusiness) WithLog(log *ulog.Log) *WsBusiness {
	o.c.WithLog(log)
	return o
}

func (o *WsBusiness) WithProxy(proxy string) *WsBusiness {
	o.c.WithProxy(proxy)
	return o
}

func (o *WsBusiness) WithLogRequest(enable bool) *WsBusiness {
	o.c.WithLogRequest(enable)
	return o
}

func (o *WsBusiness) WithLogResponse(enable bool) *WsBusiness {
	o.c.WithLogResponse(enable)
	return o
}

func (o *WsBusiness) WithOnDialError(f func(error) bool) *WsBusiness {
	o.c.WithOnDialError(f)
	return o
}

func (o *WsBusiness) WithOnConnected(f func()) *WsBusiness {
	o.onConnected = f
	return o
}

func (o *WsBusiness) WithOnDisconnected(f func()) *WsBusiness {
	o.c.WithOnDisconnected(f)
	return o
}

func (o *WsBusiness) Run() {
	o.c.WithOnConnected(func() {
		if o.onConnected != nil {
			o.onConnected()
		}
		o.subscriptions.subscribeAll()
	})
	o.c.WithOnResponse(o.onResponse)
	o.c.WithOnTopic(o.onTopic)
	o.c.Run()
}

func (o *WsBusiness) Connected() bool {
	return o.c.Connected()
}

func (o *WsBusiness) Ready() bool {
	return o.Connected()
}

func (o *WsBusiness) subscribe(topicArgs SubscriptionArgs) {
	o.c.Subscribe(topicArgs)
}

func (o *WsBusiness) unsubscribe(topicArgs SubscriptionArgs) {
	o.c.Unsubscribe(topicArgs)
}

func (o *WsBusiness) onResponse(r WsResponse) error {
	r.Log(o.c.Log())
	return nil
}

func (o *WsBusiness) onTopic(data []byte) error {
	return o.subscriptions.processTopic(data)
}

func (o *WsBusiness) CandleStick(instId string, candleType CandleStickType) *Executor[CandleSticksData] {

	args := SubscriptionArgs{
		Channel: string(candleType),
		InstId:  instId,
	}
	executor := NewExecutor[CandleSticksData](args, o.subscriptions)
	executor.SetUnMarshaller(unmarshallCandleStickData)

	return executor
}

func unmarshallCandleStickData(raw RawTopic) (Topic[CandleSticksData], error) {
	candleData := *new(Topic[CandleSticksData])
	topic, err := UnmarshalRawTopic[[][]string](raw)
	if err != nil {
		return candleData, err
	}
	data := topic.Data[0]
	//copy topic
	candleData.Arg = topic.Arg
	candleData.Action = topic.Action

	var ticks int64
	var timestamp time.Time
	var o, h, l, c, v, vCcy, vCcyQuote float64
	var confirm bool
	ticks, err = strconv.ParseInt(data[0], 10, 64)
	if err != nil {
		return candleData, err
	}
	timestamp = time.UnixMilli(ticks)
	o, err = strconv.ParseFloat(data[1], 64)
	if err != nil {
		return candleData, err
	}
	h, err = strconv.ParseFloat(data[2], 64)
	if err != nil {
		return candleData, err
	}
	l, err = strconv.ParseFloat(data[3], 64)
	if err != nil {
		return candleData, err
	}
	c, err = strconv.ParseFloat(data[4], 64)
	if err != nil {
		return candleData, err
	}
	v, err = strconv.ParseFloat(data[5], 64)
	if err != nil {
		return candleData, err
	}
	vCcy, err = strconv.ParseFloat(data[6], 64)
	if err != nil {
		return candleData, err
	}
	vCcyQuote, err = strconv.ParseFloat(data[7], 64)
	if err != nil {
		return candleData, err
	}
	confirm, err = strconv.ParseBool(data[8])
	if err != nil {
		return candleData, err
	}

	candleData.Data = CandleSticksData{
		timestamp,
		o,
		c,
		h,
		l,
		v,
		vCcy,
		vCcyQuote,
		confirm,
	}
	return candleData, nil
}
