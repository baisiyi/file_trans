import axios from 'axios'

// 打印环境变量用于调试
console.log('Current ENV:', import.meta.env)
console.log('API URL:', import.meta.env.VITE_API_URL)

// 根据环境获取基础 URL
const baseURL = import.meta.env.PROD 
  ? `${import.meta.env.VITE_API_URL}/v1`  // 生产环境
  : '/api'  // 开发环境

console.log('Current API baseURL:', baseURL) // 用于调试

const api = axios.create({
  baseURL,
  timeout: 5000
})

export interface FileItem {
  id: string
  fileName: string
  fileSize: number
  uploadTime: string
}

export const fileApi = {
  // 获取文件列表
  async getFileList(search?: string) {
    const response = await api.get<{files: FileItem[]}>('/files', {
      params: { search }
    })
    return response.data.files
  },

  // 上传文件
  async uploadFile(file: File) {
    const formData = new FormData()
    formData.append('content', file)
    formData.append('fileName', file.name)
    
    const response = await api.post<{fileId: string}>('/files', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    return response.data.fileId
  },

  // 下载文件
  async downloadFile(fileId: string, fileName: string) {
    const response = await api.get(`/files/${fileId}/download`, {
      responseType: 'blob'
    })
    
    const url = window.URL.createObjectURL(new Blob([response.data]))
    const link = document.createElement('a')
    link.href = url
    link.download = fileName
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
    window.URL.revokeObjectURL(url)
  }
}