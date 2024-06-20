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

            $responseProfile =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/profile');

            $responseConfig =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/configs');

            $responseTodayAttendance =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/attendance/today');

            $responseMotivation =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/motivation');

            $isManager = ($responseProfile['data']['as'] === 'Manager');

            $responseEmployeePeformance = null;
            $responseEmployeeAttendanceChart = null;

            if ($isManager) {
                $responseEmployeePeformance =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/users/me/department/users/attendance/stats');

                $responseEmployeeAttendanceChart = null;

                $param = trim($request->get('period', ''), ' ');

                if (!empty($param)) {
                    $responseEmployeeAttendanceChart =
                        \Http::withHeaders($httpHeaders)
                            ->get(BackendServer::url() . '/api/users/me/department/users/attendance/chart/' . $param);
                } else {
                    $responseEmployeeAttendanceChart =
                        \Http::withHeaders($httpHeaders)
                            ->get(BackendServer::url() . '/api/users/me/department/users/attendance/chart/3');
                }

                if (
                    $responseConfig->successful() && $responseTodayAttendance->successful() &&
                    ($responseMotivation->successful() || $responseMotivation->tooManyRequests()) && $responseEmployeeAttendanceChart->successful() && $responseEmployeePeformance->successful() &&
                    $responseProfile->successful()
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
                        'old_period' => (intval($param) !== 0 ? intval($param) : 3),
                        'employee_attendance_chart' => $responseEmployeeAttendanceChart['data'],
                        'employee_peformance' => $responseEmployeePeformance['data'] ?? [],
                        'profile' => $responseProfile['data'],
                        'is_manager' => $isManager
                    ]);
                } else if (
                    $responseConfig->unauthorized() || $responseTodayAttendance->unauthorized()
                    || $responseMotivation->unauthorized() || $responseEmployeeAttendanceChart->unauthorized() || $responseEmployeePeformance->unauthorized() || $responseProfile->unauthorized()
                ) {
                    return redirect()->intended(route('user.login'));
                } else if (
                    $responseConfig->serverError() || $responseTodayAttendance->serverError()
                    || $responseMotivation->serverError() || $responseEmployeeAttendanceChart->serverError() || $responseEmployeePeformance->serverError() || $responseProfile->serverError()
                ) {
                    return abort(500);
                }
            } else {
                if (
                    $responseConfig->successful() && $responseTodayAttendance->successful() &&
                    ($responseMotivation->successful() || $responseMotivation->tooManyRequests()) && $responseProfile->successful()
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
                        'is_manager' => $isManager
                    ]);
                } else if (
                    $responseConfig->unauthorized() || $responseTodayAttendance->unauthorized()
                    || $responseMotivation->unauthorized() || $responseProfile->unauthorized()
                ) {
                    return redirect()->intended(route('user.login'));
                } else if (
                    $responseConfig->serverError() || $responseTodayAttendance->serverError()
                    || $responseMotivation->serverError() || $responseProfile->serverError()
                ) {
                    return abort(500);
                }
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
