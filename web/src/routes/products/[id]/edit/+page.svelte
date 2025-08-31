<script lang="ts">
  import { onMount } from 'svelte';
  import ProductForm from '$lib/components/ProductForm.svelte';
  import { productClient } from '$lib/grpc';
  import type { Product } from '../../../types/proto/api/v1/product';

  let product: Partial<Product> | undefined;
  let loading = true;

  onMount(async () => {
    const segments = typeof window !== 'undefined' ? window.location.pathname.split('/') : [];
    const id = Number(segments[2]);
    const res = await productClient.listProducts({});
    product = res.products.find((p) => p.id === id);
    loading = false;
  });

  async function handleSubmit(e: CustomEvent) {
    await productClient.updateProduct(e.detail);
    if (typeof window !== 'undefined') window.location.href = '/products';
  }
</script>

{#if loading}
  <p>Loading…</p>
{:else if !product}
  <p>Product not found</p>
{:else}
  <ProductForm mode="update" initial={product} on:submit={handleSubmit} />
{/if}
