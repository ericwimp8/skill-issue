import { defineConfig } from "vite";

export default defineConfig({
  resolve: { conditions: ["widget"] },
  css: { preprocessorOptions: { scss: { api: "legacy" } } },
  build: { lib: { entry: "src/index.ts", name: "Widget" }, minify: "terser" },
});
