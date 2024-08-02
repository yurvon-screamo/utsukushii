<template>
  <h3 ref="titleRef" :class="cn('text-2xl font-semibold leading-none tracking-tight', className)" v-bind="$attrs">
    <slot></slot>
  </h3>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const titleRef = ref<HTMLHeadingElement | null>(null);
</script>