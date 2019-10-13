<template>
  <auto-step :menu="menu" />
</template>

<script>
import GlobalBlock from '@/components/steps/customize/blocks/Global'
import AutoStep from '@/components/steps/AutoStep'

export default {
  components: {
    AutoStep
  },
  data() {
    return {
      menu: [
        {
          component: GlobalBlock,
          href: '#global',
          title: 'Global'
        }
      ]
    }
  }
}
</script>
