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
        $credentials = \Validator::make($request->all(), [
            'username' => ['required'],
            'password' => ['required']
        ]);

        if ($credentials->fails()) {
            return redirect()->back()->withErrors($credentials->errors());
        }

        try {
            // Send request to be server
            $response =
            Http::withHeaders([
                'Content-type' => 'application/json'
            ])->post(BackendServer::url() . '/auth/admin/login', [
                'username' => $request['username'],
                'password' => $request['password'],
            ]);

            if ($response->successful()) {
                if ($response['status'] == 200) { // Ok
                    // Create session token
                    session(['token' => $response['token']]);

                    return redirect()->intended(route('admin.employees'));
                }

                return redirect()->back()
                    ->withErrors(['error' => $response['message']]);
            }

            return redirect()->back()->withErrors(['error' => 'Client error']);
        } catch (\Exception $e) {
            return redirect()->back()->withErrors(['error' => 'Cannot establish connection with backend server']);
        }
    }
}
