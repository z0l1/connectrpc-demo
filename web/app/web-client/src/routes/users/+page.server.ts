import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = (evnt) => {
	const username = evnt.url.searchParams.get("username");
	return { username };
};
