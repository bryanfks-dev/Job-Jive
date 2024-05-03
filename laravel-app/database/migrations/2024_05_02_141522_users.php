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
            $table->increments('User_ID');
            $table->string('Full_Name');
            $table->string('Email')->unique();
            $table->string('Password');
            $table->integer('Manager_ID')->unsigned()->nullable();
            $table->string('Address');
            $table->string('NIK', 16);
            $table->string('Gender', 10);
            $table->string('Phone_Number', 13);
            $table->integer('Department_ID')->unsigned();
            $table->date('First_Login')->nullable();

            $table->foreign('Manager_ID')->references('User_ID')->on('users')->onDelete('cascade');
            $table->foreign('Department_ID')->references('Department_ID')->on('divisions')->onDelete('cascade');
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
