<?php

use App\Http\Controllers\user;
use App\Http\Controllers\admin;
use Illuminate\Support\Facades\Route;

// User routes
Route::group(['prefix' => '/'], function () {
    // Login route
    Route::group(['prefix' => '/login'], function () {
        Route::get('/', [user\LoginController::class, 'index'])
            ->name('user.login');

        Route::post('/', [user\LoginController::class, 'login']);
    });

    Route::post('/logout', [user\LogoutController::class, 'logout'])
        ->name('user.logout');

    // Dashboard / default route
    Route::get('/', [user\DashboardController::class, 'index'])
        ->name('user.dashboard');

    Route::post('/attend', [user\AttendanceController::class, 'attend'])
        ->name('user.attend');

    // Profile route
    Route::get('/profile', [user\ProfileController::class, 'index'])
        ->name('user.profile');

    // Attendances route
    Route::get('/attendance', [user\AttendanceController::class, 'index'])
        ->name('user.attendance');

    Route::get('/employees', function () {
        return view('user.employees');
    })->name('employees');

});

// Admin routes
Route::group(['prefix' => '/admin'], function () {
    // Default route
    Route::get('/', function () {
        return redirect(route('admin.users'));
    });

    // Login route
    Route::group(['prefix' => '/login'], function () {
        Route::get('/', [admin\LoginController::class, 'index'])
            ->name('admin.login');

        Route::post('/', [admin\LoginController::class, 'login']);
    });

    Route::post('/logout', [admin\LogoutController::class, 'logout'])
        ->name('admin.logout');

    // Users route
    Route::group(['prefix' => '/users'], function () {
        Route::get('/', [admin\UsersController::class, 'index'])
            ->name('admin.users');

        Route::post('/', [admin\UsersController::class, 'create'])
            ->name('admin.users.create');

        Route::put('/update/{id}', [admin\UsersController::class, 'update'])
            ->where(['id' => '[1-9][0-9]*'])->name('admin.users.update');

        Route::delete('/delete/{id}', [admin\UsersController::class, 'delete'])
            ->where(['id' => '[1-9][0-9]*'])->name('admin.users.delete');
    });

    // Departments route
    Route::group(['prefix' => '/departments'], function () {
        Route::get('/', [admin\DepartmentsController::class, 'index'])
            ->name('admin.departments');

        Route::post('/create', [admin\DepartmentsController::class, 'create'])
            ->name('admin.departments.create');

        Route::put('/update/{id}', [admin\DepartmentsController::class, 'update'])
            ->where(['id' => '[1-9][0-9]*'])->name('admin.departments.update');

        Route::delete('/delete/{id}', [admin\DepartmentsController::class, 'delete'])
            ->where(['id' => '[1-9][0-9]*'])->name('admin.departments.delete');
    });

    // Config route
    Route::group(['prefix' => '/configs'], function () {
        Route::get('/', [admin\ConfigController::class, 'index'])
            ->name('admin.configs');

        Route::put('/', [admin\ConfigController::class, 'save'])
            ->name('admin.configs.save');
    });
});

Route::get('/employees/view', function () {
    return view('user.view');
})->name('view');

Route::get('/sandbox', function () {
    return view('user.sandbox');
})->name('sandbox');
