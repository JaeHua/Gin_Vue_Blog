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

    <div class="toc-navigation">
      <ul>
        <li v-for="toc in tocs" :key="toc.id" @click="jumpToToc(toc)" :style="{ paddingLeft: (toc.level - 2) * 20 + 'px' }">
          {{ toc.title }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import Clipboard from 'clipboard'
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import { scrollToToc } from '../plugins/catalog.js'

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
          return '/markdown/markdown.css'
        }
      },
      tocs: [],
      activeToc: null
    }
  },
  created () {
    this.getArt()
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeDestroy () {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    async getArt () {
      const { data: res } = await this.$http(`article/info/${this.id}`)
      this.article = res.data
      this.addCopyBtn()
      this.extractTocs()
    },
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
            const code = e.target.parentElement.children[0]
            const codeText = code.innerText
            that.shareCopy(codeText)
          })
          element.appendChild(btn)
        })
      })
    },
    shareCopy (codeText) {
      const clipboard = new Clipboard('.code-copy', {
        text: function () {
          return codeText
        }
      })
      clipboard.on('success', e => {
        this.$toast.success('复制成功', {
          timeout: 1000
        })
        clipboard.destroy()
      })
      clipboard.on('error', e => {
        clipboard.destroy()
      })
    },
    extractTocs () {
      this.$nextTick(() => {
        const headers = this.$refs.markdown.$el.querySelectorAll('h2, h3, h4')
        const tocList = Array.from(headers).map(header => {
          const id = header.id || header.textContent.trim().replace(/\s+/g, '-')
          header.id = id
          return {
            id: id,
            title: header.textContent,
            level: parseInt(header.tagName.charAt(1)) // 获取级别
          }
        })
        this.tocs = tocList.slice(0, Math.ceil(tocList.length / 2)) // 截取一半的目录
      })
    },
    jumpToToc (toc) {
      this.activeToc = toc.id
      scrollToToc(toc)
    },
    handleScroll () {
      const toc = document.querySelector('.toc-navigation')
      if (window.scrollY > 400) {
        toc.style.top = '120px'
      } else {
        toc.style.top = '500px'
      }
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
  object-fit: cover;
}
.article-content {
  max-width: 800px;
  margin: 0 auto;
  border-radius: 8px;
}
.article-title {
  font-size: 2.5em;
  color: #333;
  text-align: center;
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
::v-deep .hljs {
  background-color: #151618 !important;
}
::v-deep pre.hljs,
::v-deep code.hljs {
  background-color: #151618 !important;
}
.toc-navigation {
  position: fixed;
  top: 500px;
  right: 10px;
  padding: 15px;
  border-radius: 10px;
  border-left: 3px solid #007bff;
  transition: top 0.3s, box-shadow 0.3s;
}
.toc-navigation ul {
  list-style: none;
  padding: 0;
  margin: 0;
}
.toc-navigation li {
  cursor: pointer;
  margin: 10px 0;
  color: #007bff;
  font-family: 'Arial', sans-serif;
  font-size: 14px;
  transition: color 0.3s, transform 0.3s;
}
.toc-navigation li:hover {
  color: #0056b3;
  transform: translateX(5px);
}
@media (max-width: 768px) {
  .toc-navigation {
    display: none;
  }
}
</style>
