<script lang="ts">
	import AddStudent from './add-student.svelte';
	import DeleteStudent from './delete-student.svelte';
	import EditStudent from './edit-student.svelte';
	import ViewStudent from './view-student.svelte';

	let searchQuery = $state('');
	let { students } = $props();

	let filteredStudents = $derived(
		students.filter((student: Record<string, any>) => {
			const studentName = student.Name || '';
			const studentID =
				student.ID !== undefined && student.ID !== null ? student.ID.toString() : '';
			const nameMatch = studentName.toLowerCase().includes(searchQuery.toLowerCase());
			const idMatch = studentID.includes(searchQuery);
			return nameMatch || idMatch;
		})
	);

	function escapeRegExp(string: string): string {
		return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
	}

	function highlightText(text: string | number, query: string): string {
		const textString = String(text || '');
		const queryLower = query.toLowerCase();

		if (!queryLower) {
			return textString;
		}

		const escapedQuery = escapeRegExp(queryLower);
		const regex = new RegExp(`(${escapedQuery})`, 'gi');
		return textString.replace(regex, '<span class="bg-yellow-500">$1</span>');
	}

	// Define the 10 fields you want to print
	const printFields = [
		'Name',
		'Sex',
		'Age',
		'Birthday',
		'Address',
		'PhoneNo',
		'Email',
		'Username',
		'EducationLevel'
	];

	function printStudents() {
		let html = `
			<html>
			<head>
				<title>Print Students</title>
				<style>
					body { font-family: Arial, sans-serif; padding: 1rem; }
					table { border-collapse: collapse; width: 100%; }
					th, td { border: 1px solid #000; padding: 8px; text-align: left; }
					th { background-color: #333; color: white; }
				</style>
			</head>
			<body>
				<h1>ናዝሬት ደ/ታቦር ምስራቅ ፀሀይ ቅዱስ እግዚአብሔር አብ ቤተ-ክርስቲያን</h1>
				<table>
					<thead>
						<tr>
							${printFields.map(field => `<th>${field}</th>`).join('')}
						</tr>
					</thead>
					<tbody>
						${filteredStudents
							.map(
								(								student: { [x: string]: any; }) => `<tr>
									${printFields
										.map(field => `<td>${student[field] ?? ''}</td>`)
										.join('')}
								</tr>`
							)
							.join('')}
					</tbody>
				</table>
			</body>
			</html>
		`;

		const printWindow = window.open('', '', 'width=900,height=700');
		if (printWindow) {
			printWindow.document.write(html);
			printWindow.document.close();
			printWindow.focus();
			printWindow.print();
			// Optionally close after printing:
			// printWindow.close();
		} else {
			alert('Unable to open print window. Please allow pop-ups for this site.');
		}
	}
</script>

<div class="w-screen h-screen p-6 overflow-auto bg-white">
	<header class="mb-6">
		<div class="flex items-center gap-4">
			<img src="/logo.jpg" alt="Company Logo" class="w-20 h-20 object-contain" />
			<div>
				<h1 class="text-xl md:text-2xl font-bold">የናዝሬት ደ/ታቦር ምስራቅ ፀሀይ ቅዱስ እግዚአብሔር አብ ቤተ-ክርስቲያን</h1>
				<p class="text-sm md:text-base">Manage students below.</p>
			</div>
		</div>
	</header>

	<div class="flex items-center justify-between mb-4 flex-wrap gap-2">
		<input
			type="text"
			placeholder="Search students by name or ID..."
			bind:value={searchQuery}
			class="border py-2 px-4 w-full max-w-md text-sm"
		/>

		<div class="flex gap-2">
			<AddStudent />
			<button
	on:click={printStudents}
	class="bg-black text-white font-bold px-4 py-2 rounded hover:bg-gray-800 transition"
	type="button"
>
	Print
</button>
 
		</div>
	</div>

	<div class="overflow-auto max-h-[calc(100vh-250px)]">
		<table class="w-full border-collapse border border-black text-sm">
			<thead class="sticky top-0 bg-black text-white">
				<tr>
					<th class="border-r border-white p-2 text-left">ID</th>
					<th class="border-l border-white p-2 text-left">Name</th>
					<th class="border-l border-white p-2 text-left">Sex</th>
					<th class="border-l border-white p-2 text-left">Actions</th>
				</tr>
			</thead>
			<tbody>
				{#each filteredStudents as student (student.ID)}
					<tr class="hover:bg-gray-100">
						<td class="border border-black p-2">{@html highlightText(student.ID, searchQuery)}</td>
						<td class="border border-black p-2">{@html highlightText(student.Name, searchQuery)}</td>
						<td class="border border-black p-2">{student.Sex}</td>
						<td class="border border-black p-2 space-x-2">
							<ViewStudent {student} />
							<EditStudent {student} />
							<DeleteStudent {student} />
						</td>
					</tr>
				{:else}
					<tr>
						<td colspan="4" class="text-center p-4">No students found.</td>
					</tr>
				{/each}
			</tbody>
		</table>
	</div>
</div>
