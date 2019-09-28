<template>
  <a-anchor
    :offsetTop="20"
    :offsetBottom="20"
    class="component_menu"
  >
    <menu-item
      v-for="(item, index) in items"
      :key="index"
      :item="item"
    />
  </a-anchor>
</template>

<script>
import MenuItem from '@/components/menu/Item'

export default {
  components: {
    MenuItem
  },
  props: {
    items: {
      type: Array,
      required: true
    }
  }
}
</script>
