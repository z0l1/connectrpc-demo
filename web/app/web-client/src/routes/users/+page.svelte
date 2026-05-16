<script lang="ts">
	import { bigintReplacer } from "$lib/utils/json-replacer";
	import type { PageServerData } from "./$types";

	const { data }: { data: PageServerData } = $props();
</script>

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

<pre class="p-2 bg-blue-200">{JSON.stringify(data, bigintReplacer, 2)}</pre>
