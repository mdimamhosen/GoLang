# Secure Password Generator

A command-line application that generates secure passwords with customizable options. This project demonstrates random number generation, string manipulation, and user input validation in Go.

## 🎯 Features

- 🔐 Generate passwords with custom length (4-128 characters)
- 🎛️ Choose character types: lowercase, uppercase, numbers, special characters
- 🔍 Password strength analysis with scoring system
- 💡 Built-in security tips and best practices
- 🎨 User-friendly interface with emojis and clear formatting
- ✅ Input validation and error handling

## 🚀 How to Run

1. Navigate to the project directory:
   ```bash
   cd PasswordGenerator
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Follow the prompts to configure and generate your password!

## 📚 Learning Objectives

This project teaches:
- **Random Number Generation**: Using `math/rand` for secure randomness
- **String Manipulation**: Building and analyzing strings
- **Structs**: Organizing configuration data
- **Constants**: Defining character sets
- **Input Validation**: Ensuring user input is correct
- **Control Flow**: Loops, conditionals, and user interaction
- **Functions**: Modular code organization
- **Type Conversion**: Converting between string and int

## 🎮 Usage Example

```
==================================================
🔐 SECURE PASSWORD GENERATOR 🔐
==================================================
Generate strong, secure passwords for your accounts!

🛠️  Password Configuration:
Enter password length (4-128): 16
Include lowercase letters (a-z)? (y/N): y
Include uppercase letters (A-Z)? (y/N): y
Include numbers (0-9)? (y/N): y
Include special characters (!@#$...)? (y/N): y

🔄 Generating password...
--------------------------------------------------
🔑 Generated Password: K3#mP9$wX2@nQ7!z
--------------------------------------------------

🔐 Password Strength: Very Strong (Score: 7/7)
Includes: lowercase, uppercase, numbers, special characters
```

## 🛡️ Security Features

- **Character Set Diversity**: Supports multiple character types
- **Strength Analysis**: Evaluates password security
- **Best Practices**: Provides security tips
- **Customizable Length**: Allows passwords up to 128 characters
- **Fallback Protection**: Defaults to secure options if no types selected

## 💡 Possible Enhancements

Try adding these features to practice more Go concepts:

1. **Password History**: Remember recently generated passwords
2. **Batch Generation**: Generate multiple passwords at once
3. **File Export**: Save passwords to encrypted files
4. **Pronounceable Passwords**: Generate easier-to-remember passwords
5. **Exclude Similar Characters**: Option to avoid confusing characters (0/O, 1/l)
6. **Custom Character Sets**: Allow users to define their own character sets
7. **Entropy Calculation**: Show password entropy in bits
8. **Copy to Clipboard**: Automatically copy generated passwords

## 🔧 Code Structure

- `PasswordConfig`: Struct holding generation preferences
- `generatePassword()`: Core password generation logic
- `checkPasswordStrength()`: Password security analysis
- `getYesNoInput()` / `getIntInput()`: Input validation helpers
- `main()`: Application flow and user interface

## 🎓 Go Concepts Demonstrated

- Constants and character set definitions
- Struct-based configuration
- Random number generation and seeding
- String building and manipulation
- Input validation and error handling
- Function design and modularity
- User interface design for console applications
- Boolean logic and conditional statements

This project provides excellent practice with Go's approach to randomness, string handling, and user interaction while building a practically useful tool.