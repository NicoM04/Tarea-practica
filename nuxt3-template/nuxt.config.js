import { defineNuxtConfig } from 'nuxt/config'

export default defineNuxtConfig({
  typescript: {
    strict: true
  },
  app: {
    head: {
      title: 'Nuxt3 Test App',
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Template de prueba para postulantes' }
      ]
    }
  },
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.BASE_URL_BACKEND || 'http://localhost:8080'
    }
  }
})
