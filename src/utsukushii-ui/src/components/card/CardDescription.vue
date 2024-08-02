<template>
  <p ref="descriptionRef" :class="cn('text-sm text-muted-foreground', className)" v-bind="$attrs">
    <slot></slot>
  </p>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const descriptionRef = ref<HTMLParagraphElement | null>(null);
</script>