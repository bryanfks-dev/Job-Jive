<?php

namespace App\Http\Controllers\user;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Facades\Validator;
use Symfony\Component\HttpKernel\Exception\HttpException;

class PeoplesController extends Controller
{
    public function index(Request $request)
    {
        try {
            $response = null;

            $httpHeaders = [
                'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            $param = trim($request->get('query', ''), ' ');

            if (! empty($param)) {
                $response =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url().'/api/users/me/department/users/search/'.$param);
            } else {
                $response =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url().'/api/users/me/department/users');
            }

            if ($response->successful()) {
                $paginatedEmployees =
                    $this->paginate($response['data']['employees'] ?? [], 9);

                return view('user.peoples', [
                    'department_name' => $response['data']['name'],
                    'manager' => $response['data']['manager'],
                    'is_manager' => $response['data']['is_manager'],
                    'employees' => $paginatedEmployees,
                ]);
            } elseif ($response->unauthorized()) {
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

    public function update(Request $request, int $id)
    {
        $id = intval($id);

        if ($id <= 0) {
            return abort(404);
        }

        $validator = Validator::make($request->all(), [
            'initial_salary' => ['required', 'min:1'],
            'current_salary' => ['required', 'min:1'],
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'update-error-'.$id => $validator->errors()->first(),
            ]);
        }

        try {
            $initial_salary = floatval(str_replace('.', '', $request->input('initial_salary')));
            $current_salary = floatval(str_replace('.', '', $request->input('current_salary')));

            // dd($request->all());

            $response = Http::withHeaders([
                'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                'Content-type' => 'application/json',
                'Accept' => 'application/json',
            ])->put(BackendServer::url().'/api/users/me/department/users/update/'.$id, [
                'initial_salary' => $initial_salary,
                'current_salary' => $current_salary,
            ]);

            if ($response->successful()) {
                return redirect()->back()->with('update-success-'.$id, 'Salary updated successfully');
            } elseif ($response->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } else {
                Log::error('Failed to update employee salary', [
                    'status' => $response->status(),
                    'response' => $response->body(),
                ]);

                return redirect()->intended(route('user.peoples'))->withErrors([
                    'update-error-'.$id => $response['error'] ?? 'An error occurred',
                ])->withInput([
                    'initial_salary' => $request['initial_salary'],
                    'current_salary' => $request['current_salary'],
                ]);
            }
        } catch (\Exception $e) {
            Log::error('Exception while updating employee salary', [
                'error' => $e->getMessage(),
            ]);

            return abort(500);
        }
    }

    public function getAttendance(Request $request, int $id)
    {
        try {
            $httpHeaders = [
                'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            $responseAttendance =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url().'/api/users/me/department/users/employee-attendance/'.$id);

            $responseConfig =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url().'/api/configs');

            $responseAttendenceStats =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url().'/api/users/me/attendance/stats');

            $responseManager =
             \Http::withHeaders($httpHeaders)
                 ->get(BackendServer::url().'/api/users/me/department/users');

            if ($responseAttendance->successful() && $responseConfig->successful()) {
                $attendances = $responseAttendance['data'][0];

                // Check for filter key
                $param =
                    trim($request->get('month', ''), ' ');

                if (! empty($param)) {
                    $maxMonth = $attendances['month'];

                    // Validate param value, so index won't be out of bound
                    if ($param <= $maxMonth && $param >= $maxMonth - 2) {
                        $attendances =
                            $responseAttendance['data'][$maxMonth - $param];
                    }
                }

                $attendances['records'] =
                    $this->paginate($attendances['records'] ?? [], 7);
                // dd($responseEmployees['data']['employees']);

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
                    'is_manager' => $responseManager['data']['is_manager'],
                    'employee' => $id,
                ]);
            } elseif ($responseAttendance->unauthorized() || $responseConfig->unauthorized()) {
                return redirect()->intended(route('user.login'));
            } elseif ($responseAttendance->serverError() || $responseConfig->serverError()) {
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
        try {

            $response = \Http::withHeaders([
                'Authorization' => 'Bearer '.$request->cookie('auth_token'),
                'Accept' => 'application/json',
                'Content-Type' => 'application/json',
            ])->post(BackendServer::url().'/api/users/me/department/users/update-employee-attendance/'.$id, [
                'user_id' => $request['user_id'],
                'check_in_time' => $request['check_in_time'],
                'check_out_time' => $request['check_out_time'],
            ]);

            if ($response->successful()) {
                return redirect()->intended(route('user.peoples.attendance', $id));
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
