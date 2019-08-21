package main
import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

func main() {
	person := Person { "tom", 12 }
	if b, err := json.Marshal(person); err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("value: %s", b)
	}
}