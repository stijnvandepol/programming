#!/bin/bash
echo "Waiting for MySQL to start..."
while ! mysqladmin ping -h"db" --silent; do
    echo "MySQL is not yet available. Waiting..."
    sleep 1
done
echo "MySQL is now available. Executing SQL query..."

mysql -h"db" -u"user" -p"Rookworst31!" -e "CREATE TABLE IF NOT EXISTS chatbase (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
"
echo "SQL query executed successfully."
