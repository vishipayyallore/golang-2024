// utilities/CustomersUtils.go

package utilities

import (
	ent "e1shoppingcartservice/entities"
	"encoding/csv"
	"io"
	"strconv"
)

func ReadCustomers(r io.Reader) ([]ent.Customer, error) {
	customers := make([]ent.Customer, 0)
	csvReader := csv.NewReader(r)
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
