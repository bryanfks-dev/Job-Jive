<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Symfony\Component\HttpKernel\Exception\HttpException;

class DepartmentsController extends Controller
{
    public function index(Request $request)
    {
        try {
            $responseDepartment = [];

            if ($request->has('query')) {
                $param = trim($request->get('query'), ' ');

                if (!empty($param)) {
                    $responseDepartment =
                        \Http::withHeaders([
                            'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                            'Accept' => 'application/json'
                        ])->get(BackendServer::url() . '/api/department/search/' . $request->get('query'));
                }
            } else {
                $responseDepartment =
                    \Http::withHeaders([
                        'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                        'Accept' => 'applications/json'
                    ])->get(BackendServer::url() . '/api/departments');
            }

            $responseUser =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/users');

            if ($responseDepartment->serverError() || $responseUser->serverError()) {
                return abort(500);
            }

            if ($responseDepartment->successful() && $responseUser->successful()) {
                $paginatedDepartments =
                    $this->paginate($responseDepartment['data'] ?? []);

                $paginatedUsers =
                    $this->paginate($responseUser['data'] ?? []);

                return view("admin.departments", [
                    'departments' => $paginatedDepartments,
                    'users' => $paginatedUsers
                ]);
            } else if ($responseDepartment->unauthorized() || $responseUser->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($responseDepartment->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($responseDepartment->status());
            }

            return abort(500);
        }
    }

    public function create(Request $request)
    {
        $validator = \Validator::make($request->all(), [
            'department_name' => ['required']
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'create-error' => $validator->errors()->first()
            ])
                ->withInput([
                    'department_name' => $request['department_name']
                ]);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->post(BackendServer::url() . '/api/department/create', [
                            'department_name' => $request['department_name']
                        ]);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->badRequest()) {
                return redirect()->back()->withErrors([
                    'create-error' => $response['error']
                ])
                    ->withInput([
                        'department_name' => $request['department_name']
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }

    public function update(Request $request, int $id)
    {
        $id = intval($id);

        $validator = \Validator::make($request->all(), [
            'manager_id' => ['required']
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'update-error-' . $id => $validator->errors()->first(),
            ])
                ->withInput([
                    'manager_id' => $request['manager_id']
                ]);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->put(BackendServer::url() . '/api/department/update/' . $id, [
                            'manager_id' => intval($request['manager_id'])
                        ]);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->badRequest()) {
                return redirect()->intended(route('admin.departments'))->withErrors([
                    'update-error' . $id => $response['message']
                ])
                    ->withInput([
                        'department_name' => $request['department_name'],
                        'manager_id' => $request['manager_id']
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }

    public function delete(int $id)
    {
        $id = intval($id);

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->delete(BackendServer::url() . '/api/department/delete/' . $id);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }
}
