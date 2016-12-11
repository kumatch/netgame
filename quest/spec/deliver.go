package spec

type Verifier struct {
	req chan struct{}
	res chan bool
}

func (v *Verifier) Request() {
	v.req <- struct{}{}
}

func (v *Verifier) ReceiveRequest() <-chan struct{} {
	return v.req
}

func (v *Verifier) Result(r bool) {
	v.res <- r
}

func (v *Verifier) ReceiveResult() <-chan bool {
	return v.res
}

func newVerifier() *Verifier {
	return &Verifier{
		req: make(chan struct{}),
		res: make(chan bool),
	}
}

type Deliver struct {
	box    chan *Sheet
	signal chan struct{}
}

func (d *Deliver) Delivery(s *Sheet) {
	d.box <- s
}

func (d *Deliver) ReceiveSheet() <-chan *Sheet {
	return d.box
}

func (d *Deliver) ReceiveSignel() <-chan struct{} {
	return d.signal
}

func newDeliver() *Deliver {
	return &Deliver{
		box:    make(chan *Sheet),
		signal: make(chan struct{}),
	}
}

func NewSpecDeliver() (*Deliver, *Verifier) {
	d := newDeliver()
	v := newVerifier()

	go runServer(d, v)
	return d, v
}
