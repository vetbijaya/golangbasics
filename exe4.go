package main

import "fmt"

func main() {
	/*anonymus struct syntax
		t := struct{}{}
		Topic: Anonymous struct
		 i) struct without name
		ii) syntax variable := struct{"define attributes"}{"insert value"}
	 	iii) one can define an anonymus struct only inside function or method*/

	x := struct {
		firstname string
		lastname  string
	}{
		firstname: "Peter",
		lastname:  "Holland",
	}

	fmt.Println(x.firstname)
	fmt.Println(x.lastname)
}
