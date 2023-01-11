package events

import (
	"sync"
	"time"
)

//Evento (Carregar dados)
//Operações que serão executadas quando um evento é chamado
//Gerenciador dos nossos eventos/operações
//Registrar os eventos e suas operações
//Despachar / Fire no evento para suas operações sejam executadas

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
