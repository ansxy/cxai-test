# Use the official Node.js image
FROM node:16-alpine

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy all source files
COPY . .

# Generate Prisma client

RUN npx prisma generate
# Build TypeScript files
RUN npm run build

# Expose port
EXPOSE 3000

# Start the application
CMD ["node", "dist/app.js"]
