package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password as input and returns its bcrypt hash.
// It uses the bcrypt.GenerateFromPassword function to generate the hash with a cost of 14.
// The function returns the hashed password as a string and any error encountered during the hashing process.
func HashPassword(password string) (string, error) {
    // Convert the password string to a byte slice and generate the bcrypt hash with a cost of 14.
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    
    // Convert the resulting byte slice to a string and return it along with any error.
    return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its possible plaintext equivalent.
// It returns true if the password matches the hash, and false otherwise.
//
// Parameters:
//   - password: The plaintext password to be compared.
//   - hash: The hashed password to compare against.
//
// Returns:
//   - bool: True if the password matches the hash, false otherwise.
func CheckPasswordHash(password ,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
} 