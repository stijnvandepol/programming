package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Wachtwoord lezen uit password.txt
	password, err := readPasswordFromFile("password.txt")
	if err != nil {
		fmt.Println("Fout bij het lezen van het wachtwoord:", err)
		return
	}

	// Verbinding maken met de database
	db, err := sql.Open("mysql", fmt.Sprintf("%s@tcp(studmysql01.fhict.local:3306)/dbi530600", password))
	if err != nil {
		fmt.Println("Fout bij het verbinden met de database:", err)
		return
	}
	defer db.Close()

	// Array van toegestane kentekens
	var allowedPlates []string

	// Query uitvoeren om de toegestane kentekens op te halen
	rows, err := db.Query("SELECT plate FROM plates")
	if err != nil {
		fmt.Println("Fout bij het uitvoeren van de query:", err)
		return
	}
	defer rows.Close()

	// Resultaten verwerken
	for rows.Next() {
		var plate string
		err := rows.Scan(&plate)
		if err != nil {
			fmt.Println("Fout bij het scannen van de rij:", err)
			return
		}
		allowedPlates = append(allowedPlates, plate)
	}

	// Foutcontrole
	err = rows.Err()
	if err != nil {
		fmt.Println("Fout bij het verwerken van de resultaten:", err)
		return
	}

	currentTime := time.Now()
	hour := currentTime.Hour()

	groet(hour)
	for {
		fmt.Println("\n1. Kenteken registreren\n2. Toegang parkeerplaats\n3. Afsluiten\nKeuze: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			err := addPlate(db)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Voer uw kenteken in:")
			var inputPlate string
			fmt.Scanln(&inputPlate)
			checkPlate(db, inputPlate)
		case 3:
			shutdown()
			return // Toegevoegd om het programma te beëindigen na het afdrukken van het afscheid.
		default:
			fmt.Println("Ongeldige keuze. Probeer opnieuw.")
		}
	}
}

func readPasswordFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}

// Functie om een kenteken toe te voegen aan de database
func addPlate(db *sql.DB) error {
	fmt.Println("Voer het kenteken in om te registreren:")
	var plate string
	fmt.Scanln(&plate)

	// Voorbereid SQL-statement
	stmt, err := db.Prepare("INSERT INTO plates (plate, active) VALUES (?, true)")
	if err != nil {
		return fmt.Errorf("fout bij het voorbereiden van het SQL-statement: %v", err.Error())
	}
	defer stmt.Close()

	// Voer het SQL-statement uit
	_, err = stmt.Exec(plate)
	if err != nil {
		return fmt.Errorf("fout bij het uitvoeren van het SQL-statement: %v", err)
	}
	fmt.Println("Kenteken succesvol geregistreerd.")
	return nil
}

// Functie om te controleren of een kenteken in de lijst van toegestane kentekens staat en of het actief is
func checkPlate(db *sql.DB, inputPlate string) {
	// Query om te controleren of het kenteken actief is
	query := "SELECT COUNT(*) FROM plates WHERE plate = ? AND active = true"
	var count int
	err := db.QueryRow(query, inputPlate).Scan(&count)
	if err != nil {
		fmt.Println("Fout bij het uitvoeren van de query:", err)
		return
	}

	if count > 0 {
		fmt.Println("U heeft toegang tot de parkeerplaats.")
	} else {
		fmt.Println("U heeft geen toegang tot de parkeerplaats of het kenteken is niet actief. Excuses voor het ongemak.")
	}
}

// Functie om het programma af te sluiten
func shutdown() {
	fmt.Println("Tot ziens!")
}

// Functie om de begroeting af te handelen op basis van het tijdstip van de dag
func groet(hour int) {
	morgen := "Goedemorgen"
	middag := "Goedemiddag"
	avond := "Goedeavond"

	switch {
	case hour >= 7 && hour < 12:
		fmt.Println(morgen)
	case hour >= 12 && hour < 18:
		fmt.Println(middag)
	case hour >= 18 && hour < 24:
		fmt.Println(avond)
	default:
		fmt.Println("De parkeerplaats is 's nachts gesloten. Excuus voor het ongemak.")
	}
}
