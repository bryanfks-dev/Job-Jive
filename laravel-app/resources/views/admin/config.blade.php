@extends('layouts.admin-app')

@section('content')
    <form class="flex flex-col gap-4" method="POST" action="{{ route('admin.configs.save') }}">
        @csrf
        @method('PUT')
        <div class="grid gap-4 mb-4 grid-cols-2">
            <div class="col-span-2 sm:col-span-1">
                <label for="checkin_time" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Check-In Time</label>
                <input type="time" name="check_in_time" id="check_in_time"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    placeholder="ex. Sales" value="{{ $check_in_time }}" required="">
                @error('check_in_time')
                    <p class="text-red-500 text-xs mt-1">{{ $errors->first('check_in_time') }}</p>
                @enderror
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="check_out_time" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Check-Out Time</label>
                <input type="time" name="check_out_time" id="check_out_time"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    placeholder="ex. Sales" value="{{ $check_out_time }}" required="">
                @error('check_out_time')
                    <p class="text-red-500 text-xs mt-1">{{ $errors->first('check_out_time') }}</p>
                @enderror
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="absence_quota" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Absence Quota</label>
                <input type="number" name="absence_quota" id="absence_quota"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ $absence_quota }}" required="" min="1">
                @error('absence_quota')
                    <p class="text-red-500 text-xs mt-1">{{ $errors->first('absence_quota') }}</p>
                @enderror
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="daily_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Daily Work Hours</label>
                <input type="number" name="daily_work_hours" id="daily_work_hours"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ $daily_work_hours }}" required="" min="1" max="24">
                @error('daily_work_hours')
                    <p class="text-red-500 text-xs mt-1">{{ $errors->first('daily_work_hours') }}</p>
                @enderror
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="weekly_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Weekly Work Hours</label>
                <input type="number" name="weekly_work_hours" id="weekly_work_hours"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ $weekly_work_hours }}" required="" min="1" max="168">
                @error('weekly_work_hours')
                    <p class="text-red-500 text-xs mt-1">{{ $errors->first('weekly_work_hours') }}</p>
                @enderror
            </div>
        </div>

        <div class="mb-4">
            <button type="submit"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full sm:rounded-lg text-sm px-4 py-3 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" width="1.5em" height="1.5em" viewBox="0 0 16 16">
                    <path fill="currentColor"
                        d="M9.405 1.05c-.413-1.4-2.397-1.4-2.81 0l-.1.34a1.464 1.464 0 0 1-2.105.872l-.31-.17c-1.283-.698-2.686.705-1.987 1.987l.169.311c.446.82.023 1.841-.872 2.105l-.34.1c-1.4.413-1.4 2.397 0 2.81l.34.1a1.464 1.464 0 0 1 .872 2.105l-.17.31c-.698 1.283.705 2.686 1.987 1.987l.311-.169a1.464 1.464 0 0 1 2.105.872l.1.34c.413 1.4 2.397 1.4 2.81 0l.1-.34a1.464 1.464 0 0 1 2.105-.872l.31.17c1.283.698 2.686-.705 1.987-1.987l-.169-.311a1.464 1.464 0 0 1 .872-2.105l.34-.1c1.4-.413 1.4-2.397 0-2.81l-.34-.1a1.464 1.464 0 0 1-.872-2.105l.17-.31c.698-1.283-.705-2.686-1.987-1.987l-.311.169a1.464 1.464 0 0 1-2.105-.872zM8 10.93a2.929 2.929 0 1 1 0-5.86a2.929 2.929 0 0 1 0 5.858z" />
                </svg>
                <span class="hidden sm:block">Save Config</span>
            </button>
        </div>
    </form>
@endsection
