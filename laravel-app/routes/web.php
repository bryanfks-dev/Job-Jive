<?php

use App\Http\Controllers\UserProfileController;
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

    // Profile route
    Route::get('/profile', [UserProfileController::class, 'index'])
        ->name('user.profile');
});

// Admin routes
Route::group(['prefix'=> '/admin'], function () {
    // Login route
    Route::group(['prefix'=> '/login'], function () {
        Route::get('/', [AdminLoginController::class, 'index'])
            ->name('admin.login');
        Route::post('/', [AdminLoginController::class, 'login']);
    });

    // After login
    Route::get('/departements', function () {
        return view('admin.department');
    })->name('admin.department');
    Route::get('/employees', function () {
        return view('admin.employee');
    })->name('admin.employee');
    Route::get('/config', function () {
        return view('admin.config');
    })->name('admin.config');
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
})->name('user.profile');


