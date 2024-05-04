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
        Schema::create('employee_work_assessments', function (Blueprint $table) {
            $table->integer('User_ID')->unsigned();
            $table->text('Feedback');
            $table->timestamp('Date_Time')->default(now());

            $table->foreign('User_ID')->references('User_ID')->on('users')->onDelete('cascade');
        });
    }

    /**
     * Reverse the migrations.
     */
    public function down(): void
    {
        Schema::dropIfExists('employee_work_assessment');
    }
};
