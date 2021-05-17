f=index.d.ts

npx grpc-toolkit build-typescript-bindings \
    --protofile ../api/api.proto \
    --destfile $f

cat >> $f <<EOF

export const protoFile: string

export const serviceDefinition: tokencardGrpcToolkit.ServiceDefinition
EOF