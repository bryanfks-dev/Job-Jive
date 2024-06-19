@extends('layouts.guest')

@section('content')
    <section class="bg-gray-50 dark:bg-gray-900">
        <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0 gap-8">
            <div class="flex flex-col gap-3 items-center">
                <img class="w-11 h-11 mr-2" src="{{ asset('/img/logo.svg') }}" alt="Job Jive logo">
                <h1 class="text-xl font-bold mb-4 text-gray-900 md:text-2xl dark:text-white">
                    Admin Panel
                </h1>
            </div>
            <div
                class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
                <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
                    <form class="space-y-4 md:space-y-6" method="POST">
                        @csrf
                        <div>
                            <label for="username"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Username</label>
                            <input name="username" id="username"
                                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                placeholder="ex. admin" required="">
                        </div>
                        <div>
                            <label for="password"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
                            <input type="password" name="password" id="password" placeholder="••••••••"
                                class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                required="">
                        </div>

                        {{-- Catch login failed --}}
                        @error('error')
                            <div class="flex gap-2">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-500" width="1em" height="1em"
                                    viewBox="0 0 56 56">
                                    <path
                                        d="M28 51.906c13.055 0 23.906-10.828 23.906-23.906c0-13.055-10.875-23.906-23.93-23.906C14.899 4.094 4.095 14.945 4.095 28c0 13.078 10.828 23.906 23.906 23.906m-.023-20.39c-1.243 0-1.922-.727-1.97-1.97L25.68 17.97c-.047-1.29.937-2.203 2.273-2.203c1.313 0 2.32.937 2.274 2.226l-.329 11.555c-.047 1.265-.75 1.969-1.921 1.969m0 8.625c-1.36 0-2.626-1.078-2.626-2.532c0-1.453 1.243-2.53 2.626-2.53c1.359 0 2.624 1.054 2.624 2.53c0 1.477-1.289 2.532-2.624 2.532" />
                                </svg>
                                <p class="text-red-500 text-xs">{{ $errors->first('error') }}</p>
                            </div>
                        @enderror

                        <button type="submit"
                            class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign
                            in</button>
                    </form>
                </div>
            </div>
        </div>
    </section>
@endsection
