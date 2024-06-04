<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;
use Illuminate\Pagination\LengthAwarePaginator;

class UsersController extends Controller
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
            $userResponse =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'applications/json'
                ])->get(BackendServer::url() . '/api/users');

            $departmentResponse =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'applications/json'
                ])->get(BackendServer::url() . '/api/departments');

            if ($userResponse->successful() && $departmentResponse) {
                if ($userResponse['status'] == 200 && $departmentResponse['status'] == 200) {
                    if ($request->has('query')) {
                        $query = $request->get('query');

                        $results = [];

                        foreach ($userResponse['data'] as $user) {
                            if (in_array(strtolower($query), array_map('strtolower', $user))) {
                                $results[] = $user;
                            }
                        }

                        $paginatedUsers = $this->paginate($results ?? []);

                        return view("admin.users", [
                            'users' => $paginatedUsers,
                            'departments' => $departmentResponse['data'] ?? []
                        ]);
                    }

                    $paginatedUsers = $this->paginate($userResponse['data'] ?? []);

                    return view("admin.users", [
                        'users' => $paginatedUsers,
                        'departments' => $departmentResponse['data'] ?? []
                    ]);
                } else if ($userResponse['status'] == 401 || $departmentResponse['status'] == 401) {
                    return redirect()->intended(route('admin.login'));
                }

                return abort($userResponse['status']);
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
            'full_name' => ['required'],
            'email' => ['required', 'email'],
            'date_of_birth' => ['required', 'date_format:m-d-Y'],
            'address' => ['required'],
            'nik' => ['required'],
            'gender' => ['required', 'in:male,female'],
            'department_id' => ['required', 'integer']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors([
                'error' => 'Invalid field value'
            ]);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->post(BackendServer::url() . '/api/users/create', [
                            'full_name' => $request['full_name'],
                            'email' => $request['email'],
                            'date_of_birth' => $request['date_of_birth'],
                            'address' => $request['address'],
                            'nik' => $request['nik'],
                            'gender' => $request['gender'],
                            'department_id' => intval($request['department_id'])
                        ]);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.users'));

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
