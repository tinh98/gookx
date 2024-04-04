package okxv5

type Executor[T any] struct {
	Args          SubscriptionArgs
	subscriptions *Subscriptions

	unmarshalRawTopic func(data RawTopic) (Topic[T], error)
}

func NewExecutor[T any](args SubscriptionArgs, subscriptions *Subscriptions) *Executor[T] {
	o := new(Executor[T])
	o.Args = args
	o.subscriptions = subscriptions
	return o
}

func (o *Executor[T]) Subscribe(onShot func(Topic[T])) {
	o.subscriptions.subscribe(o.Args, func(raw RawTopic) error {

		// out, _ := json.Marshal(raw)
		// fmt.Print(string(out) + "\n\n")

		if o.unmarshalRawTopic != nil {
			topic, err := o.unmarshalRawTopic(raw)
			if err == nil {
				onShot(topic)
			}
			return err
		}
		topic, err := UnmarshalRawTopic[T](raw)

		if err == nil {
			onShot(topic)
		}
		return err
	})
}

func (o *Executor[T]) Unsubscribe() {
	o.subscriptions.unsubscribe(o.Args)
}

func (o *Executor[T]) SetUnMarshaller(unMarshaller func(data RawTopic) (Topic[T], error)) {
	o.unmarshalRawTopic = unMarshaller
}
