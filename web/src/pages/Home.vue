<template>
  <div class="page home-page">
    <div class="center text--center">
      <logo-svg class="logo" />
    </div>
    <header class="max--width">
      <div class="header__wrap">
        <div class="left">
          <h1>Create Your Portfolio Online</h1>
          <p>Upload your data from social networks or fill in manually</p>
          <div class="ant-row-flex actions">
            <a-button
              type="primary"
              size="large"
              @click="onClickTryIt"
            >
              Try it Free
            </a-button>
            <a-button
              size="large"
              icon="github"
              @click="onClickGithub"
            >
              Github
            </a-button>
          </div>
          <!--TODO https://vue.ant.design/components/avatar/-->
          <div class="contributors">
            <a
              v-for="i in 50"
              :key="i"
              href="#"
            >
              <a-avatar icon="user" alt />
            </a>
            <a-popover placement="bottom">
              <template slot="content">
                <span>Become a Contributor</span>
              </template>
              <a href="https://github.com/GPortfolio/GPortfolio" target="_blank">
                <a-avatar type="primary" style="backgroundColor: #3f3d56" icon="plus" alt />
              </a>
            </a-popover>
          </div>
        </div>
        <div class="right">
          <hire-svg />
        </div>
      </div>
    </header>
    <wave-svg class="wave" />
  </div>
</template>

<script>
export default {
  components: {
    HireSvg: () => import('@/components/svg/Hire'),
    WaveSvg: () => import('@/components/svg/Wave'),
    LogoSvg: () => import('@/components/svg/Logo'),
    AButton: () => import('ant-design-vue/lib/button'),
    AAvatar: () => import('ant-design-vue/lib/avatar'),
    APopover: () => import('ant-design-vue/lib/popover')
  },
  data() {
    return {
      rows: [
        { text: 'Choose Template' },
        { text: 'Customize' },
        { text: 'Preview' },
        { text: 'Deploy' }
      ]
    }
  },
  methods: {
    onClickGithub() {
      window.open('https://github.com/GPortfolio/GPortfolio', '_blank')
    },
    onClickTryIt() {
      this.$router.push({ name: 'steps' })
    }
  }
}
</script>

<style lang="scss" scoped>
$primary: #3f3d56;

$laptopL: 1440px;
$table: 768px;
$mobileL: 425px;
$mobileM: 375px;

header {
  .header__wrap {
    display: flex;
    justify-content: center;
    margin: 0 50px;
  }
  .left {
    display: flex;
    max-width: 750px;
    margin-right: 50px;
    flex-direction: column;
    align-items: flex-start;
    margin-top: 100px;
    h1 {
      color: $primary;
      font-weight: 900;
      font-size: 3rem;
      margin-bottom: 0;
    }
    p {
      font-size: 1.5rem;
      color: #949494;
    }
    button {
      width: 200px;
      height: 45px;
      border-radius: 10px;
      font-weight: bold;
      margin: 0 10px 20px;
    }
    .actions {
      justify-content: center;
    }
  }
  .right {
    svg {
      width: 700px;
    }
  }
}

.home-page {
  &:before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: url(https://gw.alipayobjects.com/zos/rmsportal/QepZGBUWvzoyvsBnyLJd.svg) no-repeat;
    background-position-x: right;
    background-size: 400px;
    transform: rotateX(0) rotateY(180deg);
    z-index: -1;
  }
}

.contributors {
  display: flex;
  flex-wrap: wrap;
  margin-top: 50px;
  > * {
    margin: 0 3px 5px;
  }
}

.ant-avatar {
  opacity: .7;
  transition: .2s opacity;
  &:hover {
    opacity: 1;
  }
}

.logo {
  width: 250px;
  padding: 30px;
  margin: 0 auto;
}

.wave {
  margin-top: -8%;
}

.max--width {
  @extend .center;
  max-width: 2560px;
}

.center {
  margin-left: auto;
  margin-right: auto;
}

.text--center {
  text-align: center;
}

@media only screen and (max-width: $laptopL) {
  header {
    .header__wrap {
      display: block;
      margin: 0 auto;
      max-width: 900px;
      text-align: center;
    }
    .left {
      align-items: center;
      margin: 0 0 50px;
      max-width: none;
      button {
        margin-bottom: 20px;
        &:last-child {
          margin-bottom: 0;
        }
      }
    }
    .right {
      svg {
        max-width: 650px;
        width: 85%;
        height: 100%;
      }
    }
  }

  .home-page {
    > div, > header {
      margin: 0 30px;
    }
  }

  .contributors {
    justify-content: center;
    margin-top: 30px;
  }

  .logo {
    width: 70%;
    max-width: 350px;
    padding: 20px 0;
  }

  .wave {
    margin-top: 0;
  }
}

@media only screen and (max-width: $table) {
  header .left {
    h1 {
      font-size: 2.3rem;
    }
    p {
      font-size: 1.2rem;
    }
  }

  .home-page:before {
    content: none;
  }
}

@media only screen and (max-width: $mobileM) {
  header .left p {
    margin-top: 15px;
  }
}
</style>
