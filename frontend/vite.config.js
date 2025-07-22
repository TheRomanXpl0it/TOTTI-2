import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'


export default defineConfig(() => {
  return {
    plugins: [react()],
    server: {
      port: 7000,
      host: true,
      proxy: {
        '/api': {
          target: 'http://127.0.0.1:5000',
          changeOrigin: true,
          secure: false,
          // rewrite: (path) => path.replace(/^\/api/, '')
        }
      }
    },
    build: {
      outDir: 'dist',
      emptyOutDir: true
    }
  }
});
