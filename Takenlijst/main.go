package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Task struct {
	Description string `json:"description"`
}

func main() {
	tasks := loadTasks()

	for {
		listTasks(tasks)
		fmt.Println("\nWat wil je doen?\n1. Taak toevoegen\n2. Taak verwijderen\n3. Afsluiten\nKeuze: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addTask(&tasks)
		case 2:
			deleteTask(&tasks)
		case 3:
			saveTasks(tasks)
			fmt.Println("Tot ziens!")
			return
		default:
			fmt.Println("Ongeldige keuze. Probeer opnieuw.")
		}
	}
}

func loadTasks() []Task {
	filename := filepath.Join(".", "taken.json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Fout bij het lezen van taken:", err)
		return []Task{}
	}
	return tasks
}

func saveTasks(tasks []Task) {
	filename := filepath.Join(".", "taken.json")
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Fout bij het opslaan van taken:", err)
		return
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Fout bij het opslaan van taken:", err)
	}
}

func addTask(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Voer de nieuwe taak in: ")
	scanner.Scan()
	task := Task{Description: scanner.Text()}
	*tasks = append(*tasks, task)
	fmt.Println("Taak toegevoegd: ", task.Description)
}

func deleteTask(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welke taak verwijderen? (geef het nummer): ")
	scanner.Scan()
	input := scanner.Text()

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(*tasks) {
		fmt.Println("Ongeldige index.")
		return
	}

	index--
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
	fmt.Println("Taak verwijderd.")
}

func listTasks(tasks []Task) {
	fmt.Println("===== Lijst met taken =====")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Description)
	}
}
