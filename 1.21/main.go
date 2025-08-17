package main

import "fmt"

type OldInter interface {
	Connect() string
}

type NewInter interface {
	NewConnect() string
}

type OldSystem struct{}

func (old *OldSystem) Connect() string {
	return "old"
}

type Adapter struct {
	OldSystem OldInter
}

func (a *Adapter) NewConnect() string {
	return a.OldSystem.Connect()
}

func main() {
	old := &OldSystem{}
	adapter := &Adapter{
		OldSystem: old,
	}
	fmt.Println(adapter.NewConnect())
}



// Чем хорош:
// 1. Можем коннектить несовместимые интерфейсы, например при подключении сторонних библиотек
// 2. Реалзация инкапсуляции -  юзеру адаптера не нужно знать детали реализации старого интерфейса.


// Чем плох: 
// 1. Снижает читабельность
// 2. Если в адаптере есть доп логика может снижать производительностью

// Когда используем:
// 1. При смене формата данных в API
// 2. При сборке сервисов из несогласованных частей
// 3. При подключении сторонних библиотек