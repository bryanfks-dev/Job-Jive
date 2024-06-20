<div id="update-modal-{{ $employee['id'] }}" tabindex="-1" aria-hidden="true"
    class="hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
    <div class="relative w-full max-w-4xl max-h-full p-4">
        <!-- Modal content -->
        <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
            <!-- Modal header -->
            <div class="flex items-center justify-between p-4 border-b rounded-t md:p-5 dark:border-gray-600">
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    Update {{ $employee['full_name'] }}
                </h3>
                <button type="button"
                    class="inline-flex items-center justify-center w-8 h-8 text-sm text-gray-400 bg-transparent rounded-lg hover:bg-gray-200 hover:text-gray-900 ms-auto dark:hover:bg-gray-600 dark:hover:text-white"
                    data-modal-toggle="update-modal-{{ $employee['id'] }}">
                    <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 14 14">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
                    </svg>
                    <span class="sr-only">Close modal</span>
                </button>
            </div>
            <!-- Modal body -->
            <form class="p-4 md:p-5" action="{{ route('user.peoples.update', ['id' => $employee['id']]) }}"
                method="POST">
                @csrf
                @method('PUT')
                <div class="mb-4">
                    <div class="grid grid-cols-2 gap-4">
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Full
                                Name</label>
                            <input type="text"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. Budiono Santoso" value="{{ $employee['full_name'] }}" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email</label>
                            <input type="email"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. hello@example.com" value="{{ $employee['email'] }}" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Phone
                                Number</label>
                            <input type="tel"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 081234567890" value="{{ $employee['phone_number'] }}" minlength="11"
                                maxlength="13" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Birth
                                Date</label>
                            <input type="date"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                value="{{ $employee['birth_date'] }}" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Address</label>
                            <input type="text"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. Jl. Merpati No. 17" value="{{ $employee['address'] }}" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">NIK</label>
                            <input type="text"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 35-76-01-44-03-91-0003" value="{{ $employee['nik'] }}" minlength="16"
                                maxlength="16" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Gender</label>
                            <input type="text"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 35-76-01-44-03-91-0003" value="{{ $employee['gender'] }}" readonly>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Department</label>
                            <input type="text"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 35-76-01-44-03-91-0003" value="{{ $department_name }}" readonly>
                        </div>
                    </div>
                    <div class="grid mt-4">
                        <label for="initial_salary"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Initial Salary
                            <span class="text-red-600">(You can only edit this field)</span></label>
                        <input type="text" name="initial_salary" id="initial_salary"
                            class="salary-input bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                            placeholder="Input this employee's new salary"
                            value="{{ $employee['salary']['initial'] }}">
                        <label for="current_salary"
                            class="block mt-4 mb-2 text-sm font-medium text-gray-900 dark:text-white">Current Salary
                            <span class="text-red-600">(You can only edit this field)</span></label>
                        <input type="text" name="current_salary" id="current_salary"
                            class="salary-input bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                            placeholder="Input this employee's new salary"
                            value="{{ $employee['salary']['current'] }}">
                        {{-- Display success message --}}
                        @if (session('update-success-' . $employee['id']))
                            {{-- Show modal after redirect --}}
                            <script type="module">
                                addEventListener('load', () => {
                                    const id = {{ $employee['id'] }};
                                    const modal = FlowbiteInstances.getInstance('Modal', `update-modal-${id}`);

                                    modal.toggle();
                                });
                            </script>

                            <div class="flex gap-2 mt-4">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-green-500" width="1em"
                                    height="1em" viewBox="0 0 56 56">
                                    <path
                                        d="M28 51.906c13.055 0 23.906-10.828 23.906-23.906c0-13.055-10.875-23.906-23.93-23.906C14.899 4.094 4.095 14.945 4.095 28c0 13.078 10.828 23.906 23.906 23.906m-.023-20.39c-1.243 0-1.922-.727-1.97-1.97L25.68 17.97c-.047-1.29.937-2.203 2.273-2.203c1.313 0 2.32.937 2.274 2.226l-.329 11.555c-.047 1.265-.75 1.969-1.921 1.969m0 8.625c-1.36 0-2.626-1.078-2.626-2.532c0-1.453 1.243-2.53 2.626-2.53c1.359 0 2.624 1.054 2.624 2.53c0 1.477-1.289 2.532-2.624 2.532" />
                                </svg>
                                <p class="text-xs text-green-500">{{ session('update-success-' . $employee['id']) }}
                                </p>
                            </div>
                        @endif

                        {{-- Catch create failed --}}
                        @error('update-error-' . $employee['id'])
                            {{-- Show modal after redirect --}}
                            <script type="module">
                                addEventListener('load', function() {
                                    const id = {{ $employee['id'] }};
                                    const modal = FlowbiteInstances.getInstance('Modal', `update-modal-${id}`);

                                    modal.toggle();
                                });
                            </script>

                            <div class="flex gap-2 mt-4">
                                <svg xmlns="http://www.w3.org/2000/svg" class="fill-red-500" width="1em"
                                    height="1em" viewBox="0 0 56 56">
                                    <path
                                        d="M28 51.906c13.055 0 23.906-10.828 23.906-23.906c0-13.055-10.875-23.906-23.93-23.906C14.899 4.094 4.095 14.945 4.095 28c0 13.078 10.828 23.906 23.906 23.906m-.023-20.39c-1.243 0-1.922-.727-1.97-1.97L25.68 17.97c-.047-1.29.937-2.203 2.273-2.203c1.313 0 2.32.937 2.274 2.226l-.329 11.555c-.047 1.265-.75 1.969-1.921 1.969m0 8.625c-1.36 0-2.626-1.078-2.626-2.532c0-1.453 1.243-2.53 2.626-2.53c1.359 0 2.624 1.054 2.624 2.53c0 1.477-1.289 2.532-2.624 2.532" />
                                </svg>
                                <p class="text-xs text-red-500">{{ $errors->first('update-error-' . $employee['id']) }}
                                </p>
                            </div>
                        @enderror
                    </div>

                </div>
                <button type="submit"
                    class="text-white inline-flex items-center bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    Update employee
                </button>
            </form>
        </div>
    </div>
</div>
<script type="module">
    function formatSalaryInput(ele) {
        let value = ele.value;
        value = value.replace(/\D/g, ''); // Remove non-numeric characters
        value = value.replace(/\B(?=(\d{3})+(?!\d))/g, '.'); // Add dots

        ele.value = value;
    }

    const salaryInput = document.querySelectorAll('.salary-input');

    [...salaryInput].forEach((ele) => {
        formatSalaryInput(ele);
        ele.addEventListener('input', (e) => {
            formatSalaryInput(ele);
        });
    });
</script>
