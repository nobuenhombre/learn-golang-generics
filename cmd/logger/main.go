package main

import (
	"learn/generics/internal/app/logger/mylogger"
	"learn/generics/internal/app/logger/yourlogger"
)

// GenericLogger
// При помощи GENERICS
// описывается интерфейс реализующий - Цепь вызовов.
//
// Что такое цепь вызовов?
// https://learn.javascript.ru/task/chain-calls
// https://refactoring.guru/ru/smells/message-chains
//
// реализации mylogger, yourlogger не ссылаются на данный интерфейс
type GenericLogger[T any] interface {
	WithField(string, string) T
	Info(string)
}

// DoStuff
// при помощи GENERICS можно работать с разными реализациями
// пока смущает только то,
// что при явном указании интерфейса для реализации
// IDE показывает для методов к какому интерфейсу они относятся и в коде просто ориентироваться,
// а данный подход IDE пока не до конца поддерживает и в реализациях не видно,
// что они используют данный интерфейс
//
// в данный момент я вижу реализацию работы через GENERICS.
// Но не понимаю в чем плюсы данной реализации по сравнению с прямым указанием интерфейса
// особенно учитывая что интерфейс сам по себе универсален.
func DoStuff[T GenericLogger[T]](t T) {
	t.WithField("go", "1.18").Info("is awesome")
}

func main() {
	DoStuff(mylogger.New())
	DoStuff(yourlogger.New())
}
