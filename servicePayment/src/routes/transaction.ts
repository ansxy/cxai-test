import { FastifyInstance } from "fastify";
async function TransactionRoutes(fastify: FastifyInstance) {
  fastify.get("/", async (request, reply) => {
    return { message: "Transaction Created" };
  });
}

export default TransactionRoutes;
