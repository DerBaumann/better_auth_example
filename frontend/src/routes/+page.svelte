<script lang="ts">
	interface Todo {
		name: string;
		completed: boolean;
	}

	let todos = $state<Todo[]>([]);
	let error = $state<string | null>(null);
	let name = $state('');

	async function getAllTodos() {
		const response = await fetch('/api/todos');
		const data = await response.json();
		if (!response.ok) {
			error = data.error;
			return;
		}
		todos = data;
	}

	async function saveTodo() {
		const body = JSON.stringify({ name });
		const response = await fetch('/api/todos', {
			method: 'POST',
			body: body,
			headers: { 'Content-Type': 'application/json' }
		});

		const data = await response.json();

		if (!response.ok) {
			error = data.error;
			return;
		}

		todos = [...todos, data];
	}
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

{#if error}
	<p class="text-red-500">{error}</p>
{/if}

<div><button onclick={getAllTodos}>Get all Todos</button></div>

<div><input type="text" bind:value={name} /> <button onclick={saveTodo}>Save</button></div>

<ul>
	{#each todos as todo, i (i)}
		<li>{i} {todo.name} <input type="checkbox" checked={todo.completed} /></li>
	{:else}
		<p>No todos found!</p>
	{/each}
</ul>
