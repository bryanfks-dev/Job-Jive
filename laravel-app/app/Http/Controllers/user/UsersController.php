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
                // Memastikan data user ada
                if (isset($responseUser['data']) && count($responseUser['data']) > 0) {
                    // Mengambil elemen pertama sebagai manager
                    $manager = $responseUser['data'][0];

                    // Mengambil elemen setelah elemen pertama untuk paginasi
                    $users = array_slice($responseUser['data'], 1);

                    // Melakukan paginasi pada array users
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
}
