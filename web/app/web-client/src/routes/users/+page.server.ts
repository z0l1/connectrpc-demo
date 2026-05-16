import { userClient } from "$lib/client/client";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (evnt) => {
	const username = evnt.url.searchParams.get("username");

	if (!username) {
		return { data: undefined };
	}

	const resp = await userClient.getUserByUsername({ username });

	return { data: resp };
};
