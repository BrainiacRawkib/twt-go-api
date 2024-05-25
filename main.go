package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Quantity   int    `json:"quantity"`
	BookFormat string `json:"bookFormat"`
}

var books = []book{
	{ID: "1", Title: "I  Search of Lost Time", Author: "Marcel Proust", Quantity: 2, BookFormat: "PDF"},
	{ID: "2", Title: "The Great Gatsby", Author: "Azeez Khan", Quantity: 1, BookFormat: "EPUB"},
	{ID: "3", Title: "War and Peace", Author: "Adam Kareem", Quantity: 4, BookFormat: "DOCX"},
	{ID: "4", Title: "You don't know Javascript", Author: "Kyle Thompson", Quantity: 3, BookFormat: "PDF"},
	{ID: "5", Title: "You don't know Go", Author: "Ken Thompson", Quantity: 6, BookFormat: "EPUB"},
	{ID: "6", Title: "Building APIs with Django and MongoDB", Author: "Brainiac", Quantity: 5, BookFormat: "PDF"},
	{ID: "7", Title: "Eloquent Python-Go-JS", Author: "Brainiac", Quantity: 7, BookFormat: "EPUB"},
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book does not exist")
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func updateBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter."})
		return
	}

	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	router.GET("/books", getBooks)
	router.POST("/books", addBook)
	router.GET("/books/:id", bookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/update", updateBook)
	router.Run("localhost:9000")
}
