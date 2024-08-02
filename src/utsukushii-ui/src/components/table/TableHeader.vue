<template>
    <thead ref="headerRef" :class="cn('[&_tr]:border-b', className)" v-bind="$attrs">
        <slot></slot>
    </thead>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const headerRef = ref<HTMLTableSectionElement | null>(null);

</script>
