<template>
  <div class="page home-page">
    <div class="center text--center">
      <logo-svg class="logo" />
    </div>
    <header class="max--width">
      <div class="left">
        <h1>Create Your Portfolio Online</h1>
        <p>Upload your data from social networks or fill in manually</p>
        <div class="ant-row-flex">
          <a-button
            type="primary"
            size="large"
          >
            Try it Free
          </a-button>
          <a-button
            size="large"
            icon="github"
          >
            Github
          </a-button>
        </div>
      </div>
      <div class="right">
        <hire-svg />
      </div>
    </header>
    <wave-svg class="wave" />
    <section class="section-steps">

    </section>
  </div>
</template>

<script>
export default {
  components: {
    HireSvg: () => import('@/components/svg/Hire'),
    WaveSvg: () => import('@/components/svg/Wave'),
    LogoSvg: () => import('@/components/svg/Logo'),
    AButton: () => import('ant-design-vue/lib/button')
  }
}
</script>

<style lang="scss" scoped>
header {
  display: flex;
  justify-content: center;
  .left {
    display: flex;
    min-width: 700px;
    margin-right: 70px;
    flex-direction: column;
    align-items: flex-start;
    margin-top: 130px;
    h1 {
      color: #3f3d56;
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
      &:first-child {
        margin-right: 20px;
      }
    }
  }
  .right {
    svg {
      width: 800px;
    }
  }
}

.home-page {
  &:before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 400px;
    height: 100%;
    background: url(https://gw.alipayobjects.com/zos/rmsportal/QepZGBUWvzoyvsBnyLJd.svg) no-repeat;
    transform: rotateX(0) rotateY(180deg);
    z-index: -1;
  }
}

.section-steps {
  margin-top: -80px;
  background: #3f3d56;
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

.logo {
  width: 250px;
  padding: 30px;
  margin: 0 auto;
}

.wave {
  margin-top: -80px;
}
</style>
