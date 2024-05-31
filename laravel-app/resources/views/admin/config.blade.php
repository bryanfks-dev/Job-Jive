@extends('layouts.admin-app')

@section('content')

<div class="flex flex-col gap-4">
    <div class="grid gap-4 mb-4 grid-cols-2">
        <div class="col-span-2 sm:col-span-1">
            <label for="checkin_time" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Check-In Time</label>
            <input type="time" name="checkin_time" id="checkin_time"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="ex. Sales" required="">
        </div>
        <div class="col-span-2 sm:col-span-1">
            <label for="checkout_hour" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Check-Out Time</label>
            <input type="time" name="checkout_hour" id="checkout_hour"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="ex. Sales" required="">
        </div>
        <div class="col-span-2 sm:col-span-1">
            <label for="absence_quota" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Absence Quota</label>
            <input type="number" name="absence_quota" id="absence_quota"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                required="" min="0">
        </div>
        <div class="col-span-2 sm:col-span-1">
            <label for="daily_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Daily Work Hours</label>
            <input type="number" name="daily_work_hours" id="daily_work_hours"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                required="" min="0">
        </div>
        <div class="col-span-2 sm:col-span-1">
            <label for="weekly_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                Weekly Work Hours</label>
            <input type="number" name="weekly_work_hours" id="weekly_work_hours"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                required="" min="0">
        </div>
    </div>

    <div class="mb-4">
        <button type="submit" class="bg-blue-700 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded">
            Submit Config
        </button>
    </div>
</div>

@endsection
