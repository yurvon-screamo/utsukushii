<template>
    <caption ref="captionRef" :class="cn('mt-4 text-sm text-muted-foreground', className)" v-bind="$attrs">
        <slot></slot>
    </caption>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const captionRef = ref<HTMLTableCaptionElement | null>(null);
</script>