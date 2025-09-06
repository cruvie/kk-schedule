<template>
  <el-container>
    <el-main>
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <el-text tag="h1" size="large">Welcome to kk-schedule</el-text>
          </div>
        </template>
        <el-text>
          This is a frontend interface for kk-schedule.
        </el-text>
        <el-text>
          You can manage your tasks and schedules here.
        </el-text>
        <div style="margin-top: 20px;">
          <el-link type="primary" href="https://github.com/cruvie/kk-scheduler" target="_blank">
            Visit GitHub Repository
          </el-link>
        </div>
      </el-card>
    </el-main>
  </el-container>
</template>

<style scoped>
.box-card {
  width: 480px;
  margin: 50px auto;
  text-align: center;
}

.card-header {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
