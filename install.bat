@echo off

cd ./laravel-app

:: Install Composer
start composer install
:: Create .env file
start copy .env.example .env
:: Generate laravel app key
start php artisan key:generate
:: Migrate laravel database
start php artisan migrate
:: Install node modules
start npm install

echo Installation Done..