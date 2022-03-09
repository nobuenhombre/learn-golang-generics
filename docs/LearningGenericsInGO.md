GENERICS IN GO 1.18
===================

translate article :: https://programmingpercy.tech/blog/learning-generics-in-go
Other examples :: https://github.com/mier85/rcu

Изучение дженериков в Go
========================

Вышел Go 1.18. В нем появились дженерики. Пора узнать, как использовать эту новую функцию.

Хотя добавлены и другие приятные вещи, нет никаких сомнений в том, что реализация дженериков затмила все остальное.

Это тема, которая обсуждается очень давно. 

Многие разработчики за нее, 
НО многие против.

Многие люди находят дженерики сложными и хитрыми, но давайте раскроем тайну. 
Их не так сложно использовать, как только вы с ними ознакомитесь.

Что такое дженерики и зачем они нужны Go
----------------------------------------

дженерики — это способ, позволяющий функции принимать несколько типов данных 
в качестве одного и того же входного параметра. 
Представьте, что у вас есть функция, которая должна принимать входной параметр 
и вычитать значение второго входного параметра.

Вам нужно будет решить, какой тип данных использовать: int, int64 или float. 
Это заставит любого разработчика, использующего функцию вычитания, 
преобразовать свои значения в правильный тип данных перед ее использованием.

Другим решением было бы иметь функцию вычитания для int, 
другую функцию для чисел с плавающей запятой и т. д., 
оставляя вам несколько функций, выполняющих одну и ту же работу.

Вычитание с использованием той же функции Subtract, 
путем приведения чисел с плавающей запятой к целым числам.

Что, если бы это приведение типов или дублирование функций можно было бы пропустить? 
В приведенном выше примере мы получаем неверный результат из результата с плавающей запятой, 
поскольку мы удаляем 0,5 при преобразовании его в целое число.

Таким образом, единственным правильным решением будет дублирование функций, 

- Subtract(a,b int) int 
- Subtractfloat(a,b float32) float32.

Я говорю, что дублирующая функция - единственное решение, 
я знаю, что вы можете использовать решение для interface{} ввода и вывода. 
Мне не нравится этот хак, он подвержен ошибкам и неуклюж. 

Вы также теряете проверку ошибок во время компиляции, 
поскольку компилятор не будет знать, что вы делаете с этим интерфейсом.

Это быстро становится непригодным для сопровождения и добавляет много дополнительного кода. 
Это проблема, которую призваны решить дженерики, 
и поэтому так много разработчиков очень хотят, чтобы она была выпущена.

Первый черновик дженериков выпущен вместе с Go 1.18, и решение вышеуказанной проблемы простое.

Дженерики и как они работают — основы универсальных функций
-----------------------------------------------------------

Давайте погрузимся и изучим основы использования дженериков.

Мы начнем с обычной Subtract функции и добавим к ней общие функции, 
в то время как мы узнаем, что именно мы добавляем и почему.

Синтаксис функций заключается в том, что мы определяем функцию, указав 

```go
func FunctionName() {
	
}
```

, за которой следуют параметры функции.

Параметры функции определены внутри() и их может быть сколько угодно. 
Параметр функции определяется путем объявления имени, за которым следует тип данных. 
Например, (a int) определяет, что локальная область функций будет иметь целое число с именем a.

Метод Subtract имеет параметры функции a и b тип данных int.

```go
func Subtract(a, b int) int {
    return a - b
}
```

Обычная функция вычитания, которую мы используем в качестве основы для изучения дженериков.


Помимо параметров функции, существуют также **параметры типа**. 
**Параметры типа** определяются внутри квадратных скобок[] 
и должны быть определены до параметров функции, [](a,b int).

Вы можете определить **параметр типа** так же, как и параметр функции: за именем следует тип данных.

**Параметры типа** обычно пишутся с заглавной буквы, чтобы их было легче обнаружить.

Пример, где мы объявляем, что параметр V является целым числом, [V int]

Добавлен **параметр типа** V, этот код еще не запускается.

```go
func Subtract[V int](a, b int) int {
	return a - b
}
```

Разница между **параметром функции** и **параметром типа** заключается в том, 
что **параметр функции** доступен в локальной области действия функции.

Если вы определите V так, как мы сделали выше, вы не сможете использовать V его как переменную в функции. 
**Параметр типа** указывает только, какие типы данных представляет V.

