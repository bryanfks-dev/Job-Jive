<?php

namespace App\Http\Controllers\admin;

use Http;
use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;

class DepartmentsController extends Controller
{
    public function index()
    {
        return view("admin.departments");
    }

    public function create(Request $request)
    {
        $fields = \Validator::make($request->all(), [
            'department-name' => ['required', 'unique:departments,department_name']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors([
                'error' => 'The department name already exist'
            ]);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json'
                ])->post(BackendServer::url() . '/api/department/create', [
                    'department-name' => $request['department-name']
                ]);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.departments'));

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
                }

                abort($response['status']);
            }

            abort(400); // Bad request
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpExceptionInterface) {
                $e->getStatusCode();
            }
        }
    }
}
