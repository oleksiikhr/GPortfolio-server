'use strict'

import Generate from '@/pages/Generate'
import Home from '@/pages/Home'

// List of pages
export const GENERATE_PAGE = 'generate'
export const HOME_PAGE = 'home'

export default [
  {
    path: '/',
    name: HOME_PAGE,
    component: Home
  },
  {
    path: '/generate',
    name: GENERATE_PAGE,
    component: Generate
  }
]
