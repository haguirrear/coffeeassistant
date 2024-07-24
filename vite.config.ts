import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  base: "/public",
  plugins: [react()],
  publicDir: "./public",
  server: {
    strictPort: true,
    origin: "http://localhost:8000",
  },
  build: {
    outDir: "dist",
    emptyOutDir: true,
    manifest: "manifest.json",
    rollupOptions: {
      input: "./frontend/src/main.tsx",
    },
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./frontend/src"),
    },
  },
});
