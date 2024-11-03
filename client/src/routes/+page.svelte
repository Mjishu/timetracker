<script>
	import AddTask from '../components/addTask.svelte';
	import { isModalOpen } from '../shared';

	function handleOpenModal() {
		isModalOpen.set(true);
	}
</script>

<button class="add-task-button" onclick={handleOpenModal}>Add Task</button>
<AddTask />
