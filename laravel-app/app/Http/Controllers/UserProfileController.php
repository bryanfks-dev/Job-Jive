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
                'Authorization' => 'Bearer ' . 'askjdoasd'
            ])->get(BackendServer::url() . '/api/user/profile');

        if ($response->successful()) {

            return view("profile");
        }

        return "owkow";
    }
}
