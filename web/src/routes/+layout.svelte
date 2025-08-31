<script lang="ts">
  import '../app.css';
  import favicon from '$lib/assets/favicon.svg';
  let { children } = $props();
</script>

<svelte:head>
  <link rel="icon" href={favicon} />
</svelte:head>

<header class="w-full border-b">
  <div class="mx-auto max-w-5xl px-4 py-3 flex items-center justify-between">
    <a href="/" class="font-semibold text-lg">Atlas</a>
    <nav class="text-sm">
      <a class="hover:underline" href="/products">Products</a>
    </nav>
  </div>
</header>

<main class="mx-auto max-w-5xl px-4 py-6">
  {@render children?.()}
</main>
