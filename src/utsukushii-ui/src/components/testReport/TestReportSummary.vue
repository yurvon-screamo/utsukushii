<template>
    <CardHeader class="flex items-left justify-between">
        <div class="items-center grid grid-cols-3 gap-4">
            <div class="flex items-center gap-2">
                <CalendarIcon class="h-3 w-3 text-muted-foreground" />
                <span class="text-xs text-muted-foreground">{{ content.timestamp }}</span>
            </div>
            <CardTitle>{{ content.title }}</CardTitle>
        </div>
    </CardHeader>

    <CardContent>
        <div class="grid grid-cols-3 gap-4">
            <div class="flex items-center gap-2">
                <TestTubeIcon class="h-6 w-6 text-primary" />
                <div>
                    <div class="text-lg font-medium">{{ content.total }}</div>
                    <div class="text-sm text-muted-foreground">Total Tests</div>
                </div>
            </div>
            <div class="flex items-center gap-2">
                <PercentIcon class="h-6 w-6 text-primary" />
                <div>
                    <div class="text-lg font-medium">{{ content.coverage }}%</div>
                    <div class="text-sm text-muted-foreground">Test Coverage</div>
                </div>
            </div>
            <div class="flex flex-col items-start gap-2">
                <div class="flex items-center gap-2">
                    <CheckIcon class="h-5 w-5 text-green-500" />
                    <div class="text-sm text-muted-foreground">{{ content.success }} Success</div>
                </div>
                <div class="flex items-center gap-2">
                    <XIcon class="h-5 w-5 text-red-500" />
                    <div class="text-sm text-muted-foreground">{{ content.dropped }} Dropped</div>
                </div>
                <div class="flex items-center gap-2">
                    <PauseIcon class="h-5 w-5 text-yellow-500" />
                    <div class="text-sm text-muted-foreground">{{ content.skipped }} Skipped</div>
                </div>
            </div>
        </div>
    </CardContent>
</template>

<script setup lang="ts">
import { CardHeader, CardTitle, CardContent } from '@/components/card';
import { CalendarIcon, CheckIcon, PauseIcon, PercentIcon, TestTubeIcon, XIcon } from '@/components/icons';
import type { ReportContent } from '@/lib/model';

interface TestReportSummaryProps {
    content: ReportContent;
}

const props = defineProps<TestReportSummaryProps>();
const content = props.content;
</script>