<script lang="ts">
  import type { CreateProductRequest, UpdateProductRequest, Product } from "../../types/proto/api/v1/product";
  import { createEventDispatcher } from 'svelte';
  import Button from './ui/button.svelte';
  import Input from './ui/input.svelte';
  import Textarea from './ui/textarea.svelte';

  type Mode = 'create' | 'update';

  let { mode = 'create' as Mode, initial }: { mode?: Mode; initial?: Partial<Product> } = $props();
  const dispatch = createEventDispatcher<{ submit: CreateProductRequest | UpdateProductRequest }>();

  let name = initial?.name ?? '';
  let description = initial?.description ?? '';
  let price: number = initial?.price ?? 0;
  let cover = initial?.cover ?? '';
  let id = initial?.id ?? 0;

  function onSubmit(e: Event) {
    e.preventDefault();
    if (mode === 'create') {
      dispatch('submit', { name, description, price: Number(price), cover });
    } else {
      dispatch('submit', { id: Number(id), name, description, price: Number(price), cover });
    }
  }
</script>

<form class="max-w-md mx-auto border rounded-md p-6 space-y-4" on:submit={onSubmit}>
  <h2 class="font-semibold text-lg">{mode === 'create' ? 'Create Product' : 'Update Product'}</h2>

  {#if mode === 'update'}
    <input type="hidden" bind:value={id} />
  {/if}

  <div class="space-y-1">
    <label class="text-sm font-medium">Name</label>
    <Input required bind:value={name} />
  </div>

  <div class="space-y-1">
    <label class="text-sm font-medium">Description</label>
    <Textarea class="h-28" bind:value={description} />
  </div>

  <div class="space-y-1">
    <label class="text-sm font-medium">Price</label>
    <Input type="number" step="0.01" bind:value={price} />
  </div>

  <div class="space-y-1">
    <label class="text-sm font-medium">Cover URL</label>
    <Input bind:value={cover} />
  </div>

  <div class="flex gap-2 pt-2">
    <Button type="submit">Save</Button>
    <Button href="/products" variant="outline">Cancel</Button>
  </div>
</form>
