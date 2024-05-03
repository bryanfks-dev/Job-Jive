<?php

use Illuminate\Support\Facades\Route;

Route::get('/', function () {
    return view('dashboard');
})->name('dashboard');

Route::get('/login-user', function () {
    return view('login-user');
})->name('login-user');

Route::get('/login-admin', function () {
    return view('login-admin');
})->name('login-admin');

Route::get('/register', function () {
    return view('register');
})->name('register');

Route::get('/attendance', function () {
    return view('attendance');
})->name('attendance');

Route::get('/employees', function () {
    return view('employees');
})->name('employees');

Route::get('/sandbox', function () {
    return view('sandbox');
})->name('sandbox');
