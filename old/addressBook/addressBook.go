package addressBook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type addressBook struct {
	addresses map[string]int
}

func (adr *addressBook) SetAddress(name string, number int) {

	const dir = "C:\\Go\\"

	adr.addresses[name] = number
	jsont, err := json.Marshal(adr.addresses)
	if err != nil {
		log.Println(err)
		return
	}

	content := []byte(string(jsont))

	tmpfile, err := ioutil.TempFile(dir, "addressBook")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

}

func ExecAddressBook() {
	var addressBook = addressBook{addresses: make(map[string]int)}

	addressBook.SetAddress("Alex", 89991234567)

	fmt.Println("addressBook", addressBook.addresses)

}
