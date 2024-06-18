<?php

namespace App\Http\Controllers\admin;

use App\Http\Controllers\Controller;
use App\Models\BackendServer;
use Illuminate\Http\Request;
use Intervention\Image\Drivers\Gd\Driver;
use Intervention\Image\ImageManager;
use Symfony\Component\HttpKernel\Exception\HttpException;

class UsersController extends Controller
{
    public function index(Request $request)
    {
        try {
            $responseUser = null;

            $param = trim($request->get('query', ''), ' ');

            $httpHeaders = [
                'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                'Accept' => 'application/json',
            ];

            if (!empty($param)) {
                $responseUser =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/user/search/' . $param);
            } else {
                $responseUser =
                    \Http::withHeaders($httpHeaders)
                        ->get(BackendServer::url() . '/api/users');
            }

            $responseDepartment =
                \Http::withHeaders($httpHeaders)
                    ->get(BackendServer::url() . '/api/departments');

            if ($responseDepartment->successful() && $responseUser->successful()) {
                $paginatedUsers =
                    $this->paginate($responseUser['data'] ?? []);

                return view('admin.users', [
                    'users' => $paginatedUsers,
                    'departments' => $responseDepartment['data'] ?? [],
                ]);
            } else if ($responseDepartment->unauthorized() || $responseUser->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            } else if ($responseDepartment->serverError() || $responseUser->serverError()) {
                return abort(500);
            }

            return abort(400);
        } catch (\Exception $e) {
            return abort(500);
        }
    }

    public function create(Request $request)
    {
        $validator = \Validator::make($request->all(), [
            'full_name' => ['required'],
            'email' => ['required', 'email'],
            'phone_number' => ['required', 'max:13', 'regex:/^\d+$/'],
            'date_of_birth' => ['required', 'date'],
            'address' => ['required'],
            'nik' => ['required', 'digits:16'],
            'gender' => ['required', 'in:Male,Female'],
            'department_id' => ['required', 'integer'],
            'photo.*' => ['required', 'mimes:png,jpg,jpeg'],
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'create-error' => $validator->errors()->first(),
            ])
                ->withInput([
                    'full_name' => $request['full_name'],
                    'email' => $request['email'],
                    'date_of_birth' => $request['date_of_birth'],
                    'phone_number' => $request['phone_number'],
                    'address' => $request['address'],
                    'nik' => $request['nik'],
                    'gender' => $request['gender'],
                    'department_id' => $request['department_id'],
                    'photo' => $request['photo'],
                ]);
        }

        try {
            $photo = $request->file('photo');
            $fileName = time() . '_user_' . $request['full_name'] . '.jpeg';

            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->post(BackendServer::url() . '/api/user/create', [
                            'full_name' => $request['full_name'],
                            'email' => $request['email'],
                            'date_of_birth' => $request['date_of_birth'],
                            'phone_number' => $request['phone_number'],
                            'address' => $request['address'],
                            'nik' => $request['nik'],
                            'gender' => $request['gender'],
                            'department_id' => intval($request['department_id']),
                            'photo' => $fileName,
                        ]);

            if ($response->successful()) {
                $imgManager = new ImageManager(new Driver());
                $img =
                    $imgManager->read($photo->getRealPath());

                $img->resize(500, 500)->toPng();

                \Storage::put(
                    '/public/img/user_profile/' . $fileName,
                    (string) $img->encode()
                );

                return redirect()->intended(route('admin.users'));
            } else if ($response->badRequest()) {
                return redirect()->back()->withErrors([
                    'create-error' => $response['error'],
                ])
                    ->withInput([
                        'full_name' => $request['full_name'],
                        'email' => $request['email'],
                        'date_of_birth' => $request['date_of_birth'],
                        'phone_number' => $request['phone_number'],
                        'address' => $request['address'],
                        'nik' => $request['nik'],
                        'gender' => $request['gender'],
                        'department_id' => $request['department_id'],
                        'photo' => $request['photo'],
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }

    public function update(Request $request, int $id)
    {
        $id = intval($id);

        $validator = \Validator::make($request->all(), [
            'full_name' => ['required'],
            'email' => ['required', 'email'],
            'phone_number' => ['required', 'max:13', 'regex:/^\d+$/'],
            'date_of_birth' => ['required', 'date'],
            'address' => ['required'],
            'nik' => ['required', 'digits:16'],
            'gender' => ['required', 'in:Male,Female'],
            'department_id' => ['required', 'integer'],
        ]);

        if ($validator->fails()) {
            return redirect()->back()->withErrors([
                'update-error-' . $id => $validator->errors()->first(),
            ])
                ->withInput([
                    'full_name' => $request['full_name'],
                    'email' => $request['email'],
                    'date_of_birth' => $request['date_of_birth'],
                    'phone_number' => $request['phone_number'],
                    'address' => $request['address'],
                    'nik' => $request['nik'],
                    'gender' => $request['gender'],
                    'department_id' => $request['department_id'],
                ]);
        }

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->put(BackendServer::url() . '/api/user/update/' . $id, [
                            'full_name' => $request['full_name'],
                            'email' => $request['email'],
                            'date_of_birth' => $request['date_of_birth'],
                            'phone_number' => $request['phone_number'],
                            'address' => $request['address'],
                            'nik' => $request['nik'],
                            'gender' => $request['gender'],
                            'department_id' => intval($request['department_id']),
                            'new_password' => $request['new_password'],
                        ]);

            if ($response->successful()) {
                return redirect()->intended(route('admin.users'));
            } else if ($response->badRequest()) {
                return redirect()->intended(route('admin.users'))->withErrors([
                    'update-error-' . $id => $response['error'],
                ])
                    ->withInput([
                        'full_name' => $request['full_name'],
                        'email' => $request['email'],
                        'date_of_birth' => $request['date_of_birth'],
                        'phone_number' => $request['phone_number'],
                        'address' => $request['address'],
                        'nik' => $request['nik'],
                        'gender' => $request['gender'],
                        'department_id' => $request['department_id'],
                    ]);
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
            }

            return abort($response->status());
        } catch (\Exception $e) {
            if ($e instanceof HttpException) {
                throw new HttpException($response->status());
            }

            return abort(500);
        }
    }

    public function delete(Request $request, int $id)
    {
        $id = intval($id);

        try {
            $response =
                \Http::withHeaders([
                    'Authorization' => 'Bearer ' . $request->cookie('auth_token'),
                    'Content-type' => 'application/json',
                    'Accept' => 'application/json',
                ])->delete(BackendServer::url() . '/api/user/delete/' . $id);

            if ($response->successful()) {
                return redirect()->intended(route('admin.users'));
            } else if ($response->unauthorized()) {
                return redirect()->intended(route('admin.login'));
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
