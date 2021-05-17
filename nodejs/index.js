'use strict'
const path = require('path')
const { loadProtoFile } = require('@tokencard/grpc-toolkit')
const { default: monolithLoadingOptions } = require('@tokencard/monolith-toolkit')

const protoFile = path.join(__dirname, '../api/api.proto')
exports.protoFile = protoFile

const {
  services: {
    'api.API': serviceDefinition,
  },
} = loadProtoFile({ protoFile, ...monolithLoadingOptions })
exports.serviceDefinition = serviceDefinition