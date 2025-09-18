package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Todo represents a single todo item
type Todo struct {
	ID        int
	Task      string
	Completed bool
}

// TodoList manages a collection of todos
type TodoList struct {
	todos  []Todo
	nextID int
}

// NewTodoList creates a new TodoList
func NewTodoList() *TodoList {
	return &TodoList{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

// AddTodo adds a new todo to the list
func (tl *TodoList) AddTodo(task string) {
	todo := Todo{
		ID:        tl.nextID,
		Task:      task,
		Completed: false,
	}
	tl.todos = append(tl.todos, todo)
	tl.nextID++
	fmt.Printf("✅ Added: %s\n", task)
}

// ListTodos displays all todos
func (tl *TodoList) ListTodos() {
	if len(tl.todos) == 0 {
		fmt.Println("📝 No todos yet! Add some tasks to get started.")
		return
	}

	fmt.Println("\n📋 Your Todo List:")
	fmt.Println(strings.Repeat("=", 40))
	
	for _, todo := range tl.todos {
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("%s [%d] %s\n", status, todo.ID, todo.Task)
	}
	fmt.Println(strings.Repeat("=", 40))
}

// CompleteTodo marks a todo as completed
func (tl *TodoList) CompleteTodo(id int) {
	for i, todo := range tl.todos {
		if todo.ID == id {
			if tl.todos[i].Completed {
				fmt.Printf("⚠️  Task %d is already completed!\n", id)
				return
			}
			tl.todos[i].Completed = true
			fmt.Printf("🎉 Completed: %s\n", todo.Task)
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found\n", id)
}

// DeleteTodo removes a todo from the list
func (tl *TodoList) DeleteTodo(id int) {
	for i, todo := range tl.todos {
		if todo.ID == id {
			// Remove todo from slice
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			fmt.Printf("🗑️  Deleted: %s\n", todo.Task)
			return
		}
	}
	fmt.Printf("❌ Todo with ID %d not found\n", id)
}

// ShowStats displays todo statistics
func (tl *TodoList) ShowStats() {
	total := len(tl.todos)
	completed := 0
	
	for _, todo := range tl.todos {
		if todo.Completed {
			completed++
		}
	}
	
	pending := total - completed
	
	fmt.Println("\n📊 Todo Statistics:")
	fmt.Printf("Total: %d | Completed: %d | Pending: %d\n", total, completed, pending)
	
	if total > 0 {
		percentage := float64(completed) / float64(total) * 100
		fmt.Printf("Progress: %.1f%% complete\n", percentage)
	}
}

func printMenu() {
	fmt.Println("\n🎯 Todo List Manager")
	fmt.Println("1. Add Todo")
	fmt.Println("2. List Todos")
	fmt.Println("3. Complete Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Show Statistics")
	fmt.Println("6. Exit")
	fmt.Print("Choose an option (1-6): ")
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	todoList := NewTodoList()
	
	fmt.Println("🎉 Welcome to Todo List Manager!")
	fmt.Println("Organize your tasks and boost your productivity!")
	
	for {
		printMenu()
		choice := readInput()
		
		switch choice {
		case "1":
			fmt.Print("Enter todo task: ")
			task := readInput()
			if task == "" {
				fmt.Println("❌ Task cannot be empty!")
				continue
			}
			todoList.AddTodo(task)
			
		case "2":
			todoList.ListTodos()
			
		case "3":
			todoList.ListTodos()
			if len(todoList.todos) == 0 {
				continue
			}
			fmt.Print("Enter todo ID to complete: ")
			idStr := readInput()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("❌ Please enter a valid number!")
				continue
			}
			todoList.CompleteTodo(id)
			
		case "4":
			todoList.ListTodos()
			if len(todoList.todos) == 0 {
				continue
			}
			fmt.Print("Enter todo ID to delete: ")
			idStr := readInput()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("❌ Please enter a valid number!")
				continue
			}
			todoList.DeleteTodo(id)
			
		case "5":
			todoList.ShowStats()
			
		case "6":
			fmt.Println("👋 Thank you for using Todo List Manager!")
			fmt.Println("Stay productive! 🚀")
			return
			
		default:
			fmt.Println("❌ Invalid option! Please choose 1-6.")
		}
	}
}