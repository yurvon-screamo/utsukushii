<template>
  <div v-if="reportContent" class="container mx-auto p-4">
    <TestReport :content="reportContent" />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { TestReport } from '@/components/testReport';
import type { ReportContent } from '@/lib/model';

const reportContent = ref<ReportContent | null>(null);

onMounted(async () => {
  const response = await fetch('/data.json');
  if (response.ok) {
    reportContent.value = await response.json();
  }
});
</script>

<style scoped>
.container {
  max-width: 800px;
}
</style>
