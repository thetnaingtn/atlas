import { useState } from 'react'
import { Input } from './ui/input'
import { Label } from './ui/label'
import { Button } from './ui/button'
import type { Product } from '../gen/api/v1/product_pb'

interface Props {
  initial?: Product
  onSubmit: (product: Product) => void
}

export function ProductForm({ initial, onSubmit }: Props) {
  const [name, setName] = useState(initial?.name ?? '')
  const [price, setPrice] = useState(initial?.price ?? 0)

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    onSubmit({ id: initial?.id ?? '', name, price })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="name">Name</Label>
        <Input id="name" value={name} onChange={e => setName(e.target.value)} />
      </div>
      <div className="space-y-2">
        <Label htmlFor="price">Price</Label>
        <Input id="price" type="number" value={price} onChange={e => setPrice(parseFloat(e.target.value))} />
      </div>
      <Button type="submit">Save</Button>
    </form>
  )
}
