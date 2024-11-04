<script>
	import AddTask from '../components/tasks/addTask.svelte';
	import { isModalOpen } from '../shared';
	import ShowTasks from '../components/tasks/showTasks.ts.svelte';

	function handleOpenModal() {
		isModalOpen.set(true);
	}
</script>

<button class="add-task-button" onclick={handleOpenModal}>Add Task</button>
<AddTask />

<ShowTasks />