Что мы определяем с параметром типа, 
так это то, что существует тип данных с именем V, 
который является типом int. 

Это позволяет нам использовать V как замену int в параметре функции, 
так и внутри области действия функции.

Теперь мы можем заменить тип данных для a и b на V, а также вывод функции на V.

Subtract теперь использует параметр типа V.

```go
func Subtract[V int](a, b V) V {
	return a - b
}
```

Теперь вы можете подумать, что мы ничего не добились, мы просто заменили int более сложным решением.

И вы правы, мы не закончили, приведенный выше код не скомпилируется. 
Вы не можете заменить один тип данных в параметре типа, как это сделали мы, 
если только не поместите его в интерфейс.

Если вы попытаетесь скомпилировать текущую функцию Subtract, 
вы увидите сообщение об ошибке, указывающее, что int is not an interface.

Причина этого заключается в том, что **параметр типа** ожидает в качестве значения **ограничение типа**, а не тип данных.
Ограничение - это интерфейс, которому должны соответствовать параметры функции.

Мы добавим второй тип данных в **параметр типа**  помощью | символа.

Символ вертикальной черты используется для обозначения or, что означает, что мы можем добавить к V параметру множество различных опций типа данных.

Использование | также является сокращенным способом создания нового встроенного interface.

Мы добавим int32 и float32 в качестве возможных типов данных в V. Это автоматически создаст объект interface, 
который будет использовать **ограничение типа** компилятором.

Введите параметры, используя который создает для нас ограничение типа (интерфейс)

```go
func Subtract[V int | int32 | float32](a, b V) V {
	return a - b
}
```

Чтобы было понятнее, мы можем разбить ограничение типа. 
Просто чтобы упростить понимание, а также разрешить повторное использование ограничения.

Чтобы создать ограничение, мы просто объявляем интерфейс Subtractable с типами данных, 
которые нужно включить. Это тот же синтаксис, что и сокращенное определение, 
которое мы использовали в параметре типа, поэтому мы можем его скопировать. 
Я также добавил еще несколько типов данных.

Добавление ограничения типа Subtractable к универсальной функции

```go
type Subtractable interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

func Subtract[V Subtractable](a, b V) V {
    return a - b
}
```

И последнее, прежде чем вы станете мастером дженериков. 
Вы можете применить параметр типа при вызове универсальной функции, 
чтобы установить используемый тип данных, это будет иметь больше смысла, 
когда мы перейдем к более продвинутому использованию.

Способ добавить это — снова использовать параметр типа, 
но на этот раз перед вызовом функции.

```go
result := Subtract[int](a, b)
```

Ура! Теперь мы можем даже сказать функции, какой тип данных использовать, отлично!

Теперь здесь еще есть проблема. Что делать, если у вас есть типизированные данные, 
которые являются псевдонимом/производным от любого из типов внутри Subtractable?

Вы не сможете использовать свои типы данных так, как мы в настоящее время объявили вычитаемыми.

```go
// create a custom int derived from int
type MyOwnInteger int

var myInt MyOwnInteger

myInt = 10

Subtract(myInt, 20) // This will Crash, Since myInt is not Subtractable
```

Чтобы решить эту проблему, команда Go также добавила ~ограничение (тильда), 
указывающее, что разрешен любой тип, производный от данного типа. 
В приведенном ниже примере мы позволяем MyOwnInteger быть частью Subtractable, 
применяя ~перед ним.

Использование ~(тильда), чтобы позволить типам псевдонимов быть частью ограничения типа

```go
type Subtractable interface {
	~ int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

type MyOwnInteger int

func Subtract[V Subtractable](a, b V) V {
	return a - b
}

func main() {
    var myint MyOwnInteger
    myint = 10

    result := Subtract(myint, 20)
}
```

Общие типы, любые и сопоставимые (Any & Comparable)
---------------------------------------------------

Мы разобрались в **дженерик функциях**.
Пора разобраться в **дженерик типах**.

Чтобы создать общий тип, вы должны определить новый тип и указать для него **параметр типа**. 
Чтобы узнать это, мы создадим Results тип, который является срезом любого типа данных, 
найденного в ранее созданном Subtractable **ограничении**.

