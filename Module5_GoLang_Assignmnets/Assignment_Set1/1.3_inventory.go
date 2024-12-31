// Inventory Management System
package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var inventory = []Product{
	{ID: 1, Name: "Laptop", Price: 800.50, Stock: 10},
	{ID: 2, Name: "Mobile Phone", Price: 500.00, Stock: 15},
	{ID: 3, Name: "Tablet", Price: 300.00, Stock: 8},
	{ID: 4, Name: "Smart Watch", Price: 200.00, Stock: 20},
	{ID: 5, Name: "Headphones", Price: 50.75, Stock: 50},
}

// AddProduct adds a new product to the inventory
func AddProduct(id int, name string, price float64, stock int) {
	inventory = append(inventory, Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	})
	fmt.Println("Product added successfully!")
}

// UpdateStock updates the stock of a product
func UpdateStock(id int, newStock int) {
	for i, product := range inventory {
		if product.ID == id {
			if newStock < 0 {
				fmt.Println("Stock cannot be negative.")
				return
			}
			inventory[i].Stock = newStock
			fmt.Println("Stock updated successfully!")
			return
		}
	}
	fmt.Println("Product not found.")
}

// SearchProduct searches for a product by ID or Name
func SearchProduct(query interface{}) {
	switch v := query.(type) {
	case int:
		for _, product := range inventory {
			if product.ID == v {
				fmt.Printf("Product Found: ID: %d, Name: %s, Price: %.2f, Stock: %d\n", product.ID, product.Name, product.Price, product.Stock)
				return
			}
		}
	case string:
		for _, product := range inventory {
			if strings.EqualFold(product.Name, v) {
				fmt.Printf("Product Found: ID: %d, Name: %s, Price: %.2f, Stock: %d\n", product.ID, product.Name, product.Price, product.Stock)
				return
			}
		}
	default:
		fmt.Println("Invalid query type.")
		return
	}
	fmt.Println("Product not found.")
}

// DisplayInventory displays all products in the inventory
func DisplayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventory sorts products by price or stock
func SortInventory(by string) {
	if strings.EqualFold(by, "price") {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
		fmt.Println("Inventory sorted by price.")
	} else if strings.EqualFold(by, "stock") {
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
		fmt.Println("Inventory sorted by stock.")
	} else {
		fmt.Println("Invalid sort key. Use 'price' or 'stock'.")
		return
	}
	DisplayInventory()
}

func main() {
	fmt.Println("Displaying initial inventory:")
	DisplayInventory()

	fmt.Println("\nAdding a new product...")
	AddProduct(6, "Keyboard", 25.00, 30)

	fmt.Println("\nUpdating stock for Product ID 2...")
	UpdateStock(2, 25)

	fmt.Println("\nSearching for 'Tablet'...")
	SearchProduct("Tablet")

	fmt.Println("\nSearching for Product ID 4...")
	SearchProduct(4)

	fmt.Println("\nSorting inventory by stock...")
	SortInventory("stock")

	fmt.Println("\nSorting inventory by price...")
	SortInventory("price")
}
