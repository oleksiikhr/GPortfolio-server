<template>
  <div>
    <a-button @click="onClick">Click</a-button>
    Github
  </div>
</template>

<script>
import { REGEX_PASS } from '@/data/auth'
import { popup } from '@/scripts/window'

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
  methods: {
    onClick() {
      // TODO Move logic outside
      const w = popup('/api/github/oauth/redirect', 'Github')

      w.addEventListener('beforeunload', (evt) => {
        const response = evt.target.body.innerText

        console.log('code', response)
        if (REGEX_PASS.test(response)) {
          const [login, password] = response.split('@')
          localStorage.setItem('github-website', JSON.stringify({ login, password }))
          console.log('Success!')
          // TODO Success, store to localStorage
        } else {
          console.log('Bad')
        }
      })
    }
  }
}
</script>
