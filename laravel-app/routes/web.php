<?php

use Illuminate\Support\Facades\Route;

use App\Http\Controllers\UserLoginController;
use App\Http\Controllers\AdminLoginController;

// User routes
Route::group(['prefix'=> '/'], function () {
    // Login route
    Route::group(['prefix'=> '/login'], function () {
        Route::get('/', [UserLoginController::class, 'index'])
            ->name('user.login');

        Route::post('/', [UserLoginController::class, 'login']);
    });

    // Dashboard route
    Route::get('/', function () {
        return view('dashboard');
    })->name('user.dashboard');
});

// Admin routes
Route::group(['prefix'=> '/admin'], function () {
    // Login route
    Route::group(['prefix'=> '/login'], function () {
        Route::get('/', [AdminLoginController::class, 'index'])
            ->name('admin.login');
    });
});

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
