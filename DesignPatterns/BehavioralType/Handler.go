package BehavioralType

import "fmt"

const (
	// Difficulty levels decide which handler can process a request.
	DIFFICULTY_LEVEL_1 = 1
	DIFFICULTY_LEVEL_2 = 2
	DIFFICULTY_LEVEL_3 = 3
)

// HandleMessage is the chain traversal strategy shared by concrete handlers.
type HandleMessage func(hand Handler, request IRequest)

// IRequest is the request abstraction passed through the handler chain.
type IRequest interface {
	GetRequestLevel() int

	GetRequest() string
}

// Request carries a difficulty level plus the formatted request text.
type Request struct {
	level   int
	request string
}

// InitRequset creates a request and prefixes its text with level metadata. The
// misspelled function name is kept for compatibility with existing tests.
func InitRequset(level int, request string) *Request {
	r := &Request{
		level:   level,
		request: request,
	}
	switch r.level {
	case 1:
		r.request = "level 1 request: " + request
	case 2:
		r.request = "level 2 request: " + request
	case 3:
		r.request = "level 3 request: " + request
	}
	return r
}

func (r Request) GetRequestLevel() int {
	return r.level
}

func (r Request) GetRequest() string {
	return r.request
}

// Handler is the chain node interface. A node may handle the request or pass it
// to the next handler.
type Handler interface {
	HandleMessage(request IRequest, handler Handler, message HandleMessage)
	SetNextHandler(handler Handler)
	Response(request IRequest)
	GetLevel() int
	GetNext() Handler
}

// Primary is the first-level handler in the chain.
type Primary struct {
	level   int
	request string
	next    Handler
}

func (p *Primary) GetNext() Handler {
	return p.next
}

func (p *Primary) GetLevel() int {
	return p.level
}

func (p *Primary) HandleMessage(request IRequest, handler Handler, message HandleMessage) {
	message(p, request)
}

func (p *Primary) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Primary) Response(request IRequest) {
	fmt.Println("---level 1 request---")
	fmt.Printf(request.GetRequest())
	fmt.Println("junior engineer handled the request")
}

// InitPrimary returns the entry handler for the chain.
func InitPrimary() Handler {
	return &Primary{
		level:   DIFFICULTY_LEVEL_1,
		request: "",
	}
}

// Middle is the second-level handler in the chain.
type Middle struct {
	level   int
	request string
	next    Handler
}

// HandleMess walks the chain until it finds a handler whose level is high
// enough for the request.
func HandleMess(hand Handler, request IRequest) {

	if request.GetRequestLevel() <= hand.GetLevel() {
		hand.Response(request)
	} else {
		if hand.GetNext() != nil {
			HandleMess(hand.GetNext(), request)
		} else {
			fmt.Println("---request level ", request.GetRequestLevel(), "cannot be handled")
		}
	}
}

func (p *Middle) HandleMessage(request IRequest, handler Handler, message HandleMessage) {
	handler = p
	message(handler, request)
}

func (p *Middle) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Middle) Response(request IRequest) {
	fmt.Println("---level 2 request---")
	fmt.Printf(request.GetRequest())
	fmt.Println("mid-level engineer handled the request")
}

func (p *Middle) GetLevel() int {
	return p.level
}

func (p *Middle) GetNext() Handler {
	return p.next
}

// Senior is the third-level handler in the chain.
type Senior struct {
	level   int
	request string
	next    Handler
}

func (p *Senior) HandleMessage(request IRequest, handler Handler, message HandleMessage) {
	handler = p
	message(handler, request)
}

func (p *Senior) SetNextHandler(handler Handler) {
	p.next = handler
}

func (p *Senior) Response(request IRequest) {
	fmt.Println("---level 3 request---")
	fmt.Printf(request.GetRequest())
	fmt.Println("senior engineer handled the request")
}

func (p *Senior) GetLevel() int {
	return p.level
}

func (p *Senior) GetNext() Handler {
	return p.next
}
