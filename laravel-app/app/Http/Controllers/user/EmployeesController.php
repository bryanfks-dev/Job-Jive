<?php

namespace App\Http\Controllers\user;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;

class EmployeesController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response = \Http::withHeaders([
                'Authoration' => 'Bearer ' . session('token'),
                'Accept' => 'application/json',
            ])->get(BackendServer::url() . '/api/');
        } catch(\Exception $e) {

        }
    }
}
