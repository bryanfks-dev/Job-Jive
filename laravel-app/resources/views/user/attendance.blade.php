@extends('layouts.app')

@section('content')
<div class="antialiased bg-gray-50 dark:bg-gray-900">
    <!-- Attendance -->
    <div class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
        <span class="text-lg font-semibold text-gray-900 dark:text-white">Attendance Statistics</span>
        <div class="mt-2 py-5 grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6">
            <div class="space-y-4">
                <!-- Today -->
                <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">Today</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white">4 / 8
                            <span class="font-normal">hrs</span></span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-green-600 dark:bg-green-500" style="width: 45%"></div>
                    </div>
                </div>
                <!-- This Week -->
                <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">This Week</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white">32 / 40
                            <span class="font-normal">hrs</span></span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-red-600 dark:bg-red-500" style="width: 80%"></div>
                    </div>
                </div>
            </div>
            <div class="space-y-4">
                <!-- This Month -->
                <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">This Month</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white">136 / 160
                            <span class="font-normal">hrs</span></span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-yellow-400" style="width: 85%"></div>
                    </div>
                </div>
                <!-- Annual Leaves -->
                <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">Annual Leaves</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white">8 / 14
                            <span class="font-normal">days</span></span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-indigo-600 dark:bg-indigo-500" style="width: 57.15%"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- History -->
    <div class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
        <span class="text-lg font-semibold text-gray-900 dark:text-white">History</span>
        <div class="mt-2">
            <div class="py-5">
                <form class="">
                    <label for="months" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select Month</label>
                    <select id="months" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer">
                    <option value="March">March</option>
                    <option value="April">April</option>
                    <option value="May">May</option>
                    </select>
                </form>
            </div>
            <div>
                @include('layouts.attendances.today_card')
                
                <div class="p-5 mb-4 border border-gray-300 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-600 divide-y divider-gray-200 dark:divide-gray-700">
                    <time class="text-lg font-semibold text-gray-900 dark:text-white">Wednesday, March 24th 2024</time>
                    <ol class="mt-3 flex flex-wrap md:gap-6 text-gray-900 dark:text-white">
                        <li class="items-center block p-3 sm:flex">
                            <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z"></path></svg>
                            <span class="font-medium">Succeeded Check-In</span>
                        </li>
                    </ol>
                </div>

                <div class="p-5 mb-4 border border-gray-300 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-600 divide-y divider-gray-200 dark:divide-gray-700">
                    <time class="text-lg font-semibold text-gray-900 dark:text-white">Tuesday, March 23th 2024</time>
                    <ol class="mt-3 flex flex-wrap md:gap-6 text-gray-900 dark:text-white">
                        <li class="items-center block p-3 sm:flex">
                            <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z"></path></svg>
                            <span class="font-medium">Succeeded Check-In</span>
                        </li>
                    </ol>
                </div>

                <div class="p-5 mb-4 border border-gray-300 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-600 divide-y divider-gray-200 dark:divide-gray-700">
                    <time class="text-lg font-semibold text-gray-900 dark:text-white">Monday, March 22th 2024</time>
                    <ol class="mt-3 flex flex-wrap md:gap-6 text-gray-900 dark:text-white">
                        <li class="items-center block p-3 sm:flex">
                            <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z"></path></svg>
                            <span class="font-medium">Succeeded Check-In</span>
                        </li>
                    </ol>
                </div>

                <div class="p-5 mb-4 border border-gray-300 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-600 divide-y divider-gray-200 dark:divide-gray-700">
                    <time class="text-lg font-semibold text-gray-900 dark:text-white">Friday, March 19th 2024</time>
                    <ol class="mt-3 flex flex-wrap md:gap-6 text-gray-900 dark:text-white">
                        <li class="items-center block p-3 sm:flex">
                            <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z"></path></svg>
                            <span class="font-medium">Succeeded Check-In</span>
                        </li>
                    </ol>
                </div>

            </div>

            <!-- Pagination -->
            <div class="flex items-center justify-between dark:bg-gray-800 px-4 py-3 sm:px-6">
                <div class="flex flex-1 justify-between sm:hidden">
                  <a href="#" class="relative inline-flex items-center rounded-md border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 px-4 py-2 text-sm font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700">Previous</a>
                  <a href="#" class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 px-4 py-2 text-sm font-medium text-gray-700 dark:text-gray-200 hover:bg-gray-50 dark:hover:bg-gray-700">Next</a>
                </div>
                <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
                  <div>
                    <p class="text-sm text-gray-700 dark:text-gray-200">
                      Showing
                      <span class="font-medium">1</span>
                      of
                      <span class="font-medium">4</span>
                      pages
                    </p>
                  </div>
                  <div>
                    <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
                      <a href="#" class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 focus:z-20 focus:outline-offset-0">
                        <span class="sr-only">Previous</span>
                        <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                          <path fill-rule="evenodd" d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z" clip-rule="evenodd" />
                        </svg>
                      </a>
                      <a href="#" aria-current="page" class="relative z-10 inline-flex items-center bg-indigo-600 px-4 py-2 text-sm font-semibold text-white focus:z-20 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">1</a>
                      <a href="#" class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 focus:z-20 focus:outline-offset-0">2</a>
                      <a href="#" class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 focus:z-20 focus:outline-offset-0 md:inline-flex">3</a>
                      <a href="#" class="relative hidden items-center px-4 py-2 text-sm font-semibold text-gray-900 dark:text-white ring-1 ring-inset ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 focus:z-20 focus:outline-offset-0 md:inline-flex">4</a>
                      <a href="#" class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 dark:ring-gray-600 hover:bg-gray-50 dark:hover:bg-gray-700 focus:z-20 focus:outline-offset-0">
                        <span class="sr-only">Next</span>
                        <svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                          <path fill-rule="evenodd" d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z" clip-rule="evenodd" />
                        </svg>
                      </a>
                    </nav>
                  </div>
                </div>
            </div>
        </div>
    </div>
</div>
@endsection
