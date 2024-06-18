<?php

namespace App\Models;

class BackendServer
{
    static public function url(): string
    {
        return config('backend_server.address') . ':' . config('backend_server.port');
    }
}
