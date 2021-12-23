FROM php:8-cli

COPY . /var/www/app

WORKDIR /var/www/app

RUN apt-get update && \
    apt-get install zip -y

RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');" && \
    mv composer.phar /usr/local/bin/composer

ENTRYPOINT ["php", "artisan", "serve"]

CMD ["--host=0.0.0.0"]