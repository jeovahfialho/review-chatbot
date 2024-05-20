# Build stage
FROM node:14-alpine as build-stage

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .

# Build the Vue.js application
RUN npm run build

# Production stage
FROM nginx:alpine

# Copy the built application from the build stage to Nginx's html directory
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose port 80 to allow access to the Nginx server
EXPOSE 80

# Start Nginx in the foreground to keep the container running
CMD ["nginx", "-g", "daemon off;"]
