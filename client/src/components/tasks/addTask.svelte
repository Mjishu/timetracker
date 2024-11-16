<script>
	import { isModalOpen, addTaskDatabase } from '../../helper/shared';

	let newTask = $state({
		id: '2',
		title: '',
		time: '',
		date: '',
		type: ''
	});

	const types = [
		{
			id: 1,
			text: 'Reading'
		},
		{
			id: 2,
			text: 'Exercise'
		},
		{
			id: 3,
			text: 'coding'
		},
		{
			id: 4,
			text: 'Development'
		}
	];

	/**
	 * @param {{ preventDefault: () => void; }} e
	 */
	async function handleSubmit(e) {
		e.preventDefault();
		console.log($state.snapshot(newTask));
		isModalOpen.set(false);
		await addTaskDatabase(newTask);
	}
</script>

{#if $isModalOpen}
	<div class="modal">
		<form class="modal-content" onsubmit={handleSubmit}>
			<label for="title">Title</label>
			<input type="text" bind:value={newTask.title} id="title" placeholder="walking the dog" />

			<label for="task">Time</label>
			<input type="text" bind:value={newTask.time} id="time" placeholder="00:00 (hours:minutes)" />

			<label for="date">Date</label>
			<input type="Date" bind:value={newTask.date} id="date" />

			<select bind:value={newTask.type}>
				{#each types as type}
					<option value={type.text}>{type.text}</option>
				{/each}
			</select>

			<button onclick={() => isModalOpen.set(false)}>cancel</button>
			<button type="submit">Add task</button>
		</form>
	</div>
{/if}

<style>
	.modal {
		display: block;
		align-items: center;
		position: fixed;
		z-index: 1;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		overflow: auto;
		background-color: rgb(0, 0, 0);
		background-color: rgba(0, 0, 0, 0.4);
	}

	.modal-content {
		display: flex;
		flex-direction: column;
		align-items: center;
		background: #fefefe;
		margin: 15% auto;
		padding: 20px;
		border: 1px solid #888;
		width: 25%;
		height: 20rem;
		border-radius: 10px;
		border: none;
	}

	.modal-content input,
	.modal-content button {
		width: fit-content;
	}
</style>
