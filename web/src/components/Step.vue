<template>
  <div class="component_step">
    <aside class="cs_aside">
      <slot name="before-sidebar" />
      <slot name="sidebar">
        <menu-list :items="menu" />
      </slot>
      <slot name="after-sidebar" />
    </aside>
    <div class="cs_wrap">
      <slot />
    </div>
    <a-back-top class="cs_back_top">
      <a-icon type="up-square" />
    </a-back-top>
  </div>
</template>

<script>
import MenuList from '@/components/menu/Index'

export default {
  components: {
    MenuList
  },
  props: {
    menu: {
      type: Array,
      default: () => []
    }
  }
}
</script>
