import { createChannel, createClient } from 'nice-grpc-web';
import { ProductServiceDefinition } from '../types/proto/api/v1/product';

// Use a relative base URL so requests always go through the dev proxy
// and production origin without hardcoding host.
const baseUrl = '';

export const channel = createChannel(baseUrl);

export const productClient = createClient(ProductServiceDefinition, channel);
