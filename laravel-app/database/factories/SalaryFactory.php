<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\Salary>
 */
class SalaryFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            //
            'User_ID' => mt_rand(1, 6),
            'Initial_Salary' => mt_rand(10000, 20000),
            'Final_Salary' => mt_rand(10000, 20000)
        ];
    }
}
