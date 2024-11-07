import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd())
  
  return {
    plugins: [vue()],
    
    // 添加路径解析配置
    resolve: {
      alias: {
        '@': path.resolve(__dirname, './src')  // 确保这里的路径正确
      }
    },
    
    server: {
      host: '0.0.0.0',
      port: 5173,
      proxy: {
        '/api': {
          target: env.VITE_API_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '/v1')
        }
      }
    },

    build: {
      outDir: 'dist',
      minify: 'esbuild',
      esbuild: {
        drop: ['console', 'debugger']
      },
      // 添加分包配置
      rollupOptions: {
        output: {
          manualChunks: {
            'vue': ['vue', 'vue-router'],
            'element-plus': ['element-plus'],
            'vendor': ['axios', '@element-plus/icons-vue'],
            // 如果使用了其他大型依赖，也可以单独分包
          },
          // 自定义 chunk 文件名
          chunkFileNames: 'assets/js/[name]-[hash].js',
          // 自定义入口文件名
          entryFileNames: 'assets/js/[name]-[hash].js',
          // 自定义静态资源文件名
          assetFileNames: 'assets/[ext]/[name]-[hash].[ext]'
        }
      },
      // 提高警告阈值（如果你不想看到这个警告）
      chunkSizeWarningLimit: 1000
    },

    preview: {
      port: 5173,
      host: '0.0.0.0',
      strictPort: true
    }
  }
})