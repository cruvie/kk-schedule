<template>
  <el-container style="height: 100vh;">
    <el-aside width="200px" style="background-color: #545c64;">
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b"
        router
      >
        <el-menu-item index="/">
          <i class="el-icon-s-home"></i>
          <template #title>Home</template>
        </el-menu-item>
        <el-menu-item index="/job">
          <i class="el-icon-s-tools"></i>
          <template #title>Job List</template>
        </el-menu-item>
        <el-menu-item index="/service">
          <i class="el-icon-s-platform"></i>
          <template #title>Service List</template>
        </el-menu-item>
      </el-menu>
    </el-aside>
    <el-container>
      <el-main>
        <slot />
      </el-main>
    </el-container>
  </el-container>
</template>

<style>
html, body, #__nuxt {
  height: 100%;
  margin: 0;
  padding: 0;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>