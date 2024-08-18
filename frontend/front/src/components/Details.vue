<template>
  <div>
    <v-img class="bg" :src="article.img"></v-img>
    <div>
      <div class="article-content">
        <h1 class="article-title">{{ article.title }}</h1>
        <p class="article-desc">概要：{{ article.desc }}</p>
        <div class="ml-n6 markdown-body">
          <mavonEditor
            class="md"
            :value="this.article.content"
            :subfield="false"
            :defaultOpen="'preview'"
            :toolbarsFlag="false"
            :editable="false"
            :external-link="externalLink"
            :scrollStyle="true"
            :codeStyle="codeStyle"
            :ishljs="true"
            box-shadow-style="0 0 0 0 rgba(0,0,0,0)"
            previewBackground="white"
            ref="markdown"
          ></mavonEditor>
        </div>
      </div>
    </div>
    <v-snackbar v-model="showSnackbar"  variant="outlined" color="success"  :timeout="3000" top>
      复制成功
    </v-snackbar>
  </div>
</template>

<script>
import Clipboard from 'clipboard'
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
export default {
  props: ['id'],
  components: { mavonEditor },
  data () {
    return {
      article: {
        title: '',
        desc: '',
        content: '',
        img: '',
        category: ''
      },
      codeStyle: 'atom-one-dark',
      externalLink: {
        markdown_css: function () {
          // 这是你的markdown css文件路径
          return '/markdown/markdown.css'
        }
      },
      tocs: [],
      showSnackbar: false
    }
  },
  created () {
    this.getArt()
  },
  methods: {
    async getArt () {
      const { data: res } = await this.$http(`article/info/${this.id}`)
      this.article = res.data
      this.addCopyBtn()
    },
    // 动态添加复制按钮以及复制方法（获取到文档content成功的时候调用此方法）
    addCopyBtn () {
      const that = this
      this.$nextTick(() => {
        const content = document.getElementsByClassName('v-show-content')[0]
        const precodes = content.getElementsByTagName('pre')
        const arr = Array.from(precodes)
        arr.forEach((element, index) => {
          const btn = document.createElement('div')
          btn.setAttribute('class', 'code-copy')
          btn.addEventListener('click', function (e) {
            const code = e.target.parentElement.children[0] // code 标签
            const codeText = code.innerText
            // 复制函数
            that.shareCopy(codeText)
          })
          element.appendChild(btn)
        })
      })
    },
    // 复制代码函数
    shareCopy (codeText) {
      // const that = this
      const clipboard = new Clipboard('.code-copy', {
        text: function () {
          return codeText
        }
      })
      clipboard.on('success', e => {
      // 复制成功
        this.showSnackbar = true
        clipboard.destroy()
      })
      clipboard.on('error', e => {
      // 复制失败
        clipboard.destroy()
      })
    }
  }
}
</script>

<style scoped>
.bg {
  width: 100%;
  height: 400px;
  margin: 0;
  padding: 0;
  object-fit: cover; /* 确保图片完全覆盖容器 */
}
.article-content {
  max-width: 800px;
  margin: 0 auto;
  border-radius: 8px; /* 添加圆角 */
}
.article-title {
  font-size: 2.5em;
  color: #333; /* 标题使用深色 */
  text-align: center; /* 居中显示 */
}
.article-desc {
  color: #797979;
  font-family: 'Arial', sans-serif;
  font-size: 16px;
  margin-top: 10px;
  line-height: 30px;
}
.article-body {
  font-size: 1em;
  line-height: 1.6;
  color: #797979;
}
/* 使用深度选择器 */
::v-deep .hljs {
  background-color: #151618 !important;
}

::v-deep pre.hljs,
::v-deep code.hljs {
  background-color: #151618 !important;
}

</style>
