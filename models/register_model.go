package models

type Model_Interface struct {
	Model interface{}
}

func RegisterModels() []Model_Interface {
	return []Model_Interface{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: Category{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Order{}},
		{Model: OrderCustomer{}},
		{Model: OrderItem{}},
		{Model: Payment{}},
		{Model: Shipment{}},
	}
}
