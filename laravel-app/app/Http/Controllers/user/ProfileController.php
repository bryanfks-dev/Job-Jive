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
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/user/profile');

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return view('user.profile', $response['data']);

                    case 401: // Unauthorized
                        return redirect()->intended(route('user.login'));
                }

                return abort($response['status']);
            }

            return abort(400); // Bad request
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return abort(500);
        }
    }
}
