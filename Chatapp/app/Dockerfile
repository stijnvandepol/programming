# Gebruik een officiële PHP + Apache-image als basis
FROM php:7.4-apache

# Installeer mysqli-extensie
RUN docker-php-ext-install mysqli

# Kopieer de PHP-bestanden naar de juiste locatie in de container
COPY src/ /var/www/html/

# Expose de poort waarop de webserver draait
EXPOSE 80

# Start de Apache-webserver wanneer de container wordt gestart
CMD ["apache2-foreground"]