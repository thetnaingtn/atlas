<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  const props = $props();
  let { variant = 'default', size = 'default', type = 'button', href = '', disabled = false, class: className = '' } = props;
  const dispatch = createEventDispatcher<{ click: MouseEvent }>();

  const base = 'inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 disabled:opacity-50 disabled:pointer-events-none';
  const variants: Record<string, string> = {
    default: 'bg-black text-white hover:bg-black/90',
    destructive: 'bg-red-600 text-white hover:bg-red-600/90',
    outline: 'border',
    ghost: 'hover:bg-gray-100',
    link: 'underline-offset-4 hover:underline',
  };
  const sizes: Record<string, string> = {
    default: 'h-9 px-3 py-2',
    sm: 'h-8 rounded-md px-2',
    lg: 'h-10 rounded-md px-4',
  };
  const cls = `${base} ${variants[variant] ?? variants.default} ${sizes[size] ?? sizes.default} ${className}`;
</script>

{#if href}
  <a href={href} class={cls} on:click={(e)=>dispatch('click', e)}><slot /></a>
{:else}
  <button type={type} class={cls} disabled={disabled} on:click={(e)=>dispatch('click', e)}><slot /></button>
{/if}
