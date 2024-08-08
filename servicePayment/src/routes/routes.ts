import { FastifyInstance } from "fastify";
import TransactionRoutes from "./transaction";

async function Routes(fastify: FastifyInstance) {
  fastify.register(TransactionRoutes, { prefix: "/transaction" });
}

export default Routes;
