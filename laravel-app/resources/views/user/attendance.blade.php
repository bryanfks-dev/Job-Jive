@extends('layouts.app')

@section('content')
    <div class="antialiased bg-gray-50 dark:bg-gray-900">
        <!-- Attendance -->
        <div
            class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
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
        <div
            class="p-5 mb-4 border border-gray-200 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-700 divide-y divider-gray-200 dark:divide-gray-700">
            <span class="text-lg font-semibold text-gray-900 dark:text-white">History</span>
            <div class="mt-2">
                <div class="py-5">
                    <form id="months-form" method="GET" action="">
                        <label for="month" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select
                            Month</label>
                        <select id="months-selector" name="month"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer">
                            @foreach ($months as $month)
                                <option value="{{ $month['id'] }}" {{ $old_month_id === $month['id'] ? 'selected' : '' }}>{{ $month['name'] }}</option>
                            @endforeach
                        </select>
                    </form>
                </div>
                <div>
                    @forelse ($attendances['records'] as $record)
                        @if (date('Y-m-d') == $record['date'])
                            @include('partials.user.attendance.today-record')
                        @else
                            @include('partials.user.attendance.regular-record')
                        @endif
                    @empty
                        <div class="flex items-center justify-center h-60 col-span-3">
                            <div class="text-center">
                                <h2 class="text-2xl text-gray-600 dark:text-gray-400">No attendance found</h2>
                                <p class="text-gray-500 dark:text-gray-400">Please add an attendance</p>
                            </div>
                        </div>
                    @endforelse
                </div>

                <!-- Pagination -->
                <div class="p-4 mt-4">
                    {{ $attendances['records']->withPath(url()->current())->links() }}
                </div>
            </div>
        </div>
    </div>
    <script type="module">
        const monthsForm = document.querySelector('#months-form');
        const monthsSelector = document.querySelector('#months-selector');

        monthsSelector.addEventListener('change', (e) => {
            monthsForm.submit();
        })
    </script>
@endsection