Результаты общего типа, которые представляют собой срез вычитаемых типов данных.

```go
type Results[T Subtractable] []T
```

Синтаксис не должен быть новым, и я надеюсь, что теперь вы понимаете, 
что происходит в определении универсального типа. Он такой же, как и в предыдущих примерах.

В примере мы создаем срез, в котором записи данных будут иметь тип данных Subtractable, 
тем самым сообщая компилятору, что при инициализации среза результатов необходимо определить тип, 
включенный в интерфейс Subtractable.

Чтобы использовать результаты, мы обновим основную функцию и заметим, 
как мы должны определить тип данных, который Results срез будет использовать при создании переменной.

Использование универсального типа результатов (срез универсальных типов данных)

```go
type Results[T Subtractable] []T

func main(){
	var a int = 20
	var b int = 10

	result := Subtract(a, b)

	var c float32 = 20.5
	var d float32 = 10.3

	result2 := Subtract[float32](c, d)
	result3 := Subtract(c, d)
	
	fmt.Println("Result: ", result)
	fmt.Println("Result2: ", result2)
	fmt.Println("Result3: ", result3)
	
	// Create a generic Results type, and set the instantitation to int
	var resultStorage Results[int]

	resultStorage = append(resultStorage, result)

	fmt.Println("ResultStorage: ", resultStorage)
}
```

Теперь я уже знаю, о чем вы думаете. Можем ли мы сказать, что Resultsследует использовать Subtractableограничение типа?

```go
var resultStorage Results[Subtractable]
```

К сожалению, Subtractable — это ограничение типа, и мы не можем использовать его при инициализации Results.

Это вызовет ошибку компиляции interface contains type constraints.

Что мы можем сделать, так это использовать недавно введенный any, чтобы позволить Results хранить any тип данных. 

**any является псевдонимом для interface{}**

ResultStorage теперь может хранить все значения из Subtract.

```go
// Results is a array of results, the data types can be any
type Results[T any] []T

// Subtractable is a type constraint that defines subtractable datatypes to be used in generic functions
type Subtractable interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

func main(){
	var a int = 20
	var b int = 10

	result := Subtract(a, b)

	var c float32 = 20.5
	var d float32 = 10.3

	result2 := Subtract[float32](c, d)
	result3 := Subtract(c, d)
	
	fmt.Println("Result: ", result)
	fmt.Println("Result2: ", result2)
	fmt.Println("Result3: ", result3)
	// Create a generic Results type, and set the instantitation to int
	var resultStorage Results[any]
	// We can now append all values
	resultStorage = append(resultStorage, result, result2, result3)

	fmt.Println("ResultStorage: ", resultStorage)
}
```

Он еще не идеален, это первый набросок. 
Я надеюсь, что мы увидим способ добавления Subtractable в будущем, 
чтобы избежать необходимости использовать, any поскольку это позволяет нам добавлять все виды типов данных.

Существует еще один новый тип с именем comparable, 
который представляет собой ограничение типа, 
которое выполняется любым типом данных, который можно сравнивать с помощью == или !=.

Важно ознакомиться с этими словами, так как вы, вероятно, увидите, 
что они появляются повсюду по мере того, как дженерики становятся более знакомыми сообществу.


Ограничения интерфейса и общие структуры (Interface Constraints & Generic Structs)
----------------------------------------------------------------------------------
До сих пор мы использовали только одиночные ограничения и выходы. 
Мы только использовали, 

- type constraints однако можно использовать 
- и интерфейсы в качестве ограничений.

Давайте попробуем создать общую функцию, которая принимает интерфейс с именем Moveable. 
Функция просто активирует Move для типов ввода, любая структура, выполняющая этот интерфейс, 
должна иметь возможность Move.

Move — общая функция с ограничением на то, что вход v должен быть подвижным.

```go
type Moveable interface{
	func Move(meters int)
}

func Move[V Moveable](v V, meters int) {
	v.Move(meters)
}
```

Вы должны быть знакомы с синтаксисом, мы создаем ограничение, 
говорящее, что тип V должен быть Moveable 
и что входной параметр v является типом V. 
Мы также создадим a Person и a, Car чтобы попробовать это.

Общий способ перемещения нескольких структур.

