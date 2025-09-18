package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers          = "0123456789"
	SpecialChars     = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

// PasswordConfig holds the configuration for password generation
type PasswordConfig struct {
	Length           int
	IncludeLowercase bool
	IncludeUppercase bool
	IncludeNumbers   bool
	IncludeSpecial   bool
}

// generatePassword creates a password based on the given configuration
func generatePassword(config PasswordConfig) string {
	var charset string
	
	// Build character set based on configuration
	if config.IncludeLowercase {
		charset += LowercaseLetters
	}
	if config.IncludeUppercase {
		charset += UppercaseLetters
	}
	if config.IncludeNumbers {
		charset += Numbers
	}
	if config.IncludeSpecial {
		charset += SpecialChars
	}
	
	// If no character types selected, default to lowercase
	if charset == "" {
		charset = LowercaseLetters
		fmt.Println("⚠️  No character types selected, using lowercase letters only.")
	}
	
	// Generate password
	password := make([]byte, config.Length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	
	return string(password)
}

// checkPasswordStrength evaluates password strength
func checkPasswordStrength(password string) string {
	score := 0
	checks := []string{}
	
	// Length check
	if len(password) >= 8 {
		score += 2
	} else if len(password) >= 6 {
		score += 1
	}
	
	// Character type checks
	hasLower := strings.ContainsAny(password, LowercaseLetters)
	hasUpper := strings.ContainsAny(password, UppercaseLetters)
	hasNumber := strings.ContainsAny(password, Numbers)
	hasSpecial := strings.ContainsAny(password, SpecialChars)
	
	if hasLower {
		score++
		checks = append(checks, "lowercase")
	}
	if hasUpper {
		score++
		checks = append(checks, "uppercase")
	}
	if hasNumber {
		score++
		checks = append(checks, "numbers")
	}
	if hasSpecial {
		score++
		checks = append(checks, "special characters")
	}
	
	// Determine strength level
	var strength string
	var emoji string
	
	switch {
	case score >= 6:
		strength = "Very Strong"
		emoji = "🔐"
	case score >= 5:
		strength = "Strong"
		emoji = "🔒"
	case score >= 3:
		strength = "Medium"
		emoji = "🔓"
	default:
		strength = "Weak"
		emoji = "⚠️"
	}
	
	fmt.Printf("\n%s Password Strength: %s (Score: %d/7)\n", emoji, strength, score)
	fmt.Printf("Includes: %s\n", strings.Join(checks, ", "))
	
	return strength
}

// getYesNoInput gets a yes/no response from user
func getYesNoInput(prompt string) bool {
	fmt.Print(prompt + " (y/N): ")
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes"
}

// getIntInput gets an integer input from user with validation
func getIntInput(prompt string, min, max int) int {
	for {
		fmt.Printf("%s (%d-%d): ", prompt, min, max)
		var input string
		fmt.Scanln(&input)
		
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("❌ Please enter a valid number!\n")
			continue
		}
		
		if value < min || value > max {
			fmt.Printf("❌ Please enter a number between %d and %d!\n", min, max)
			continue
		}
		
		return value
	}
}

func printHeader() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("🔐 SECURE PASSWORD GENERATOR 🔐")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Generate strong, secure passwords for your accounts!")
	fmt.Println()
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	printHeader()
	
	for {
		// Get password configuration from user
		config := PasswordConfig{}
		
		fmt.Println("🛠️  Password Configuration:")
		
		// Get password length
		config.Length = getIntInput("Enter password length", 4, 128)
		
		// Get character type preferences
		config.IncludeLowercase = getYesNoInput("Include lowercase letters (a-z)?")
		config.IncludeUppercase = getYesNoInput("Include uppercase letters (A-Z)?")
		config.IncludeNumbers = getYesNoInput("Include numbers (0-9)?")
		config.IncludeSpecial = getYesNoInput("Include special characters (!@#$...)?")
		
		// Generate password
		fmt.Println("\n🔄 Generating password...")
		password := generatePassword(config)
		
		// Display results
		fmt.Println(strings.Repeat("-", 50))
		fmt.Printf("🔑 Generated Password: %s\n", password)
		fmt.Println(strings.Repeat("-", 50))
		
		// Check password strength
		checkPasswordStrength(password)
		
		// Security tips
		fmt.Println("\n💡 Security Tips:")
		fmt.Println("• Never reuse passwords across multiple accounts")
		fmt.Println("• Store passwords in a secure password manager")
		fmt.Println("• Enable two-factor authentication when available")
		fmt.Println("• Change passwords regularly for important accounts")
		
		// Ask if user wants to generate another password
		fmt.Println()
		if !getYesNoInput("Generate another password?") {
			break
		}
		fmt.Println()
	}
	
	fmt.Println("\n👋 Thank you for using Secure Password Generator!")
	fmt.Println("Stay secure! 🛡️")
}