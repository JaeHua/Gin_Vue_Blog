<template>
    <span>
      <span
        v-for="(char, index) in displayedCharacters"
        :key="index"
        :style="{ color: getColor(index), animationDelay: `${index * speed}ms` }"
        class="fade-in"
      >
        {{ char }}
      </span>
    </span>
  </template>

<script>
export default {
  name: 'TypewriterEffect',
  props: {
    text: {
      type: String,
      required: true
    },
    speed: {
      type: Number,
      default: 100 // 调整速度为 100ms，速度变慢
    },
    colors: {
      type: Array,
      default: () => ['#FF5733', '#33FF57', '#3357FF', '#FF33A1', '#F8FF33']
    },
    loop: {
      type: Boolean,
      default: true
    },
    pauseTime: {
      type: Number,
      default: 1000 // Pause time between loops in milliseconds
    }
  },
  data () {
    return {
      displayedCharacters: [],
      currentIndex: 0
    }
  },
  mounted () {
    this.typeText()
  },
  methods: {
    typeText () {
      this.displayedCharacters = []
      this.currentIndex = 0

      const timer = setInterval(() => {
        if (this.currentIndex < this.text.length) {
          this.displayedCharacters.push(this.text.charAt(this.currentIndex))
          this.currentIndex++
        } else {
          clearInterval(timer)
          if (this.loop) {
            setTimeout(() => {
              this.resetText()
            }, this.pauseTime)
          }
        }
      }, this.speed)
    },
    resetText () {
      this.displayedCharacters = []
      this.currentIndex = 0
      this.typeText()
    },
    getColor (index) {
      return this.colors[index % this.colors.length]
    }
  }
}
</script>

  <style scoped>
  span {
    font-family: monospace;
    font-size: 24px;
  }

  .fade-in {
    opacity: 0;
    animation: fadeIn 0.5s forwards;
  }

  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
  </style>
