<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class LogoutController extends Controller {
    public function logout(Request $request) {
        if (\Auth::check()) {
            \Auth::logout();

            $request->session()->flush();
        }

        return redirect()->intended(route('user.login'));
    }
}
