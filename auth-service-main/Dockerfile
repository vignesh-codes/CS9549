# Use an official Node.js runtime (Alpine version) as a base image
FROM node:22-alpine

# Set the working directory inside the container
WORKDIR /app

# Install dependencies needed for TypeScript and Node.js native dependencies
RUN apk add --no-cache \
    python3 \
    make \
    g++

# Copy package.json and package-lock.json files to the working directory
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the entire project into the working directory
COPY . .

# Compile TypeScript files
RUN npm run build

# Expose the port your app runs on
EXPOSE 5000

# Start the application
CMD ["npm", "start"]
