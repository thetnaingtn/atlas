<script lang="ts">
  const props = $props();
  let { class: className = '', placeholder = undefined } = props;
  let value = $bindable(props.value ?? '');
</script>

<textarea
  class={`flex w-full rounded-md border px-3 py-2 text-sm ${className}`}
  {placeholder}
  bind:value
/>
