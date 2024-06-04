@extends('layouts.admin-app')

@section('content')
    <div>
        <div class="flex justify-between gap-16 mb-6">
            {{-- Create button --}}
            <button type="button" data-modal-target="crud-modal" data-modal-toggle="crud-modal"
                class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full sm:rounded-lg text-sm px-4 py-2 text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 gap-2">
                <svg xmlns="http://www.w3.org/2000/svg" width="2em" height="2em" viewBox="0 0 256 256">
                    <path fill="currentColor"
                        d="M128 24a104 104 0 1 0 104 104A104.13 104.13 0 0 0 128 24m40 112h-32v32a8 8 0 0 1-16 0v-32H88a8 8 0 0 1 0-16h32V88a8 8 0 0 1 16 0v32h32a8 8 0 0 1 0 16" />
                </svg>
                <span class="hidden sm:block">Add User</span>
            </button>

            {{-- Search bar --}}
            <form class="flex items-center max-w-sm" method="GET">
                @csrf
                <label for="simple-search" class="sr-only">Search</label>
                <div class="relative w-full">
                    <input type="text" id="simple-search"
                        class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full  p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                        placeholder="Search..." name="query" />
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
                        <th scope="col" class="px-6 py-3" width="20%">
                            Name
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Email
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Birth Date
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Phone Number
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Gender
                        </th>
                        <th scope="col" class="px-6 py-3">
                            Department
                        </th>
                        <th scope="col" class="text-right px-6 py-3">
                            Action
                        </th>
                    </tr>
                </thead>
                <tbody>
                    @forelse ($users as $user)
                        @include('partials.admin.users.record')
                    @empty
                        <tr class="h-10"></tr>
                        <tr class="bg-gray-50 dark:bg-gray-800 dark:border-none">
                            <td class="text-center" colspan="8">No user is found</td>
                        </tr>
                    @endforelse
                </tbody>
            </table>

            @include('partials.admin.users.create-modal')

            @foreach ($users as $user)
                @include('partials.admin.users.update-modal')
                @include('partials.admin.users.delete-modal')
            @endforeach

            <div class="p-4 mt-4">
                {{ $users->withPath(url()->current())->links() }}
            </div>
        </div>
    </div>
@endsection
