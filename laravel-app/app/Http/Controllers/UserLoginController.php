<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class UserLoginController extends Controller
{
    public function index()
    {
        return view('login-user');
    }

    public function login(Request $request)
    {
        $credentials = $request->validate([
            'email' => ['required', 'email'],
            'password' => ['required']
        ]);

        // Send request to be server
        $response =
            Http::post('http://127.0.0.1:5000/auth/user/login', $credentials);

        switch ($response['status']) {
            case 401:
                return redirect()->back();

            case 200:
                return redirect()->intended(route('user.dashboard'));
        }
    }
}
