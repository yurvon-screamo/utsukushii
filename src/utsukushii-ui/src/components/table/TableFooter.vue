<template>
    <tfoot ref="footerRef" :class="cn('border-t bg-muted/50 font-medium [&>tr]:last:border-b-0', className)"
        v-bind="$attrs">
        <slot></slot>
    </tfoot>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { cn } from '@/lib/merge';

const props = defineProps<{ className?: string }>();
const footerRef = ref<HTMLTableSectionElement | null>(null);
</script>
