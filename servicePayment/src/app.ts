import fastify from "fastify";
import initConsumer from "./clients/kafka/client";
import { LoadConfig } from "./config/config";
import Routes from "./routes/routes";

const startServer = async () => {
  const server = fastify();
  const cnf = LoadConfig();
  const port = Number(cnf.API_PORT) || 3001;
  const host = String(cnf.API_HOST) || "localhost";

  server.register(Routes, { prefix: "/api" });
  server.get("/ping", async (request, reply) => {
    return { ping: "pong" };
  });

  //temp
  // Start server
  try {
    server.listen({
      port,
      host,
    });
  } catch (err) {
    server.log.error(err);
    process.exit(1);
  }
};

initConsumer().catch((err) => {
  console.error(err);
});
startServer();
