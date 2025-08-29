import { useEffect, useState } from 'react'
import { Button } from './ui/button'
import type { Product } from '../gen/api/v1/product_pb'
import { productClient } from '../lib/grpc'

interface Props {
  onSelect: (product: Product) => void
  onCreate: () => void
}

export function ProductList({ onSelect, onCreate }: Props) {
  const [products, setProducts] = useState<Product[]>([])

  useEffect(() => {
    async function fetchProducts() {
      try {
        const res = await productClient.listProducts({})
        setProducts(res.products)
      } catch (e) {
        console.error(e)
      }
    }
    fetchProducts()
  }, [])

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h2 className="text-xl font-bold">Products</h2>
        <Button onClick={onCreate}>Create</Button>
      </div>
      <table className="w-full border">
        <thead>
          <tr className="bg-gray-100 text-left">
            <th className="p-2 border">Name</th>
            <th className="p-2 border">Price</th>
          </tr>
        </thead>
        <tbody>
          {products.map(p => (
            <tr key={p.id} className="hover:bg-gray-50 cursor-pointer" onClick={() => onSelect(p)}>
              <td className="p-2 border">{p.name}</td>
              <td className="p-2 border">{p.price}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
