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
      <div class="cs_wrap_content">
        <slot />
      </div>
      <div class="cs_wrap_navigation">
        <a-button
          type="primary"
          class="cs_wrap_navigation__prev"
          @click="onClickPrev"
        >
          <a-icon type="left" />Prev
        </a-button>
        <a-button
          type="primary"
          class="cs_wrap_navigation__next"
          @click="onClickNext"
        >
          Next<a-icon type="right" />
        </a-button>
      </div>
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
    onClickPrev() {
      this.current -= 1
    },
    onClickNext() {
      this.current += 1
    }
  }
}
</script>
