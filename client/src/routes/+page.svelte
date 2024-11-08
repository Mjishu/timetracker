<!-- So I need to make 4 different functions? one for each timePeriod, and then call the function that relates to teh currentPeriod state and return true or false based on the dates -->

<script lang="ts">
	import { onMount } from 'svelte';
	import { parse, parseISO } from 'date-fns';
	import { isModalOpen } from '../helper/shared';
	import * as timeHelper from '../helper/time';
	import * as taskType from '../types/tasks';
	import AddTask from '../components/tasks/addTask.svelte';
	import Chart from '../components/tasks/chart.svelte';
	import { removeTask } from '../helper/shared';

	let currentPeriod = $state({ period: 'week', functionToCall: timeHelper.isWeek });
	let allTasks = $state<taskType.TaskInformation[] | undefined>(undefined); //sort this based on current period before sending it to showtasks

	onMount(async () => {
		fetch('/api/tasks/all')
			.then((res) => res.json())
			.then((data) => (allTasks = data)) //set tasks to this
			.catch((err) => console.error(`There was an error trying to get tasks. ${err}`));
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
	}
</script>

<button class="add-task-button" onclick={handleOpenModal}>Add Task</button>
<div class="time-period-div">
	{#each timePeriods as period}
		<button class="time-period-button" onclick={() => callTimePeriod(period)}>{period}</button>
	{/each}
</div>
<AddTask />

{#if allTasks != undefined && allTasks.length > 0}
	<div class="task-holder">
		{#each allTasks as task}
			{#if currentPeriod.functionToCall(new Date(), parseISO(task.date))}
				<div class="task-information">
					<h3>{task.title}</h3>
					<p>{task.time}</p>
					<p>{task.date}</p>
					<p>{task.type}</p>
				</div>
				<button onclick={() => removeTask(task.id)} class="delete-task">delete</button>
			{/if}
		{/each}
	</div>
	<Chart taskData={allTasks} />
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
