<?php

use App\Http\Controllers\LoginController;
use Illuminate\Support\Facades\Route;

Route::get('/', function() {
    return view('dashboard');
});

Route::get('/login-user', [LoginController::class, 'viewUser']);

Route::get('/login-admin', function() {
    return view('login-admin');
});

Route::get('/register', function() {
    return view('register');
});

Route::get('/attendance', function() {
    return view('attendance');
});

Route::get('/employees', function() {
    return view('employees');
});

Route::get('/profile', function() {
    return view('profile');
});

