<?php

namespace App\Http\Controllers\user;

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
                'Accept' => 'applications/json',
            ])->get(BackendServer::url() . '/auth/verify-token');

            if ($response->successful()) {
                return view('user.login');
            }
            else if ($response->forbidden()) {
                if ($response['role'] == 'user') {
                    return redirect()->intended(route('admin.users'));
                }

                return redirect()->intended(route('admin.users'));
            }
            else if ($response->serverError()) {
                return abort($response->status());
            }

            return view('user.login')->withErrors(['error' => 'Client error']);
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return view('user.login')->withErrors(['error' => 'Server error']);
        }
    }

    public function login(Request $request)
    {
        $credentials = \Validator::make($request->all(), [
            'email' => ['required', 'email'],
            'password' => ['required']
        ]);

        if ($credentials->fails()) {
            return redirect()->back()->withErrors($credentials->errors());
        }

        $remember = $request['remember'] === 'on';

        try {
            // Send request to be server
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'applications/json'
                ])->post(BackendServer::url() . '/auth/user/login', [
                    'email' => $request['email'],
                    'password' => $request['password'],
                    'remember' => $remember
                ]);

            if ($response->successful()) {
                if ($response['status'] === 200) {
                    // Create session token
                    session(['token' => $response['token']]);

                    // Create cookie
                    if ($remember) {
                        cookie('token', $response['token'], 60 * 24 * 7); // 7 days expire time
                    }

                    return redirect()->intended(route('user.dashboard'));
                } else if ($response['status'] == 401) {
                    return redirect()->back()->withErrors([
                        'error' => $response['message']
                    ]);
                }

                return redirect()->back()->withErrors([
                    'error' => $response['message']
                ]);
            }

            return redirect()->back()->withErrors(['error' => 'Client Error']);
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return redirect()->back()->withErrors(['error' => 'Server Error']);
        }
    }
}
