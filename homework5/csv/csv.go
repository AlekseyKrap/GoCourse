package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Init() {

	dir := "/homework5/csv/test.csv"

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	bs, err := ioutil.ReadFile(pwd + dir)
	if err != nil {
		fmt.Println("err ReadFile:", err)
		return
	}

	r := csv.NewReader(strings.NewReader(string(bs)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}
