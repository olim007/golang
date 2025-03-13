package main

//type Car struct {
//	name string
//	model string
//}
//
//type Aircraft struct {
//	name string
//	model string
//}
//
//type Vehicle interface {
//	move()
//}

func drive(v Vehicle) {
	v.move()
}

//func (c Car) move() {
//	fmt.Println("Автомобил едет")
//}
//
//func (a Aircraft) move() {
//	fmt.Println("Самолет летит")
//}

func main2() {
	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)
}