<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\EmployeeWorkAssessment>
 */
class EmployeeWorkAssessmentFactory extends Factory
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
            'User_ID' => mt_rand(3, 6),
            'Feedback' => fake()->sentence(4),
            'Date_Time' => now()
        ];
    }
}
