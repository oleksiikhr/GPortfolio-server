'use strict'

import ChooseTemplate from '@/components/steps/choose/Index'
import Customize from '@/components/steps/customize/Index'
import Preview from '@/components/steps/preview/Index'
import Deploy from '@/components/steps/deploy/Index'

export default [{
  title: 'Choose Template',
  icon: 'inbox',
  component: ChooseTemplate
}, {
  title: 'Customize',
  icon: 'setting',
  component: Customize
}, {
  title: 'Preview',
  icon: 'eye',
  component: Preview
}, {
  title: 'Deploy',
  icon: 'database',
  component: Deploy
}]
