<script lang="ts">
	import { onMount } from 'svelte';
	import Chart from './chart.svelte';
	import * as types from '../../types/tasks';

	let tasks: types.TaskInformation[] | undefined = $state();
	let chartType = $state('pie');

	onMount(async () => {
		fetch('/api/tasks/all')
			.then((res) => res.json())
			.then((data) => (tasks = data)) //set tasks to this
			.catch((err) => console.error(`There was an error trying to get tasks. ${err}`));
	});
</script>

<!-- todo -->
<!-- get all from backend and map over each one and show it on the page with for each -->

{#if tasks != undefined}
	{#each tasks as task}
		<h3>{task.title}</h3>
		<p>{task.time}</p>
		<p>{task.date}</p>
		<p>{task.type}</p>
	{/each}

	<Chart taskData={tasks} />
{/if}

<style>
</style>
