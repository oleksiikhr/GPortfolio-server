<template>
  <div class="page page_generate">
    <app-logo />
    <div class="mv_laptop_l mc">
      <div class="pg_wrap">
        <main class="pg_main">
          <navigation-steps />
          <component :is="steps[currentStep].component" />
        </main>
      </div>
    </div>
  </div>
</template>

<script>
import NavigationSteps from '@/components/home/NavigationSteps'
import AppLogo from '@/components/global/Logo'
import steps from '@/data/steps'

export default {
  components: {
    AppLogo, NavigationSteps
  },
  data() {
    return {
      steps
    }
  },
  computed: {
    currentStep() {
      return this.$store.state.app.step
    }
  }
}
</script>
