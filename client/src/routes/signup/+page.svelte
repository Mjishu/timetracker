<script lang="ts">
	const information = $state({ username: '', password: '', email: '', confirmPassword: '' });

	async function handleSubmit(e) {
		e.preventDefault();
		try {
			const fetchParams = {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify(information)
			};

			const response = await fetch('/api/users', fetchParams);

			if (!response.ok) {
				throw new Error(`HTTP error; status: ${response.status}`);
			}

			const data = response.json();
			console.log('Data recieved is ', data);
		} catch (err) {
			console.error(`error fetching data: ${err}`);
		}
	}
</script>

<h2>Sign Up</h2>
<form onsubmit={handleSubmit}>
	<div class="authentication-div">
		<label for="username">Username</label>
		<input type="text" name="username" id="" bind:value={information.username} />
	</div>
	<div class="authentication-div">
		<label for="email">Email</label>
		<input type="email" name="email" id="" bind:value={information.email} />
	</div>
	<div class="authentication-div">
		<label for="password">Password</label>
		<input type="password" name="password" id="" bind:value={information.password} />
	</div>
	<div class="authentication-div">
		<label for="confirmPassword">Confirm Password</label>
		<input type="password" name="Confirmpassword" id="" bind:value={information.confirmPassword} />
	</div>
	<button>Log In</button>
	<a href="/">Home</a>
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
