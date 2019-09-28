'use strict'

import Generate from '@/pages/Generate'
import Home from '@/pages/Home'

export default [
  {
    path: '/',
    name: 'home',
    component: Home
  },
  {
    path: '/generate',
    name: 'generate',
    component: Generate
  }
]
