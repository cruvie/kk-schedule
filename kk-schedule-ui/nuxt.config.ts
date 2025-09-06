// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    ssr: false,
    app: {
        head: {
            title: 'kk-schedule'
        }
    },
    compatibilityDate: '2025-09-05',
    devtools: {enabled: true},
    modules: [
        '@element-plus/nuxt'
    ],
    build: {
    },
    elementPlus: { /** Options */}
})
