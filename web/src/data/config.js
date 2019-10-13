'use strict'

export default {
  global: {
    template: '',
    base: '',
    opg: {},
    pwa: {},
    domain: ''
  },
  data: {
    name: '',
    avatar_url: '',
    position: '',
    hire: false,
    socialMedia: []
  },
  modules: {
    github: {
      username: '',
      token: '',
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
      response: {
        profile: null,
        repositories: []
      }
    },
    dribbble: {
      auth: {
        client_id: '',
        client_secret: '',
        code: ''
      },
      filter: {
        shots: []
      },
      sort: {
        shots: {
          attr: 'id',
          enable: false,
          sortByDesc: true
        }
      },
      response: {
        profile: null,
        shots: []
      }
    }
  },
  templates: {
    default: {
      background: '',
      github_repositories_more: 15,
      dribbble_shots_more: 15
    }
  }
}
