<template>
  <el-dialog v-model="dialogVisible" :title="isEdit ? 'Edit Job' : 'Create Job'" width="50%">
    <el-form :model="form" label-width="120px">
      <el-form-item label="Description">
        <el-input v-model="form.Description"></el-input>
      </el-form-item>
      <el-form-item label="Function Name">
        <el-input v-model="form.FuncName" :disabled="isEdit"></el-input>
      </el-form-item>
      <el-form-item label="Service Name">
        <el-input v-model="form.ServiceName"></el-input>
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
import { JobPut_InputSchema } from '~~/gen/kk_schedule/JobPut_pb';
import {type PBJob, PBJobSchema} from '~~/gen/kk_schedule/Job_pb';
import { TimestampSchema } from '~~/gen/google/protobuf/timestamp_pb';
import {create} from "@bufbuild/protobuf";
import { PBRegisterJobSchema, type PBRegisterJob } from '~~/gen/kk_schedule/Base_pb';

const dialogVisible = ref(false);
const isEdit = ref(false);

const form = reactive<PBRegisterJob>(create(PBRegisterJobSchema, {
  Description: '',
  FuncName: '',
  ServiceName: '',
}));

const emit = defineEmits(['jobUpdated']);

const open = (job?: PBJob) => {
  dialogVisible.value = true;

  if (job) {
    isEdit.value = true;
    form.Description = job.Description;
    form.FuncName = job.FuncName;
    form.ServiceName = job.ServiceName;
  } else {
    isEdit.value = false;
    Object.assign(form, create(PBRegisterJobSchema, {
      Description: '',
      FuncName: '',
      ServiceName: '',
    }));
  }
};

const handleSave = async () => {
  try {
    if (isEdit.value) {
      const putRequest = create(JobPut_InputSchema, {
        Job: create(PBRegisterJobSchema, {
          Description: form.Description,
          ServiceName: form.ServiceName,
          FuncName: form.FuncName,
        }),
      });
      await clientKKSchedule.jobPut(putRequest);
      ElMessage.success('Job updated successfully!');
    } else {
      const request = create(JobPut_InputSchema, {
        Job: create(PBRegisterJobSchema, {
          Description: form.Description,
          ServiceName: form.ServiceName,
          FuncName: form.FuncName,
        }),
      });
      await clientKKSchedule.jobPut(request);
      ElMessage.success('Job created successfully!');
    }
    dialogVisible.value = false;
    emit('jobUpdated');
  } catch (error) {
    ElMessage.error('Error saving job: ' + error);
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