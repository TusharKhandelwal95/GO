package main
import "fmt"
import "maps"

// Maps --> dict, map, hash table

func main() {

	//creating a map
	m := make(map[int]string)  // key->int  value->string

	m[1] = "one"
	m[2] = "two"

	fmt.Println(m)
	fmt.Println(m[1]) //one

	// if key is not present in the map, it returns the zero value of the value type i.e ""-->string 0-->int
	fmt.Println(m[3]) //""

	//length of map
	fmt.Println(len(m)) //2

	//deleting a key from map
	delete(m, 1)
	fmt.Println(m) //map[2:two]

	//clear the map
	clear(m)


	// Another way to declare mao
	m2 := map[int]string{ 1: "happy", 2: "sad"}
	fmt.Println(m2) //map[1:happy 2:sad]

	v,ok := m2[1] //ok tell if key present and v is the value of that key

	if ok{
		fmt.Println("Good day")
		fmt.Println(v) 
	}else{
		fmt.Println("Bad day")
	}


	//Iterating over map
	for k,v := range m2{
		fmt.Println(k,v) //1 happy 2 sad
	}

	//Check if maps are equal
	m3 := map[int]string{ 1: "happy", 2: "sad"}

	// fmt.Println(m2 == m3)  --------> not like this

	fmt.Println(maps.Equal(m2, m3))
}