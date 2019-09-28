<template>
  <div class="page page_generate">
    <div class="mc tc">
      <logo-svg class="logo" />
    </div>
    <div class="mv_laptop_l mc">
      <div class="pg_wrap">
        <main class="pg_main">
          <a-steps :current="current">
            <a-step
              v-for="(step, index) in steps"
              :key="index"
              :title="step.title"
              @click="onClickStep"
            >
              <a-icon
                :type="step.icon"
                slot="icon"
              />
            </a-step>
          </a-steps>
          <component :is="steps[current].component" />
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import ChooseTemplate from '@/components/steps/ChooseTemplate'
import Customize from '@/components/steps/Customize'
import Preview from '@/components/steps/Preview'
import Deploy from '@/components/steps/Deploy'
import LogoSvg from '@/components/svg/Logo'

export default {
  components: {
    LogoSvg
  },
  data() {
    return {
      current: 0,
      steps: [{
        title: 'Choose Template',
        icon: 'inbox',
        component: ChooseTemplate
      }, {
        title: 'Customize',
        icon: 'setting',
        content: Customize
      }, {
        title: 'Preview',
        icon: 'eye',
        content: Preview
      }, {
        title: 'Deploy',
        icon: 'database',
        content: Deploy
      }]
    }
  },
  methods: {
    onClickStep() {
      console.log('onClickStep')
    }
  }
}
</script>
