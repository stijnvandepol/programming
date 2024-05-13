<?php
include 'dbconnection.php';
// Maak verbinding met de database
$mysqli = new mysqli("$hostname", "$username", "$password", "$database");

// Controleer op fouten bij het maken van verbinding
if ($mysqli->connect_errno) {
    die("Failed to connect to MySQL: " . $mysqli->connect_error);
}

// Query om de laatste 10 berichten op te halen
$result = $mysqli->query("SELECT * FROM chatbase ORDER BY timestamp DESC LIMIT 100");

// Array om de berichten op te slaan
$messages = array();
while ($row = $result->fetch_assoc()) {
    $messages[] = $row;
}

// Stuur de berichten terug als JSON
echo json_encode($messages);
?>
