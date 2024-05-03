<?php

use App\Http\Controllers\LoginController;
use Illuminate\Support\Facades\Route;

Route::get('/', function() {
    return view('dashboard');
});

Route::get('/login', [LoginController::class, 'view']);

Route::get('/register', function() {
    return view('register');
});

Route::get('/profile', function() {
    return view('profile');
});
