<div id="update-modal-{{ $user['id'] }}" tabindex="-1" aria-hidden="true"
    class="hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
    <div class="relative p-4 w-full max-w-4xl max-h-full">
        <!-- Modal content -->
        <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
            <!-- Modal header -->
            <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
                <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                    Update User
                </h3>
                <button type="button"
                    class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                    data-modal-toggle="update-modal-{{ $user['id'] }}">
                    <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none"
                        viewBox="0 0 14 14">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6" />
                    </svg>
                    <span class="sr-only">Close modal</span>
                </button>
            </div>
            <!-- Modal body -->
            <form class="p-4 md:p-5" action="{{ route('admin.users.update', ['id' => $user['id']]) }}" method="POST">
                @csrf
                @method('PUT')
                <div class="mb-4">
                    <div class="grid gap-4 grid-cols-2">
                        <div class="col-span-2 sm:col-span-1">
                            <label for="full_name"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Full
                                Name</label>
                            <input type="text" name="full_name" id="full_name"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. Budiono Santoso"
                                value="{{ empty(old('full_name')) ? $user['full_name'] : old('full_name') }}"
                                required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="email"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Email</label>
                            <input type="email" name="email" id="email"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. hello@example.com"
                                value="{{ empty(old('email')) ? $user['email'] : old('email') }}" required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="phone_number"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Phone
                                Number</label>
                            <input type="tel" name="phone_number" id="phone_number"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 081234567890"
                                value="{{ empty(old('phone_number')) ? $user['phone_number'] : old('phone_number') }}"
                                minlength="11" maxlength="13" required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="date_of_birth"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Birth Date</label>
                            <input type="date" name="date_of_birth" id="date_of_birth"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                value="{{ empty(old('date_of_birth')) ? $user['birth_date'] : old('date_of_birth') }}"
                                required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="address"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Address</label>
                            <input type="text" name="address" id="address"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. Jl. Merpati No. 17"
                                value="{{ empty(old('address')) ? $user['address'] : old('address') }}" required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="nik"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">NIK</label>
                            <input type="text" name="nik" id="nik"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                placeholder="ex. 35-76-01-44-03-91-0003"
                                value="{{ empty(old('nik')) ? $user['nik'] : old('nik') }}" minlength="16"
                                maxlength="16" required="">
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="gender"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Gender</label>
                            <select name="gender" id="gender"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                required>
                                @php
                                    $genders = ['Male', 'Female'];
                                @endphp
                                @foreach ($genders as $gender)
                                    <option value="{{ $gender }}"
                                        {{ $gender == (empty(old('gender')) ? $user['gender'] : old('gender')) ? 'selected' : '' }}>
                                        {{ $gender }}</option>
                                @endforeach
                            </select>
                        </div>
                        <div class="col-span-2 sm:col-span-1">
                            <label for="department_id"
                                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Department</label>
                            <select name="department_id" id="department_id"
                                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                                required>
                                @foreach ($departments as $department)
                                    <option value="{{ $department['id'] }}"
                                        {{ $department['id'] == (isset($user['department']) ? $user['department']['id'] : old('department_id')) ? 'selected' : '' }}>
                                        {{ $department['name'] }}</option>
                                @endforeach
                            </select>
                        </div>
                    </div>
                    <div class="grid mt-4">
                        <label for="new_password"
                            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">New Password</label>
                        <input type="password" name="new_password" id="new_password"
                            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                            placeholder="•••••••••">
                        {{-- Catch create failed --}}
                        @error('update-error-' . $user['id'])
                            {{-- Show modal after redirect --}}
                            <script type="module">
                                addEventListener('load', function() {
                                    const id = {{ $user['id'] }};
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
                                <p class="text-red-500 text-xs">{{ $errors->first('update-error-' . $user['id']) }}</p>
                            </div>
                        @enderror
                    </div>
                </div>
                <button type="submit"
                    class="text-white inline-flex items-center bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
                    Update user
                </button>
            </form>
        </div>
    </div>
</div>
