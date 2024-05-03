<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function() {
    return view('dashboard');
});

Route::get('/login-user', function() {
    return view('login-user');
});

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

