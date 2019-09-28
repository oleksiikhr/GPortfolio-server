'use strict'

import Vue from 'vue'

/*
 * We import only used components/other for size reduction
 */
import AAnchorLink from 'ant-design-vue/lib/anchor/AnchorLink'
import AStep from 'ant-design-vue/lib/vc-steps/Step'
import ABackTop from 'ant-design-vue/lib/back-top'
import APopover from 'ant-design-vue/lib/popover'
import AButton from 'ant-design-vue/lib/button'
import AAnchor from 'ant-design-vue/lib/anchor'
import AAvatar from 'ant-design-vue/lib/avatar'
import ASteps from 'ant-design-vue/lib/steps'
import AIcon from 'ant-design-vue/lib/icon'

Vue.component('AAnchorLink', AAnchorLink)
Vue.component('ABackTop', ABackTop)
Vue.component('APopover', APopover)
Vue.component('AAvatar', AAvatar)
Vue.component('AAnchor', AAnchor)
Vue.component('AButton', AButton)
Vue.component('ASteps', ASteps)
Vue.component('AStep', AStep)
Vue.component('AIcon', AIcon)
