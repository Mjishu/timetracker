<script lang="ts">
	const information = $state({ username: '', password: '' });

	async function handleSubmit(e) {
		e.preventDefault();
		try {
			const response = await fetch(`/api/users/${information.username}`);
			if (!response.ok) {
				throw new Error(`http error, status: ${response.status}`);
			}
			const data = await response.json();

			console.log(data);
		} catch (err) {
			console.error(`error fetching data: ${err}`);
		}
	}
</script>

<h2>Log In</h2>
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
