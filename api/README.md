# API Definition & Code Generation

The `swagger.yaml` API spec defines the APIs exposed by the service.

Tooling is provided to generate portions of the server code, specifically the models leveraged for reading/writing client requests/responses.

## Model Generation

To automatically generate model content execute the following:

```
cd api
./codegen.sh
```

This will use `swagger-codegen-cli` to generate a go-server implmenetation in `api/generated`.  The `model_*` files generated will be extracted and copied to `pkg/models` for use within the server implementation.

The remainder of the generated code is not used and will be automatically removed by the script.
