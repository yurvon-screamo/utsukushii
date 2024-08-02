<template>
  <div class="relative w-full overflow-auto">
    <table ref="tableRef" :class="cn('w-full caption-bottom text-sm', className)" v-bind="$attrs">
      <slot></slot>
    </table>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

defineProps<{ className?: string }>();
const tableRef = ref<HTMLTableElement | null>(null);

</script>
