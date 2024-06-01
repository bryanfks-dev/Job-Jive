<?php

namespace App\Http\Controllers\admin;

use Carbon\Carbon;
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
            'check_in_time' => ['required', 'date_format:H:i', 'before:' . $request['check_out_time']],
            'check_out_time' => ['required', 'date_format:H:i', 'after:' . $request['check_in_time']],
            'absence_quota' => ['required', 'integer'],
            'daily_work_hours' => ['required', 'integer', 'lte:weekly_work_hours', 'min:1', 'max:24'],
            'weekly_work_hours' => ['required', 'integer', 'gte:daily_work_hours', 'min:1', 'max:168']
        ]);

        if ($fields->fails()) {
            return redirect()->back()->withErrors($fields->errors());
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
                    'absence_quota' => intval($request['absence_quota']),
                    'daily_work_hours' => intval($request['daily_work_hours']),
                    'weekly_work_hours' => intval($request['weekly_work_hours'])
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
