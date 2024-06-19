
:: Frontend
start cd laravel-app && composer install && copy .env.example .env 
    && php artisan key:generate && php artisan migrate:fresh && npm install 
    && echo Frontend Installation Done..

:: Backend
start cd go-worker && copy .env.example .env && go build -o addAdmin.exe utils/addAdmin/main.go 
    && echo Backend Installation Done..
