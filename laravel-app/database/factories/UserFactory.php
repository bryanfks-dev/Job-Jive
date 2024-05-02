<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Str;

/**
 * @extends \Illuminate\Database\Eloquent\Factories\Factory<\App\Models\User>
 */
class UserFactory extends Factory
{
    /**
     * The current password being used by the factory.
     */
    protected static ?string $password;

    /**
     * Define the model's default state.
     *
     * @return array<string, mixed>
     */
    public function definition(): array
    {
        return [
            'Full_Name' => fake()->name(),
            'Email' => fake()->unique()->safeEmail(),
            'Password' => static::$password ??= Hash::make('Password'),
            'Manager_ID' => mt_rand(1, 2),
            'Address' => fake()->address(),
            'NIK' => fake()->unique()->numberBetween(1000000000, 9999999999),
            'Gender' => fake()->randomElement(['Laki-laki', 'Perempuan']),
            'Phone_Number' => fake()->phoneNumber(),
            'Department_ID' => mt_rand(1, 5),
            'First_Login' => now()
        ];
    }

    /**
     * Indicate that the model's email address should be unverified.
     */
    
}
