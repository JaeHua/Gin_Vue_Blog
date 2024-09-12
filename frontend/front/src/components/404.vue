<template>
    <div class="_404">
      <p @click="goHome" class="font-weight-bold">
        抱歉，页面未找到，<span>{{ countDown }}</span>s后自动跳转到首页
      </p>
      <img src="@/assets/404.jpg" alt="页面未找到">
    </div>
  </template>

<script>
export default {
  data () {
    return {
      countDown: 5
    }
  },
  methods: {
    goHome () {
      this.$router.push('/')
    },
    startCountDown () {
      this.timer = setInterval(() => {
        this.countDown--
        if (this.countDown === 0) {
          clearInterval(this.timer)
          this.$router.push('/')
        }
      }, 1000)
    }
  },
  mounted () {
    this.startCountDown()
    document.body.style.backgroundColor = '#dceebc'
  },
  beforeDestroy () {
    clearInterval(this.timer)
    document.body.style.backgroundColor = ''
  }
}
</script>

  <style scoped>
  ._404 {
    width: 30%;
    margin: 5rem auto;
  }

  ._404 img {
    width: 30rem;
  }

  ._404 a {
    color: #010101;
  }

  ._404 a:hover {
    color: skyblue;
  }
  </style>
