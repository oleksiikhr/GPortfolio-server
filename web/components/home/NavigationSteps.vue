<template>
  <a-steps :current="currentStep">
    <a-step
      v-for="(step, index) in steps"
      :key="index"
    >
      <span
        slot="title"
        @click="onClickStep(index)"
      >
        {{ step.title }}
      </span>
      <a-icon
        slot="icon"
        :type="step.icon"
        @click="onClickStep(index)"
      />
    </a-step>
  </a-steps>
</template>

<script>
import steps from '@/data/steps'

export default {
  data() {
    return {
      steps
    }
  },
  computed: {
    currentStep() {
      return this.$store.state.app.step
    }
  },
  methods: {
    onClickStep(index) {
      this.$store.commit('app/SET_STEP', index)
    }
  }
}
</script>
