<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { Chart, type ChartConfiguration, registerables } from 'chart.js';
	import * as types from '../../types/tasks';

	export let taskData: types.TaskInformation[];
	let chartType = 'pie';

	Chart.register(...registerables);

	let chartTypes = ['pie', 'line', 'bar', 'polarArea'];

	function changeChart(type: string) {
		chartType = type;
		updateChart();
	}

	function timeToMinutes(timeStr: string) {
		const [hours, minutes] = timeStr.split(':').map(Number);
		return hours * 60 + minutes;
	}

	const aggregatedData = taskData.reduce<types.AggregatedTime>(
		(acc: { [x: string]: number }, item: types.TaskInformation) => {
			const type = item.type;
			const timeInMinutes = timeToMinutes(item.time);

			if (!acc[type]) {
				acc[type] = 0;
			}
			acc[type] += timeInMinutes;

			return acc;
		},
		{}
	);

	const labels = Object.keys(aggregatedData);
	const data = Object.values(aggregatedData);

	let chartCanvas: HTMLCanvasElement;

	const chartConfig: ChartConfiguration = {
		type: chartType,
		data: {
			labels: labels,
			datasets: [
				{
					label: 'Time by Type (in minutes)',
					data: data,
					backgroundColor: ['#FF6384', '#36A2EB', '#FFCE56', '#4BC0C0', '#9966FF']
				}
			]
		},
		options: {
			responsive: true,
			plugins: {
				legend: {
					display: true
				},
				tooltip: {
					callbacks: {
						label: (tooltipItem) => {
							const label = tooltipItem.label || '';
							const value = tooltipItem.raw || 0;
							return `${label}: ${value} minutes`;
						}
					}
				}
			}
		}
	};

	let chart: Chart;
	onMount(() => {
		console.log('we have mounted');
		chart = new Chart(chartCanvas, chartConfig);
	});

	function updateChart() {
		if (chart) {
			chart.destroy();
		}
		chartConfig.type = chartType;

		chart = new Chart(chartCanvas, chartConfig);
	}

	onDestroy(() => {
		if (chart) {
			chart.destroy();
		}
	});
</script>

{#each chartTypes as chart}
	<button onclick={() => changeChart(chart)}>{chart}</button>
{/each}

<div>
	<canvas bind:this={chartCanvas}></canvas>
</div>

<style>
	canvas {
		max-width: 100%;
		height: auto;
	}

	div {
		width: 400px;
		height: 400px;
	}
</style>
