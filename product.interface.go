package main

type product interface {
	getId() int
	getName() string
	getQuantity() int
	getPrice() float64
	getTotalValue() float64
}