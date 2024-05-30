<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;

class LoginController extends Controller
{
    public function index()
    {
        return view('admin.login');
    }

    public function login(Request $request)
    {
        $credentials = $request->validate([
            'username' => ['required'],
            'password' => ['required']
        ]);

        try {
            // Send request to be server
            $response =
            Http::withHeaders([
                'Content-type' => 'application/json'
            ])->post(BackendServer::url() . '/auth/admin/login', $credentials);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 401:
                        return redirect()->back();

                    case 500:
                        return redirect()->back();

                    case 200:
                        // Create session token
                        session(['token' => $response['token']]);

                        return redirect()->intended(route('admin.dashboard'));
                }
            }

            return redirect()->back();
        } catch (\Exception $e) {
            return redirect()->back();
        }
    }
}
