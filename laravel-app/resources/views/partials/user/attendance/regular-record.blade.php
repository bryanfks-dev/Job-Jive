<div
    class="p-5 mb-4 border border-gray-300 rounded-lg shadow bg-white dark:bg-gray-800 dark:border-gray-600 divide-y divider-gray-200 dark:divide-gray-700">
    <time class="text-lg font-semibold text-gray-900 dark:text-white">{{ date('l, d F Y', strtotime($record['date'])) }}</time>
    <ol class="mt-3 flex flex-wrap md:gap-6 text-gray-900 dark:text-white">
        <li class="items-center block p-3 sm:flex">
            <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg"
                width="24" height="24" viewBox="0 0 24 24">
                <path
                    d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z">
                </path>
            </svg>
            <span class="font-medium">Check-In</span>
        </li>
        <li class="items-center block p-3 sm:flex">
            @if (isset($record['check_out_time']))
                <svg class="fill-green-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg"
                    width="24" height="24" viewBox="0 0 24 24">
                    <path
                        d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm-1.999 14.413-3.713-3.705L7.7 11.292l2.299 2.295 5.294-5.294 1.414 1.414-6.706 6.706z">
                    </path>
                </svg>
            @else
                <svg class="fill-red-500 w-11 h-11 mb-3 me-3 rounded-full sm:mb-0" xmlns="http://www.w3.org/2000/svg"
                    width="24" height="24" viewBox="0 0 24 24">
                    <path
                        d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10 10-4.486 10-10S17.514 2 12 2zm4.207 12.793-1.414 1.414L12 13.414l-2.793 2.793-1.414-1.414L10.586 12 7.793 9.207l1.414-1.414L12 10.586l2.793-2.793 1.414 1.414L13.414 12l2.793 2.793z">
                    </path>
                </svg>
            @endif
            <span class="font-medium text-gray-900 dark:text-white">Check-Out</span>
        </li>
    </ol>
</div>
