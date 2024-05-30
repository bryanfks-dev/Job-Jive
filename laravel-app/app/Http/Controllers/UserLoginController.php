<?php

namespace App\Http\Controllers;

use App\Models\BackendServer;
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
            'password' => ['required'],
        ]);

        $remember = $request->input('remember') === "on";

        $credentials['remember'] = $remember;

        try {
            // Send request to be server
            $response =
                Http::withHeaders([
                    'Content-type' => 'application/json'
                ])->post(BackendServer::url() . '/auth/user/login', $credentials);

            if ($response->successful()) {
                if ($response['status'] === 200) {
                    // Create session token
                    session(['token' => $response['token']]);

                    // Create cookie
                    if ($remember) {
                        cookie('token', $response['token'], 60 * 24 * 7); // 7 days expire time
                    }

                    return redirect()->intended(route('user.dashboard'));
                }

                return redirect()->intended(route('user.login'))
                    ->withErrors(['error' => $response['message']]);
            }

            return redirect()->back()->withErrors(['error' => 'Server error']);
        } catch (\Exception $e) {
            return redirect()->back()->withErrors(['error' => 'Cannot resolve host']);
        }
    }
}
