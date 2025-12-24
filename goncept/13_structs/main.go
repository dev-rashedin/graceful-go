// STRUCTS {STRUCTURES}
// Allows us to store multiple values of different types in a single variable called a Struct




package main

import "fmt"


func main(){

	type Book struct {
	Title string
	Author string
	Subject string
	BookId int
	Price float64
	AddedToCart bool
}

	book1 := Book{
		Title: "The Lord of the Rings",
		Author: "JRR Tolkien",
		Subject: "Fantasy",
		BookId: 123456,
		Price: 19.99,
		AddedToCart: true,
	}


fmt.Println("The title of the book is", book1.Title)
fmt.Println("The author of the book is", book1.Author)
fmt.Println("The subject of the book is", book1.Subject)
fmt.Println("The book id is", book1.BookId)
fmt.Println("The price of the book is", book1.Price)
fmt.Println("The book is added to cart", book1.AddedToCart)

}