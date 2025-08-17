package main

// произвольная структура Human
type Human struct {
	Name string
	Age  int
}

// структура Action которая содержит указатель на Human
// это позволяет Action использовать методы Human без копирования данных
// и изменять оригинальный объект Human если это необходимо
type Action struct {
	*Human
}

func (h *Human) CallName() string {
	return h.Name
}

func (h *Human) CallAge() int {
	return h.Age
}

// конструктор для создания human
func NewHuman(
	name string,
	age int,
) *Human {
	return &Human{
		Name: name,
		Age:  age,
	}
}

// конструктор для создания action
func NewAction(
	h *Human,
) *Action {
	return &Action{
		Human: h,
	}
}

func main() {
	h := NewHuman("John", 30) // создаем новый объект Human

	a := NewAction(h) // создаем новый объект Action, который содержит указатель на Human

	h.Age = 32 // изменяем возраст в оригинальном объекте Human

	println("Name:", a.CallName())       // John
	println("Age:", a.CallAge())         // 32, мы изменили оригинал human, и т.к action содержит указатель на него, то изменения видны и в action
	a.Age = 35                           // меняем возраст в action, что также изменяет оригинальный объект Human
	println("Updated Age:", a.CallAge()) // 35

	println("Age from Human:", h.CallAge()) // 35
}
