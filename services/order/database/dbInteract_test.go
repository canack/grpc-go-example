package database

import (
	"github.com/google/uuid"
	"testing"
)

func TestCreateTables(t *testing.T) {
	DBStart()
	CreateTables()
}

func TestOrderGenerate(t *testing.T) {
	DBStart()
	CreateTables()
	CreateTestOrders()
}

func TestOrderDelete(t *testing.T) {
	DBStart()
	err := DeleteTestOrders()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func BenchmarkOrderGenerateAndGet(b *testing.B) {
	DBStart()
	CreateTables()

	for i := 0; i < b.N; i++ {
		id := uuid.NewString()
		username := uuid.NewString()
		pid := uuid.NewString()

		o := Order{OrderUUID: id, CustomerUUID: username,
			Product: Product{
				ProductUUID: pid,
				Name:        "Toy car",
				ImageUrl:    "https://example.com/deneme.jpg",
			}}
		o.CreateOrder()
		o.GetOrder()
	}
}
