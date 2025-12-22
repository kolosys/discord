# Use Neuron CLI to generate the internal REST client

echo "Installing Neuron CLI"
go install github.com/kolosys/neuron/cmd/neuron-cli@v0.6.0
echo "Verifying Neuron CLI installation"
neuron-cli version
echo "Cleaning up old client"

rm -rf ./rest/internal
rm -rf ./models

sleep 1

echo "Generating REST client"
neuron-cli generate --url https://raw.githubusercontent.com/discord/discord-api-spec/refs/heads/main/specs/openapi.json -o rest/internal -p internal -m "*Response:,*Request:Options"
echo "Moving models to /models"
if [ -d "./rest/internal/models" ]; then
	mv ./rest/internal/models models
	echo "Models moved successfully"
else
	echo "Warning: rest/internal/models directory not found"
fi

echo "Copying /rest/internal/client to /rest/internal"
if [ -d "./rest/internal/client" ]; then
	mv ./rest/internal/client/* ./rest/internal/
    rm -d ./rest/internal/client
	echo "Client copied successfully"
else
	echo "Warning: rest/internal/client directory not found"
fi

echo "Updating imports in client files"
find rest/internal -name "*.go" -type f -exec sed -i 's|"models"|"github.com/kolosys/discord/models"|g' {} +

echo "Patching models manually"

# Patch "type Error" to "type ErrorResponse" in error.go
if [ -f ./models/error.go ]; then
	sed -i 's/type Error /type ErrorResponse /' ./models/error.go
fi

# Patch "type Ratelimited" to "type RatelimitedResponse" in ratelimited.go
if [ -f ./models/ratelimited.go ]; then
	sed -i 's/type Ratelimited /type RatelimitedResponse /' ./models/ratelimited.go
fi

echo "Done"