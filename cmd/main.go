package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {

}

func isOperator(a string) bool {
	return (a == "+" || a == "-" || a == "*" || a == "/" || a == "(" || a == ")")
}

func operationPriority(a string) int {
	priority := 0
	if (a == "*") || (a == "/") {
		priority = 5
	}
	if (a == "+") || (a == "-") {
		priority = 4
	}
	if a == "(" {
		priority = 1
	}
	if a == ")" {
		priority = 2
	}
	return priority
}
func parseString(data string) *list.List {
	fmt.Println("===============")
	fmt.Println(data)
	data = strings.Replace(data, " ", "", -1)
	fmt.Println(data)
	l := list.New()
	stack := list.New()
	var numberTemp string
	for _, c := range data {
		currentChar := string(c)
		fmt.Println("Checking:" + currentChar)
		if isOperator(currentChar) {
			if len(numberTemp) > 0 {
				l.PushBack(numberTemp)
				numberTemp = ""
			}
			if currentChar == "(" {
				stack.PushBack(currentChar)
			} else if currentChar == ")" {
				for {
					l.PushBack(stack.Back().Value.(string))
					stack.Remove(stack.Back())
					if stack.Back().Value.(string) == "(" {
						break
					}
					stack.Remove(stack.Back())
				}
			} else {
				if stack.Len() > 0 {
					if operationPriority(currentChar) <= operationPriority(stack.Back().Value.(string)) {
						for {
							l.PushBack(stack.Back().Value.(string))
							stack.Remove(stack.Back())
							if stack.Len() == 0 {
								break
							}
						}
					}
				}
				stack.PushBack(currentChar)
			}
		} else {
			numberTemp += currentChar
		}
	}
	// 	fmt.Println("Checking:" + currentChar)
	// 	if isOperator(currentChar) {
	// 		if len(numberTemp) > 0 {
	// 			l.PushBack(numberTemp)
	// 			numberTemp = ""
	// 		}
	// 		l.PushBack(currentChar)
	// 	} else {
	// 		numberTemp += currentChar
	// 	}
	// }
	if len(numberTemp) > 0 {
		l.PushBack(numberTemp)
	}
	if stack.Len() > 0 {
		for {
			l.PushBack(stack.Back().Value.(string))
			stack.Remove(stack.Back())
			if stack.Len() == 0 {
				break
			}
		}
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // print out the elements
	}
	fmt.Println("===============")
	return l
}

func calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)
	functors := parseString(params["expression"])
	stack := list.New()
	for e := functors.Front(); e != nil; e = e.Next() {
		current := e.Value.(string)
		if !isOperator(current) {
			fmt.Println("pushed to stack:" + current)
			stack.PushBack(current)
		} else {
			fmt.Println("Performing operation:")
			b, _ := strconv.ParseFloat(stack.Back().Value.(string), 64)
			stack.Remove(stack.Back())
			a, _ := strconv.ParseFloat(stack.Back().Value.(string), 64)
			stack.Remove(stack.Back())
			fmt.Println("Operandds:" + fmt.Sprint(a) + " and " + fmt.Sprint(b))
			switch current {
			case "+":
				stack.PushBack(fmt.Sprintf("%f", a+b))
			case "-":
				stack.PushBack(fmt.Sprintf("%f", a-b))
			case "*":
				stack.PushBack(fmt.Sprintf("%f", a*b))
			case "/":
				stack.PushBack(fmt.Sprintf("%f", a/b))
			}
		}
	}
	final, _ := strconv.ParseFloat(stack.Back().Value.(string), 64)
	fmt.Println(final)
	fmt.Fprintf(w, stack.Back().Value.(string))
	//stack := list.New()
	return
}

func main() {
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/v1/health", healthCheck).Methods("HEAD")
	r.HandleFunc("/api/v1/arithmetic/{expression}", calculate).Methods("POST")
	//log.Fatal(http.ListenAndServe(":8000", r))
	port, exists := os.LookupEnv("PORT")
	if exists == false {
		fmt.Printf("NO PORT")
	}

	http.ListenAndServe(":"+port, nil)
}
