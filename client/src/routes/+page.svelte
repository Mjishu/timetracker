<!-- So I need to make 4 different functions? one for each timePeriod, and then call the function that relates to teh currentPeriod state and return true or false based on the dates -->

<script lang="ts">
	import { onMount } from 'svelte';
	import { parse, parseISO } from 'date-fns';
	import { isModalOpen, getAllTasks } from '../helper/shared';
	import * as timeHelper from '../helper/time';
	import * as taskType from '../types/tasks';
	import AddTask from '../components/tasks/addTask.svelte';
	import Chart from '../components/tasks/chart.svelte';
	import { removeTask } from '../helper/shared';

	//periods
	let currentPeriod = $state({ period: 'week', functionToCall: timeHelper.isWeek });
	let allTasks = $state<taskType.TaskInformation[] | undefined>(undefined); //sort this based on current period before sending it to showtasks
	let filteredTasks = $state<taskType.TaskInformation[] | undefined>(undefined);

	onMount(async () => {
		//might not be the way, probably need to do $state(getalltasks) outsided
		const data = await getAllTasks();
		allTasks = data;
		updateFilteredTasks();
	});

	function handleOpenModal() {
		isModalOpen.set(true);
	}

	let timePeriods = ['Day', 'Week', 'Month', 'Year'];

	function callTimePeriod(period: string): void {
		//I could maybe use a store so that I can keep the call up here and have it update down the treeS?
		period = period.toLowerCase();
		if (period == currentPeriod.period) return;

		if (period == 'day') {
			currentPeriod.period = period;
			currentPeriod.functionToCall = timeHelper.isDay;
		} else if (period == 'week') {
			currentPeriod.period = period;
			currentPeriod.functionToCall = timeHelper.isWeek;
		} else if (period == 'month') {
			currentPeriod.period = period;
			currentPeriod.functionToCall = timeHelper.isMonth;
		} else if (period == 'year') {
			currentPeriod.period = period;
			currentPeriod.functionToCall = timeHelper.isYear;
		}
		updateFilteredTasks();
	}

	function updateFilteredTasks() {
		filteredTasks = allTasks?.filter((task) =>
			currentPeriod.functionToCall(new Date(), parseISO(task.date))
		);
	}
</script>

<button class="add-task-button" onclick={handleOpenModal}>Add Task</button>
<div class="time-period-div">
	{#each timePeriods as period}
		<button class="time-period-button" onclick={() => callTimePeriod(period)}>{period}</button>
	{/each}
</div>
<AddTask />

{#if filteredTasks && filteredTasks.length > 0}
	<div class="task-holder">
		{#each filteredTasks as task}
			<div class="task-information">
				<h3>{task.title}</h3>
				<p>{task.time}</p>
				<p>{task.date}</p>
				<p>{task.type}</p>
			</div>
			<button onclick={() => task.id !== undefined && removeTask(task.id)} class="delete-task"
				>delete</button
			>
		{/each}
	</div>
	<Chart {filteredTasks} />
{/if}

<style>
	.time-period-div {
		display: flex;
		gap: 1rem;
	}
	.time-period-button {
		width: 6rem;
		height: 2rem;
		border: none;
		border-bottom: 2px solid black;
		background: transparent;
	}
	.time-period-button:hover {
		border: 1px solid black;
		border-radius: 5px;
		background-color: hsla(0, 0, 0, 0.2);
	}
	.task-holder {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		width: 25rem;
	}
	.delete-task {
		width: 3rem;
		height: 1.5rem;
	}
</style>
