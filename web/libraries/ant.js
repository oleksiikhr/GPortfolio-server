'use strict'

import Vue from 'vue'

/*
 * We import only used components/other for size reduction
 */
import AAnchorLink from 'ant-design-vue/lib/anchor/AnchorLink'
import ACardMeta from 'ant-design-vue/lib/card/Meta'
import AStep from 'ant-design-vue/lib/vc-steps/Step'
import ABackTop from 'ant-design-vue/lib/back-top'
import APopover from 'ant-design-vue/lib/popover'
import AButton from 'ant-design-vue/lib/button'
import AAnchor from 'ant-design-vue/lib/anchor'
import AAvatar from 'ant-design-vue/lib/avatar'
import ASteps from 'ant-design-vue/lib/steps'
import AAlert from 'ant-design-vue/lib/alert'
import ABadge from 'ant-design-vue/lib/badge'
import ACard from 'ant-design-vue/lib/card'
import AIcon from 'ant-design-vue/lib/icon'

Vue.component('AAnchorLink', AAnchorLink)
Vue.component('ACardMeta', ACardMeta)
Vue.component('ABackTop', ABackTop)
Vue.component('APopover', APopover)
Vue.component('AAvatar', AAvatar)
Vue.component('AAnchor', AAnchor)
Vue.component('AButton', AButton)
Vue.component('ASteps', ASteps)
Vue.component('AAlert', AAlert)
Vue.component('ABadge', ABadge)
Vue.component('AStep', AStep)
Vue.component('ACard', ACard)
Vue.component('AIcon', AIcon)
