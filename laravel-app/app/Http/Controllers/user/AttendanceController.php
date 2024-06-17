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
            $responseAttendance =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json'
                ])->get(BackendServer::url() . '/api/user/attendance');

            $responseConfig =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'applicaton/json'
                ])->get(BackendServer::url() . '/api/configs');

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

            $attendances['records'] =
                $this->paginate($attendances['records'] ?? [], 7);

            if ($responseAttendance->successful() && $responseConfig->successful()) {
                return view('user.attendance', [
                    'old_month_id' => intval($request->get('month')),
                    'configs' => $responseConfig['data'],
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
                    'attendances' => $attendances
                ]);
            } else if ($responseAttendance->unauthorized() || $responseConfig->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } else if ($responseAttendance->serverError() || $responseConfig->serverError()) {
                return abort(500);
            }

            return abort(400);
        } catch (\Exception $e) {
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
