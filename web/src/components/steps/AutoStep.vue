<template>
  <step :menu="menu">
    <auto-step-item
      v-for="(item, index) in activeMenu"
      :key="index"
      :item="item"
    />
  </step>
</template>

<script>
import AutoStepItem from '@/components/steps/AutoStepItem'
import Step from '@/components/steps/Step'

export default {
  components: {
    Step, AutoStepItem
  },
  props: {
    menu: {
      type: Array,
      required: () => []
    }
  },
  computed: {
    activeMenu() {
      return this.menu.filter((item) => {
        if (typeof item.enable === 'undefined') {
          return true
        }

        return item.enable
      })
    }
  }
}
</script>
