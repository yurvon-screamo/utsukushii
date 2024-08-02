<template>
  <div ref="footerRef" :class="cn('flex items-center p-6 pt-0', className)" v-bind="$attrs">
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const footerRef = ref<HTMLDivElement | null>(null);
</script>