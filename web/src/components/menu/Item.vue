<template>
  <a-anchor-link
    :children="null"
    :component="null"
    :enable="null"
    v-bind="item"
    v-on="item"
  >
    <template v-if="hasChildren">
      <menu--item
        v-for="(childItem, index) in item.children"
        :key="index"
        :item="childItem"
      />
    </template>
  </a-anchor-link>
</template>

<script>
export default {
  name: 'Menu-Item',
  props: {
    item: {
      type: Object,
      required: true
    }
  },
  computed: {
    hasChildren() {
      return Array.isArray(this.item.children) && !!this.item.children.length
    }
  }
}
</script>
