# Hola

## Instalación
```bash
go get -u github.com/CristianAngarita/greetings

```

## Ejemplo
```go
package main

import (
	"fmt"
	"log"

	"github.com/CristianAngarita/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Cristian", "Ana", "Michael"}
	msgs, err2 := greetings.Hellos(names)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(msgs)

	fmt.Println("*****************************************")
	message, err := greetings.Hello("Cristian")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```