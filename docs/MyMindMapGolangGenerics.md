Дженерики (Обобщения) — это способ, 
позволяющий функции принимать несколько типов данных 
в качестве одного и того же входного параметра.


Синтаксис дженериков
--------------------
1. Дженерик типы (ограничители типов - type constraint)

**any** - является псевдонимом для interface{}

**comparable** - является псевдонимом для любых типов данных, которые можно сравнивать с помощью == или !=.

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
цель - создать интерфейсы с дженерик типами
```go
type Active[A GenericTypeA, B GenericTypeB, C GenericTypeC] interface {
    Action(a A, b B) C
}
```