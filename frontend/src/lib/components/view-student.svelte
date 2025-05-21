<script lang="ts">
	import { X } from '@lucide/svelte';
	import { Dialog, Label, Separator } from 'bits-ui';

	let { student } = $props();

	let isOpen = $state(false);
</script>

<Dialog.Root open={isOpen} onOpenChange={(open) => (isOpen = open)}>
	<Dialog.Trigger class="text-green-500 hover:underline cursor-pointer">View</Dialog.Trigger>

	<Dialog.Portal>
		<Dialog.Overlay
			class="data-[state=open]:animate-in data-[state=closed]:animate-out fixed inset-0 z-50 bg-black/80"
		/>
		<Dialog.Content
			class="fixed inset-0 z-50 bg-white p-6 overflow-auto text-sm"
		>
			<div class="flex justify-between items-center mb-4">
				<h2 class="text-xl font-semibold">View Student</h2>
				<Dialog.Close class="cursor-pointer">
					<X class="text-gray-600 hover:text-red-600 size-5" />
				</Dialog.Close>
			</div>

			<Separator.Root class="bg-gray-300 h-px mb-4" />

			<!-- Grid Layout for Fields -->
			<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
				{#each [
					['ID', 'ID', student.ID],
					['Name', 'Name', student.Name],
					['Sex', 'Sex', student.Sex],
					['Age', 'Age', student.Age],
					['Birthday', 'Birthday', student.Birthday],
					['Address', 'Address', student.Address],
					['BirthPlace', 'Birth Place', student.BirthPlace],
					['NameOfChrist', 'Name of Christ', student.NameOfChrist],
					['MotherName', "Mother's Name", student.MotherName],
					['Kebele', 'Kebele', student.Kebele],
					['HouseNo', 'House No', student.HouseNo],
					['PhoneNo', 'Phone No', student.PhoneNo],
					['Email', 'Email', student.Email],
					['Username', 'Username', student.Username],
					['PriviesSchool', 'Previous School', student.PriviesSchool],
					['EducationLevel', 'Education Level', student.EducationLevel],
					['WorkPosition', 'Work Position', student.WorkPosition],
					['ChristFatherName', 'Christ Father Name', student.ChristFatherName],
					['Location', 'Location', student.Location],
					['EmergencyName', 'Emergency Contact Name', student.EmergencyName],
					['EmergencyPhoneNo', 'Emergency Phone No', student.EmergencyPhoneNo]
				] as [id, label, value]}
					<div class="flex flex-col gap-1">
						<Label.Root for={id} class="text-xs font-medium">{label}</Label.Root>
						<input
							id={id}
							readonly
							value={value}
							class="border p-1 bg-gray-100 rounded text-xs"
						/>
					</div>
				{/each}
			</div>
			
		</Dialog.Content>
	</Dialog.Portal>
</Dialog.Root>
