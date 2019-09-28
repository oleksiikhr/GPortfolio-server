<template>
  <div class="content">
    <aside class="sidebar">
      <slot name="before-sidebar" />
      <slot name="sidebar">
        Content
      </slot>
      <slot name="after-sidebar" />
    </aside>
    <div class="content__wrap">
      <slot />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    menu: {
      type: Array,
      default: () => []
    }
  }
}
</script>

<style lang="scss" scoped>
.content {
  display: flex;
  .sidebar {
    width: 300px;
    margin-right: 20px;
  }
  .content__wrap {
    flex: 1;
    background: #fff;
  }
}
</style>
