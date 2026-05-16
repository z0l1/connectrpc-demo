<script lang="ts">
	import type { PageServerData } from "./$types";

	const { data }: { data: PageServerData } = $props();
</script>

{JSON.stringify(data)}

<form method="GET" class="mt-10 p-4 bg-blue-100 flex flex-col gap-4">
	<p>search for a user</p>

	<div>
		<div>
			<label for="">username</label>
		</div>
		<div>
			<input
				autofocus
				id="username"
				name="username"
				type="text"
				minlength="4"
			/>
		</div>
	</div>

	<div>
		<button type="submit">search</button>
	</div>
</form>
