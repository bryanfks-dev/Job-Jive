@extends('layouts.app')

@section('content')
    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
        @include('partials.user.dashboard.attendance-checker')
        <div class="rounded-lg dark:border-gray-600 h-full md:h-full">
            <div class="w-full bg-white rounded-lg shadow dark:bg-gray-800 p-4 md:p-6">
                <div class="flex justify-between items-start w-full">
                    <div class="flex-col items-center">
                        <div class="flex items-center mb-1">
                            <h5 class="text-xl font-bold leading-none text-gray-900 dark:text-white me-1">Employee
                                Attendance
                            </h5>
                        </div>
                        <div class="flex justify-between items-center">
                            <!-- Button -->
                            <button id="dropdownDefaultButton" data-dropdown-toggle="lastDaysdropdown"
                                data-dropdown-placement="bottom"
                                class="text-sm font-medium text-gray-500 dark:text-gray-400 hover:text-gray-900 text-center inline-flex items-center dark:hover:text-white"
                                type="button">
                                Last 7 days
                                <svg class="w-2.5 m-2.5 ms-1.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                                    fill="none" viewBox="0 0 10 6">
                                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"
                                        stroke-width="2" d="m1 1 4 4 4-4" />
                                </svg>
                            </button>
                            <div id="lastDaysdropdown"
                                class="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700">
                                <ul class=text-sm text-gray-700 dark:text-gray-200" aria-labelledby="dropdownDefaultButton">
                                    <li>
                                        <a href="#"
                                            class="block px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Yesterday</a>
                                    </li>
                                    <li>
                                        <a href="#"
                                            class="block px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Today</a>
                                    </li>
                                    <li>
                                        <a href="#"
                                            class="block px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Last
                                            7 days</a>
                                    </li>
                                    <li>
                                        <a href="#"
                                            class="block px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Last
                                            30 days</a>
                                    </li>
                                    <li>
                                        <a href="#"
                                            class="block px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Last
                                            90 days</a>
                                    </li>
                                </ul>
                            </div>
                        </div>
                        <div id="dateRangeDropdown"
                            class="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-80 lg:w-96 dark:bg-gray-700 dark:divide-gray-600">
                            <div class="p-3" aria-labelledby="dateRangeButton">
                                <div date-rangepicker datepicker-autohide class="flex items-center">
                                    <div class="relative">
                                        <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                                            <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true"
                                                xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                                                <path
                                                    d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z" />
                                            </svg>
                                        </div>
                                        <input name="start" type="text"
                                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                            placeholder="Start date">
                                    </div>
                                    <span class="mx-2 text-gray-500 dark:text-gray-400">to</span>
                                    <div class="relative">
                                        <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                                            <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true"
                                                xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                                                <path
                                                    d="M20 4a2 2 0 0 0-2-2h-2V1a1 1 0 0 0-2 0v1h-3V1a1 1 0 0 0-2 0v1H6V1a1 1 0 0 0-2 0v1H2a2 2 0 0 0-2 2v2h20V4ZM0 18a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8H0v10Zm5-8h10a1 1 0 0 1 0 2H5a1 1 0 0 1 0-2Z" />
                                            </svg>
                                        </div>
                                        <input name="end" type="text"
                                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full ps-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                            placeholder="End date">
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Line Chart -->
                <div class="py-6" id="pie-chart"></div>

                <div
                    class="grid grid-cols-2 items-center border-gray-200 border-t dark:border-gray-700 justify-between pt-4">
                    <div>
                        <h1 class="font-semibold text-gray-500 mb-2">Total Attendance</h1>
                        <p class="text-2xl">92</p>
                    </div>

                    <div>
                        <h1 class="font-semibold text-gray-500 mb-2">Total Absence</h1>
                        <p class="text-2xl">8</p>
                    </div>
                </div>
            </div>
        </div>

        <div class="border bg-white rounded-lg shadow-sm p-4 md:p-6 text-gray-900">
            <h1 class="font-bold text-xl mb-8">Employee Performance</h1>

            {{-- Headline Top Attendance --}}
            <div class="flex gap-3 mb-2">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none">
                    <circle cx="12" cy="12" r="12" fill="#6DB653" />
                    <path
                        d="M17.7459 15.3512L13.0824 7.27389C12.4871 6.24204 11.5129 6.24204 10.9176 7.27389L6.25407 15.3512C5.65877 16.3831 6.14583 17.2259 7.33644 17.2259H16.6636C17.8542 17.2259 18.3412 16.3817 17.7459 15.3512Z"
                        fill="white" />
                </svg>
                <h2 class="font-semibold opacity-80 text-base">Top Attendance</h2>
            </div>

            {{-- Table Top Attendance --}}
            <div class="w-full overflow-x-auto">
                <table class="w-full table-fixed text-sm text-left">
                    <thead class="opacity-50 border-b-[1.5px]">
                        <tr>
                            <th scope="col" class="w-[12%] py-3 font-medium">
                                No
                            </th>
                            <th scope="col" class="w-[48%] py-3 font-medium">
                                Name
                            </th>
                            <th scope="col" class="w-[20%] py-3 font-medium">
                                Att/Abs
                            </th>
                            <th scope="col" class="w-[20%] py-3 font-medium">
                                Percentage
                            </th>
                        </tr>
                    </thead>
                    <tr class="h-4"></tr>
                    <tbody>
                        {{-- Empty State --}}
                        <tr class="h-5"></tr>
                        <tr class= "dark:bg-gray-800 dark:border-none">
                            <td class="text-center" colspan="4">No employee yet</td>
                        </tr>
                        <tr class="h-5"></tr>

                        {{-- Employees --}}
                        <tr>
                            <th class="font-medium py-1">
                                1
                            </th>
                            <th class="font-medium py-1">
                                Aldo Arista
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">90</span>/<span class="text-red-500">10</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">98.4%</span>
                            </th>
                        </tr>

                        <tr>
                            <th class="font-medium py-1">
                                2
                            </th>
                            <th class="font-medium py-1">
                                Budi Wahyu
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">89</span>/<span class="text-red-500">11</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">96.7%</span>
                            </th>
                        </tr>

                        <tr>
                            <th class="font-medium py-1">
                                3
                            </th>
                            <th class="font-medium py-1">
                                Clara
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">88</span>/<span class="text-red-500">12</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">95.6%</span>
                            </th>
                        </tr>
                    </tbody>
                </table>

                {{-- Headline Worst Attendance --}}
                <div class="flex gap-3 mb-2 mt-8">
                    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                        fill="none">
                        <circle cx="12" cy="12" r="12" transform="rotate(180 12 12)" fill="#C64747" />
                        <path
                            d="M6.25407 9.64876L10.9176 17.7261C11.5129 18.758 12.4871 18.758 13.0824 17.7261L17.7459 9.64876C18.3412 8.6169 17.8542 7.7741 16.6636 7.7741L7.33644 7.7741C6.14583 7.7741 5.65877 8.61835 6.25407 9.64876Z"
                            fill="white" />
                    </svg>
                    <h2 class="font-semibold opacity-80 text-base">Worst Attendance</h2>
                </div>

                {{-- Table Worst Attendance --}}
                <table class="w-full table-fixed text-sm text-left">
                    <thead class="opacity-50 border-b-[1.5px]">
                        <tr>
                            <th scope="col" class="w-[12%] py-3 font-medium">
                                No
                            </th>
                            <th scope="col" class="w-[48%] py-3 font-medium">
                                Name
                            </th>
                            <th scope="col" class="w-[20%] py-3 font-medium">
                                Att/Abs
                            </th>
                            <th scope="col" class="w-[20%] py-3 font-medium">
                                Percentage
                            </th>
                        </tr>
                    </thead>
                    <tr class="h-4"></tr>
                    <tbody>
                        {{-- Empty State --}}
                        <tr class="h-5"></tr>
                        <tr class= "dark:bg-gray-800 dark:border-none">
                            <td class="text-center" colspan="4">No employee yet</td>
                        </tr>
                        <tr class="h-5"></tr>

                        {{-- Employees --}}
                        <tr>
                            <th class="font-medium py-1">
                                1
                            </th>
                            <th class="font-medium py-1">
                                Aldo Arista
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">90</span>/<span class="text-red-500">10</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">98.4%</span>
                            </th>
                        </tr>

                        <tr>
                            <th class="font-medium py-1">
                                2
                            </th>
                            <th class="font-medium py-1">
                                Budi Wahyu
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">89</span>/<span class="text-red-500">11</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">96.7%</span>
                            </th>
                        </tr>

                        <tr>
                            <th class="font-medium py-1">
                                3
                            </th>
                            <th class="font-medium py-1">
                                Clara
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">88</span>/<span class="text-red-500">12</span>
                            </th>
                            <th class="font-medium py-1">
                                <span class="text-green-500">95.6%</span>
                            </th>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
@endsection
