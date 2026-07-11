import { userClient } from "$lib/client/client";
import type { GetUserByUsernameResponse } from "$lib/client/gen/grpc/user/v1/user_pb";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async (evnt) => {
	const username = evnt.url.searchParams.get("username");

	if (!username) {
		return { data: undefined };
	}

	let resp: GetUserByUsernameResponse | undefined = undefined;
	try {
		resp = await userClient.getUserByUsername({ username });
	} catch (err) {
		console.log("rip");
		console.log(err);
	}

	if (!resp) {
		return { data: undefined };
	}

	return { data: resp };
};
