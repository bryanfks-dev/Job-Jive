<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class EmployeesController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response = \Http::withHeaders([
                'Authoration' => 'Bearer ' . session('token'),
                'Accept' => 'application/json',
            ])->get(BackendServer::url() . '/api/')
        }   
        catch(\Exception $e) {

        }
        
    }
}
