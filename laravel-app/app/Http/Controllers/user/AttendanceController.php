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
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/user/attendance');

            $attendances = $response['data'][0];

            if ($request->has('month')) {
                $param = trim($request->get('month'));

                if (!empty($param)) {
                    $param = intval($param);

                    $maxMonth = $attendances['month'];

                    // Validate param value, so index won't be out of bound
                    if ($param <= $maxMonth && $param >= $maxMonth - 2) {
                        $attendances = $response['data'][$maxMonth - $param];
                    }
                }
            }

            $attendances['records'] =
                $this->paginate($attendances['records'] ?? [], 7);

            if ($response->successful()) {
                return view('user.attendance', [
                    'old_month_id' => intval($request->get('month')),
                    'stats' => [

                    ],
                    'months' =>
                        array_map(function ($v) {
                            return [
                                'id' => $v['month'],
                                'name' => date(
                                    'F',
                                    mktime(0, 0, 0, $v['month'], 1)
                                )
                            ];
                        }, $response['data']),
                    'attendances' => $attendances
                ]);
            }

        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
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
