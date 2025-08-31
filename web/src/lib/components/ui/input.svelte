<script lang="ts">
  const props = $props();
  let { class: className = '', type = 'text', placeholder = undefined, required = false } = props;
  let value = $bindable(props.value ?? '');
</script>

<input
  class={`flex h-9 w-full rounded-md border px-3 py-2 text-sm ${className}`}
  {type}
  {placeholder}
  {required}
  bind:value
/>
