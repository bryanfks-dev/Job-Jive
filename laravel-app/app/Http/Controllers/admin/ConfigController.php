<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;

class ConfigController extends Controller
{
    public function index()
    {
        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/configs');

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return view('admin.config', $response['data']);

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
                }
                abort($response['status']);
            }

            abort(400); // Bad request
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpExceptionInterface) {
                return abort($e->getStatusCode());
            }

            return abort(500);
        }
    }

    public function save(Request $request)
    {
        $fields = \Validator::make(request()->all(), [
            'check_in_time' => ['required'],
            'check_out_time' => ['required'],
            'absence_quota' => ['required'],
            'daily_work_hours' => ['required'],
            'weekly_work_hours' => ['required']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors(['error' => 'Do not leave input fields empty']);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json'
                ])->put(BackendServer::url() . '/api/configs/save', [
                    'check_in_time' => $request['check_in_time'],
                    'check_out_time' => $request['check_out_time'],
                    'absence_quota' => $request['absence_quota'],
                    'daily_work_hours' => $request['daily_work_hours'],
                    'weekly_work_hours' => $request['weekly_work_hours']
                ]);

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        return redirect()->intended(route('admin.configs'));

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
                }

                return abort($response['status']);
            }

            return abort(400); // Bad request
        } catch (\Exception $e) {
            return abort(500);
        }
    }
}
