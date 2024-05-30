@extends('layouts.admin-app')

@section('content')
    <div>
        <div class="flex justify-between gap-16 mb-6">
            {{-- Create button --}}
            <button type="button"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full sm:rounded-lg text-sm px-4 py-2 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 256 256"><path fill="currentColor" d="M128 24a104 104 0 1 0 104 104A104.13 104.13 0 0 0 128 24m40 112h-32v32a8 8 0 0 1-16 0v-32H88a8 8 0 0 1 0-16h32V88a8 8 0 0 1 16 0v32h32a8 8 0 0 1 0 16"/></svg>
                <span class="hidden sm:block">Add Employee</span>
            </button>

            {{-- Search bar --}}
            <form class="flex items-center max-w-sm">
                <label for="simple-search" class="sr-only">Search</label>
                <div class="relative w-full">
                    <input type="text" id="simple-search"
                        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full  p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                        placeholder="Search..." required />
                </div>
                <button type="submit"
                    class="p-2.5 ms-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 20 20">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
                    </svg>
                    <span class="sr-only">Search</span>
                </button>
            </form>

        </div>

        <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
            <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
                <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                    <tr>
                        <th scope="col" class="px-6 py-3">
                            #
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Name
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Email
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Phone Number
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Gender
                        </th>
                        <th scope="col" class="text-right px-6 py-3">
                            Action
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {{-- Employee 1 --}}
                    <tr
                        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                        <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            1
                        </th>
                        <td class="px-6 py-4">
                            Aldo Arista
                        </td>
                        <td class="px-6 py-4">
                            aldoarista@gmail.com
                        </td>
                        <td class="px-6 py-4">
                            081338556780
                        </td>
                        <td class="px-6 py-4">
                            Male
                        </td>
                        <td class="flex px-6 py-4 justify-end items-center gap-2">
                            {{-- Edit button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-blue-700 hover:fill-blue-800"
                                    width="2em" height="2em" viewBox="0 0 24 24">
                                    <g fill-rule="evenodd" clip-rule="evenodd">
                                        <path
                                            d="M11.32 6.176H5c-1.105 0-2 .949-2 2.118v10.588C3 20.052 3.895 21 5 21h11c1.105 0 2-.948 2-2.118v-7.75l-3.914 4.144A2.46 2.46 0 0 1 12.81 16l-2.681.568c-1.75.37-3.292-1.263-2.942-3.115l.536-2.839c.097-.512.335-.983.684-1.352z" />
                                        <path
                                            d="M19.846 4.318a2.2 2.2 0 0 0-.437-.692a2 2 0 0 0-.654-.463a1.92 1.92 0 0 0-1.544 0a2 2 0 0 0-.654.463l-.546.578l2.852 3.02l.546-.579a2.1 2.1 0 0 0 .437-.692a2.24 2.24 0 0 0 0-1.635M17.45 8.721L14.597 5.7L9.82 10.76a.54.54 0 0 0-.137.27l-.536 2.84c-.07.37.239.696.588.622l2.682-.567a.5.5 0 0 0 .255-.145l4.778-5.06Z" />
                                    </g>
                                </svg>
                            </a>

                            {{-- Delete button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-700 hover:fill-red-800"
                                    width="1.5em" height="1.5em" viewBox="0 0 15 15">
                                    <path fill-rule="evenodd"
                                        d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3zM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM7 7v5h1V7zm-3 5V9h1v3zm6-3v3h1V9z"
                                        clip-rule="evenodd" />
                                </svg>
                            </a>
                        </td>
                    </tr>

                    {{-- Employee 2 --}}
                    <tr
                        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                        <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            2
                        </th>
                        <td class="px-6 py-4">
                            Aldo Arista
                        </td>
                        <td class="px-6 py-4">
                            aldoarista@gmail.com
                        </td>
                        <td class="px-6 py-4">
                            081338556780
                        </td>
                        <td class="px-6 py-4">
                            Male
                        </td>
                        <td class="flex px-6 py-4 justify-end items-center gap-2">
                            {{-- Edit button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-blue-700 hover:fill-blue-800"
                                    width="2em" height="2em" viewBox="0 0 24 24">
                                    <g fill-rule="evenodd" clip-rule="evenodd">
                                        <path
                                            d="M11.32 6.176H5c-1.105 0-2 .949-2 2.118v10.588C3 20.052 3.895 21 5 21h11c1.105 0 2-.948 2-2.118v-7.75l-3.914 4.144A2.46 2.46 0 0 1 12.81 16l-2.681.568c-1.75.37-3.292-1.263-2.942-3.115l.536-2.839c.097-.512.335-.983.684-1.352z" />
                                        <path
                                            d="M19.846 4.318a2.2 2.2 0 0 0-.437-.692a2 2 0 0 0-.654-.463a1.92 1.92 0 0 0-1.544 0a2 2 0 0 0-.654.463l-.546.578l2.852 3.02l.546-.579a2.1 2.1 0 0 0 .437-.692a2.24 2.24 0 0 0 0-1.635M17.45 8.721L14.597 5.7L9.82 10.76a.54.54 0 0 0-.137.27l-.536 2.84c-.07.37.239.696.588.622l2.682-.567a.5.5 0 0 0 .255-.145l4.778-5.06Z" />
                                    </g>
                                </svg>
                            </a>

                            {{-- Delete button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-700 hover:fill-red-800"
                                    width="1.5em" height="1.5em" viewBox="0 0 15 15">
                                    <path fill-rule="evenodd"
                                        d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3zM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM7 7v5h1V7zm-3 5V9h1v3zm6-3v3h1V9z"
                                        clip-rule="evenodd" />
                                </svg>
                            </a>
                        </td>
                    </tr>

                    {{-- Employee 3 --}}
                    <tr
                        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                        <th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
                            3
                        </th>
                        <td class="px-6 py-4">
                            Aldo Arista
                        </td>
                        <td class="px-6 py-4">
                            aldoarista@gmail.com
                        </td>
                        <td class="px-6 py-4">
                            081338556780
                        </td>
                        <td class="px-6 py-4">
                            Male
                        </td>
                        <td class="flex px-6 py-4 justify-end items-center gap-2">
                            {{-- Edit button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-blue-700 hover:fill-blue-800"
                                    width="2em" height="2em" viewBox="0 0 24 24">
                                    <g fill-rule="evenodd" clip-rule="evenodd">
                                        <path
                                            d="M11.32 6.176H5c-1.105 0-2 .949-2 2.118v10.588C3 20.052 3.895 21 5 21h11c1.105 0 2-.948 2-2.118v-7.75l-3.914 4.144A2.46 2.46 0 0 1 12.81 16l-2.681.568c-1.75.37-3.292-1.263-2.942-3.115l.536-2.839c.097-.512.335-.983.684-1.352z" />
                                        <path
                                            d="M19.846 4.318a2.2 2.2 0 0 0-.437-.692a2 2 0 0 0-.654-.463a1.92 1.92 0 0 0-1.544 0a2 2 0 0 0-.654.463l-.546.578l2.852 3.02l.546-.579a2.1 2.1 0 0 0 .437-.692a2.24 2.24 0 0 0 0-1.635M17.45 8.721L14.597 5.7L9.82 10.76a.54.54 0 0 0-.137.27l-.536 2.84c-.07.37.239.696.588.622l2.682-.567a.5.5 0 0 0 .255-.145l4.778-5.06Z" />
                                    </g>
                                </svg>
                            </a>

                            {{-- Delete button --}}
                            <a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-700 hover:fill-red-800"
                                    width="1.5em" height="1.5em" viewBox="0 0 15 15">
                                    <path fill-rule="evenodd"
                                        d="M11 3V1.5A1.5 1.5 0 0 0 9.5 0h-4A1.5 1.5 0 0 0 4 1.5V3H0v1h1v9.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V4h1V3zM5 1.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5V3H5zM7 7v5h1V7zm-3 5V9h1v3zm6-3v3h1V9z"
                                        clip-rule="evenodd" />
                                </svg>
                            </a>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
@endsection
