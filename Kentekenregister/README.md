# Parkeerplaats Toegangscontrolesysteem

Dit programma biedt een eenvoudig toegangscontrolesysteem voor een parkeerplaats. Het maakt gebruik van een MySQL-database om kentekens op te slaan en te beheren.


## Installatie

1. Zorg ervoor dat Go is geïnstalleerd op uw systeem. Meer informatie over het installeren van Go is te vinden op [golang.org](https://golang.org/doc/install).

2. Clone de repository naar uw lokale machine:

   git clone https://github.com/stijnvandepol/programming.git

3. Navigeer naar de map van het project:

   cd programming/kentekenregister

4. Voer het programma uit:

   go run main.go


## Gebruik

Het programma biedt de volgende functionaliteit:

1. **Kenteken registreren**: Voer een nieuw kenteken in om het te registreren in het systeem.

2. **Toegang parkeerplaats**: Controleer of een kenteken toegang heeft tot de parkeerplaats.

3. **Afsluiten**: Sluit het programma af.

## Vereisten

- Go 1.16 of hoger
- MySQL-database

## Configuratie

Voordat u het programma uitvoert, zorg ervoor dat u een MySQL-database heeft geconfigureerd en de verbindingsspecificaties in het programma hebt bijgewerkt om te verwijzen naar uw database.
