@extends('layouts.app')

@section('content')
    <div class="grid grid-cols-1 lg:grid-cols-4 gap-y-4 lg:gap-4 mb-4">

        <div class="w-full bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
            <div class="flex justify-end px-4 pt-4">

            </div>
            <div class="flex flex-col items-center pb-10 w-full">
                <img class="w-24 h-24 mb-3 rounded-full shadow-lg object-cover" src="https://source.unsplash.com/v6771a4avV4"
                    alt="Chinese man image" />
                <span
                    class="mb-2 bg-blue-100 text-blue-800 text-xs font-semibold px-2.5 py-0.5 rounded dark:bg-blue-200 dark:text-blue-800">Manager</span>
                <h5 class="mb-1 text-xl font-medium text-gray-900 dark:text-white">Aldo Arista</h5>
                <span class="text-sm text-gray-500 dark:text-gray-400 text-wrap">aristaaldo2@gmail.com</span>
            </div>
        </div>
        <a href="#"
            class="flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow md:flex-row hover:bg-gray-100 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-700"
            style="grid-column: span 3; background-image: url('https://images.unsplash.com/photo-1523841589119-b55aee0f66e7?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'); background-size: cover;">
            <div class="flex flex-col justify-between p-4 leading-normal">
                <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">IT Department</h5>
                <p class="mb-3 font-normal text-gray-700 dark:text-gray-400">Jl. Merpati Putih No. 16A, Malang, Indonesia
                </p>
            </div>
        </a>
    </div>

    <div class="flex flex-col items-center p-7">
        <h4 class="text-2xl text-gray-900 dark:text-white">Employees</h4>
    </div>

    <form class="max-w-md mx-auto mb-3 w-full">
        <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
        <div class="relative">
            <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
                    fill="none" viewBox="0 0 20 20">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
                </svg>
            </div>
            <input type="search" id="default-search"
                class="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                placeholder="Search Employee by Name" required />
            <button type="submit"
                class="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Search</button>
        </div>
    </form>

    {{-- Employee Content --}}
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mb-4 mt-4 w-full">
        {{-- @if (!$employees->count()) --}}
        <div class="flex items-center justify-center h-60 md:col-span-3">
            <div class="text-center">
                <h2 class="text-2xl text-gray-600 dark:text-gray-400">No employees found</h2>
                <p class="text-gray-500 dark:text-gray-400">Please add an employee</p>
            </div>
        </div>
        {{-- @endif --}}

        @include('partials.user.employees.employee-card')
        {{-- @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card')
        @include('partials.user.employees.employee-card') --}}
    </div>

    {{-- @include('partials.user.employees.feedback-modal') --}}
@endsection
