package mqtt

////////////////////Interface//////////////////////////////

type Stack interface {
	CreateTransport(network string, address string, port int) Transport
	GetTransports() []Transport
	DeleteTransport(t Transport)

	CreateProvider() Provider
	GetProviders() []Provider
	DeleteProvider(p Provider)

	Run()
	Stop()
}

////////////////////Implementation////////////////////////

var stackSingleton Stack

func GetStack() Stack {
	if stackSingleton == nil {
		return newStack()
	} else {
		return stackSingleton
	}
}

type stack struct {
	transports map[Transport]*transport
	providers  map[Provider]*provider
}

func newStack() Stack {
	this := &stack{}

	this.transports = make(map[Transport]*transport)
	this.providers = make(map[Provider]*provider)

	return this
}

func (this *stack) CreateTransport(network string, address string, port int) Transport {
	t := newTransport(network, address, port)

	this.transports[t] = t

	return t
}

func (this *stack) GetTransports() []Transport {
	transports := make([]Transport, len(this.transports))

	l := 0
	for _, value := range this.transports {
		transports[l] = value
		l++
	}

	return transports
}

func (this *stack) DeleteTransport(t Transport) {
	delete(this.transports, t)
}

func (this *stack) CreateProvider() Provider {
	p := newProvider()

	this.providers[p] = p

	return p
}

func (this *stack) GetProviders() []Provider {
	providers := make([]Provider, len(this.providers))

	l := 0
	for _, value := range this.providers {
		providers[l] = value
		l++
	}

	return providers
}

func (this *stack) DeleteProvider(p Provider) {
	delete(this.providers, p)
}

func (this *stack) Run() {
	//TODO
}

func (this *stack) Stop() {
	//TODO
}
