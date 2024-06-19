<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Illuminate\Pagination\LengthAwarePaginator;
use Illuminate\Pagination\Paginator;
use Illuminate\Support\Facades\Http;
use Symfony\Component\HttpKernel\Exception\HttpException;

class UsersController extends Controller
{
    public function index(Request $request)
    {
        try {

            $responseUser = [];
            if ($request->has('query')) {
                $param = trim($request->get('query'));

                if (! empty($param)) {
                    $responseUser = Http::withHeaders([
                        'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                        'Accept' => 'application/json',
                    ])->get(BackendServer::url().'/api/user/search/'.$request->get('query'));
                }
            } else {
                $responseUser = Http::withHeaders([
                    'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                    'Accept' => 'application/json',
                ])->get(BackendServer::url().'/api/manager/users');
            }

            $responseDepartment = Http::withHeaders([
                'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                'Accept' => 'application/json',
            ])->get(BackendServer::url().'/api/manager/departments');

            if ($responseDepartment->serverError() || $responseUser->serverError()) {
                return abort(500);
            }

            if ($responseDepartment->successful() && $responseUser->successful()) {
                // Ensuring that there is at least one user
                if (isset($responseUser['data']) && count($responseUser['data']) > 0) {

                    // Taking the first element
                    $manager = $responseUser['data'][0];

                    // Taking all elements except the first
                    $users = array_slice($responseUser['data'], 1);

                    // Doing pagination
                    $perPage = 9;
                    $currentPage = Paginator::resolveCurrentPage('page');
                    $currentItems = array_slice($users, ($currentPage - 1) * $perPage, $perPage);
                    $paginatedUsers = new LengthAwarePaginator($currentItems, count($users), $perPage, $currentPage, [
                        'path' => Paginator::resolveCurrentPath(),
                        'pageName' => 'page',
                    ]);
                } else {
                    $manager = null;
                    $paginatedUsers = new LengthAwarePaginator([], 0, 9);
                }
                // dd($paginatedUsers);

                return view('user.employees', [
                    'manager' => $manager,
                    'users' => $paginatedUsers,
                    'departments' => $responseDepartment['data'] ?? [],
                ]);
            } elseif ($responseDepartment->unauthorized() || $responseUser->unauthorized()) {
                return redirect()->intended(route('user.login'));
            }

            return abort($responseUser->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($e->getStatusCode(), $e->getMessage());
            }

            return abort(500);
        }
    }

    public function update(Request $request, int $id)
    {
        $id = intval($id);

        try {
            // Convert salary inputs to float
            $initial_salary = floatval(str_replace('.', '', $request->input('initial_salary')));
            $current_salary = floatval(str_replace('.', '', $request->input('current_salary')));

            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->put(BackendServer::url().'/api/manager/update-user/'.$id, [
                    'initial_salary' => $initial_salary,
                    'current_salary' => $current_salary,
                ]);

            if ($response->successful()) {
                return redirect()->back()
                    ->with('update-success-'.$id, 'Salary updated successfully');
            } elseif ($response->badRequest()) {
                return redirect()->intended(route('user.employees'))->withErrors([
                    'update-error-'.$id => $response['error'],
                ])->withInput([
                    'initial_salary' => $request['initial_salary'],
                    'current_salary' => $request['current_salary'],
                ]);
            } elseif ($response->unauthorized()) {
                return redirect()->intended(route('user.login'));
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
