FROM php:8.2-fpm

# Install PHP extensions for Laravel
RUN docker-php-ext-install mysqli pdo pdo_mysql

# Install Redis extension
RUN pecl install redis && docker-php-ext-enable redis

# Set working directory
WORKDIR /var/www/html

# Copy composer (optional if you want to install packages inside container)
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

# Permissions
RUN chown -R www-data:www-data /var/www/html
