import dotenv from "dotenv";
import path from "path";

type Config = {
  API_PORT: number;
  API_HOST: string;
};

export const LoadConfig = (): Config => {
  const envPath = path.join(__dirname, "../../.env");
  const result = dotenv.config({ path: envPath });
  if (result.error) {
    throw new Error(`Error loading .env file: ${result.error}`);
  }

  const schema: Config = {
    API_PORT: Number(process.env.API_PORT) || 3001,
    API_HOST: String(process.env.API_HOST) || "localhost",
  };

  return schema;
};
