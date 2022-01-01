FROM composer as composer

FROM php:7.4

COPY --from=composer /usr/bin/composer /usr/bin/composer

COPY . /var/www/app

WORKDIR /var/www/app

RUN apt-get update && \
    apt-get install zip -y && \
    apt-get install vim -y && \
    apt-get install -y libicu-dev && \
    docker-php-ext-configure intl && \
    docker-php-ext-install intl
