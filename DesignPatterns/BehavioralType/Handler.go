package BehavioralType

import "fmt"

const (
	DIFFICULTY_LEVEL_1 = 1
	DIFFICULTY_LEVEL_2 = 2
	DIFFICULTY_LEVEL_3 = 3
)

type HandleMessage func(hand Handler, request IRequest)

type IRequest interface {
	GetRequestLevel() int

	GetRequest() string
}

type Request struct {
	level   int
	request string
}

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

type Handler interface {
	HandleMessage(request IRequest, handler Handler, message HandleMessage)
	SetNextHandler(handler Handler)
	Response(request IRequest)
	GetLevel() int
	GetNext() Handler
}

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

func InitPrimary() Handler {
	return &Primary{
		level:   DIFFICULTY_LEVEL_1,
		request: "",
	}
}

type Middle struct {
	level   int
	request string
	next    Handler
}

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
