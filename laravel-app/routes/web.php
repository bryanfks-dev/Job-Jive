<?php

use Illuminate\Support\Facades\Route;

use App\Http\Controllers\LoginController;

Route::get('/', function () {
    return view('dashboard');
})->name('dashboard');

Route::get('/login-user', [LoginController::class, 'viewUser']
)->name('login-user');

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

Route::get('/employees/view', function () {
    return view('view');
})->name('view');

Route::get('/sandbox', function () {
    return view('sandbox');
})->name('sandbox');

Route::get('/profile', function() {
    return view('profile');
})->name('profile');
