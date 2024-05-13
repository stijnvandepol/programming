<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Chat</title>
    <link rel="stylesheet" href="css/styles.css"> <!-- Koppel het CSS-bestand -->
</head>
<body>
    <div class="container">
        <h2>Chat</h2>
        <div id="chat"></div>
        <div class="input-area">
            <input type="text" id="message" class="input-field" placeholder="Type your message here" onkeydown="if(event.keyCode===13) sendMessage()">
            <button onclick="sendMessage()" class="send-button">Send</button>
        </div>
    </div>

    <script>
        // Functie om de gebruikersnaam op te vragen en op te slaan in de sessie
        function askUsername() {
            // Vraag de gebruiker om een gebruikersnaam
            let username = prompt("Please enter your username:");

            // Als de gebruiker een gebruikersnaam heeft ingevoerd, sla deze dan op in de sessie
            if (username) {
                sessionStorage.setItem('username', username);
            } else {
                // Als de gebruiker geen gebruikersnaam heeft ingevoerd, blijf vragen tot er een is ingevoerd
                askUsername();
            }
        }

        // Functie om de gebruikersnaam uit de sessie op te halen
        function getUsername() {
            let username = sessionStorage.getItem('username');
            if (!username) {
                // Als er geen gebruikersnaam in de sessie is opgeslagen, vraag de gebruiker om een te invoeren
                askUsername();
                // Haal de gebruikersnaam opnieuw op
                username = sessionStorage.getItem('username');
            }
            return username;
        }

        // Call getMessages immediately when the page loads
        getMessages();

        // Then call getMessages every 5 seconds
        setInterval(getMessages, 5000);

        function getMessages() {
            fetch('includes/get_messages.php')
            .then(response => response.json())
            .then(messages => {
                const chatDiv = document.getElementById('chat');
                chatDiv.innerHTML = '';
                messages.reverse();
                messages.forEach(message => {
                    chatDiv.innerHTML += `<div class="chat-message"><strong>${message.username}:</strong> ${message.message}</div>`;
                });
                chatDiv.scrollTop = chatDiv.scrollHeight;
            });
        }

        function sendMessage() {
            const messageInput = document.getElementById('message');
            const message = messageInput.value;
            messageInput.value = '';

            const username = getUsername(); // Haal de gebruikersnaam op uit de sessie

            // Controleer of het bericht niet leeg is voordat het wordt verzonden
            if (message.trim() !== '') {
                // Verzend het bericht naar de server
                fetch('includes/send_message.php', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ message, username }), // Stuur de gebruikersnaam mee met het bericht
                })
                .then(response => {
                    if (response.ok) {
                        // Geef een bericht weer in de console als het bericht succesvol is verzonden
                        console.log('Message sent successfully');
                        // Haal de berichten opnieuw op om de nieuw toegevoegde bericht weer te geven
                        getMessages();
                    } else {
                        throw new Error('Failed to send message');
                    }
                })
                .catch(error => {
                    console.error('Error:', error);
                });
            }
        }

        getMessages();

        setInterval(getMessages, 2000);
    </script>
</body>
</html>
