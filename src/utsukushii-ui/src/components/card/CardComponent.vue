<template>
    <div ref="cardRef" :class="cn('rounded-lg border bg-card text-card-foreground shadow-sm', className)"
        v-bind="$attrs">
        <slot></slot>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const cardRef = ref<HTMLDivElement | null>(null);
</script>