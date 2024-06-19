<?php

namespace App\Http\Controllers\user;

use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Symfony\Component\HttpKernel\Exception\HttpException;

class ProfileController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/users/me/profile');

            if ($response->successful()) {
                return view('user.profile', [
                    'user' => $response['data']
                ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('user.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($response->status());
            }

            return abort(500);
        }
    }
}
