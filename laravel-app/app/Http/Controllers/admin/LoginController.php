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
        try {
            // Send request to be server
            $response =
            Http::withHeaders([
                'Authorization' => 'Bearer ' . session('token'),
                'Accept' => 'applications/json'
            ])->get(BackendServer::url() . '/auth/verify-token');

            if ($response->successful()) {
                if ($response['status'] == 200) { // Ok
                    return view('admin.login');
                } else if ($response['status'] == 403) { // Forbidden
                    if ($response['role'] == 'user') {
                        return redirect()->intended(route('user.dashboard'));
                    }

                    return redirect()->intended(route('admin.employees'));
                }

                return abort($response['status']);
            }

            return view('admin.login')->withErrors(['error' => 'Client error']);
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return view('admin.login')->withErrors(['error' => 'Server error']);
        }
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
                'Authorization' => 'Bearer ' . session('token'),
                'Content-type' => 'application/json',
                'Accept' => 'applications/json'
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

                return abort($response['status']);
            }

            return redirect()->back()->withErrors(['error' => 'Client error']);
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return redirect()->back()->withErrors(['error' => 'Server error']);
        }
    }
}
