<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class LoginController extends Controller
{
    public function view(Request $req) {
        $msg = $req->query('msg');

        return view('login', [
            'msg' => $msg
        ]);
    }
}
