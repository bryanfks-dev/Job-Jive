<?php

namespace App\Http\Controllers;

use Http;
use App\Models\BackendServer;

class UserProfileController extends Controller
{
    public function index()
    {
        $response =
            Http::withHeaders([
                'Authorization' => 'Bearer ' . ''
            ])->get(BackendServer::url() . '/api/user/profile');

        if ($response->successful()) {
            switch ($response['status']) {
                case 401: // Unauthorized
                    return redirect()->intended(route('user-login'));

                case 500: // Internal server error
                    abort(500);
                    return;

                case 200:
                    return view('profile', []);
            }
        }

        return abort(500);
    }
}
