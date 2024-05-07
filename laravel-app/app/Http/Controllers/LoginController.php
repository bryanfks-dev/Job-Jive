<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class LoginController extends Controller
{
    public function viewUser(Request $req) {
        $msg = $req->query('msg');

        return view('login-user', [
            'msg' => $msg
        ]);
    }
}
