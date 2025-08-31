<script lang="ts">
  import { onMount } from 'svelte';
  import { productClient } from '$lib/grpc';
  import type { Product } from '../../types/proto/api/v1/product';
  import Button from '$lib/components/ui/button.svelte';
  import { Pencil, Trash2 } from 'lucide-svelte';

  let products: Product[] = [];
  let loading = true;
  let error = '';

  async function load() {
    try {
      const res = await productClient.listProducts({});
      products = res.products;
    } catch (e: any) {
      error = e?.message ?? String(e);
    } finally {
      loading = false;
    }
  }

  async function remove(id: number) {
    if (!confirm('Delete this product?')) return;
    await productClient.deleteProduct({ id });
    await load();
  }

  onMount(load);
</script>

<section class="space-y-4">
  <div class="flex items-center justify-between">
    <h2 class="font-semibold text-lg">Products</h2>
    <Button href="/products/new">Create</Button>
  </div>

  {#if loading}
    <p>Loading…</p>
  {:else if error}
    <p class="text-red-600">{error}</p>
  {:else}
    <div class="border rounded-md overflow-hidden">
      <table class="w-full text-sm">
        <thead class="bg-gray-50">
          <tr>
            <th class="text-left p-3 w-40">Cover</th>
            <th class="text-left p-3">Name</th>
            <th class="text-left p-3">Description</th>
            <th class="text-left p-3 w-24">Price</th>
            <th class="text-left p-3 w-56">Actions</th>
          </tr>
        </thead>
        <tbody>
          {#each products as p}
            <tr class="border-t">
              <td class="p-3">
                {#if p.cover}
                  <img class="h-20 w-28 object-cover" src={p.cover} alt={p.name} />
                {/if}
              </td>
              <td class="p-3">{p.name}</td>
              <td class="p-3">{p.description}</td>
              <td class="p-3">{p.price}</td>
              <td class="p-3">
                <Button href={`/products/${p.id}/edit`}>
                  <span class="inline-flex items-center gap-2"><Pencil size={16} /> Update</span>
                </Button>
                <Button variant="destructive" class="ml-2" on:click={() => remove(p.id)}>
                  <span class="inline-flex items-center gap-2"><Trash2 size={16} /> Delete</span>
                </Button>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  {/if}
</section>
