<template>
  <div class="file-list">
    <!-- 搜索和上传区域 -->
    <div class="operations">
      <el-input
        v-model="searchText"
        placeholder="搜索文件..."
        class="search-input"
        clearable
        @input="handleSearch"
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
      </el-input>

      <el-upload
        :auto-upload="false"
        :show-file-list="false"
        :on-change="handleFileChange"
      >
        <el-button type="primary">
          <el-icon><Upload /></el-icon>
          上传文件
        </el-button>
      </el-upload>
    </div>

    <!-- 文件列表 -->
    <el-table
      v-loading="loading"
      :data="fileList"
      style="width: 100%"
    >
      <el-table-column prop="fileName" label="文件名">
        <template #default="{ row }">
          <el-button 
            link 
            type="primary" 
            @click="handleDownload(row)"
          >
            {{ row.fileName }}
          </el-button>
        </template>
      </el-table-column>
      <el-table-column prop="fileSize" label="大小">
        <template #default="{ row }">
          {{ formatFileSize(row.fileSize) }}
        </template>
      </el-table-column>
      <el-table-column prop="uploadTime" label="上传时间">
        <template #default="{ row }">
          {{ formatDate(row.uploadTime) }}
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Search, Upload } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { fileApi, type FileItem } from '../api/file'
import { formatFileSize, formatDate } from '../utils/format'

const loading = ref(false)
const fileList = ref<FileItem[]>([])
const searchText = ref('')

// 获取文件列表
const fetchFileList = async () => {
  loading.value = true
  try {
    const files = await fileApi.getFileList(searchText.value)
    fileList.value = files
  } catch (error) {
    ElMessage.error('获取文件列表失败')
  } finally {
    loading.value = false
  }
}

// 搜索处理
const handleSearch = () => {
  fetchFileList()
}

// 文件上传处理
const handleFileChange = async (file: any) => {
  try {
    await fileApi.uploadFile(file.raw)
    ElMessage.success('上传成功')
    fetchFileList() // 刷新文件列表
  } catch (error) {
    ElMessage.error('上传失败')
  }
}

// 文件下载处理
const handleDownload = async (file: FileItem) => {
  try {
    await fileApi.downloadFile(file.id, file.fileName)
  } catch (error) {
    ElMessage.error('下载失败')
  }
}

onMounted(() => {
  fetchFileList()
})
</script>

<style scoped>
.file-list {
  padding: 20px;
}

.operations {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.search-input {
  width: 300px;
}
</style>