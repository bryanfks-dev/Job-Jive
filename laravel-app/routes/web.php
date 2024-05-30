<?php

use App\Http\Controllers\admin;
use App\Http\Controllers\user;
use Illuminate\Support\Facades\Route;

// User routes
Route::group(['prefix'=> '/'], function () {
    // Login route
    Route::group(['prefix'=> '/login'], function () {
        Route::get('/', [user\LoginController::class, 'index'])
            ->name('user.login');

        Route::post('/', [user\LoginController::class, 'login']);
    });

    // Dashboard route
    Route::get('/', function () {
        return view('user.dashboard');
    })->name('user.dashboard');

    // Profile route
    Route::get('/profile', [user\ProfileController::class, 'index'])
        ->name('user.profile');
});

// Admin routes
Route::group(['prefix'=> '/admin'], function () {
    // Login route
    Route::group(['prefix'=> '/login'], function () {
        Route::get('/', [admin\LoginController::class, 'index'])
            ->name('admin.login');

        Route::post('/', [admin\LoginController::class, 'login']);
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
    return view('user.attendance');
})->name('attendance');

Route::get('/employees', function () {
    return view('user.employees');
})->name('employees');

Route::get('/employees/view', function () {
    return view('user.view');
})->name('view');

Route::get('/sandbox', function () {
    return view('user.sandbox');
})->name('sandbox');
