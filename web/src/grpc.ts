import { createChannel, createClientFactory, FetchTransport } from 'nice-grpc-web'
import { ProductServiceDefinition } from './types/proto/api/v1/product'

const channel = createChannel(
    window.location.origin,
    FetchTransport({
        credentials:"include"
    })
)

const clientFactory = createClientFactory()

export const productClient = clientFactory.create(ProductServiceDefinition, channel)
