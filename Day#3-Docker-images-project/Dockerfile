# Use the official Nginx image with Alpine
FROM nginx:alpine

# Add metadata to document the image
LABEL maintainer="you@example.com" \
      version="1.0" \
      description="A lightweight Nginx container for static website hosting"

# Copy your website files
COPY index.html /usr/share/nginx/html

# Expose the default HTTP port
EXPOSE 80

# Start Nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
