<?php

namespace App\Http\Controllers\admin;

use Http;
use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Pagination\LengthAwarePaginator;

class DepartmentsController extends Controller
{
    private function paginate(array $items, int $perPage = 10, ?int $page = null, $options = []): LengthAwarePaginator
    {
        $page = $page ?: (LengthAwarePaginator::resolveCurrentPage() ?: 1);
        $items = collect($items);

        return new LengthAwarePaginator(
            $items->forPage($page, $perPage),
            $items->count(),
            $perPage,
            $page,
            $options
        );
    }

    public function index(Request $request)
    {
        try {
            $departmentsResponse =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'applications/json'
                ])->get(BackendServer::url() . '/api/departments');

            $userResponse =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'applications/json'
                ])->get(BackendServer::url() . '/api/users');

            if ($departmentsResponse->successful() && $userResponse->successful()) {
                if ($departmentsResponse['status'] == 200 && $userResponse['status'] == 200) {
                    if ($request->has('query')) {
                        $query = $request->get('query');

                        if (!empty($query)) {
                            $results = [];

                            foreach ($departmentsResponse['data'] as $department) {
                                if (in_array(strtolower($query), array_map('strtolower', $department))) {
                                    $results[] = $department;
                                }
                            }

                            $paginatedDepartments = $this->paginate($results ?? []);

                            return view("admin.departments", [
                                'departments' => $paginatedDepartments,
                                'users' => $userResponse['data'] ?? []
                            ]);
                        }
                    }

                    $paginatedDepartments = $this->paginate($departmentsResponse['data'] ?? []);

                    return view("admin.departments", [
                        'departments' => $paginatedDepartments,
                        'users' => $userResponse['data'] ?? []
                    ]);
                } else if ($departmentsResponse['status'] == 401 || $userResponse['status'] == 401) {
                    return redirect()->intended(route('admin.login'));
                }

                if ($departmentsResponse['status'] == 200) {
                    return abort($userResponse['status']);
                }

                return abort($departmentsResponse['status']);
            }

            return abort(400); // Bad request
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return abort(500);
        }
    }

    public function create(Request $request)
    {
        $fields = \Validator::make($request->all(), [
            'department_name' => ['required']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors([
                'create-error' => $fields->errors()->first()
            ])
                ->withInput([
                    'department_name' => $request['department_name']
                ]);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->post(BackendServer::url() . '/api/department/create', [
                            'department_name' => $request['department_name']
                        ]);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.departments'));

                    case 400: // Bad request
                        return redirect()->back()->withErrors([
                            'create-error' => $response['message']
                        ])
                            ->withInput([
                                'department_name' => $request['department_name']
                            ]);

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
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

    public function update(Request $request, int $id)
    {
        $id = intval($id);

        $fields = \Validator::make($request->all(), [
            'manager_id' => ['required']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors([
                'update-error-' . $id => $fields->errors()->first(),
            ])
                ->withInput([
                    'department_name' => $request['department_name'],
                    'manager_id' => $request['manager_id']
                ]);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->put(BackendServer::url() . '/api/department/update/' . $id, [
                            'manager_id' => intval($request['manager_id'])
                        ]);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.departments'));

                    case 400: // Bad request
                        return redirect()->intended(route('admin.departments'))->withErrors([
                            'update-error' . $id => $response['message']
                        ])
                            ->withInput([
                                'department_name' => $request['department_name'],
                                'manager_id' => $request['manager_id']
                            ]);

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
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

    public function delete(int $id)
    {
        $id = intval($id);

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->delete(BackendServer::url() . '/api/department/delete/' . $id);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.departments'));

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
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
