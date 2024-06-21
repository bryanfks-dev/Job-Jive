<?php

namespace App\Http\Controllers\admin;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Symfony\Component\HttpKernel\Exception\HttpException;

class DepartmentsController extends Controller
{
    public function index(Request $request)
    {
        try {
            $responseDepartment = null;

            $param = trim($request->get('query', ''), ' ');

            $httpHeaders = [
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            if (!empty($param)) {
                $responseDepartment =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/departments/search/' . $param);
            } else {
                $responseDepartment =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/departments');
            }

            $responseDepartmentsUser =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/departments/users');

            if ($responseDepartment->successful() && $responseDepartmentsUser->successful()) {
                $paginatedDepartments =
                    $this->paginate($responseDepartment['data'] ?? []);

                return view('admin.departments', [
                    'departments' => $paginatedDepartments,
                    'users' => $responseDepartmentsUser['data'],
                ]);
            } else if ($responseDepartment->unauthorized() || $responseDepartmentsUser->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            } else if ($responseDepartment->serverError() || $responseDepartmentsUser->serverError()) {
                return abort(500);
            }

            return abort($responseDepartment->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($responseDepartment->status());
            }

            return abort(500);
        }
    }

    public function create(Request $request)
    {
        $validator =
            \Validator::make($request->all(), [
                'department_name' => ['required'],
            ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'create-error' => $validator->errors()->first(),
            ])
                ->withInput([
                    'department_name' => $request['department_name'],
                ]);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->post(BackendServer::url() . '/api/departments/create', [
                            'department_name' => $request['department_name'],
                        ]);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->badRequest()) {
                return redirect()->back()->withErrors([
                    'create-error' => $response['error'],
                ])
                    ->withInput([
                        'department_name' => $request['department_name'],
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($response->status());
            }

            return abort(500);
        }
    }

    public function update(Request $request, int $id)
    {
        if ($id <= 0) {
            return abort(404);
        }

        $validator =
            \Validator::make($request->all(), [
                'manager_id' => ['required'],
            ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'update-error-' . $id => $validator->errors()->first(),
            ])
                ->withInput([
                    'manager_id' => $request['manager_id'],
                ]);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->put(BackendServer::url() . '/api/departments/update/' . $id, [
                            'manager_id' => intval($request['manager_id']),
                        ]);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->badRequest()) {
                return redirect()->intended(route('admin.departments'))->withErrors([
                    'update-error' . $id => $response['message'],
                ])
                    ->withInput([
                        'department_name' => $request['department_name'],
                        'manager_id' => $request['manager_id'],
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($response->status());
            }

            return abort(500);
        }
    }

    public function delete(Request $request, int $id)
    {
        if ($id <= 0) {
            return abort(404);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->delete(BackendServer::url() . '/api/departments/delete/' . $id);

            if ($response->successful()) {
                return redirect()->intended(route('admin.departments'));
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
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
