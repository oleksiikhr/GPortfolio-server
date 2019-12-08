<template>
  <div class="component_contributors">
    <a
      v-for="(contributor, index) in list"
      :key="index"
      :href="`${GITHUB_DOMAIN}/${contributor.login}`"
      :title="contributor.name"
      target="_blank"
      rel="noreferrer"
    >
      <a-avatar
        icon="user"
        :src="contributor.image"
        :alt="contributor.name"
      />
    </a>
    <a-popover placement="bottom">
      <template slot="content">
        <span>Become a Contributor</span>
      </template>
      <a
        :href="ORGANIZATION"
        rel="noreferrer"
        target="_blank"
      >
        <a-avatar
          type="primary"
          icon="plus"
          alt="Become a contributor"
        />
      </a>
    </a-popover>
  </div>
</template>

<script>
import { ORGANIZATION, GITHUB_DOMAIN } from '@/scripts/links'
import contributors from '@/data/contributors'
import { shuffle } from '@/scripts/arr'

/**
 * How many users (random) we display on page
 * @type {number}
 */
const COUNT_USERS = 50

export default {
  data() {
    return {
      ORGANIZATION,
      GITHUB_DOMAIN
    }
  },
  computed: {
    list() {
      const users = Object.entries(contributors)
        .reduce((arr, [login, contributor]) => {
          arr.push({ ...contributor, login })
          return arr
        }, [])

      return shuffle(users).slice(0, COUNT_USERS)
    }
  }
}
</script>
