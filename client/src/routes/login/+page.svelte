<script lang="ts">
	const information = $state({ username: '', password: '' });

	function handleSubmit(e) {
		e.preventDefault();
		console.log($state.snapshot(information));
	}
</script>

<h2>Log In!</h2>
<form onsubmit={handleSubmit}>
	<div class="authentication-div">
		<label for="username">Username</label>
		<input type="text" name="username" id="" bind:value={information.username} />
	</div>
	<div class="authentication-div">
		<label for="password">Password</label>
		<input type="password" name="password" id="" bind:value={information.password} />
	</div>
	<button>Log In</button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
	}

	button {
		width: fit-content;
	}
</style>
