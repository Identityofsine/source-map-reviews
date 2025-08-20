
echo "Building and pushing Docker image..."
docker build -t ghcr.io/identityofsine/de-archive-api:latest .
if [ $? -ne 0 ]; then
    echo "Docker build failed. Exiting."
    exit 1
fi

echo "Pushing Docker image to GitHub Container Registry..."
docker push ghcr.io/identityofsine/de-archive-api:latest
if [ $? -ne 0 ]; then
    echo "Docker push failed. Exiting."
    exit 1
fi
