<div
    class="flex flex-col gap-8 sm:flex-row justify-between bg-white w-full h-full rounded-lg col-span-full p-4 md:p-6 shadow dark:shadow-lg dark:border-gray-600">
    {{-- Images (kalau mau dipake) --}}
    {{-- <div class="sm:hidden flex justify-center items-center">
    <img src="{{ asset('img/not_attend.webp') }}" alt="">
</div> --}}

    <div class="flex flex-col-reverse gap-6 sm:flex-row justify-between w-full">
        {{-- Buttons --}}
        <form class="flex items-center justify-center gap-4" method="POST" action="{{ route('user.attend') }}">
            @csrf
            {{-- Not attend --}}
            <button type="submit"
                class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-yellow-400 rounded-xl sm:rounded-lg hover:bg-yellow-500 focus:ring-4 focus:outline-none focus:ring-yellow-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                <div class="text-left">
                    <div>Check In</div>
                    <div class="text-xs font-normal" id="clock">00:00:00</div>
                </div>
                <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 512 512">
                    <path fill="currentColor"
                        d="M392 80H232a56.06 56.06 0 0 0-56 56v104h153.37l-52.68-52.69a16 16 0 0 1 22.62-22.62l80 80a16 16 0 0 1 0 22.62l-80 80a16 16 0 0 1-22.62-22.62L329.37 272H176v104c0 32.05 33.79 56 64 56h152a56.06 56.06 0 0 0 56-56V136a56.06 56.06 0 0 0-56-56M80 240a16 16 0 0 0 0 32h96v-32Z" />
                </svg>
            </button>

            {{-- DON'T DELETE: On time --}}

            {{-- <a href="#awda"
        class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-green-500 hover:bg-green-600 rounded-xl ">
        <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
            <path fill="currentColor"
                d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
        </svg>
        <div class="text-left">
            <div>Check Out</div>
            <div class="text-xs font-normal">On-Time</div>
        </div>
    </a> --}}

            {{-- DON'T DELETE: Late --}}

            {{-- <a href="#awda"
        class="gap-3 inline-flex items-center w-full justify-between sm:justify-normal sm:w-auto px-6 py-3 text-sm font-medium text-center text-white bg-red-700 hover:bg-red-800 rounded-xl">
        <svg xmlns="http://www.w3.org/2000/svg" width="2.5em" height="2.5em" viewBox="0 0 16 16">
            <path fill="currentColor"
                d="M3 2.75C3 1.784 3.784 1 4.75 1h6.5c.966 0 1.75.784 1.75 1.75v3.457A5.5 5.5 0 0 0 7.257 15H4.75A1.75 1.75 0 0 1 3 13.25zM6 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10 2.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0m-6.853-.354l-.003.003a.5.5 0 0 0-.144.348v.006a.5.5 0 0 0 .146.35l2 2a.5.5 0 0 0 .708-.707L10.707 12H13.5a.5.5 0 0 0 0-1h-2.793l1.147-1.146a.5.5 0 0 0-.708-.708z" />
        </svg>
        <div class="text-left">
            <div>Check Out</div>
            <div class="text-xs font-normal">Late</div>
        </div>
    </a> --}}
        </form>


        <div class="flex justify-between items-center gap-4">
            <div class="flex flex-col gap-1">
                {{-- Warnanya diganti aja:
            Kalau 'Late': bg-red-700
            Kalau 'On Time': bg-green-700
            Kalau 'Belum attendance': bg-yellow-400
            --}}
                <div class="bg-yellow-400 rounded-full py-[2px] px-4 text-white text-sm text-center">Date</div>
                <div class="bg-yellow-400 rounded-full py-[2px] px-4 text-white text-sm text-center">Timelimit</div>
            </div>
            <div class="flex flex-col gap-1">
                <div>:</div>
                <div>:</div>
            </div>
            <div class="flex flex-col gap-1">
                <div>Tuesday, 29 Mei 2024</div>
                {{-- Kalau sudah check in, value nya berubah jadi timelimit Check-Out --}}
                <div>12.00 WIB</div>
            </div>
        </div>
    </div>
</div>
<script type="module">
    // Clock script
    const clock = document.querySelector('#clock');

    const initClock = () => {
        const time = new Date();

        const dataTime = {
            hour: time.getHours(),
            minute: time.getMinutes(),
            second: time.getSeconds()
        };

        // Format dataTime
        dataTime.hour = dataTime.hour < 10 ? '0' + dataTime.hour : dataTime.hour;
        dataTime.minute = dataTime.minute < 10 ? '0' + dataTime.minute : dataTime.minute;
        dataTime.second = dataTime.second < 10 ? '0' + dataTime.second : dataTime.second;

        clock.textContent = `${dataTime.hour}:${dataTime.minute}:${dataTime.second}`;
    };

    setInterval(initClock, 1000);
</script>
