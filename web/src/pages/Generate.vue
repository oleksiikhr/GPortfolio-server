<template>
  <div class="page page_generate">
    <app-logo />
    <div class="mv_laptop_l mc">
      <div class="pg_wrap">
        <main class="pg_main">
          <a-steps :current="current">
            <a-step
              v-for="(step, index) in steps"
              :key="index"
              :title="step.title"
              @click="onClickStep(step, index)"
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
import AppLogo from '@/components/Logo'

export default {
  components: {
    AppLogo
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
    }
  },
  methods: {
    onClickStep(_, index) {
      this.current = index
    }
  }
}
</script>
