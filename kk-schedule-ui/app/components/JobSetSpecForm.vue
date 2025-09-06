<template>
  <el-dialog v-model="dialogVisible" title="Set Job Spec" width="50%">
    <el-form :model="form" label-width="120px">
      <el-form-item label="Service Name">
        <el-input v-model="form.ServiceName" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="Function Name">
        <el-input v-model="form.FuncName" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="Spec">
        <el-input v-model="form.Spec" type="textarea" :rows="5"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="handleSave">Confirm</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { ElMessage } from 'element-plus';
import { clientKKSchedule } from '~/utils/api/client';
import { JobSetSpec_InputSchema } from '~~/gen/kk_schedule/JobSetSpec_pb';
import { create } from "@bufbuild/protobuf";
import type { PBJob } from '~~/gen/kk_schedule/Job_pb';

const dialogVisible = ref(false);

const form = reactive({
  ServiceName: '',
  FuncName: '',
  Spec: '',
});

const emit = defineEmits(['jobUpdated']);

const open = (job: PBJob) => {
  dialogVisible.value = true;
  form.ServiceName = job.ServiceName;
  form.FuncName = job.FuncName;
  form.Spec = job.Spec;
};

const handleSave = async () => {
  try {
    const request = create(JobSetSpec_InputSchema, {
      serviceName: form.ServiceName,
      funcName: form.FuncName,
      spec: form.Spec,
    });
    await clientKKSchedule.jobSetSpec(request);
    ElMessage.success('Job spec updated successfully!');
    dialogVisible.value = false;
    emit('jobUpdated');
  } catch (error) {
    ElMessage.error('Error setting job spec: ' + error);
  }
};

defineExpose({
  open,
});
</script>

<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>