<template>
  <td ref="cellRef" :class="cn('p-4 align-middle [&:has([role=checkbox])]:pr-0', className)" v-bind="$attrs">
    <slot></slot>
  </td>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const cellRef = ref<HTMLTableCellElement | null>(null);
</script>