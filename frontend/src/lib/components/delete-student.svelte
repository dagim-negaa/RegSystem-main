<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { DeleteStudent } from '$lib/wailsjs/go/main/App';
	import { AlertDialog } from 'bits-ui';

	let { student } = $props();

	let error = $state('');

	let isOpen = $state(false);

	async function deleteStudent() {
		const deletedStudent = await DeleteStudent(student.ID);
		if (deletedStudent) {
			error = '';
			isOpen = false;
			invalidateAll();
		} else {
			error = 'Something went wrong';
		}
	}
</script>

<AlertDialog.Root open={isOpen} onOpenChange={(open) => (isOpen = open)}>
	<AlertDialog.Trigger class="text-red-500 hover:underline cursor-pointer">
		Delete
	</AlertDialog.Trigger>
	<AlertDialog.Portal>
		<AlertDialog.Overlay
			class="data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 fixed inset-0 z-50 bg-black/80"
		/>
		<AlertDialog.Content
			class="bg-background shadow-popover data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 outline-hidden fixed left-[50%] top-[50%] z-50 grid w-full max-w-[calc(100%-2rem)] translate-x-[-50%] translate-y-[-50%] gap-4 border p-7 sm:max-w-lg md:w-full "
		>
			<div class="flex flex-col gap-4 pb-6">
				<AlertDialog.Title class="text-lg font-semibold tracking-tight">
					Are you sure you want to delete this student?
				</AlertDialog.Title>
				<AlertDialog.Description class="text-foreground-alt text-sm">
					This action cannot be undone. Please confirm that you want to delete this student.
				</AlertDialog.Description>
			</div>
			<div class="flex w-full items-center justify-center gap-2">
				<AlertDialog.Cancel
					class="bg-dark w-full cursor-pointer text-background hover:bg-dark/95 inline-flex items-center justify-center px-4 py-2 font-semibold"
				>
					Cancel
				</AlertDialog.Cancel>
				<AlertDialog.Action
					onclick={deleteStudent}
					class="bg-destructive w-full cursor-pointer text-background hover:bg-destructive/95 inline-flex items-center justify-center px-4 py-2 font-semibold"
				>
					Delete
				</AlertDialog.Action>
			</div>
		</AlertDialog.Content>
	</AlertDialog.Portal>
</AlertDialog.Root>
