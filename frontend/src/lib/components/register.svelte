<script lang="ts">
	import { Register } from '$lib/wailsjs/go/main/App';
	import { invalidateAll } from '$app/navigation';

	let username = $state('');
	let password = $state('');
	let error = $state('');

	async function register() {
		if (username === '' || password === '') {
			error = 'Please fill in all fields';
			return;
		} else {
			const newAdmin = await Register(username, password);
			if (newAdmin) {
				invalidateAll();
				error = '';
			} else {
				error = 'Something went wrong';
			}
		}
	}
</script>

<div class="h-full flex items-center justify-center flex-col bg-background">
	<div class="p-6 flex flex-col">
		<h2 class="text-2xl font-semibold text-center mb-4 text-gray-800">Register</h2>

		<div class="mb-4 min-w-md">
			<label for="username" class="block text-gray-700 text-sm font-bold mb-2">Username</label>
			<input
				type="text"
				id="username"
				bind:value={username}
				class="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				placeholder="Enter your username"
			/>
		</div>
		<div class="mb-4 min-w-md">
			<label for="password" class="block text-gray-700 text-sm font-bold mb-2">Password</label>
			<input
				type="password"
				id="password"
				bind:value={password}
				class="shadow appearance-none border w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
				placeholder="Enter your password"
			/>
		</div>
		<div class="flex items-center justify-between min-w-md">
			<button
				onclick={register}
				class="bg-dark cursor-pointer w-full text-background hover:bg-dark/95 inline-flex h-12 items-center justify-center px-4 font-semibold"
			>
				Register
			</button>
		</div>
		{#if error != ''}
			<p class="text-red-500 text-center">{error}</p>
		{/if}
	</div>
</div>
