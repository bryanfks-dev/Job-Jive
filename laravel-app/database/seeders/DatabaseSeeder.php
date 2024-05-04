<?php

use App\Models\User;
use App\Models\Division;
use App\Models\Salary;
use App\Models\EmployeeWorkAssessment;
use App\Models\Attendance;
// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        Division::factory(5)->create();

        User::create([
            'Full_Name' => 'Admin',
            'Email' => 'admin@localhost',
            'Password' => bcrypt('456'),
            'Manager_ID' => null,
            'Address' => 'Jl. Cangcimen',
            'NIK' => '123456789',
            'Gender' => 'Laki-laki',
            'Phone_Number' => '08123456789',
            'Department_ID' => 1,
            'First_Login' => now()
        ]);

        User::create([
            'Full_Name' => 'Kanz',
            'Email' => 'Kanz@localhost',
            'Password' => bcrypt('456'),
            'Manager_ID' => null,
            'Address' => 'Jl. Cangcimen',
            'NIK' => '123456789',
            'Gender' => 'Laki-laki',
            'Phone_Number' => '08123456789',
            'Department_ID' => 1,
            'First_Login' => now()
        ]);

        User::factory(4)->create();

        Attendance::factory(15)->create();
        EmployeeWorkAssessment::factory(13)->create();
        Salary::factory(10)->create();
    }
}
