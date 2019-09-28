<template>
  <div class="component_contributors">
    <!--TODO Load from file and use random-->
    <a
      v-for="i in 50"
      :key="i"
      href="#"
    >
      <a-avatar
        icon="user"
        alt="TODO"
      />
    </a>
    <a-popover placement="bottom">
      <template slot="content">
        <span>Become a Contributor</span>
      </template>
      <a
        :href="ORGANIZATION"
        target="_blank"
      >
        <a-avatar
          type="primary"
          icon="plus"
          alt="TODO"
        />
      </a>
    </a-popover>
  </div>
</template>

<script>
import { ORGANIZATION } from '@/scripts/links'

export default {
  data() {
    return {
      ORGANIZATION
    }
  }
}
</script>
