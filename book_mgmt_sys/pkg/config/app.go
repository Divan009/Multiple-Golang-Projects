package config

import "github.com/jinzhu/gorm"

// Declare a package-level variable 'db' which is a pointer to a gorm.DB object.
// This will hold the database connection once it is initialized.
var (
	db *gorm.DB
)

// Connect is a function that initializes the database connection using GORM.
// This function opens a connection to the MySQL database with the provided connection string.
func Connect() {
	// Attempt to open a connection to the MySQL database using the provided DSN (Data Source Name).
	d, err := gorm.Open("mysql", "diva:DivaPswd1@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		// If there's an error opening the connection, panic and terminate the program.
		panic(err)
	}
	// If the connection is successful, assign the database connection to the package-level variable 'db'.
	db = d
}

// GetDB is a function that returns the current database connection.
// This function allows other parts of the program to access the database connection.
func GetDB() *gorm.DB {
	// Return the database connection stored in the package-level variable 'db'.
	return db
}
