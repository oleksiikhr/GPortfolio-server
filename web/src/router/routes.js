'use strict'

export default [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/Home')
  },
  {
    path: '/steps',
    name: 'steps',
    component: () => import('@/pages/Steps')
  }
]
