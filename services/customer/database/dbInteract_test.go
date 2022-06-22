package database

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateTables(t *testing.T) {
	DBStart()
	CreateTables()
}

func TestCustomerGenerate(t *testing.T) {
	DBStart()
	CreateTables()
	CreateTestCustomers()
}

func TestCustomerDelete(t *testing.T) {
	DBStart()
	err := DeleteTestCustomers()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func BenchmarkCustomerGenerateAndGet(b *testing.B) {
	DBStart()
	CreateTables()

	for i := 0; i < b.N; i++ {
		id := uuid.NewString()

		c := Customer{CustomerUUID: id}
		c.CreateCustomer()
		c.GetCustomer()
	}
}
