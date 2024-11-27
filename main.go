package main

func main() {
	age := getAge()
	canDrink(age)
}

func getAge() *int {
	age := 18
	return &age
}

func canDrink(age *int) bool {
	return *age >= 18
}
