<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class LogoutController extends Controller
{
    public function logout(Request $request)
    {
        // Forget cookie
        if ($request->cookie('auth_token') !== null) {
            cookie()->queue(
                cookie()->forget('auth_token')
            );
        }

        return redirect()->intended(route('user.login'));
    }
}
