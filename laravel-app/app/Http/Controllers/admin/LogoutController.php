<?php

namespace App\Http\Controllers\admin;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class LogoutController extends Controller {
    public function logout(Request $request) {
        // Forget cookie
        if (cookie()->hasQueued('auth_token')) {
            cookie()->queue(
                cookie()->forget('auth_token')
            );
        }

        return redirect()->intended(route('admin.login'));
    }
}
