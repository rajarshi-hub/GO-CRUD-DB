package models

import (
	"go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Title       string `json:"title"`
	Publication string `json:"publication"`
}

var db *gorm.DB

// Executed Automatically on Import: When a package is imported, any init() function in that package will run automatically, even before the main() function starts.
// Run Once Per Package: If a package is imported multiple times (directly or indirectly), init() will only run once for that package.
// Executed Before main(): The init() function in each imported package will run before the main() function of the main package is executed.
// Multiple init() Functions: If a package has multiple init() functions (e.g., across multiple files), Go will run them in the order they appear within the package. However, all init() functions in a package will finish before main() starts.
// Execution Order of Imports: The init() functions execute based on the dependency order of imported packages. For example, if package A imports package B, then B's init() will run before A's.

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func GetBookByID(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?", id).Find(&book) // Infers table from the type of strut
	return &book, db
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func PostBook(book *Book) *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}
