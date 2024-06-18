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
                            <span class="text-sm font-medium text-gray-900 dark:text-white" id="today-hour">0 /
                                <span class="font-normal">{{ $configs['daily_work_hours'] }} hrs</span></span>
                        </div>
                        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                            <div class="h-2.5 rounded-full bg-green-600 dark:bg-green-500" style="width: 1%" id="today-bar">
                            </div>
                        </div>
                    </div>
                    <!-- This Week -->
                    <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                        <div class="flex justify-between mb-1">
                            <span class="text-base font-medium text-gray-900 dark:text-white">This Week</span>
                            <span class="text-sm font-medium text-gray-900 dark:text-white" id="week-hour">0 /
                                {{ $configs['weekly_work_hours'] }} hrs</span>
                        </div>
                        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                            <div class="h-2.5 rounded-full bg-red-600 dark:bg-red-500" style="width: 1%" id="week-bar">
                            </div>
                        </div>
                    </div>
                </div>
                <div class="space-y-4">
                    <!-- This Month -->
                    <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                        <div class="flex justify-between mb-1">
                            <span class="text-base font-medium text-gray-900 dark:text-white">This Month</span>
                            <span class="text-sm font-medium text-gray-900 dark:text-white" id="month-hour">0 /
                                {{ 4 * $configs['weekly_work_hours'] }} hrs</span>
                        </div>
                        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                            <div class="h-2.5 rounded-full bg-yellow-400" style="width: 1%" id="month-bar"></div>
                        </div>
                    </div>
                    <!-- Annual Leaves -->
                    <div class="border rounded-xl p-5 border-gray-300 dark:border-gray-600">
                        <div class="flex justify-between mb-1">
                            <span class="text-base font-medium text-gray-900 dark:text-white">Annual Leaves</span>
                            <span class="text-sm font-medium text-gray-900 dark:text-white" id="absence-count">0 /
                                {{ $configs['absence_quota'] }} days</span>
                        </div>
                        <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                            <div class="h-2.5 rounded-full bg-indigo-600 dark:bg-indigo-500" style="width: 1%"
                                id="absence-bar"></div>
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
                        @csrf
                        <label for="month" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select
                            Month</label>
                        <select id="months-selector" name="month"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer">
                            @foreach ($months as $month)
                                <option value="{{ $month['id'] }}" {{ $old_month_id === $month['id'] ? 'selected' : '' }}>
                                    {{ $month['name'] }}</option>
                            @endforeach
                        </select>
                    </form>
                </div>
                <div>
                    @foreach ($attendances['records'] as $record)
                        @if (date('Y-m-d') === $record['date'])
                            @include('partials.user.attendance.today-record')
                        @else
                            @include('partials.user.attendance.regular-record')
                        @endif
                    @endforeach
                </div>

                <!-- Pagination -->
                <div class="p-4 mt-4">
                    {{ $attendances['records']->withPath(url()->current())->links() }}
                </div>
            </div>
        </div>
    </div>
    @if (isset($attendances['records'][0]['check_in_time']))
        {{-- Script stats handler --}}
        <script type="module">
            const today = {
                span: document.querySelector('#today-hour'),
                max: {{ $configs['daily_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#today-bar')
            };

            const week = {
                span: document.querySelector('#week-hour'),
                max: {{ $configs['weekly_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#week-bar')
            };

            const month = {
                span: document.querySelector('#month-hour'),
                max: {{ 4 * $configs['weekly_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#month-bar')
            };

            const absence = {
                span: document.querySelector('#absence-count'),
                max: {{ $configs['absence_quota'] }},
                unit: 'days',
                bar: document.querySelector('#absence-bar')
            };

            const uDate = '2016-01-02';

            const tz = '{{ Config::get('app.timezone') }}';

            let checkInTime = null;

            if ({{ count($attendances['records']) }} > 0) {
                if ('{{ date('Y-m-d') }}' === '{{ $attendances['records'][0]['date'] }}') {
                    checkInTime = new Date(`${uDate} {{ $attendances['records'][0]['check_in_time'] }}`);
                }
            }

            function initProgress(obj) {
                if (checkInTime !== null) {
                    let now = new Date().toLocaleString('en-GB', {
                        timeZone: tz,
                        hour: '2-digit',
                        minute: '2-digit',
                        second: '2-digit',
                        hour12: false
                    });

                    now = new Date(`${uDate} ${now}`);

                    // Calculate time diff
                    const timeDiff = now.getTime() - checkInTime.getTime();

                    let diff = Math.floor(timeDiff / (1000 * 3600));

                    diff = Math.min(Math.max(diff, 0), obj.max);

                    obj.span.textContent = `${diff} / ${obj.max} ${obj.unit}`;
                    obj.bar.style.width = `${100 * diff / obj.max}%`;
                }
            }

            (function startProgress() {
                initProgress(today);
                initProgress(week);
                initProgress(month);

                setTimeout(startProgress, 60000);
            })();
        </script>
    @endif
    {{-- Script form submit handler --}}
    <script type="module">
        const monthsForm = document.querySelector('#months-form');
        const monthsSelector = document.querySelector('#months-selector');

        monthsSelector.addEventListener('change', (e) => {
            monthsForm.submit();
        })
    </script>
@endsection
