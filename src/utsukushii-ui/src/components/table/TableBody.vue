<template>
    <tbody ref="bodyRef" :class="cn('[&_tr:last-child]:border-0', className)" v-bind="$attrs">
        <slot></slot>
    </tbody>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const bodyRef = ref<HTMLTableSectionElement | null>(null);
</script>
