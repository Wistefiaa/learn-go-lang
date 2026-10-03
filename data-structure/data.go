package datastructure

import "fmt"

func main() {
	fmt.Println("Hello world")
	var a [5]int //array var name [size]type
	a[0] += 5
	var s []int //slice the same with vector you don't have to type a size
	fmt.Println(s)
	fmt.Println(len(s)) // length of slice just output of slice
	fmt.Println(cap(s)) // capacity of slice all memory of slice
	//ex
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)      // output: [10, 20, 30 ,40, 50]
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) //5

	subSlice := slice[1:3]
	fmt.Println(subSlice)      // [20, 30]
	fmt.Println(len(subSlice)) // 2
	fmt.Println(cap(subSlice)) // 4; cap will count the number after slice in this case is [20, 30, 40, 50] which is 4 that slice allocate the memory
	var mySlice []int
	mySlice = append(mySlice, 10) // append 10 in slice

	//converting array to slice
	myArray := [2]int{10, 20}
	newSlice := myArray[:]
	newSlice = append(newSlice, 30)
	fmt.Println(newSlice) // [10, 20, 30]
	/*
		map:
		1. the key must be the unique key (only 1 value per key)
		2. map will not sort the key by dictionary like c++
		3. declaration [name] := make(map[key type]value type) ex. myMap := make(map[string]int)
	*/
	myMap := make(map[string]int)
	myMap["Apple"] = 5
	//acess value from key
	fmt.Println(myMap["Apple"])
	//iterate over the map
	for key, value := range myMap {
		fmt.Printf("%s -> %d", key, value)
	}
	delete(myMap, "Apple")    // delete key
	val, ok := myMap["Apple"] // ok is check status if the key is exist in map
	if ok {
		fmt.Println(val)
	}
	/*
		struct:
		1. declaration type Student struct {
							name type
							name type
							....
							name type
						}
		2. when use the struct: var name [structname] ex. var User1 Student
		3. User1.name = [by type] ex User1.Name = "Wistefia"
		4. We can print the whole struct ex. fmt.Println(Student)
		5. or just one of their field ex. fmt.Println(Student.Name)
		6. we can use with array with [] ex. var user [3]Student, user[0].Name = "Wistefiaa"
		7. or use with map ex. students := make(map[string]Student), students["st01"] = Student{Name: "Wistefiaa", Weight: 60}
		8. use can use struct in struct ex. type Address struct{
												zipcode int
											}
											type User struct{
												Name string
												Height int
												Live Address(struct type that we already declare)
											}
	*/

}
