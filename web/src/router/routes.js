'use strict'

export default [
  {
    path: '/',
    name: 'home',
    component: () => import('@/pages/Home')
  },
  {
    path: '/test',
    name: 'test',
    component: () => import('@/pages/Home')
  }
]
