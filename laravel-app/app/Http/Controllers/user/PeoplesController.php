<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Validator;
use Symfony\Component\HttpKernel\Exception\HttpException;

class PeoplesController extends Controller
{
    public function index(Request $request)
    {
        try {
            $responseDepartmentUsers = null;

            $httpHeaders = [
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            $param = trim($request->get('query', ''), ' ');

            if (!empty($param)) {
                $responseDepartmentUsers =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/users/me/department/users/search/' . $param);
            } else {
                $responseDepartmentUsers =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/users/me/department/users');
            }

            $responseProfile =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/profile');

            if ($responseDepartmentUsers->successful() && $responseProfile->successful()) {
                $paginatedEmployees =
                    $this->paginate($responseDepartmentUsers['data']['employees'] ?? [], 9);

                return view('user.peoples', [
                    'department_name' => $responseDepartmentUsers['data']['name'],
                    'manager' => $responseDepartmentUsers['data']['manager'],
                    'is_manager' => ($responseProfile['data']['as'] === 'Manager'),
                    'people_counts' => count($responseDepartmentUsers['data']['employees'] ?? []) +
                        (isset($responseDepartmentUsers['data']['manager']) ? 1 : 0),
                    'employees' => $paginatedEmployees,
                    'profile' => $responseProfile['data']
                ]);
            } else if ($responseDepartmentUsers->unauthorized() || $responseProfile->unauthorized()) {
                return redirect()->intended(route('user.login'));
            }

            return abort($responseDepartmentUsers->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($responseDepartmentUsers->status());
            }

            return abort(500);
        }
    }

    public function update(Request $request, int $id)
    {
        if ($id <= 0) {
            return abort(404);
        }

        $validator = Validator::make($request->all(), [
            'initial_salary' => ['required', 'min:1'],
            'current_salary' => ['required', 'min:1'],
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'update-error-' . $id => $validator->errors()->first(),
            ]);
        }

        try {
            $initial_salary = floatval(str_replace('.', '', $request->input('initial_salary')));
            $current_salary = floatval(str_replace('.', '', $request->input('current_salary')));

            $response = Http::withHeaders([
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Content-type' => 'application/json',
                'Accept' => 'application/json',
            ])->put(BackendServer::url() . '/api/users/me/department/users/' . $id . '/update/salary', [
                        'initial_salary' => $initial_salary,
                        'current_salary' => $current_salary,
                    ]);

            if ($response->successful()) {
                return redirect()->back()
                    ->with('update-success-' . $id, 'Salary updated successfully');
            } else if ($response->badRequest()) {
                return redirect()->intended(route('user.employees'))
                    ->withErrors([
                        'update-error-' . $id => $response['error']
                    ])->withInput([
                            'initial_salary' => $request['initial_salary'],
                            'current_salary' => $request['current_salary'],
                        ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } else {
                return redirect()->intended(route('user.peoples'))->withErrors([
                    'update-error-' . $id => $response['error'] ?? 'An error occurred',
                ])->withInput([
                            'initial_salary' => $request['initial_salary'],
                            'current_salary' => $request['current_salary'],
                        ]);
            }
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                return abort($response->status());
            }

            return abort(500);
        }
    }

    public function getAttendance(Request $request, int $id)
    {
        try {
            $httpHeaders = [
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            $responseAttendance =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/department/users/' . $id . '/attendance/data');

            $responseConfig =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/configs');

            $responseAttendenceStats =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/' . $id . '/attendance/stats');

            $responseProfile =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/users/me/profile');

            if (
                $responseAttendance->successful() && $responseConfig->successful() &&
                $responseAttendenceStats->successful() && $responseProfile->successful()
            ) {
                $attendances = $responseAttendance['data'][0];

                // Check for filter key
                $param =
                    trim($request->get('month', ''), ' ');

                if (!empty($param)) {
                    $maxMonth = $attendances['month'];

                    // Validate param value, so index won't be out of bound
                    if ($param <= $maxMonth && $param >= $maxMonth - 2) {
                        $attendances =
                            $responseAttendance['data'][$maxMonth - $param];
                    }
                }

                $attendances['records'] =
                    $this->paginate($attendances['records'] ?? [], 7);

                return view('user.employees.attendance', [
                    'configs' => $responseConfig['data'],
                    'attendance_stats' => $responseAttendenceStats['data'],
                    'old_month_id' => intval($request->get('month')),
                    'months' => array_map(function ($v) {
                        return [
                            'id' => $v['month'],
                            'name' => date(
                                'F',
                                mktime(0, 0, 0, $v['month'], 1)
                            ),
                        ];
                    }, $responseAttendance['data']),
                    'check_in_time' => $responseAttendance['data'][0]['records'][0]['check_in_time'],
                    'attendances' => $attendances,
                    'is_manager' => ($responseProfile['data']['as'] === 'Manager'),
                    'employee' => $id,
                    'profile' => $responseProfile['data']
                ]);
            } else if (
                $responseAttendance->unauthorized() || $responseConfig->unauthorized()
                || $responseAttendenceStats->unauthorized() || $responseProfile->unauthorized()
            ) {
                return redirect()->intended(route('user.login'));
            } else if (
                $responseAttendance->serverError() || $responseConfig->serverError()
                || $responseAttendenceStats->serverError() || $responseProfile->serverError()
            ) {
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

    public function updateAttendance(Request $request, int $id)
    {
        if ($id <= 0) {
            return abort(404);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Accept' => 'application/json',
                    'Content-Type' => 'application/json',
                ])->post(BackendServer::url() . '/api/users/me/department/users/' . $id . '/attendance/update', [
                            'user_id' => $request['user_id'],
                            'check_in_time' => $request['check_in_time'],
                            'check_out_time' => $request['check_out_time'],
                        ]);

            if ($response->successful() || $response->unprocessableEntity()) {
                return redirect()->intended(route('user.peoples.attendance', $id));
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('user.login'));
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
