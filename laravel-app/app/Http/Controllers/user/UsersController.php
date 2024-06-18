<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
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
                $paginatedUsers = $this->paginate($responseUser['data'] ?? []);

                return view('user.employees', [
                    'users' => $paginatedUsers,
                    'departments' => $responseDepartment['data'] ?? [],
                ]);
            } elseif ($responseDepartment->unauthorized() || $responseUser->unauthorized()) {
                return redirect()->intended(route('user.login'));
            }

            return abort($responseUser->status());
        } catch (\Exception $e) {
            // dd($e);

            if ($e instanceof HttpException) {
                throw new HttpException($e->getStatusCode(), $e->getMessage());
            }

            return abort(500);
        }
    }
}
