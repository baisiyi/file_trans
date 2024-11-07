<template>
  <el-container class="file-manager">
    <el-header>
      <h2>文件管理系统</h2>
    </el-header>
    <el-main>
      <FileList />
    </el-main>
  </el-container>
</template>

<script setup lang="ts">
import FileList from '@/components/FileList.vue'
</script>

<style>
/* 全局样式 */
html, body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
}

#app {
  width: 100%;
  height: 100%;
}
</style>

<style scoped>
/* 组件样式 */
.file-manager {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
}

.el-header {
  width: 100%;
  background-color: #f5f7fa;
  color: #333;
  line-height: 60px;
  padding: 0 20px;
  border-bottom: 1px solid #e6e6e6;
  /* 使用 box-sizing 确保 padding 不会增加实际宽度 */
  box-sizing: border-box;
  /* 固定在顶部 */
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
}

.el-main {
  /* 为固定的 header 留出空间 */
  margin-top: 60px;
  flex: 1;
  padding: 20px;
  box-sizing: border-box;
}

h2 {
  margin: 0;
}

.content {
  margin-top: 20px;
}
</style>