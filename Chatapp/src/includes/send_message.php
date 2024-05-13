<?php
include 'dbconnection.php';
// Maak verbinding met de database
$mysqli = new mysqli("$hostname", "$username", "$password", "$database");

// Controleer op fouten bij het maken van verbinding
if ($mysqli->connect_errno) {
    die("Failed to connect to MySQL: " . $mysqli->connect_error);
}

// Ontvang het bericht en de gebruikersnaam van de frontend
$data = json_decode(file_get_contents("php://input"));

// Controleer of het bericht en de gebruikersnaam zijn ontvangen
if (!empty($data->message) && !empty($data->username)) {
    // Voeg het bericht en de gebruikersnaam toe aan de database
    $message = $mysqli->real_escape_string($data->message);
    $username = $mysqli->real_escape_string($data->username);
    $mysqli->query("INSERT INTO chatbase (username, message) VALUES ('$username', '$message')");

    // Geef een succesbericht terug
    echo json_encode(array("message" => "Message sent successfully."));
} else {
    // Geef een foutbericht terug als er geen bericht of gebruikersnaam is ontvangen
    echo json_encode(array("error" => "No message or username received."));
}
?>
