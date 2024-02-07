// handlers/FileHandlers.go

package handlers

import (
	ent "c1jsonmessages/entities"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	CustomersFilePath = "../data/customers.csv"
)

func GetCustomersInJsonHandler(w http.ResponseWriter, req *http.Request) {
	customers, err := readCustomers()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(customers)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(data)
}

// TODO: Implement this function in Util.go
func readCustomers() ([]ent.Customer, error) {
	f, err := os.Open(CustomersFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	customers := make([]ent.Customer, 0)
	csvReader := csv.NewReader(f)
	csvReader.Read() // throw away header
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			return customers, nil
		}
		if err != nil {
			return nil, err
		}
		var c ent.Customer
		id, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}
		c.ID = id
		c.FirstName = fields[1]
		c.LastName = fields[2]
		c.Address = fields[3]
		customers = append(customers, c)
	}
}
