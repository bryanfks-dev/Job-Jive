<div
    class="flex flex-col gap-8 sm:flex-row justify-between bg-white w-full h-full rounded-lg col-span-full p-4 md:p-6 shadow dark:shadow-lg dark:border-gray-600">
    <div class="flex flex-col-reverse gap-6 sm:flex-row justify-between w-full">
        {{-- Buttons --}}
        <form class="flex items-center justify-center gap-4" method="POST" action="{{ route('user.attend') }}">
            @csrf
            @if ($today_attendance['needed_check_type'] === 'check_in')
                {{-- Check-in --}}
                <button type="submit"
                    class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-yellow-400 rounded-xl sm:rounded-lg hover:bg-yellow-500 focus:ring-4 focus:outline-none focus:ring-yellow-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    <div class="text-left">
                        <div>Check In</div>
                        <div class="text-xs font-normal" id="time">00:00:00</div>
                    </div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 512 512">
                        <path fill="currentColor"
                            d="M392 80H232a56.06 56.06 0 0 0-56 56v104h153.37l-52.68-52.69a16 16 0 0 1 22.62-22.62l80 80a16 16 0 0 1 0 22.62l-80 80a16 16 0 0 1-22.62-22.62L329.37 272H176v104c0 32.05 33.79 56 64 56h152a56.06 56.06 0 0 0 56-56V136a56.06 56.06 0 0 0-56-56M80 240a16 16 0 0 0 0 32h96v-32Z" />
                    </svg>
                </button>
            @else
                @if ($today_attendance['is_late'])
                    {{-- Late check in --}}
                    @if (!isset($today_attendance['needed_check_type']))
                        <button type="submit" id="check-out-btn" disabled
                            class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-red-800 rounded-xl">
                            <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
                                <path fill="currentColor"
                                    d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
                            </svg>
                            <div class="text-left">
                                <div>Check Out</div>
                                <div class="text-xs font-normal">Late</div>
                            </div>
                        </button>
                    @else
                        <button type="submit" id="check-out-btn"
                            class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-red-700 hover:bg-red-800 rounded-xl">
                            <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
                                <path fill="currentColor"
                                    d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
                            </svg>
                            <div class="text-left">
                                <div>Check Out</div>
                                <div class="text-xs font-normal">Late</div>
                            </div>
                        </button>
                    @endif
                @else
                    {{-- On time check in --}}
                    @if (!isset($today_attendance['needed_check_type']))
                        <button type="submit" id="check-out-btn" disabled
                            class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-green-600 rounded-xl ">
                            <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
                                <path fill="currentColor"
                                    d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
                            </svg>
                            <div class="text-left">
                                <div>Check Out</div>
                                <div class="text-xs font-normal">On-Time</div>
                            </div>
                        </button>
                    @else
                        <button type="submit" id="check-out-btn"
                            class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-green-500 hover:bg-green-600 rounded-xl ">
                            <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
                                <path fill="currentColor"
                                    d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
                            </svg>
                            <div class="text-left">
                                <div>Check Out</div>
                                <div class="text-xs font-normal">On-Time</div>
                            </div>
                        </button>
                    @endif
                @endif
            @endif
        </form>

        <div class="flex justify-between items-center gap-4">
            <div class="flex flex-col gap-1">
                @if ($today_attendance['needed_check_type'] === 'check_in')
                    <div class="bg-yellow-400 rounded-full py-[2px] px-4 text-white text-sm text-center">Date</div>
                    <div class="bg-yellow-400 rounded-full py-[2px] px-4 text-white text-sm text-center">Due Time</div>
                @else
                    @if ($today_attendance['is_late'])
                        <div class="bg-red-700 rounded-full py-[2px] px-4 text-white text-sm text-center">Date</div>
                        <div class="bg-red-700 rounded-full py-[2px] px-4 text-white text-sm text-center">Due Time</div>
                    @else
                        <div class="bg-green-500 rounded-full py-[2px] px-4 text-white text-sm text-center">Date</div>
                        <div class="bg-green-500 rounded-full py-[2px] px-4 text-white text-sm text-center">Due Time
                        </div>
                    @endif
                @endif
            </div>
            <div class="flex flex-col gap-1">
                <div>:</div>
                <div>:</div>
            </div>
            <div class="flex flex-col gap-1">
                <div id="date">{{ \Carbon\Carbon::today()->format('l, d M Y') }}</div>
                <div>
                    {{ $today_attendance['needed_check_type'] === 'check_in' ? $configs['check_in_time'] : $configs['check_out_time'] }}
                </div>
            </div>
        </div>
    </div>
</div>
@if ($today_attendance['needed_check_type'] === 'check_in')
    <script type="module">
        // Clock script
        const time = document.querySelector('#time');
        const tz = '{{ Config::get('app.timezone') }}';

        (function init() {
            const currTime = new Date().toLocaleString('en-GB', {
                timeZone: tz,
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit',
                hour12: false
            });

            // Display date & time
            time.textContent = currTime;

            setTimeout(init, 1000);
        })();
    </script>
@else
    <script type="module">
        // Button submit preventer script
        const checkOutBtn = document.querySelector('#check-out-btn');

        const tz = '{{ Config::get('app.timezone') }}';
        const dumyDate = '2024-01-02';

        const minCheckOutTime = new Date(`${dumyDate} {{ $configs['check_out_time'] }}`);

        (function init() {
            let currTime = new Date().toLocaleString('en-GB', {
                timeZone: tz,
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit',
                hour12: false
            });

            currTime = new Date(`${dumyDate} ${currTime}`);

            // Prevent user to check out before check out time
            if (currTime < minCheckOutTime) {
                checkOutBtn.disabled = true;
            }

            setTimeout(init, 1000);
        })();
    </script>
@endif
