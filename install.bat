@echo off

:: Frontend
start cd laravel-app && composer install && copy .env.example .env 
    && php artisan key:generate && php artisan migrate && npm install 
    && echo Frontend Installation Done..

:: Backend
start cd server && copy .env.example .env && cd utils && go build ./addAdmin.go 
    && echo Backend Installation Done..
