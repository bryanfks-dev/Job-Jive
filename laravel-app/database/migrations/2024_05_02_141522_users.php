<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('users', function (Blueprint $table) {
            $table->integerIncrements('User_ID')->primary();
            $table->string('Full_Name');
            $table->string('Email')->unique();
            $table->string('Password');
            $table->integer('Manager_ID')->unsigned()->nullable();
            $table->date('Date_of_Birth');
            $table->string('Address');
            $table->string('NIK', 16);
            $table->enum('Gender', ['Male', 'Female']);
            $table->string('Phone_Number', 13);
            $table->tinyInteger('Department_ID')->unsigned()->nullable();
            $table->date('First_Login')->nullable();

            $table->foreign('Manager_ID')->references('User_ID')
                ->on('users')->onDelete('set null');
            $table->foreign('Department_ID')->references('Department_ID')
                ->on('departments')->onDelete('set null');

            $table->timestamps();
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('users');
    }
};
