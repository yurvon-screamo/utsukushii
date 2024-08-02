<template>
  <tr ref="rowRef" :class="cn('border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted', className)"
    v-bind="$attrs">
    <slot></slot>
  </tr>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const rowRef = ref<HTMLTableRowElement | null>(null);
</script>