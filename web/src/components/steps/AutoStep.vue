<template>
  <step :menu="menu">
    <block
      v-for="(item, index) in activeMenu"
      :key="index"
      :title="item.title"
      :href="item.href"
      :level="item.level"
    >
      <component :is="item.component" />
    </block>
  </step>
</template>

<script>
import Block from '@/components/steps/Block'
import Step from '@/components/steps/Step'

export default {
  components: {
    Step, Block
  },
  props: {
    menu: {
      type: Array,
      required: true
    }
  },
  computed: {
    activeMenu() {
      const list = []

      const each = (menu, level = 1) => {
        if (typeof menu.enable === 'undefined' || menu.enable) {
          menu.level = level
          list.push(menu)
        }

        if (Array.isArray(menu.children)) {
          menu.children.forEach((item) => each(item, level + 1))
        }
      }

      this.menu.forEach((item) => each(item))

      return list
    }
  }
}
</script>
