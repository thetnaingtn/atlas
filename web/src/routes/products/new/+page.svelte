<script lang="ts">
  import ProductForm from '$lib/components/ProductForm.svelte';
  import { productClient } from '$lib/grpc';

  async function handleSubmit(e: CustomEvent) {
    await productClient.createProduct(e.detail);
    if (typeof window !== 'undefined') window.location.href = '/products';
  }
</script>

<ProductForm mode="create" on:submit={handleSubmit} />
