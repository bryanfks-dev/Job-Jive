<?php

namespace App\Http\Controllers\admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class LogoutController extends Controller {
    public function logout(Request $request) {
        $request->session()->invalidate();

        return redirect()->intended(route('admin.login'));
    }
}
