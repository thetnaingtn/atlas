import { useState } from 'react'
import { ProductList } from './components/product-list'
import { ProductDetail } from './components/product-detail'
import { ProductForm } from './components/product-form'
import type { Product } from './gen/api/v1/product_pb'
import { productClient } from './grpc'

export default function App() {
  const [view, setView] = useState<'list' | 'detail' | 'create' | 'edit'>('list')
  const [selected, setSelected] = useState<Product | null>(null)

  const handleCreate = async (p: Product) => {
    await productClient.createProduct({ product: p })
    setView('list')
  }

  const handleUpdate = async (p: Product) => {
    await productClient.updateProduct({ product: p })
    setSelected(p)
    setView('detail')
  }

  return (
    <div className="p-4">
      {view === 'list' && (
        <ProductList
          onSelect={(p) => {
            setSelected(p)
            setView('detail')
          }}
          onCreate={() => setView('create')}
        />
      )}
      {view === 'detail' && selected && (
        <ProductDetail
          product={selected}
          onEdit={() => setView('edit')}
          onBack={() => setView('list')}
        />
      )}
      {view === 'create' && (
        <ProductForm
          onSubmit={async (p) => {
            await handleCreate(p)
          }}
        />
      )}
      {view === 'edit' && selected && (
        <ProductForm
          initial={selected}
          onSubmit={async (p) => {
            await handleUpdate(p)
          }}
        />
      )}
    </div>
  )
}
