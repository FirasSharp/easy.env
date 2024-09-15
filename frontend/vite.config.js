import { build, defineConfig, normalizePath } from "vite";
//import { viteStaticCopy } from "vite-plugin-static-copy";
import vue from "@vitejs/plugin-vue";
//import path from "node:path";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    sourcemap: "inline",
  },
  plugins: [vue()],
});
