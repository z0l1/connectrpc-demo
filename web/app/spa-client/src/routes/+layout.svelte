<script lang="ts">
	import "./layout.css";
	import favicon from "$lib/assets/favicon.svg";

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<nav class="bg-amber-200 py-2 px-4 mb-2 flex justify-between">
	<a href="/">HOME</a>
</nav>

<div class="container mx-auto">
	{@render children()}
</div>