```go
// Moveable is a interface that is used to handle many objects that are moveable
type Moveable interface {
	Move(int)
}
// Person is a person, implements Moveable
type Person struct {
	Name string
}

func (p Person) Move(meters int) {
	fmt.Printf("%s moved %d meters\n", p.Name, meters)
}
// Car is a test struct for cars, implements Moveable
type Car struct {
	Name string
}

func (c Car) Move(meters int) {
	fmt.Printf("%s moved %d meters\n", c.Name, meters)
}
// Move is a generic function that takes in a Moveable and moves it
func Move[V Moveable](v V, meters int) {
	v.Move(meters)
}

func main(){
	p := Person{Name: "John"}
	c := Car{Name: "Ferrari"}
	// Since the V paramter accepts Moveable, we can now call Move on Both Structs
	Move(p, 10)
	Move(c, 20)
}
```

До сих пор мы использовали только один общий параметр для простоты. 
Но помните, что вы можете добавить несколько и вывести несколько.

Давайте объединим ограничения и с Moveable функцией, 
что позволит пользователям добавлять значение, которое мы используем для расчета того, 
как далеко осталось до цели. SubtractableMoveDistance

Чтобы добавить больше ограничений типа, 
просто добавьте параметры внутри [] так же, как с обычными параметрами. 
Мы добавим Subtractable тип, определенный как S и вместо того, чтобы принимать meters его как Int, 
мы позволим ему быть S

Вот как наша Move функция будет выглядеть с обоими Type Constraints

```go
func Move[V Moveable, S Subtractable](v V, distance S, meters S) S
```

Однако само по себе это изменение заставило бы компилятор расплакаться, 
потому что Move функция принимает Int и Moveable определяет, 
что именно так должен работать Move. 
Итак, нам нужно сделать Moveaccept Subtractable.

Добавление Subtractable к функции Move и изменение Int на Subtractable

```go
// Moveable is a interface that is used to handle many objects that are moveable
type Moveable interface {
	Move(Subtractable)
}

// Person is a person, implements Moveable
type Person struct {
	Name string
}

func (p Person) Move(meters Subtractable) {
	fmt.Printf("%s moved %d meters\n", p.Name, meters)
}
// Car is a test struct for cars, implements Moveable
type Car struct {
	Name string
}

func (c Car) Move(meters Subtractable) {
	fmt.Printf("%s moved %d meters\n", c.Name, meters)
}
// Move is a generic function that takes in a Moveable and moves it
func Move[V Moveable, S Subtractable](v V, distance S, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}
```

Это выглядит потрясающе, правда?! 
К сожалению, приведенный выше пример не является рабочим примером, 
это всего лишь псевдокод того, чего мы хотим достичь.

Приведенный выше код не скомпилируется, 
компилятор будет сердито орать на вас, потому что мы используем Type Constraint внутреннюю часть, 
Interface которая не разрешена, помните?

Однако есть способ получить то, что мы хотим, Generic Structs. 

Общие структуры — это структуры, тип данных которых определяется во время инициализации.

Нам нужно определить S тип interface так же, как мы это делаем для универсальных функций. 
Просто добавляя [S Subtractable] к interface объявлению, мы говорим, 
что Struct не только нуждается в том же наборе методов, 
чтобы быть частью интерфейса, но также должен быть Generic Struct.

Добавление ограничения типа в интерфейс

```go
// Moveable is a interface that is used to handle many objects that are moveable
// To implement this interface you need a Generic Type with a Constraint of Subtractable
type Moveable[S Subtractable] interface {
	Move(S)
}
```

Если это правило для контракта, давайте добавим это к Car и Person также. 
Теперь эти структуры будут Generic. Это означает, что при создании объекта вы также должны определить тип данных, 
который будет использоваться для этого объекта.

Вот как объявлять общие структуры в Go

```go
// Car is a Generic Struct with the type S to be defined
type Car[S Subtractable] struct {
    Name string
}

// Person is a Generic Struct with the type S to be defined
type Person[S Subtractable] struct {
    Name string
}
```


Чтобы создать наши Car и Person, нам нужно указать, какие типы данных они используют. 
Это делается так же, как и все общие функции, с расширением []. 
Помните, они называются Type Arguments.

Инициализация двух универсальных структур, назначение типа данных с помощью []

