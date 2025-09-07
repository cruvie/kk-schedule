<template>
  <div>
    <h1>Job List</h1>
    <el-button type="primary" @click="handleCreateJob">Create Job</el-button>
    <el-table :data="jobs" border stripe style="width: 100%">
      <el-table-column prop="EntryID" label="Entry ID" width="100"></el-table-column>
      <el-table-column prop="Description" label="Description"></el-table-column>
      <el-table-column prop="FuncName" label="Function Name"></el-table-column>
      <el-table-column prop="Spec" label="Spec"></el-table-column>
      <el-table-column prop="ServiceName" label="Service Name"></el-table-column>
      <el-table-column prop="Next" label="Next">
        <template #default="scope">
          {{ scope.row.Next?.seconds ? new Date(Number(scope.row.Next.seconds) * 1000).toLocaleString() : '' }}
        </template>
      </el-table-column>
      <el-table-column prop="Prev" label="Prev">
        <template #default="scope">
          {{ scope.row.Prev?.seconds ? new Date(Number(scope.row.Prev.seconds) * 1000).toLocaleString() : '' }}
        </template>
      </el-table-column>
      <el-table-column prop="Enabled" label="Status" width="100">
        <template #default="scope">
          <el-tag :type="scope.row.Enabled ? 'success' : 'danger'">{{ scope.row.Enabled ? 'Enabled' : 'Disabled' }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="Actions" width="280">
        <template #default="scope">
          <div class="button-group">
            <el-button @click="handleEnableJob(scope.row)" :disabled="scope.row.Enabled">Enable</el-button>
            <el-button @click="handleDisableJob(scope.row)" :disabled="!scope.row.Enabled">Disable</el-button>
            <el-button @click="handleEditJob(scope.row)">Edit</el-button>
            <el-button @click="handleSetSpecJob(scope.row)">Set Spec</el-button>
            <el-button type="primary" @click="handleTriggerJob(scope.row)">Trigger</el-button>
            <el-button type="danger" @click="handleDeleteJob(scope.row)">Delete</el-button>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <JobForm ref="jobFormRef" @jobUpdated="fetchJobs" />
    <JobSetSpecForm ref="jobSetSpecFormRef" @jobUpdated="fetchJobs" /></div>
</template>


<style scoped>
.button-group .el-button {
  margin-bottom: 10px; /* Add margin to the bottom of each button */
}
</style>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { clientKKSchedule } from '~/utils/api/client';
import { JobList_InputSchema} from '~~/gen/kk_schedule/JobList_pb';
import type { PBJob } from '~~/gen/kk_schedule/Job_pb';
import {create} from "@bufbuild/protobuf";
import { JobEnable_InputSchema } from "~~/gen/kk_schedule/JobEnable_pb";
import { JobDisable_InputSchema } from "~~/gen/kk_schedule/JobDisable_pb";
import { JobDelete_InputSchema } from "~~/gen/kk_schedule/JobDelete_pb";
import { JobPut_InputSchema } from "~~/gen/kk_schedule/JobPut_pb";
import { JobTrigger_InputSchema } from "~~/gen/kk_schedule/JobTrigger_pb";
import JobForm from '~/components/JobForm.vue';
import JobSetSpecForm from '~/components/JobSetSpecForm.vue';
import { ElRow, ElCol, ElMessage, ElMessageBox } from 'element-plus';

const jobs = ref<PBJob[]>([]);
const jobFormRef = ref<InstanceType<typeof JobForm> | null>(null);
const jobSetSpecFormRef = ref<InstanceType<typeof JobSetSpecForm> | null>(null);

const fetchJobs = async () => {
  try {
    const param = create(JobList_InputSchema);
    const out = await clientKKSchedule.jobList(param);
    jobs.value = out.JobList || [];
  } catch (error) {
    ElMessage.error('Error fetching job list: ' + error);
  }
};

onMounted(async () => {
  await fetchJobs();
});

const handleDisableJob = async (job: PBJob) => {
  try {
    const request = create(JobDisable_InputSchema, { serviceName: job.ServiceName, funcName: job.FuncName });
    await clientKKSchedule.jobDisable(request);
    await fetchJobs();
  } catch (error) {
    ElMessage.error('Error disabling job: ' + error);
  }
};

const handleDeleteJob = async (job: PBJob) => {
  ElMessageBox.confirm(
    `Are you sure you want to delete job "${job.FuncName}" from service "${job.ServiceName}"?`,
    'Warning',
    {
      confirmButtonText: 'Yes',
      cancelButtonText: 'No',
      type: 'warning',
    }
  )
    .then(async () => {
      try {
        const request = create(JobDelete_InputSchema, { serviceName: job.ServiceName, funcName: job.FuncName });
        await clientKKSchedule.jobDelete(request);
        await fetchJobs();
        ElMessage.success('Job deleted successfully');
      } catch (error) {
        ElMessage.error('Error deleting job: ' + error);
      }
    })
    .catch(() => {
      ElMessage.info('Delete canceled');
    });
};

const handleEnableJob = async (job: PBJob) => {
  try {
    const request = create(JobEnable_InputSchema, { serviceName: job.ServiceName, funcName: job.FuncName });
    await clientKKSchedule.jobEnable(request);
    await fetchJobs();
  } catch (error) {
    ElMessage.error('Error enabling job: ' + error);
  }
};

const handleEditJob = async (job: PBJob) => {
  jobFormRef.value?.open(job);
};

const handleCreateJob = () => {
  jobFormRef.value?.open();
};

const handleSetSpecJob = (job: PBJob) => {
  jobSetSpecFormRef.value?.open(job);
};

const handleTriggerJob = async (job: PBJob) => {
  try {
    ElMessageBox.confirm(
      `Are you sure you want to trigger job "${job.FuncName}" from service "${job.ServiceName}" manually?`,
      'Confirmation',
      {
        confirmButtonText: 'Yes',
        cancelButtonText: 'No',
        type: 'info',
      }
    )
      .then(async () => {
        try {
          const request = create(JobTrigger_InputSchema, { serviceName: job.ServiceName, funcName: job.FuncName });
          await clientKKSchedule.jobTrigger(request);
          ElMessage.success(`Job "${job.FuncName}" triggered successfully`);
        } catch (error) {
          ElMessage.error('Error triggering job: ' + error);
        }
      })
      .catch(() => {
        ElMessage.info('Trigger canceled');
      });
  } catch (error) {
    ElMessage.error('Error: ' + error);
  }
};

</script>
