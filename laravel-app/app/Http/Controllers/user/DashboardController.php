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

            $httpHeaders = [
                'Authorization' => 'Bearer ' . $token,
                'Accept' => 'application/json'
            ];

            $responseConfig =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/configs');

            $responseTodayAttendance =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/attendance/today');

            $responseMotivation =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/motivation');

            $responseEmployeePeformance =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/department/users/attendance/stats');

            if (
                $responseConfig->successful() && $responseTodayAttendance->successful() &&
                ($responseMotivation->successful() || $responseMotivation->tooManyRequests())
            ) {
                // Decide next needed check type
                $neededCheckType = null;

                if (
                    !isset($responseTodayAttendance['data']['check_in_time']) &&
                    !isset($responseTodayAttendance['data']['check_out_time'])
                ) {
                    $neededCheckType = 'check_in';
                } else if (
                    isset($responseTodayAttendance['data']['check_in_time']) &&
                    !isset($responseTodayAttendance['data']['check_out_time'])
                ) {
                    $neededCheckType = 'check_out';
                }

                $isLate = false;

                if (isset($responseTodayAttendance['data']['check_in_time'])) {
                    $dueTime =
                        \Carbon\Carbon::parse($responseConfig['data']['check_in_time']);
                    $checkInTime =
                        \Carbon\Carbon::parse($responseTodayAttendance['data']['check_in_time']);

                    $isLate = $dueTime->lt($checkInTime);
                }

                $motivation = '';

                if (!$responseMotivation->tooManyRequests()) {
                    $motivation = $responseMotivation['data']['motivation'];
                }

                return view('user.dashboard', [
                    'configs' => $responseConfig['data'],
                    'today_attendance' => [
                        'needed_check_type' => $neededCheckType,
                        'is_late' => $isLate
                    ],
                    'motivation' => $motivation,
                    'employee_peformance' => $responseEmployeePeformance['data'] ?? []
                ]);
            } else if (
                $responseConfig->unauthorized() || $responseTodayAttendance->unauthorized()
                || $responseMotivation->unauthorized()
            ) {
                return redirect()->intended(route('user.login'));
            } else if (
                $responseConfig->serverError() || $responseTodayAttendance->serverError()
                || $responseMotivation->serverError()
            ) {
                return abort(500);
            }

            return abort($responseTodayAttendance->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($responseTodayAttendance->status());
            }

            return abort(500);
        }
    }
}
