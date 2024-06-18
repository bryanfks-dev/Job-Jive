<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration {
    /**
     * Run the migrations.
     */
    public function up(): void
    {
        Schema::create('attendance_statistics', function (Blueprint $table) {
            $table->integer('User_ID')->unsigned();
            $table->integer('Current_Week_Hours')->default(0);
            $table->integer('Current_Month_Hours')->default(0);
            $table->integer('Annual_Leaves')->default(0);

            $table->foreign('User_ID')->references('User_ID')
                ->on('users')->onDelete('cascade');

        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('attendance_statistics');
    }
};
