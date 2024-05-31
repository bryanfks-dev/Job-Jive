<?php

namespace App\Http\Controllers\user;

use Http;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;

class ProfileController extends Controller
{
    public function index()
    {
        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token')
                ])->get(BackendServer::url() . '/api/user/profile');

            dd($response['message']);

            if ($response->successful() && isset($response)) {
                switch ($response['status']) {
                    case 401: // Unauthorized
                        return redirect()->intended(route('user-login'));

                    case 500: // Internal server error
                        return abort(500);

                    case 200:
                        return view('user.profile', []);
                }
            }

            return abort(404);
        } catch (\Exception $e) {
            return abort(500);
        }
    }
}
