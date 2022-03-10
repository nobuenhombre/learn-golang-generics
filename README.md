# LEARN GENERICS IN GO LANG

Дженерики (Обобщения) — это способ,
позволяющий использовать несколько типов данных в качестве одного.

Синтаксис дженериков
--------------------
1. Дженерик типы (ограничители типов - type constraint)

**any** - является псевдонимом для interface{}

**comparable** - является псевдонимом для любых типов данных, которые можно сравнивать с помощью == или !=.

**~** - префикс для производных типов

```go
type MyInt32 int32 // тип MyInt32 является производным от int32
```

1.1 Неявное ограничение типов
```go
 ... [Name typeA | typeB | typeC] ...
```

1.2 Явное ограничение типов
```go
type GenericType interface {
    typeA | typeB | typeC
}

... [Name GenericType] ...
```

1.3 Производные дженерик типы
```go
type GenericMap[A GenericTypeA] map[int]A
```
```go
type GenericMap[A GenericTypeA] map[string]A
```
```go
type GenericMap[Key GenericTypeKey, Value GenericTypeValue] map[Key]Value
```
```go
type GenericSlice[Item GenericTypeItem] []Item
```

2. Дженерик Функции
   цель - передать в функцию дженерик типы
```go
func Action[A GenericTypeA, B GenericTypeB, C GenericTypeC](a A, b B) C {
	
}
```

вызов без указания типа
```go
c := Action(a, b)
```
вызов с явным указанием типа
```go
c := Action[int, float64](a, b)
```

3. Дженерик Структуры
   цель - передать в Методы структуры или саму структуру дженерик типы
```go
type Config[A GenericTypeA, B GenericTypeB, C GenericTypeC] struct {
	ID int
	Name string
	СС С
}

func (c Config[A,B,C]) Action(a A, b B) C {
    ...
}
``` 

4. Дженерик Интерфейсы
   цель - создать интерфейсы с методами использующими дженерик типы
```go
type Active[A GenericTypeA, B GenericTypeB, C GenericTypeC] interface {
    Action(a A, b B) C
}
```

Сравниваю Generics и Reflection
-------------------------------
Для обработки нескольких типов одного параметра также можно использовать interface{} и reflections.
Но при этом требуется дополнительная обработка ошибок которые могут возникнуть.
И обработка этих ошибок происходит в Runtime.

Generics позволяют избежать дополнительной обработки ошибок в Runtime.
Ошибки будут выявлены компилятором еще при компиляции.