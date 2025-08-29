import { createChannel, createClient } from 'nice-grpc-web'
import { ProductService } from '../gen/api/v1/product_pb'

const channel = createChannel('http://localhost:8080')

export const productClient = createClient(ProductService, channel)
