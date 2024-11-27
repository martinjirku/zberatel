import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
/* @ts-ignore */
import codegen from "vite-plugin-codegen";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), codegen()],
  server: {
    port: 3000,
    proxy: {
      "/query": "http://localhost:8081",
    },
  },
});
