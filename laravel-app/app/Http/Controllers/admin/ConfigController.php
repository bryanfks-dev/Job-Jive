<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;
use Symfony\Component\HttpKernel\Exception\HttpException;

class ConfigController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/configs');

            if ($response->successful()) {
                return view('admin.config', [
                    'configs' => $response['data']
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

    public function save(Request $request)
    {
        $validator =
            \Validator::make(request()->all(), [
                'check_in_time' => ['required', 'date_format:H:i', 'before:' . $request['check_out_time']],
                'check_out_time' => ['required', 'date_format:H:i', 'after:' . $request['check_in_time']],
                'absence_quota' => ['required', 'integer'],
                'daily_work_hours' => ['required', 'integer', 'lte:weekly_work_hours', 'min:1', 'max:24'],
                'weekly_work_hours' => ['required', 'integer', 'gte:daily_work_hours', 'min:1', 'max:168']
            ]);

        if ($validator->fails()) {
            return redirect()->back()
                ->withErrors(['error' => $validator->errors()->first()])
                ->withInput([
                    'check_in_time' => $request['check_in_time'],
                    'check_out_time' => $request['check_out_time'],
                    'absence_quota' => $request['absence_quota'],
                    'daily_work_hours' => $request['daily_work_hours'],
                    'weekly_work_hours' => $request['weekly_work_hours']
                ]);
        }

        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
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
                return redirect()->intended(route('admin.configs'));
            } else if ($response->badRequest()) {
                return redirect()->back()
                    ->withErrors(['error' => $response['error']])
                    ->withInput([
                        'check_in_time' => $request['check_in_time'],
                        'check_out_time' => $request['check_out_time'],
                        'absence_quota' => $request['absence_quota'],
                        'daily_work_hours' => $request['daily_work_hours'],
                        'weekly_work_hours' => $request['weekly_work_hours']
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
}
