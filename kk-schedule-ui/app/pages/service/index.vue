<template>
  <div>
    <div class="common-layout">
      <h1>Service List</h1>
      <el-button type="primary" @click="handleCreateService">Create Service</el-button>
      <el-table :data="services" border stripe style="width: 100%">
        <el-table-column prop="ServiceName" label="Service Name"></el-table-column>
        <el-table-column prop="Target" label="Target"></el-table-column>
        <el-table-column prop="AuthToken" label="Auth Token"></el-table-column>
        <el-table-column label="Actions" width="200">
          <template #default="scope">
            <div class="button-group">
              <el-button @click="handleEditService(scope.row)">Edit</el-button>
              <el-button type="danger" @click="handleDeleteService(scope.row)">Delete</el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>
      <ServiceForm ref="serviceFormRef" @serviceUpdated="fetchServices"/>
    </div>
  </div>
</template>

<style scoped>
.button-group .el-button {
  margin-bottom: 10px; /* Add margin to the bottom of each button */
}
</style>

<script setup lang="ts">
import {ref, onMounted} from 'vue';
import {clientKKSchedule} from '~/utils/api/client';
import ServiceForm from '~/components/ServiceForm.vue';
import type {PBRegisterService} from '~~/gen/kk_schedule/Base_pb';
import {ServiceList_InputSchema} from '~~/gen/kk_schedule/ServiceList_pb';
import {create} from '@bufbuild/protobuf';
import {ServiceDelete_InputSchema} from '~~/gen/kk_schedule/ServiceDelete_pb';
import {ElMessage} from 'element-plus';

const services = ref<PBRegisterService[]>([]);
const serviceFormRef = ref<InstanceType<typeof ServiceForm> | null>(null);

const fetchServices = async () => {
  try {
    const request = create(ServiceList_InputSchema);
    const response = await clientKKSchedule.serviceList(request);
    services.value = response.ServiceList || [];
  } catch (error) {
    ElMessage.error('Error fetching service list: ' + error);
  }
};

onMounted(async () => {
  await fetchServices();
});

const handleEditService = async (service: PBRegisterService) => {
  serviceFormRef.value?.open(service);
};

const handleCreateService = () => {
  serviceFormRef.value?.open();
};

const handleDeleteService = async (service: PBRegisterService) => {
  try {
    const request = create(ServiceDelete_InputSchema, {ServiceName: service.ServiceName});
    await clientKKSchedule.serviceDelete(request);
    await fetchServices();
  } catch (error) {
    ElMessage.error('Error deleting service: ' + error);
  }
};
</script>
