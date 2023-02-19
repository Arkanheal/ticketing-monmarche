package model

type Product struct {
    Id string
    Price float32
    Name string
}

type Products struct {
    Products []Product
}
