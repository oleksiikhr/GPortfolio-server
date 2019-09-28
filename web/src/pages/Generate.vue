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
            >
              <span
                slot="title"
                @click="onClickStep(step, index)"
              >
                {{ step.title }}
              </span>
              <a-icon
                :type="step.icon"
                slot="icon"
                @click="onClickStep(step, index)"
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
import ChooseTemplate from '@/components/steps/choose'
import Customize from '@/components/steps/customize'
import Preview from '@/components/steps/preview'
import Deploy from '@/components/steps/deploy'
import AppLogo from '@/components/global/Logo'

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
