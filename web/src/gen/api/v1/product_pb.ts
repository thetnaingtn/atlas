import type { ServiceType } from '@bufbuild/protobuf'

export interface Product {
  id: string
  name: string
  price: number
}

export interface GetProductRequest { id: string }
export interface ListProductsRequest {}
export interface ListProductsResponse { products: Product[] }
export interface CreateProductRequest { product?: Product }
export interface UpdateProductRequest { product?: Product }
export interface DeleteProductRequest { id: string }

export const ProductService = {
  typeName: 'api.v1.ProductService',
  methods: {
    getProduct: { name: 'GetProduct', I: {} as GetProductRequest, O: {} as Product },
    listProducts: { name: 'ListProducts', I: {} as ListProductsRequest, O: {} as ListProductsResponse },
    createProduct: { name: 'CreateProduct', I: {} as CreateProductRequest, O: {} as Product },
    updateProduct: { name: 'UpdateProduct', I: {} as UpdateProductRequest, O: {} as Product },
    deleteProduct: { name: 'DeleteProduct', I: {} as DeleteProductRequest, O: {} },
  },
} as unknown as ServiceType
