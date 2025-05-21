<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { EditStudent } from '$lib/wailsjs/go/main/App';
	import { X } from '@lucide/svelte';
	import { Dialog, Label, Separator } from 'bits-ui';

	let { student } = $props();

	// Use local editable state for all fields
	let name = $state(student.Name);
	let sex = $state(student.Sex);
	let age = $state(student.Age);
	let birthday = $state(student.Birthday);
	let address = $state(student.Address);
	let birthPlace = $state(student.BirthPlace);
	let nameOfChrist = $state(student.NameOfChrist);
	let motherName = $state(student.MotherName);
	let kebele = $state(student.Kebele);
	let houseNo = $state(student.HouseNo);
	let phoneNo = $state(student.PhoneNo);
	let email = $state(student.Email);
	let username = $state(student.Username);
	let previousSchool = $state(student.PriviesSchool);
	let educationLevel = $state(student.EducationLevel);
	let workPosition = $state(student.WorkPosition);
	let christFatherName = $state(student.ChristFatherName);
	let location = $state(student.Location);
	let emergencyName = $state(student.EmergencyName);
	let emergencyPhoneNo = $state(student.EmergencyPhoneNo);

	let error = $state('');
	let isOpen = $state(false);

	async function editStudent() {
		if (!name || !address || !sex || !age || !birthday) {
			error = 'Please fill in all required fields';
			return;
		}

		try {
			const editedStudent = await EditStudent(
				student.ID,
				name,
				address,
				sex,
				age,
				birthday,
				birthPlace,
				nameOfChrist,
				motherName,
				kebele,
				houseNo,
				phoneNo,
				email,
				username,
				previousSchool,
				educationLevel,
				workPosition,
				christFatherName,
				location,
				emergencyName,
				emergencyPhoneNo
			);

			if (editedStudent) {
				error = '';
				isOpen = false;
				await invalidateAll();
			} else {
				error = 'Something went wrong';
			}
		} catch (e) {
			error = 'Error: ' + e.message;
		}
	}
</script>

