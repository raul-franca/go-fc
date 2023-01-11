package events

import (
	"time"
)

//Evento (Carregar dados)
//Operações que serão executadas quando um evento é chamado
//Gerenciador dos nossos eventos/operações
//Registrar os eventos e suas operações
//Despachar / Fire no evento para suas operações sejam executadas

// EventeInterface Evento (Carregar dados)
type EventeInterface interface {
	GetName() string
	GetDataTime() time.Time
	GetPayLoad() interface{} // dados do evento
}

type EventHandlerInterface interface {
	Handle(event EventeInterface)
}

type EventDispacherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Dispacher(event EventeInterface) error
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error
}
