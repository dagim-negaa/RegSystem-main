<script lang="ts">
	import Home from '$lib/components/home.svelte';
	import Login from '$lib/components/login.svelte';
	import Register from '$lib/components/register.svelte';
	import { appState } from '../state.svelte.js';

	let { data } = $props();
</script>

{#if data.admins.length > 0}
	{#if appState.isLoggedIn}
		<Home students={data.students} />
	{:else}
		<Login user={{ username: data.admins[0].Username, password: data.admins[0].Password }} />
	{/if}
{:else}
	<Register />
{/if}
