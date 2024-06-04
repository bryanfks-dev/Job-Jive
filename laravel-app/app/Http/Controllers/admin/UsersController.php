<?php

namespace App\Http\Controllers\admin;

use Illuminate\Http\Request;
use App\Models\BackendServer;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Http;
use Illuminate\Pagination\LengthAwarePaginator;

class UsersController extends Controller
{
    private function paginate(array $items, int $perPage = 10, ?int $page = null, $options = []): LengthAwarePaginator
    {
        $page = $page ?: (LengthAwarePaginator::resolveCurrentPage() ?: 1);
        $items = collect($items);

        return new LengthAwarePaginator(
            $items->forPage($page, $perPage),
            $items->count(),
            $perPage,
            $page,
            $options
        );
    }

    public function index(Request $request)
    {
        try {
            $response =
                Http::withHeaders([
                    'Authorization' => 'Bearer ' . session('token'),
                    'Accept' => 'applications/json'
                ])->get(BackendServer::url() . '/api/users');

            if ($response->successful()) {
                switch ($response['status']) {
                    case 200: // Ok
                        if ($request->has('query')) {
                            $query = $request->get('query');

                            $results = [];

                            foreach ($response['data'] as $user) {
                                if (in_array(strtolower($query), array_map('strtolower', $user))) {
                                    $results[] = $user;
                                }
                            }

                            $paginatedUsers = $this->paginate($results ?? []);

                            return view("admin.users", [
                                'users' => $paginatedUsers
                            ]);
                        }

                        $paginatedUsers = $this->paginate($response['data'] ?? []);

                        return view("admin.users", [
                            'users' => $paginatedUsers
                        ]);

                    case 401: // Unauthorized
                        return redirect()->intended(route('admin.login'));
                }

                return abort($response['status']);
            }

            return abort(400); // Bad request
        } catch (\Exception $e) {
            if ($e instanceof \Symfony\Component\HttpKernel\Exception\HttpException) {
                return abort($e->getStatusCode());
            }

            return abort(500);
        }
    }
}
