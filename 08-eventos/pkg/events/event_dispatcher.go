package events

import (
	"errors"
	"sync"
)

var (
	ErrHandlerAlreadyRegistered = errors.New("handler already registered")
	ErrHandlerNameNotFound      = errors.New("handler name not registered")
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

// Dispatch Com uso de multiThreading Tornar uma função assíncrona.
func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	//Verificar se existe algum evento registrado.
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			//add 1 ao waitGroup
			wg.Add(1)
			//Criar uma Threading
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	//Verificar se existe algum evento registrado.
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	//Verificar se existe algum evento registrado.
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	//Verificar se existe algum evento registrado.
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] =
					//Seleciona o evento encontrado
					append(ed.handlers[eventName][:i],
						//Adiciono todos os eventos após o evento encontrado append todos apos o i
						ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return ErrHandlerNameNotFound
}

func (ed *EventDispatcher) Clear() {
	//Sobre escreve handlers com um Map vazio.
	ed.handlers = make(map[string][]EventHandlerInterface)
}
