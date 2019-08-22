package main

/*
@author Bijaya
2019 August
This go function does a map search

*/
import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	// add a record to map
	m[`no_danny`] = []string{`Being good`, `Latte`, `Beach`}

	fmt.Println(m)

	// delete a record from map
	delete(m, `no_dr`)

	fmt.Println(m)

	// Print out the map out using "range loop"
	for x, v := range m {
		fmt.Println("This is the record for", x)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
