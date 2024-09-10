<template>
  <v-footer class="footer">
    <v-col class="text-center" cols="12">
      <hr class="gradient-line" />
      <span class="footer-text" ref="textElement"></span>
    </v-col>
  </v-footer>
</template>

<script>
export default {
  data () {
    return {
      currentYear: new Date().getFullYear(),
      launchDate: new Date('2024-09-10T00:00:00'),
      timeRunning: ''
    }
  },
  mounted () {
    this.updateTimeRunning()
    setInterval(this.updateTimeRunning, 1000)
    this.typeEffect()
  },
  methods: {
    updateTimeRunning () {
      const now = new Date()
      const timeDifference = now - this.launchDate

      const days = Math.floor(timeDifference / (1000 * 60 * 60 * 24))
      const hours = Math.floor((timeDifference / (1000 * 60 * 60)) % 24)
      const minutes = Math.floor((timeDifference / (1000 * 60)) % 60)
      const seconds = Math.floor((timeDifference / 1000) % 60)

      this.timeRunning = `${days}天 ${hours}小时 ${minutes}分钟 ${seconds}秒`
    },
    typeEffect () {
      const text = `${this.currentYear} — JiangBlog\n网站已运行 ${this.timeRunning}`
      let index = 0
      const textElement = this.$refs.textElement
      textElement.innerHTML = ''

      const type = () => {
        if (index < text.length) {
          textElement.innerHTML += text[index] === '\n' ? '<br>' : text[index]
          index++
          setTimeout(type, 100)
        } else {
          index = 0
          setTimeout(() => {
            textElement.innerHTML = ''
            type()
          }, 2000)
        }
      }

      type()
    }
  }
}
</script>

<style scoped>
.footer {
  background-color: #ffffff;
  padding: 20px 0;
  transition: background-color 0.3s ease;
}

.footer-text {
  font-size: 18px;
  color: #333;
  transition: color 0.3s ease;
  display: inline-block;
  animation: colorChange 5s infinite;
}

@keyframes colorChange {
  0% { color: #61dafb; }
  25% { color: #ff6f61; }
  50% { color: #f3c623; }
  75% { color: #16c79a; }
  100% { color: #61dafb; }
}

.gradient-line {
  width: 80%;
  margin: 10px auto;
  height: 1px;
  background-image: linear-gradient(
    to right,
    rgba(255, 255, 255, 0),
    cyan,
    rgba(0, 0, 0, 0)
  );
}

strong {
  font-weight: bold;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
}
</style>
