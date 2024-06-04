<tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
    <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
        1
    </th>
    <td class="px-6 py-4">
        {{ $user['full_name'] }}
    </td>
    <td class="px-6 py-4">
        {{ $user['email'] }}
    </td>
    <td class="px-6 py-4">
        {{ $user['phone_number'] }}
    </td>
    <td class="px-6 py-4">
        {{ $user['birth_date'] }}
    </td>
    <td class="px-6 py-4">
        {{ $user['gender'] }}
    </td>
    <td class="px-6 py-4">
        {{ $user['department'] }}
    </td>
    <td class="flex px-6 py-4 justify-end items-center gap-2">
        {{-- Edit button --}}
        <a class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
            data-modal-target="update-modal-{{ $user['id'] }}" data-modal-toggle="update-modal-{{ $user['id'] }}">
            <svg xmlns="http://www.w3.org/2000/svg" class="fill-blue-700 hover:fill-blue-800" width="2em"
                height="2em" viewBox="0 0 24 24">
                <g fill-rule="evenodd" clip-rule="evenodd">
                    <path
                        d="M11.32 6.176H5c-1.105 0-2 .949-2 2.118v10.588C3 20.052 3.895 21 5 21h11c1.105 0 2-.948 2-2.118v-7.75l-3.914 4.144A2.46 2.46 0 0 1 12.81 16l-2.681.568c-1.75.37-3.292-1.263-2.942-3.115l.536-2.839c.097-.512.335-.983.684-1.352z" />
                    <path
                        d="M19.846 4.318a2.2 2.2 0 0 0-.437-.692a2 2 0 0 0-.654-.463a1.92 1.92 0 0 0-1.544 0a2 2 0 0 0-.654.463l-.546.578l2.852 3.02l.546-.579a2.1 2.1 0 0 0 .437-.692a2.24 2.24 0 0 0 0-1.635M17.45 8.721L14.597 5.7L9.82 10.76a.54.54 0 0 0-.137.27l-.536 2.84c-.07.37.239.696.588.622l2.682-.567a.5.5 0 0 0 .255-.145l4.778-5.06Z" />
                </g>
            </svg>
        </a>
        {{-- Delete button --}}
        <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
            data-modal-target="popup-modal-{{ $user['id'] }}" data-modal-toggle="popup-modal-{{ $user['id'] }}">
            <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-700 hover:fill-red-800" width="1.5em"
                height="1.5em" viewBox="0 0 15 15">
                <path fill-rule="evenodd"
                    d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3zM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM7 7v5h1V7zm-3 5V9h1v3zm6-3v3h1V9z"
                    clip-rule="evenodd" />
            </svg>
        </a>
    </td>
</tr>
