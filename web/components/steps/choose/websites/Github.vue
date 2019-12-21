<template>
  <div>
    <a-button @click="onClick">Click</a-button>
    Github
    <template v-if="store">
      <a-button @click="onClickProfile">Profile</a-button>
    </template>
    <br><br>
    {{ $store.state.githubWebsite.profile }}
  </div>
</template>

<script>
import {
  GITHUB_WEBSITE,
  getObject as getObjectStore,
  setObject as setObjectStore
} from '@/scripts/storage'
import { getAuthHeaders, windowAuth } from '@/scripts/auth'
import { popup } from '@/scripts/window'
import axios from 'axios'

/*
parse: {
  repositories: {
    type: 'owner',
    sort: 'updated',
    direction: 'desc',
    visibility: 'public',
    affiliation: 'owner,collaborator,organization_member'
  }
},
filter: {
  repositories: []
},
sort: {
  repositories: {
    attr: 'stargazers_count',
    enable: true,
    sortByDesc: true
  }
},
 */

export default {
  data() {
    return {
      store: null
    }
  },
  mounted() {
    const store = getObjectStore(GITHUB_WEBSITE)

    if (Object.keys(store).length > 0) {
      this.store = store
    }
  },
  methods: {
    onClick() {
      const w = popup('/api/github/oauth/redirect', 'Github')

      windowAuth(w)
        .then(({ login, password }) => {
          this.store = { login, password }
          setObjectStore(GITHUB_WEBSITE, this.store)
        })
    },
    onClickProfile() {
      axios.get('/api/github/profile', {
        headers: getAuthHeaders(this.store.login, this.store.password)
      })
        .then((resp) => {
          this.$store.commit('githubWebsite/SET_PROFILE', resp.data.data)
        })
    }
  }
}
</script>
