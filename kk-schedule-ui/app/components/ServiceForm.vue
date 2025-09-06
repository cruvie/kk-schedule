<template>
  <el-dialog v-model="dialogVisible" :title="isEdit ? 'Edit Service' : 'Create Service'" width="50%">
    <el-form :model="form" label-width="120px">
      <el-form-item label="Service Name">
        <el-input v-model="form.ServiceName" :disabled="isEdit"></el-input>
      </el-form-item>
      <el-form-item label="Target">
        <el-input v-model="form.Target"></el-input>
      </el-form-item>
      <el-form-item label="Auth Token">
        <el-input v-model="form.AuthToken"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="submitForm">Confirm</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { create } from "@bufbuild/protobuf";
import { clientKKSchedule } from '~/utils/api/client';
import { ServicePut_InputSchema } from '~~/gen/kk_schedule/ServicePut_pb';
import type { PBRegisterService } from '~~/gen/kk_schedule/Base_pb';
import { PBRegisterServiceSchema } from '~~/gen/kk_schedule/Base_pb';
import { ElMessage } from 'element-plus';

const dialogVisible = ref(false);
const isEdit = ref(false);
const form = reactive<PBRegisterService>(create(PBRegisterServiceSchema));

const emit = defineEmits(['serviceUpdated']);

const open = (service?: PBRegisterService) => {
  dialogVisible.value = true;
  if (service) {
    isEdit.value = true;
    Object.assign(form, service);
  } else {
    isEdit.value = false;
    // Reset form for new service
    form.ServiceName = '';
    form.Target = '';
    form.AuthToken = '';
  }
};

const submitForm = async () => {
  try {
    const request = create(ServicePut_InputSchema, {
      Service: form,
    });
    await clientKKSchedule.servicePut(request);
    dialogVisible.value = false;
    emit('serviceUpdated');
  } catch (error) {
    ElMessage.error('Error submitting service form: ' + error);
  }
};

defineExpose({ open });
</script>

<style scoped>
.dialog-footer button:first-child {
  margin-right: 10px;
}
</style>