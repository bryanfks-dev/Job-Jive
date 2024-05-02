@echo off

set frontend_path=./laravel-app
set backend_path=./server

set frontend_cmd=php artisan serve
set backend_cmd=go run .

:: Start backend server
start cmd /k "cd %backend_path% && %backend_cmd%"
:: Start frontend server
start cmd /k "cd %frontend_path% && %frontend_cmd% --host=0.0.0.0"