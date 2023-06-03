export default defineNuxtConfig({
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  css: ['~/assets/css/main.css'],
})
// DO NOT PUT ANYTHING BEFORE THE EXPORT DEFAULT.
// https://nuxt.com/docs/api/configuration/nuxt-config