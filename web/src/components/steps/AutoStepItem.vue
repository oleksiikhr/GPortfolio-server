<template>
  <div>
    <block
      :title="item.title"
      :href="item.href"
      :level="level"
    >
      <component :is="item.component" />
    </block>
    <template v-if="hasChildren">
      <auto-step-item
        v-for="(child, index) in item.children"
        :key="index"
        :item="child"
        :level="level + 1"
      />
    </template>
  </div>
</template>

<script>
import Block from '@/components/steps/Block'

export default {
  name: 'AutoStepItem',
  components: {
    Block
  },
  props: {
    item: {
      type: Object,
      required: true
    },
    level: {
      type: Number,
      default: 1
    }
  },
  computed: {
    hasChildren() {
      return this.item.children && !!this.item.children.length
    }
  }
}
</script>
