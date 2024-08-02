<template>
  <th ref="headRef"
    :class="cn('h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0', className)"
    v-bind="$attrs">
    <slot></slot>
  </th>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const headRef = ref<HTMLTableCellElement | null>(null);
</script>