<template>
    <CardContent>
        <div class="mt-6">
            <TableComponent>
                <TableBody>
                    <template v-for="(group, index) in content.groups" :key="index">
                        <TableRow @click="handleRowClick(index)"
                            :class="{ 'cursor-pointer': true, 'bg-muted': expandedRow === index }">
                            <TableCell class="font-medium">{{ group.name }}</TableCell>
                            <TableCell
                                :class="{ 'font-medium': true, 'text-green-500': group.state === State.success, 'text-red-500': group.state === State.dropped, 'text-yellow-500': group.state === State.skipped }">
                                {{ group.state }}
                            </TableCell>
                            <TableCell>{{ group.duration }}</TableCell>
                        </TableRow>
                        <TableRow v-if="expandedRow === index">
                            <TableCell colspan="3">
                                <TableComponent>
                                    <TableBody>
                                        <template v-for="(test, childIndex) in group.tests" :key="childIndex">
                                            <TableRow>
                                                <TableCell class="font-medium">{{ test.name }}</TableCell>
                                                <TableCell
                                                    :class="{ 'font-medium': true, 'text-green-500': test.state === State.success, 'text-red-500': test.state === State.dropped, 'text-yellow-500': test.state === State.skipped }">
                                                    {{ test.state }}
                                                </TableCell>
                                                <TableCell>{{ test.duration }}</TableCell>
                                            </TableRow>
                                        </template>
                                    </TableBody>
                                </TableComponent>
                            </TableCell>
                        </TableRow>
                    </template>
                </TableBody>
            </TableComponent>
        </div>
    </CardContent>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { TableComponent, TableBody, TableRow, TableCell } from '@/components/table';
import { type ReportContent, State } from '@/lib/model';

interface TestReportDetailsProps {
    content: ReportContent;
}

const props = defineProps<TestReportDetailsProps>();
const content = props.content;

const expandedRow = ref<number | null>(null);

function handleRowClick(rowIndex: number) {
    expandedRow.value = rowIndex === expandedRow.value ? null : rowIndex;
}
</script>