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
          ></mavonEditor>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
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
      }
    }
  },
  created () {
    this.getArt()
  },
  methods: {
    async getArt () {
      const { data: res } = await this.$http(`article/info/${this.id}`)
      this.article = res.data
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
