package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

type JsonStruc struct {
	Test1 string
	Test2 int
	Test3 []int
	Test4 []string
}

func main() {
	// 	SocketServer(3333)

	var mymap = make(map[ string ] string)
	mymap[ "Test1" ] = "asdf"
	mymap[ "Test2" ] = "asdasff"
	mymap[ "Test3" ] = "agasfhdfht"

	fmt.Println(mymap)

	if val, ok := mymap[ "foo" ]; ok {
		fmt.Println(val)
	}

	if str, ok := json2Str(mymap); ok {
		fmt.Println("jsonString: " + str)
	}

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(containsInt(primes, 3))

	strs := []string{"Hallo", "Blaasdf", "dsfasd", "Bla"}
	fmt.Println(strs)
	fmt.Println(containsString(strs, "Bla"))

	dynamticArray := []int{1, 2, 3, 4}
	dynamticArray = append(dynamticArray, 5)
	fmt.Println(dynamticArray)

	jsonTest := JsonStruc{}
	jsonTest.Test1 = "asdfa"
	jsonTest.Test2 = 123
	jsonTest.Test3 = dynamticArray
	jsonTest.Test4 = strs

	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := jsonfiy(jsonTest)
	if _, err := file.Write(bytes); err != nil {
		fmt.Println(err)
		return
	}

	file.Close()

	fmt.Println("Done")

	inputFile, _ := ioutil.ReadFile("output.json")
	fmt.Println("read=" + string(inputFile))

	var jsonObj map[ string ] JsonStruc

	if err := json.Unmarshal(inputFile, &jsonObj); err != nil {
		fmt.Println(err)
	}

	fmt.Println("blabla =", jsonObj)
	fmt.Println(jsonObj[ "Test3" ])

	for key, value := range jsonObj {
		fmt.Printf("%s = %s\n", key, value)
	}
}
