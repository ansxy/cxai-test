import { PrismaClient } from "@prisma/client";

export const prisma = new PrismaClient();

try {
  prisma.$connect();

  console.log("Connected to the database");
} catch (error) {
  console.log("Error", error);
}

module.exports = { prisma };
