<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { AddStudent } from '$lib/wailsjs/go/main/App';
	import { X } from '@lucide/svelte';
	import { Dialog, Label, Separator } from 'bits-ui';

	let address = $state('');
	let name = $state('');
	let sex = $state('Male');
	let age = $state(0);
	let birthday = $state('');
	let error = $state('');
	let isOpen = $state(false);
	let birthPlace = $state('');
	let nameOfChrist = $state('');
	let motherName = $state('');
	let kebele = $state('');
	let houseNo = $state('');
	let phoneNo = $state('');
	let email = $state('');
	let username = $state('');
	let priviesSchool = $state('');
	let educationLevel = $state('');
	let workPosition = $state('');
	let christFatherName = $state('');
	let location = $state('');
	let emergencyName = $state('');
	let emergencyPhoneNo = $state('');

	async function addStudent() {
		if (!name || !sex || age === 0 || !birthday || !address || !birthPlace || !nameOfChrist ||
			!motherName || !kebele || !houseNo || !phoneNo || !email || !username ||
			!priviesSchool || !educationLevel || !workPosition || !christFatherName ||
			!location || !emergencyName || !emergencyPhoneNo) {
			error = 'Please fill in all fields';
			return;
		}

		const newStudent = await AddStudent(
			birthday,
			address,
			name,
			sex,
			age,
			birthPlace,
			nameOfChrist,
			motherName,
			kebele,
			houseNo,
			phoneNo,
			email,
			username,
			priviesSchool,
			educationLevel,
			workPosition,
			christFatherName,
			location,
			emergencyName,
			emergencyPhoneNo
		);

		if (newStudent) {
			error = '';
			isOpen = false;
			invalidateAll();

			// Reset form
			name = '';
			sex = 'Male';
			age = 0;
			birthday = '';
			address = '';
			birthPlace = '';
			nameOfChrist = '';
			motherName = '';
			kebele = '';
			houseNo = '';
			phoneNo = '';
			email = '';
			username = '';
			priviesSchool = '';
			educationLevel = '';
			workPosition = '';
			christFatherName = '';
			location = '';
			emergencyName = '';
			emergencyPhoneNo = '';
		} else {
			error = 'Something went wrong';
		}
	}
</script>

<Dialog.Root open={isOpen} onOpenChange={(open) => (isOpen = open)}>
	<Dialog.Trigger class="bg-dark cursor-pointer text-background hover:bg-dark/95 inline-flex items-center justify-center px-4 py-2 font-semibold">
		Add Student
	</Dialog.Trigger>

	<Dialog.Portal>
		<Dialog.Overlay class="fixed inset-0 z-50 bg-black/80" />
		<Dialog.Content
			class="fixed left-1/2 top-1/2 z-50 w-full max-w-2xl translate-x-[-50%] translate-y-[-50%] border bg-background p-4 shadow-lg rounded-md"
		>
			<Dialog.Title class="text-center text-lg font-semibold mb-3">Add Student</Dialog.Title>

			{#if error !== ''}
				<p class="text-red-500 text-center text-sm mb-2">{error}</p>
			{/if}

			<Separator.Root class="bg-dark mb-3 h-px" />

			<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-3 text-sm">
				<!-- Input fields -->
				{#each [
					['name', 'Name'],
					['sex', 'Sex'],
					['age', 'Age'],
					['birthday', 'Birthday'],
					['address', 'Address'],
					['birthPlace', 'Birth Place'],
					['nameOfChrist', 'Name of Christ'],
					['motherName', "Mother's Name"],
					['kebele', 'Kebele'],
					['houseNo', 'House No'],
					['phoneNo', 'Phone No'],
					['email', 'Email'],
					['username', 'Username'],
					['priviesSchool', 'Previous School'],
					['educationLevel', 'Education Level'],
					['workPosition', 'Work Position'],
					['christFatherName', 'Christ Father Name'],
					['location', 'Location'],
					['emergencyName', 'Emergency Contact Name'],
					['emergencyPhoneNo', 'Emergency Phone No']
				] as [field, label]}
					<div class="flex flex-col gap-1">
						<Label.Root for={field} class="text-xs font-medium">{label}</Label.Root>

						{#if field === 'sex'}
							<select id={field} bind:value={sex} class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full">
								<option value="Male">Male</option>
								<option value="Female">Female</option>
							</select>
						{:else if field === 'birthday'}
							<input id={field} bind:value={birthday} type="date" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'age'}
							<input id={field} bind:value={age} type="number" min="0" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'email'}
							<input id={field} bind:value={email} type="email" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'phoneNo'}
							<input id={field} bind:value={phoneNo} type="tel" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'emergencyPhoneNo'}
							<input id={field} bind:value={emergencyPhoneNo} type="tel" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'name'}
							<input id={field} bind:value={name} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'address'}
							<input id={field} bind:value={address} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'birthPlace'}
							<input id={field} bind:value={birthPlace} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'nameOfChrist'}
							<input id={field} bind:value={nameOfChrist} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'motherName'}
							<input id={field} bind:value={motherName} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'kebele'}
							<input id={field} bind:value={kebele} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'houseNo'}
							<input id={field} bind:value={houseNo} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'username'}
							<input id={field} bind:value={username} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'priviesSchool'}
							<input id={field} bind:value={priviesSchool} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'educationLevel'}
							<input id={field} bind:value={educationLevel} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'workPosition'}
							<input id={field} bind:value={workPosition} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'christFatherName'}
							<input id={field} bind:value={christFatherName} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'location'}
							<input id={field} bind:value={location} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{:else if field === 'emergencyName'}
							<input id={field} bind:value={emergencyName} type="text" class="border shadow py-1.5 px-2 text-sm text-gray-700 focus:outline-none focus:ring w-full" />
						{/if}
					</div>
				{/each}
			</div>

			<!-- Save Button -->
			<div class="mt-5">
				<button
					class="w-full bg-dark text-background py-2 px-4 text-sm font-semibold hover:bg-dark/90"
					onclick={addStudent}
				>
					Save
				</button>
			</div>

			<Dialog.Close class="absolute right-4 top-4 cursor-pointer text-gray-600 hover:text-gray-900">
				<X class="size-4" />
				<span class="sr-only">Close</span>
			</Dialog.Close>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
