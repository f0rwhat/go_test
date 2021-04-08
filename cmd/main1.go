// package main

// import "fmt"

// type point struct {
// 	X int
// 	Y int
// }
// type Point interface {
// 	GetCords() (int, int)
// }

// func (p *point) GetCords() (int, int) {

// 	return p.X, p.Y
// }
// func zero(xPtr *int) {
// 	*xPtr = 0
// }
// func addOnlyPositive(a int, b int) (int, error) {
// 	if a <= 0 || b <= 0 {
// 		return 0, fmt.Errorf("something not positive")
// 	}
// 	return a + b, nil
// }
// func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println("Сумма равна ", sum)

// 	if true {
// 		fmt.Println("True")
// 	}
// 	if false {
// 		fmt.Println("a")
// 	}
// 	s, err := addOnlyPositive(3, 7)
// 	if err == nil {
// 		fmt.Println("Сумма равна ", s)
// 	}

// 	zero(&s)
// 	fmt.Println("Сумма равна ", s)

// 	point := Point{3, 3}
// 	points := [3]Point{Point{1, 1}, Point{2, 2}, point}
// 	fmt.Println(points)
// }