<Dialog.Root open={isOpen} onOpenChange={(open) => (isOpen = open)}>
	<Dialog.Trigger class="text-blue-500 hover:underline cursor-pointer">Edit</Dialog.Trigger>

	<Dialog.Portal>
		<Dialog.Overlay
			class="data-[state=open]:animate-in data-[state=closed]:animate-out fixed inset-0 z-50 bg-black/80"
		/>
		<Dialog.Content
			class="fixed inset-0 z-50 bg-white p-6 overflow-auto text-xs max-w-[95vw] max-h-[90vh] mx-auto my-auto rounded-md shadow-lg"
		>
			<div class="flex justify-between items-center mb-4">
				<h2 class="text-lg font-semibold">Edit Student</h2>
				<Dialog.Close
					class="focus-visible:ring-foreground cursor-pointer focus-visible:ring-offset-background focus-visible:outline-hidden"
				>
					<X class="text-foreground size-5" />
					<span class="sr-only">Close</span>
				</Dialog.Close>
			</div>

			{#if error}
				<p class="text-red-600 mb-4">{error}</p>
			{/if}

			<Separator.Root class="bg-dark mb-6 h-px" />

			<!-- Grid with 3 columns for fields -->
			<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">

				<!-- Each field block -->
				<div>
					<Label.Root for="name" class="text-xs font-semibold">Name *</Label.Root>
					<input
						id="name"
						bind:value={name}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Name"
					/>
				</div>

				<div>
					<Label.Root for="sex" class="text-xs font-semibold">Sex *</Label.Root>
					<select
						id="sex"
						bind:value={sex}
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
					>
						<option value="">Select sex</option>
						<option value="Male">Male</option>
						<option value="Female">Female</option>
					</select>
				</div>

				<div>
					<Label.Root for="age" class="text-xs font-semibold">Age *</Label.Root>
					<input
						id="age"
						bind:value={age}
						type="number"
						min="0"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Age"
					/>
				</div>

				<div>
					<Label.Root for="birthday" class="text-xs font-semibold">Birthday *</Label.Root>
					<input
						id="birthday"
						bind:value={birthday}
						type="date"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
					/>
				</div>

				<div>
					<Label.Root for="address" class="text-xs font-semibold">Address *</Label.Root>
					<input
						id="address"
						bind:value={address}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Address"
					/>
				</div>

				<div>
					<Label.Root for="birthPlace" class="text-xs font-semibold">Birth Place</Label.Root>
					<input
						id="birthPlace"
						bind:value={birthPlace}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Birth Place"
					/>
				</div>

				<div>
					<Label.Root for="nameOfChrist" class="text-xs font-semibold">Name of Christ</Label.Root>
					<input
						id="nameOfChrist"
						bind:value={nameOfChrist}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Name of Christ"
					/>
				</div>

				<div>
					<Label.Root for="motherName" class="text-xs font-semibold">Mother's Name</Label.Root>
					<input
						id="motherName"
						bind:value={motherName}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Mother's Name"
					/>
				</div>

				<div>
					<Label.Root for="kebele" class="text-xs font-semibold">Kebele</Label.Root>
					<input
						id="kebele"
						bind:value={kebele}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Kebele"
					/>
				</div>

				<div>
					<Label.Root for="houseNo" class="text-xs font-semibold">House No</Label.Root>
					<input
						id="houseNo"
						bind:value={houseNo}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="House No"
					/>
				</div>

				<div>
					<Label.Root for="phoneNo" class="text-xs font-semibold">Phone No</Label.Root>
					<input
						id="phoneNo"
						bind:value={phoneNo}
						type="tel"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Phone Number"
					/>
				</div>

				<div>
					<Label.Root for="email" class="text-xs font-semibold">Email</Label.Root>
					<input
						id="email"
						bind:value={email}
						type="email"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Email"
					/>
				</div>

				<div>
					<Label.Root for="username" class="text-xs font-semibold">Username</Label.Root>
					<input
						id="username"
						bind:value={username}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Username"
					/>
				</div>

				<div>
					<Label.Root for="previousSchool" class="text-xs font-semibold">Previous School</Label.Root>
					<input
						id="previousSchool"
						bind:value={previousSchool}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Previous School"
					/>
				</div>

				<div>
					<Label.Root for="educationLevel" class="text-xs font-semibold">Education Level</Label.Root>
					<input
						id="educationLevel"
						bind:value={educationLevel}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Education Level"
					/>
				</div>

				<div>
					<Label.Root for="workPosition" class="text-xs font-semibold">Work Position</Label.Root>
					<input
						id="workPosition"
						bind:value={workPosition}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Work Position"
					/>
				</div>

				<div>
					<Label.Root for="christFatherName" class="text-xs font-semibold">Christ Father Name</Label.Root>
					<input
						id="christFatherName"
						bind:value={christFatherName}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Christ Father Name"
					/>
				</div>

				<div>
					<Label.Root for="location" class="text-xs font-semibold">Location</Label.Root>
					<input
						id="location"
						bind:value={location}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Location"
					/>
				</div>

				<div>
					<Label.Root for="emergencyName" class="text-xs font-semibold">Emergency Contact Name</Label.Root>
					<input
						id="emergencyName"
						bind:value={emergencyName}
						type="text"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Emergency Contact Name"
					/>
				</div>

				<div>
					<Label.Root for="emergencyPhoneNo" class="text-xs font-semibold">Emergency Phone No</Label.Root>
					<input
						id="emergencyPhoneNo"
						bind:value={emergencyPhoneNo}
						type="tel"
						class="shadow border rounded w-full py-1 px-2 text-gray-800 focus:outline-none focus:ring-1 focus:ring-blue-500"
						placeholder="Emergency Phone Number"
					/>
				</div>

			</div>

			<div class="mt-6 flex justify-end gap-4">
				<Dialog.Close class="rounded-md px-4 py-2 border border-gray-300 hover:bg-gray-100">
					Cancel
				</Dialog.Close>

				<button
					class="rounded-md bg-blue-600 text-white px-4 py-2 hover:bg-blue-700"
					on:click|preventDefault={editStudent}
				>
					Save Changes
				</button>
			</div>
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
