<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Symfony\Component\HttpKernel\Exception\HttpException;

class DashboardController extends Controller
{
    public function index(Request $request)
    {
        try {
            $token = $request->cookie('auth_token');

            $responseConfig =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $token,
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/configs');

            $reponseLatestCheck =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $token,
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/user/attendance/today/latest');

            if ($responseConfig->successful() && $reponseLatestCheck->successful()) {
                return view('user.dashboard', [
                    'configs' => $responseConfig['data'],
                ]);
            } else if ($responseConfig->unauthorized() || $reponseLatestCheck->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } else if ($responseConfig->serverError() || $reponseLatestCheck->serverError()) {
                return abort(500);
            }
        } catch (\Exception $e) {
            return abort(500);
        }
    }
}
