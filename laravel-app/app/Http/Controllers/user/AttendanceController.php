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
            $httpHeaders = [
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json'
            ];

            $responseAttendance =
                \Http::withHeaders($httpHeaders)->get(BackendServer::url() . '/api/user/attendance');

            $responseConfig =
                \Http::withHeaders($httpHeaders)->get(BackendServer::url() . '/api/configs');

            $responseAttendenceStats =
                \Http::withHeaders($httpHeaders)->get(BackendServer::url() . '/api/user/attendance/stats');

            $attendances = $responseAttendance['data'][0];

            // Check for filter key
            if ($request->has('month')) {
                $param = trim($request->get('month'));

                if (!empty($param)) {
                    $param = intval($param);

                    $maxMonth = $attendances['month'];

                    // Validate param value, so index won't be out of bound
                    if ($param <= $maxMonth && $param >= $maxMonth - 2) {
                        $attendances = $responseAttendance['data'][$maxMonth - $param];
                    }
                }
            }

            if ($responseAttendance->successful() && $responseConfig->successful()) {
                $attendances['records'] =
                    $this->paginate($attendances['records'] ?? [], 7);

                return view('user.attendance', [
                    'configs' => $responseConfig['data'],
                    'attendance_stats' => $responseAttendenceStats['data'],
                    'old_month_id' => intval($request->get('month')),
                    'months' =>
                        array_map(function ($v) {
                            return [
                                'id' => $v['month'],
                                'name' => date(
                                    'F',
                                    mktime(0, 0, 0, $v['month'], 1)
                                )
                            ];
                        }, $responseAttendance['data']),
                    'check_in_time' => $responseAttendance['data'][0]['records'][0]['check_in_time'],
                    'attendances' => $attendances
                ]);
            } else if ($responseAttendance->unauthorized() || $responseConfig->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } else if ($responseAttendance->serverError() || $responseConfig->serverError()) {
                return abort(500);
            }

            return abort($responseAttendance->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($responseAttendance->status());
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
                return abort($response->status());
            }

            return abort(500);
        }
    }
}
