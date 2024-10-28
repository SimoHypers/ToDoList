package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var urgencyTypes = []string{"very cold", "cold", "mild", "hot", "very hot"}

type Task struct {
	Name        string
	Description string
	Date        string
	Urgency     string
}

var tasks []Task

func main() {
	ToDoListMainMenu()
}

func ToDoListMainMenu() {
	fmt.Println("Welcome To The To-Do List App")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n------------------------------------------")
		fmt.Println("Choose an action:")
		fmt.Println("1) Add New Task")
		fmt.Println("2) Edit Current Tasks")
		fmt.Println("3) Exit The Program")

		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.TrimSpace(userChoice)

		switch userChoice {
		case "1":
			addTask(reader)
		case "2":
			editTasks(reader)
		case "3":
			fmt.Println("Exiting To-Do List. (GoodBye)")
			return
		default:
			fmt.Println("Pick an existing choice [1, 2, or 3]")
		}
	}
}

func addTask(reader *bufio.Reader) {
	var task Task

	fmt.Print("Task Name (max 30 characters): ")
	task.Name, _ = reader.ReadString('\n')
	task.Name = strings.TrimSpace(task.Name)

	fmt.Print("Task Due Date (YYYY-MM-DD format): ")
	for {
		dateInput, _ := reader.ReadString('\n')
		dateInput = strings.TrimSpace(dateInput)
		_, err := time.Parse("2006-01-02", dateInput)
		if err == nil {
			task.Date = dateInput
			break
		} else {
			fmt.Println("Invalid date format! Please enter the date in YYYY-MM-DD format.")
		}
	}

	fmt.Print("Task Urgency [Very Cold, Cold, Mild, Hot, Very Hot]: ")
	for {
		urgencyInput, _ := reader.ReadString('\n')
		urgencyInput = strings.ToLower(strings.TrimSpace(urgencyInput))
		if contains(urgencyTypes, urgencyInput) {
			task.Urgency = strings.Title(urgencyInput)
			break
		} else {
			fmt.Println("Invalid Urgency Type! Please enter a valid urgency.")
		}
	}

	fmt.Print("Task Description (max 200 characters): ")
	task.Description, _ = reader.ReadString('\n')
	task.Description = strings.TrimSpace(task.Description)

	tasks = append(tasks, task)
	fmt.Println("\nTask added successfully!")
	printTasks()
}

func editTasks(reader *bufio.Reader) {
	for {
		printTasks()
		fmt.Println("\nChoose an option:")
		fmt.Println("1) Edit task details")
		fmt.Println("2) Delete tasks")
		fmt.Println("3) Sort the tasks")
		fmt.Println("4) Exit to Main Menu")

		editChoice, _ := reader.ReadString('\n')
		editChoice = strings.TrimSpace(editChoice)

		switch editChoice {
		case "1":
			editTaskDetails(reader)
		case "2":
			deleteTask() // Remove 'reader' here
		case "3":
			sortTasks(reader)
		case "4":
			return
		default:
			fmt.Println("Please enter an available option [1, 2, 3, or 4]")
		}
	}
}

func editTaskDetails(reader *bufio.Reader) {
	fmt.Print("Enter the task number to edit: ")
	var taskIndex int
	fmt.Scanln(&taskIndex)
	if taskIndex <= 0 || taskIndex > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	taskIndex -= 1

	for {
		fmt.Println("\nChoose an aspect to edit:")
		fmt.Println("1) Edit Task Name")
		fmt.Println("2) Edit Task Description")
		fmt.Println("3) Edit Task Date")
		fmt.Println("4) Edit Task Urgency")
		fmt.Println("5) Exit Edit Menu")

		editChoice, _ := reader.ReadString('\n')
		editChoice = strings.TrimSpace(editChoice)

		switch editChoice {
		case "1":
			fmt.Print("Enter the new task name: ")
			newName, _ := reader.ReadString('\n')
			tasks[taskIndex].Name = strings.TrimSpace(newName)
		case "2":
			fmt.Print("Enter the new task description: ")
			newDescription, _ := reader.ReadString('\n')
			tasks[taskIndex].Description = strings.TrimSpace(newDescription)
		case "3":
			fmt.Print("Enter the new task date (YYYY-MM-DD): ")
			for {
				newDate, _ := reader.ReadString('\n')
				newDate = strings.TrimSpace(newDate)
				_, err := time.Parse("2006-01-02", newDate)
				if err == nil {
					tasks[taskIndex].Date = newDate
					break
				} else {
					fmt.Println("Invalid date format! Please enter the date in YYYY-MM-DD format.")
				}
			}
		case "4":
			fmt.Print("Enter the new task urgency: ")
			for {
				newUrgency, _ := reader.ReadString('\n')
				newUrgency = strings.ToLower(strings.TrimSpace(newUrgency))
				if contains(urgencyTypes, newUrgency) {
					tasks[taskIndex].Urgency = strings.Title(newUrgency)
					break
				} else {
					fmt.Println("Invalid Urgency Type! Please enter a valid urgency.")
				}
			}
		case "5":
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func deleteTask() {
	reader := bufio.NewReader(os.Stdin) // initialize the reader

	fmt.Print("Enter the task number to delete: ")
	input, _ := reader.ReadString('\n') // read input until newline

	var taskIndex int
	_, err := fmt.Sscanf(input, "%d", &taskIndex)
	if err != nil || taskIndex <= 0 || taskIndex > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks = append(tasks[:taskIndex-1], tasks[taskIndex:]...)
	fmt.Println("Task deleted successfully.")
}

func sortTasks(reader *bufio.Reader) {
	fmt.Println("Sort by:")
	fmt.Println("1) Urgency")
	fmt.Println("2) Alphabetical Order")
	fmt.Println("3) Exit")

	sortChoice, _ := reader.ReadString('\n')
	sortChoice = strings.TrimSpace(sortChoice)

	switch sortChoice {
	case "1":
		sort.Slice(tasks, func(i, j int) bool {
			return indexOf(urgencyTypes, strings.ToLower(tasks[i].Urgency)) < indexOf(urgencyTypes, strings.ToLower(tasks[j].Urgency))
		})
		fmt.Println("Tasks sorted by urgency.")
	case "2":
		sort.Slice(tasks, func(i, j int) bool {
			return tasks[i].Name < tasks[j].Name
		})
		fmt.Println("Tasks sorted by alphabetical order.")
	case "3":
		return
	default:
		fmt.Println("Invalid choice.")
	}
	printTasks()
}

func printTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Printf("\n%-3s %-20s %-40s %-15s %-10s\n", "No.", "Task Name", "Description", "Due Date", "Urgency")
	for i, task := range tasks {
		fmt.Printf("%-3d %-20s %-40s %-15s %-10s\n", i+1, task.Name, task.Description, task.Date, task.Urgency)
	}
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func indexOf(slice []string, item string) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}