```go
func main(){
	// John is travelling to his Job
	// His car travel is counted in int
	// And his walking in Float32
	p := Person[float64]{Name: "John"}
	c := Car[int]{Name: "Ferrari"}
}
```

Обратите внимание, что мы сейчас работаем с Person[S Subtractable] и не только Person, 
поэтому все методы тоже должны использовать эту инициализацию.

Чтобы сделать общую структуру частью интерфейса Moveable, необходимо установить [S]

```go
// Person is a struct that accepts a type definition at initialization
// And uses that Type as the data type for meters as input
func (p Person[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", p.Name, meters)
}
func (c Car[S]) Move(meters S) {
	fmt.Printf("%s moved %d meters\n", c.Name, meters)
}
```

Мы также должны сделать так, чтобы Move функция теперь принимала a Generic Movable путем обновления, 
а также принимала ограничение Subtractable.

Перемещаемый объект должен иметь определенный тип [S]

```go
// Move is a generic function that takes in a Generic Moveable and moves it
func Move[V Moveable[S], S Subtractable](v V, distance S, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}
```

Мы готовы создать main функцию, которая использует все то, что мы сделали.

Первый Move вызов с использованием Car легко понять, 
это потому, что мы используем, int и компилятор будет использовать его по умолчанию.

Однако второй вызов более сложен, так как теперь мы хотим использовать float64 тип данных. 
Для этого нам нужно добавить к Move() вызову регулярку [] в том месте, где мы определяем аргументы типа.

В этом случае Moveable он будет Person инициализирован как float64. И тип данных для Subtractable также будет float64. 
Таким образом, определение типа будет [Person[float64], float64]Move().

Вызов функции перемещения как для автомобиля, так и для человека

```go
func main(){
	// John is travelling to his Job
	// His car travel is counted in int
	// And his walking in Float32
	p := Person[float64]{Name: "John"}
	c := Car[int]{Name: "Ferrari"}
	
	
	// John has 100 miles to his job
	milesToDestination := 100
	// John moves with the Car
	distanceLeft := Move(c, milesToDestination, 95)
	// John has 5 miles left to walk after parking (phew)
	fmt.Println(distanceLeft)
	fmt.Println("DistanceType: ", reflect.TypeOf(distanceLeft))

	// Jumps out of Car and Walks to Building
	// Again we need to define the data type to use for the Move function, or else it will default to int
	// So here we have to tell Move to initialize with a Person type, with a float64 value,
	// And that the Subtract data type is also float64
	// [Move[float64], float64]
	// distanceLeft is also a INT, since Move defaulted to int in previous call, so we need to convert it
	newDistanceLeft := Move[Person[float64], float64](p, float64(distanceLeft), 5)
	fmt.Println(newDistanceLeft)
	fmt.Println("DistanceType: ", reflect.TypeOf(newDistanceLeft))

}
```


Совет: подумайте о том, в каком порядке вы размещаете аргументы типа. 
Мы можем избежать этого Move[Person[float64], float64], изменив порядок аргументов типа в Move.

Это будет работать благодаря тому, что компилятор и среда выполнения определяют типы данных.

Порядок аргументов типа может быть синтаксически более удобным

```go
// Move is a generic function that takes in a Moveable and moves it
// Subtractable is placed infront of Moveable instead
func Move[S Subtractable, V Moveable[S]](v V, distance S, meters S) S {
	v.Move(meters)
	return Subtract(distance, meters)
}
// You can now do this instead of having to define Person 
newDistanceLeft := Move[float64](p, float64(distanceLeft), 5)
```

Заключение
----------

Поздравляем, если вы попали сюда, то теперь вы являетесь экспертом в области дженериков!

Я надеюсь, что вы найдете их полезными, 
я знаю, что многие люди очень ждали этого релиза. 
Многие создают свои библиотеки для сортировки срезов, карт с использованием очередей и т.д.

Вероятно, скоро появится много новых пакетов и новых API для старых библиотек.

Многие варианты использования хороши с дженериками, 
но иногда вместо этого легко сделать что-то слишком сложным. 
Старайтесь использовать их только тогда, когда для них есть реальный вариант использования.

Если вам нравится то, что я пишу, не пропустите мое руководство по нечеткому тестированию в Go !

Иди туда и будь универсальным!

Спасибо за чтение и не стесняйтесь обращаться ко мне в любой из моих социальных сетей