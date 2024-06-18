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
        Schema::create('department_heads', function (Blueprint $table) {
            $table->tinyInteger('Department_ID')->unsigned();
            $table->integer('Manager_ID')->unsigned()->unique()->nullable();

            $table->foreign('Department_ID')->references('Department_ID')
                ->on('departments')->onDelete('cascade');
            $table->foreign('Manager_ID')->references('User_ID')->on('users')
                ->onDelete('set null');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        //
    }
};
