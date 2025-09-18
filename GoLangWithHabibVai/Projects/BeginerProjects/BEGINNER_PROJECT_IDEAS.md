# Beginner-Level Go Projects Guide

This guide provides a comprehensive list of beginner-level projects you can build using Go (Golang). These projects are designed to help you practice fundamental Go concepts while building practical applications.

## 🎯 Why Build Projects in Go?

Go is an excellent language for beginners because:
- Simple and clean syntax
- Fast compilation
- Built-in concurrency support
- Excellent standard library
- Strong typing with type inference
- Great for system programming and web development

## 📚 Current Projects in This Repository

### ✅ Already Implemented
1. **Calculator** - Basic arithmetic operations with console interface
2. **Number Guessing Game** - Interactive game with difficulty levels
3. **Rock Paper Scissors Lizard Spock** - Extended version of the classic game

## 🚀 Beginner Project Ideas by Category

### 1. Console Applications (Easy)
Perfect for learning basic Go syntax, input/output, and control structures.

#### **Text-Based Games**
- [x] **Number Guessing Game** (Implemented)
- [x] **Rock Paper Scissors** (Implemented)
- [ ] **Tic Tac Toe** - Two-player game with board display
- [ ] **Hangman** - Word guessing game with ASCII art
- [ ] **Word Counter** - Count words, characters, and lines in text
- [ ] **Password Generator** - Generate secure passwords with options
- [ ] **Dice Roller** - Simulate rolling dice with different sides
- [ ] **Quiz Game** - Multiple choice questions with scoring

#### **Utility Tools**
- [x] **Calculator** (Implemented)
- [ ] **Unit Converter** - Convert between different units (length, weight, temperature)
- [ ] **BMI Calculator** - Calculate Body Mass Index
- [ ] **Grade Calculator** - Calculate GPA/grades from scores
- [ ] **Tip Calculator** - Calculate tips and split bills
- [ ] **Currency Converter** - Convert between currencies (with fixed rates)
- [ ] **Text Encoder/Decoder** - Base64, ROT13, Caesar cipher

#### **Data Processing**
- [ ] **Contact Book** - Store and manage contacts
- [ ] **Todo List** - Command-line task management
- [ ] **Expense Tracker** - Track daily expenses
- [ ] **Simple Banking System** - Account management with basic operations
- [ ] **Library Management** - Manage books and borrowers
- [ ] **Student Grade Manager** - Store and calculate student grades

### 2. File Operations (Easy-Medium)
Learn file I/O, data persistence, and text processing.

- [ ] **File Organizer** - Sort files by type/date
- [ ] **Log File Analyzer** - Parse and analyze log files
- [ ] **CSV Reader/Writer** - Process CSV data
- [ ] **Configuration File Parser** - Read/write config files
- [ ] **Backup Tool** - Simple file backup utility
- [ ] **Duplicate File Finder** - Find duplicate files in directories
- [ ] **Text File Merger** - Combine multiple text files

### 3. Web Applications (Medium)
Introduction to web development with Go's `net/http` package.

- [ ] **Personal Portfolio Website** - Static website with your information
- [ ] **Simple Blog** - Read-only blog with static posts
- [ ] **URL Shortener** - Create short URLs for long links
- [ ] **QR Code Generator** - Generate QR codes for text/URLs
- [ ] **Weather App** - Display weather using free APIs
- [ ] **Random Quote Generator** - Display random quotes
- [ ] **Simple Chat Room** - Real-time messaging with WebSockets

### 4. APIs and JSON (Medium)
Practice working with external APIs and JSON data.

- [ ] **News Aggregator** - Fetch and display news from APIs
- [ ] **Joke Generator** - Fetch random jokes from APIs
- [ ] **Cryptocurrency Tracker** - Track crypto prices
- [ ] **GitHub Profile Viewer** - Display GitHub user information
- [ ] **Movie Database** - Search and display movie information
- [ ] **Recipe Finder** - Search for recipes using APIs
- [ ] **Translation Tool** - Translate text using translation APIs

### 5. Network Programming (Medium)
Learn about networking, protocols, and concurrent programming.

- [ ] **Port Scanner** - Scan for open ports on hosts
- [ ] **Simple TCP Chat Server** - Multi-client chat server
- [ ] **HTTP Request Tool** - Simple alternative to curl
- [ ] **Network Speed Test** - Test download/upload speeds
- [ ] **Simple Proxy Server** - Basic HTTP proxy
- [ ] **DNS Lookup Tool** - Resolve domain names to IP addresses

### 6. System Programming (Medium-Hard)
Interact with the operating system and learn system-level programming.

- [ ] **System Monitor** - Display CPU, memory, disk usage
- [ ] **Process Manager** - List and manage running processes
- [ ] **Directory Watcher** - Monitor directory changes
- [ ] **Log Rotator** - Rotate log files based on size/date
- [ ] **Cron Job Scheduler** - Schedule and run tasks
- [ ] **Environment Variable Manager** - Manage environment variables

## 🛠 Getting Started

### Prerequisites
- Go installed (version 1.16+)
- Basic understanding of programming concepts
- Text editor or IDE (VS Code, GoLand, etc.)

### Project Structure
Each project should follow this basic structure:
```
project-name/
├── main.go          # Entry point
├── README.md        # Project documentation
├── go.mod          # Go module file (if using modules)
└── internal/       # Internal packages (if needed)
    └── ...
```

### Basic Template
Here's a minimal Go program template:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, World!")
    // Your code here
}
```

## 📖 Learning Path

### Phase 1: Foundation (Weeks 1-2)
Focus on basic syntax and concepts:
1. Complete the existing calculator project
2. Build a simple text-based game
3. Create a basic file processor

### Phase 2: Intermediate (Weeks 3-4)
Add complexity and learn new concepts:
1. Build a web application
2. Work with external APIs
3. Create a utility tool with file operations

### Phase 3: Advanced Beginner (Weeks 5-6)
Tackle more challenging projects:
1. Network programming project
2. System utility
3. Multi-component application

## 💡 Project Selection Tips

### Choose projects that:
- **Interest you** - You'll be more motivated to complete them
- **Match your skill level** - Start easy and gradually increase difficulty
- **Teach new concepts** - Each project should introduce something new
- **Have practical use** - Tools you might actually use

### Project Difficulty Indicators:
- **Easy**: Uses basic Go syntax, simple input/output, basic control structures
- **Medium**: Involves external libraries, file I/O, error handling, structs
- **Hard**: Requires concurrency, advanced networking, complex algorithms

## 🎯 Skills You'll Learn

By completing these projects, you'll master:
- **Go syntax and idioms**
- **Error handling patterns**
- **Working with packages and modules**
- **File and network I/O**
- **JSON processing**
- **HTTP client/server programming**
- **Concurrent programming with goroutines**
- **Testing and debugging**
- **Code organization and project structure**

## 📚 Additional Resources

- [Go Official Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)

## 🤝 Contributing

Feel free to contribute new project ideas or implementations! Please:
1. Follow Go best practices
2. Include clear documentation
3. Add tests where appropriate
4. Update this README if adding new projects

---

**Happy Coding! 🚀**

Start with simple projects and gradually work your way up. Remember, the best way to learn Go is by building real projects!