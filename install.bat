@echo off

cd ./laravel-app

:: Frontend
start composer install && copy .env.example .env && php artisan key:generate && php artisan migrate && npm install && echo Frontend Installation Done..

:: Backend
cd ../server
start copy .env.example .env && echo Backend Installation Done..