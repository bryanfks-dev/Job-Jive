<tr>
    <th class="font-medium py-1">
        {{ $index + 1 }}
    </th>
    <th class="font-medium py-1">
        {{ $employee['user_full_name'] }}
    </th>
    <th class="font-medium py-1">
        <span class="text-green-500">{{ $employee['attend_count'] }}</span>/<span
            class="text-red-500">{{ $employee['absence_count'] }}</span>
    </th>
    <th class="font-medium py-1">
        <span
            class="text-green-500">{{ number_format(($employee['attend_count'] / ($employee['attend_count'] + $employee['absence_count'])) * 100, 2, '.', ',') }}%</span>
    </th>
</tr>
