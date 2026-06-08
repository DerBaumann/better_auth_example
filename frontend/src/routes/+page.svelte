<script lang="ts">
	interface Todo {
		name: string;
		completed: boolean;
	}

	let todos = $state<Todo[]>([]);
	let error = $state<string | null>(null);

	async function getAllTodos() {
		const response = await fetch('/api/todos');
		const data = await response.json();
		console.log(data);
		if (!response.ok) {
			console.error('Unauthenticated');
			error = data.error;
			return;
		}
		todos = data;
	}
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

{#if error}
	<p class="text-red-500">{error}</p>
{/if}

<p><button onclick={getAllTodos}>Get all Todos</button></p>

<ul>
	{#each todos as todo, i (i)}
		<li>{i} {todo.name} <input type="checkbox" checked={todo.completed} /></li>
	{:else}
		<p>No todos found!</p>
	{/each}
</ul>
