FROM php:8.0-apache

RUN chmod 777 -R /var/www \
    && chown -R www-data:www-data /var/www \
    && usermod -u 1000 www-data \
    && chsh -s /bin/bash www-data

RUN ln -sf /dev/null /var/log/apache2/access.log \
    && ln -sf /dev/null /var/log/apache2/error.log

COPY 000-default.conf /etc/sites-available

COPY index.php /var/www/html

RUN a2ensite 000-default

CMD ["apache2-foreground"]
