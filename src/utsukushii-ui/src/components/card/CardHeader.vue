<template>
  <div ref="headerRef" :class="cn('flex flex-col space-y-1.5 p-6', className)" v-bind="$attrs">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const headerRef = ref<HTMLDivElement | null>(null);
</script>