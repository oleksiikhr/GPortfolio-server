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
                slot="icon"
                :type="step.icon"
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
import ChooseTemplate from '@/components/steps/choose/Index'
import Customize from '@/components/steps/customize/Index'
import Preview from '@/components/steps/preview/Index'
import Deploy from '@/components/steps/deploy/Index'
import AppLogo from '@/components/global/Logo'

export default {
  components: {
    AppLogo
  },
  data() {
    return {
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
  computed: {
    current: {
      get() {
        return this.$store.state.app.step
      },
      set(value) {
        this.$store.commit('app/SET_STEP', value)
      }
    }
  },
  methods: {
    onClickStep(_, index) {
      this.current = index
    }
  }
}
</script>
