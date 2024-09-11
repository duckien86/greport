package mytest

import (
	utils "2ndbrand-api/common"
	"math/rand"
	"strings"
)

func RandomCreateUserStruct(via, info string) map[string]string {

	// Generate random data for creating a user
	user := map[string]string{
		"username":   info,
		"password":   RandomString(10),
		"first_name": RandomFirstName(),
		"last_name":  RandomLastName(),
		"phone":      info,
		"email":      RandomEmail(),
		"verify":     via,
	}

	return user
}

func RandomVia() string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible verification methods
	vias := []string{"sms", "email"}

	// Generate a random verification method
	via := vias[rand.Intn(len(vias))]

	return via
}

func RandomFirstName() string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible first names
	firstNames := []string{"John", "Jane", "Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Henry"}

	// Generate a random first name
	firstName := firstNames[rand.Intn(len(firstNames))]

	return firstName
}

func RandomLastName() string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible last names
	lastNames := []string{"Smith", "Johnson", "Williams", "Jones", "Brown", "Davis", "Miller", "Wilson", "Moore", "Taylor"}

	// Generate a random last name
	lastName := lastNames[rand.Intn(len(lastNames))]

	return lastName
}

func RandomPhoneNumber() string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible phone number prefixes
	prefixes := []string{"091", "090", "098", "097", "096", "094", "093"}

	// Generate a random prefix
	prefix := prefixes[rand.Intn(len(prefixes))]

	// Generate a random number with the specified length
	number := prefix + utils.GenerateOTP(7)

	return number
}

func RandomEmail() string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible email domains
	domains := []string{"gmail.com", "yahoo.com", "hotmail.com", "outlook.com"}

	// Generate a random username
	username := RandomString(10)

	// Select a random domain
	domain := domains[rand.Intn(len(domains))]

	// Combine the username and domain to form the email address
	email := username + "@" + domain

	return email
}

func RandomString(length int) string {
	// rand.Seed(time.Now().UnixNano())

	// List of possible characters for the random string
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate a random string of the specified length
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(characters[rand.Intn(len(characters))])
	}

	return sb.String()
}
