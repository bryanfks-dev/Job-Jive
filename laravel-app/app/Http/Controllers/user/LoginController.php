<?php

namespace App\Http\Controllers\user;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;
use Symfony\Component\HttpKernel\Exception\HttpException;

class LoginController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                ])->get(BackendServer::url() . '/auth/verify-token');

            if ($response->successful()) {
                return view('user.login');
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }

    public function login(Request $request)
    {
        // Validate form
        $validator = \Validator::make($request->all(), [
            'email' => ['required', 'email'],
            'password' => ['required']
        ]);

        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors(['error' => $validator->errors()->first()]);
        }

        $rememberMe = $request['remember'] === 'on';

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->post(BackendServer::url() . '/auth/user/login', [
                            'email' => $request['email'],
                            'password' => $request['password'],
                            'remember' => $rememberMe
                        ]);

            if ($response->successful()) {
                // Create cookie
                $expireTime = 60 * 24 * 6; // 6 days expire time

                // Create cookie
                if ($rememberMe) {
                    $expireTime = 60 * 24 * 30; // 30 days expire time
                }

                return redirect()->intended(route('user.dashboard'))
                    ->cookie(
                        cookie(
                            'auth_token',
                            $response['token'],
                            $expireTime,
                            secure: true
                        )
                    );
            } else if ($response->serverError()) {
                return abort($response->status());
            }

            return redirect()->back()
                ->withErrors(['error' => $response['error']]);
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }
}
