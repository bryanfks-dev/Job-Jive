@extends('layouts.app')

@section('content')
<div class="antialiased bg-gray-50 dark:bg-gray-900">
    <!-- Attendance -->
    <div
        class="p-5 mb-4 bg-white border border-gray-200 divide-y rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 divider-gray-200 dark:divide-gray-700">
        <span class="text-lg font-semibold text-gray-900 dark:text-white">Attendance Statistics</span>
        <div class="grid grid-cols-1 gap-4 py-5 mt-2 md:grid-cols-2 md:gap-6">
            <div class="space-y-4">
                <!-- Today -->
                <div class="p-5 border border-gray-300 rounded-xl dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">Today</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white" id="today-hour">0 /
                            {{ $configs['daily_work_hours'] }} hrs</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-green-600 dark:bg-green-500" style="width: 0%" id="today-bar">
                        </div>
                    </div>
                </div>
                <!-- This Week -->
                <div class="p-5 border border-gray-300 rounded-xl dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">This Week</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white" id="week-hour">{{
                            $attendance_stats['current_week_hours'] }} /
                            {{ $configs['weekly_work_hours'] }} hrs</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-red-600 dark:bg-red-500"
                            style="width: {{ ($attendance_stats['current_week_hours'] / $configs['weekly_work_hours']) * 100 }}%"
                            id="week-bar">
                        </div>
                    </div>
                </div>
            </div>
            <div class="space-y-4">
                <!-- This Month -->
                <div class="p-5 border border-gray-300 rounded-xl dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">This Month</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white" id="month-hour">{{
                            $attendance_stats['current_month_hours'] }} /
                            {{ 4 * $configs['weekly_work_hours'] }} hrs</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-yellow-400"
                            style="width: {{ ($attendance_stats['current_month_hours'] / $configs['weekly_work_hours']) * 25 }}%"
                            id="month-bar"></div>
                    </div>
                </div>
                <!-- Annual Leaves -->
                <div class="p-5 border border-gray-300 rounded-xl dark:border-gray-600">
                    <div class="flex justify-between mb-1">
                        <span class="text-base font-medium text-gray-900 dark:text-white">Annual Leaves</span>
                        <span class="text-sm font-medium text-gray-900 dark:text-white" id="absence-count">{{
                            $attendance_stats['annual_leaves'] }} /
                            {{ $configs['absence_quota'] }} days</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-700">
                        <div class="h-2.5 rounded-full bg-indigo-600 dark:bg-indigo-500"
                            style="width: {{ ($attendance_stats['annual_leaves'] / $configs['absence_quota']) * 100 }}%"
                            id="absence-bar"></div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <!-- History -->
    <div
        class="p-5 mb-4 bg-white border border-gray-200 divide-y rounded-lg shadow dark:bg-gray-800 dark:border-gray-700 divider-gray-200 dark:divide-gray-700">
        <span class="text-lg font-semibold text-gray-900 dark:text-white">History</span>
        <div class="mt-2">
            <div class="py-5">
                <form id="months-form" method="GET" action="">
                    <label for="month" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select
                        Month</label>
                    <select id="months-selector" name="month"
                        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 cursor-pointer">
                        @foreach ($months as $month)
                        <option value="{{ $month['id'] }}" {{ $old_month_id===$month['id'] ? 'selected' : '' }}>
                            {{ $month['name'] }}</option>
                        @endforeach
                    </select>
                </form>
            </div>
            <div>
                @foreach ($attendances['records'] as $record)
                @if (date('Y-m-d') === $record['date'])
                @include('partials.user.employees.attendance-log.today-record')
                @else
                @include('partials.user.employees.attendance-log.regular-record')
                @endif
                @endforeach
            </div>

            <!-- Pagination -->
            <div class="p-4 mt-4">
                {{ $attendances['records']->withPath(url()->current())->appends(['month' =>
                request()->get('month')])->links() }}
            </div>
        </div>
    </div>
</div>
@if (isset($check_in_time))
{{-- Script stats handler --}}
<script type="module">
    const today = {
                span: document.querySelector('#today-hour'),
                curr: 0,
                max: {{ $configs['daily_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#today-bar')
            };

            const week = {
                span: document.querySelector('#week-hour'),
                curr: {{ $attendance_stats['current_week_hours'] }},
                max: {{ $configs['weekly_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#week-bar')
            };

            const month = {
                span: document.querySelector('#month-hour'),
                curr: {{ $attendance_stats['current_month_hours'] }},
                max: {{ 4 * $configs['weekly_work_hours'] }},
                unit: 'hrs',
                bar: document.querySelector('#month-bar')
            };

            const dumyDate = '2016-01-02';
            const tz = '{{ Config::get('app.timezone') }}';
            const checkInTime = new Date(`${dumyDate} {{ $check_in_time }}`);

            function initProgress(obj) {
                if (checkInTime !== null) {
                    let currTime = new Date().toLocaleString('en-GB', {
                        timeZone: tz,
                        hour: '2-digit',
                        minute: '2-digit',
                        second: '2-digit',
                        hour12: false
                    });

                    currTime = new Date(`${dumyDate} ${currTime}`);

                    // Calculate time diff
                    const timeDiff = currTime.getTime() - checkInTime.getTime();

                    let diff = (timeDiff / 3600000).toFixed(1);

                    diff = Math.min(Math.max(diff, 0), obj.max);

                    obj.span.textContent = `${obj.curr + diff} / ${obj.max} ${obj.unit}`;
                    obj.bar.style.width = `${(obj.curr + diff / obj.max) * 100}%`;
                }
            }

            (function init() {
                initProgress(today);
                initProgress(week);
                initProgress(month);

                setTimeout(init, 60000);
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