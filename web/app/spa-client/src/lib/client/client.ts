import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService } from "./gen/grpc/user/v1/user_pb";

const userTransport = createConnectTransport({
	baseUrl: "http://localhost:8080",
	// Not needed. Web browsers use HTTP/2 automatically.
	// httpVersion: "1.1"
});

const userClient = createClient(UserService, userTransport);

export const services = {
	user: userClient,
};
