@echo off

set frontend_path=./laravel-app
set backend_path=./go-worker

:: Start backend server
start cmd /k "cd %backend_path% && go run ."
:: Start frontend server
start cmd /k "cd %frontend_path% && npm run dev"
start cmd /k "cd %frontend_path% && php artisan serve"
