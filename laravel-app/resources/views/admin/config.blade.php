@extends('layouts.admin-app')

@section('content')
    <form class="flex flex-col gap-4" method="POST" action="{{ route('admin.configs.save') }}">
        @csrf
        @method('PUT')
        <div class="grid gap-4 mb-2 grid-cols-2">
            <div class="col-span-2 sm:col-span-1">
                <label for="checkin_time" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Check-In Time</label>
                <input type="time" name="check_in_time" id="check_in_time"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    placeholder="ex. Sales"
                    value="{{ empty(old('check_in_time')) ? $configs['check_in_time'] : old('check_in_time') }}" required="">
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="check_out_time" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Check-Out Time</label>
                <input type="time" name="check_out_time" id="check_out_time"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    placeholder="ex. Sales"
                    value="{{ empty(old('check_out_time')) ? $configs['check_out_time'] : old('check_out_time') }}" required="">
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="absence_quota" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Absence Quota</label>
                <input type="number" name="absence_quota" id="absence_quota"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ empty(old('absence_quota')) ? $configs['absence_quota'] : old('absence_quota') }}" required=""
                    min="1">
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="daily_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Daily Work Hours</label>
                <input type="number" name="daily_work_hours" id="daily_work_hours"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ empty(old('daily_work_hours')) ? $configs['daily_work_hours'] : old('daily_work_hours') }}"
                    required="" min="1" max="24">
            </div>
            <div class="col-span-2 sm:col-span-1">
                <label for="weekly_work_hours" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Weekly Work Hours</label>
                <input type="number" name="weekly_work_hours" id="weekly_work_hours"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ empty(old('weekly_work_hours')) ? $configs['weekly_work_hours'] : old('daily_work_hours') }}"
                    required="" min="1" max="168">
            </div>
            {{-- <div class="col-span-2 sm:col-span-1">
                <label for="salary_cut" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">
                    Salary Cut Amount (Late Attendance)</label>
                <input type="number" name="salary_cut" id="salary_cut"
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                    value="{{ empty(old('salary_cut')) ? $configs['salary_cut'] : old('daily_work_hours') }}"
                    required="" min="1">
            </div> --}}
        </div>

        {{-- Catch write config failed --}}
        @error('error')
            <div class="flex gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-500" width="1em" height="1em" viewBox="0 0 56 56">
                    <path
                        d="M28 51.906c13.055 0 23.906-10.828 23.906-23.906c0-13.055-10.875-23.906-23.93-23.906C14.899 4.094 4.095 14.945 4.095 28c0 13.078 10.828 23.906 23.906 23.906m-.023-20.39c-1.243 0-1.922-.727-1.97-1.97L25.68 17.97c-.047-1.29.937-2.203 2.273-2.203c1.313 0 2.32.937 2.274 2.226l-.329 11.555c-.047 1.265-.75 1.969-1.921 1.969m0 8.625c-1.36 0-2.626-1.078-2.626-2.532c0-1.453 1.243-2.53 2.626-2.53c1.359 0 2.624 1.054 2.624 2.53c0 1.477-1.289 2.532-2.624 2.532" />
                </svg>
                <p class="text-red-500 text-xs">{{ $errors->first('error') }}</p>
            </div>
        @enderror

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
