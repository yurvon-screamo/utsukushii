<template>
    <CardComponent class="w-full max-w-3xl">
        <TestReportSummary :content="content" />
        <TestReportDetails :content="content" />
    </CardComponent>
</template>

<script setup lang="ts">
import { CardComponent } from '@/components/card';
import TestReportSummary from './TestReportSummary.vue';
import TestReportDetails from './TestReportDetails.vue';
import type { ReportContent } from '@/lib/model';

interface TestReportProps {
    content: ReportContent;
}

const props = defineProps<TestReportProps>();
const content = props.content;
</script>