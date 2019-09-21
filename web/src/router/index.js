'use strict'

import NProgress from 'nprogress'
import router from '@/router/router'

router.beforeEach((to, from, next) => {
  NProgress.start()
  next()
})

router.afterEach(() => {
  NProgress.done()
})

export default router
