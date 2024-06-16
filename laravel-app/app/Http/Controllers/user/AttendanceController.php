<?php

namespace App\Http\Controllers\user;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Symfony\Component\HttpKernel\Exception\HttpException;

class AttendanceController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response = [];

            if ($request->has('filter')) {
                $param = trim($request->get('filter'));

                if (!empty($param)) {

                }
            } else {
                $response =
                    \Http::withHeaders([
                        'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                        'Accept' => 'application/json'
                    ])->get(BackendServer::url() . '/api/user/attendance');


            }
        } catch(\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }

        return view('user.attendance');
    }

    public function attend(Request $request)
    {
        try {
            $dateTime =
                \Carbon\Carbon::now(config('app.timezone'));

            $response = \Http::withHeaders([
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json',
                'Content-Type' => 'application/json'
            ])->post(BackendServer::url() . '/api/user/attend', [
                'date' => $dateTime->format('Y-m-d'),
                'time' => $dateTime->format('H:i:s')
            ]);

            if ($response->successful()) {
                return redirect()->intended(route('user.dashboard'));
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
