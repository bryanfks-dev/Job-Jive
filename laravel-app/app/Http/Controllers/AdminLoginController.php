<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class AdminLoginController extends Controller
{
    public function index()
    {
        return view('login-admin');
    }

    public function login(Request $request)
    {
        $credentials = $request->validate([
            'username' => ['required'],
            'password' => ['required']
        ]);

        // Send request to be server
        $response =
            Http::post('http://127.0.0.1:5000/auth/admin/login', $credentials);

        switch ($response['status']) {
            case 401:
                return redirect()->back();

            case 200:
                return redirect()->intended(route('admin.dashboard'));
        }
    }
}
