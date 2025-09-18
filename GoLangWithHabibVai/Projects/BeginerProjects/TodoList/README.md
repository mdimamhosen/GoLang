# Todo List Manager

A simple command-line todo list application built in Go. This project demonstrates struct usage, slice manipulation, and user interface design in a console application.

## 🎯 Features

- ✅ Add new todos
- 📋 List all todos with status
- 🎉 Mark todos as completed
- 🗑️ Delete todos
- 📊 View todo statistics and progress
- 🎨 User-friendly console interface with emojis

## 🚀 How to Run

1. Navigate to the project directory:
   ```bash
   cd TodoList
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Follow the on-screen menu to manage your todos!

## 📚 Learning Objectives

This project teaches:
- **Structs**: Defining custom data types (`Todo`, `TodoList`)
- **Methods**: Functions associated with structs
- **Slices**: Dynamic arrays and manipulation
- **User Input**: Reading and validating user input
- **Control Structures**: Switch statements, loops, conditionals
- **String Processing**: Formatting and manipulation
- **Error Handling**: Basic input validation
- **Code Organization**: Separating concerns into functions

## 🎮 Usage Example

```
🎯 Todo List Manager
1. Add Todo
2. List Todos
3. Complete Todo
4. Delete Todo
5. Show Statistics
6. Exit
Choose an option (1-6): 1
Enter todo task: Learn Go programming
✅ Added: Learn Go programming

Choose an option (1-6): 2

📋 Your Todo List:
========================================
❌ [1] Learn Go programming
========================================
```

## 💡 Possible Enhancements

Try adding these features to practice more Go concepts:

1. **File Persistence**: Save todos to a file
2. **Due Dates**: Add deadlines to todos
3. **Categories**: Organize todos by category
4. **Priority Levels**: Add high/medium/low priority
5. **Search Function**: Find todos by keyword
6. **Color Coding**: Use terminal colors for different states
7. **Import/Export**: Load todos from/to JSON/CSV files

## 🔧 Code Structure

- `Todo`: Struct representing a single todo item
- `TodoList`: Struct managing a collection of todos
- `main()`: Application entry point and user interface
- Helper functions for input/output and menu display

## 🎓 Go Concepts Demonstrated

- Custom types and structs
- Methods and receivers
- Slice operations (append, delete)
- Input/output with bufio
- String manipulation
- Type conversion
- Basic error handling
- Function organization

This project provides a solid foundation for understanding Go's approach to data structures, methods, and user interaction in console applications.