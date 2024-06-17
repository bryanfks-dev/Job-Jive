@extends('layouts.app')

@section('content')
    <div class="antialiased bg-gray-50 dark:bg-gray-900">
        <!-- Profile Picture -->
        <div
            class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
            <span class="text-lg font-semibold text-gray-900 dark:text-white">Profile</span>
            <div class="mt-2">
                <div class="items-center block p-3 sm:flex">
                    <img class="w-32 h-32 md:w-40 md:h-40 mb-3 me-3 rounded-full sm:mb-0 object-cover"
                        src="{{ asset('/storage/img/user_profile/' . $user['photo']) }}" alt="Jese Leos image" />
                    <div class="text-gray-600 dark:text-gray-400 ml-0 sm:ml-5 mt-4 sm:mt-0">
                        <span class="font-medium text-gray-900 dark:text-white text-2xl">{{ $user['full_name'] }}</span>
                        <p class="text-md font-normal">{{ $user['as'] }}</p>
                        <span class="text-xs font-normal text-gray-500 dark:text-gray-400">
                            Since {{ $user['first_login'] ?? 'N/A' }}
                        </span>
                    </div>
                </div>
            </div>
        </div>

        <!-- Personal Information -->
        <div
            class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
            <span class="text-lg font-semibold text-gray-900 dark:text-white">Personal Information</span>
            <div class="mt-2">
                <div class="grid grid-cols-1 md:grid-cols-2">
                    <div class="w-full py-4 sm:px-4 pb-0 text-gray-900 md:pb-4 dark:text-white">
                        <ul class="space-y-4">
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Email</span>
                                <span class="text-md font-normal truncate">{{ $user['email'] }}</span>
                            </li>
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">NIK</span>
                                <span class="text-md font-normal truncate">{{ $user['nik'] }}</span>
                            </li>
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Phone Number</span>
                                <span class="text-md font-normal truncate">{{ $user['phone_number'] }}</span>
                            </li>
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Department</span>
                                <span class="text-md font-normal truncate">{{ $user['department']['name'] }}</span>
                            </li>
                        </ul>
                    </div>
                    <div class="py-4 sm:px-4 pb-0 text-gray-900 md:pb-4 dark:text-white">
                        <ul class="space-y-4">
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Address</span>
                                <span class="text-md font-normal break-words">{{ $user['address'] }}</span>
                            </li>
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Gender</span>
                                <span class="text-md font-normal truncate">{{ $user['gender'] }}</span>
                            </li>
                            <li class="flex flex-col">
                                <span class="text-sm font-normal text-gray-600 dark:text-gray-400">Salary</span>
                                <span class="text-md font-normal truncate">Rp {{ $user['salary']['current'] }}</span>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
