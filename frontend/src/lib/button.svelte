<script lang="ts">
import type { Snippet } from 'svelte';
import type { ClassValue, HTMLButtonAttributes } from 'svelte/elements';

interface Props extends HTMLButtonAttributes {
  class?: ClassValue;
  children?: Snippet;
}

let { ...props }: Props = $props();
</script>

<button
  {...props}
  class={[
    'cursor-pointer rounded bg-purple-600 p-2 text-neutral-50 hover:bg-purple-500 active:bg-purple-700',
    props.class,
  ]}>{@render props.children?.()}</button
>
