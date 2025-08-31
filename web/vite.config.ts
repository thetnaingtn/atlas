import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

const devProxyServer = 'http://localhost:8080';

export default defineConfig({
	plugins: [tailwindcss(), sveltekit()],
	server: {
    host: "0.0.0.0",
    port: 8888,
    proxy: {
      "^/api": {
        target: devProxyServer,
        xfwd: true,
      },
      "^/api.v1.*": {
        target: devProxyServer,
        xfwd: true,
      },
      "^/file": {
        target: devProxyServer,
        xfwd: true,
      },
    },
  },
});
