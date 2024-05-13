# Kopieer het init-db.sh script naar de container
COPY init-db.sh /usr/local/bin/

# Maak het script uitvoerbaar
RUN chmod +x /usr/local/bin/init-db.sh

# Voer het script uit bij het starten van de container
ENTRYPOINT ["/usr/local/bin/init-db.sh"]
