<template>
    <div id="editor">
        <!-- @imgDel="imgDel" -->
      <mavon-editor v-model="content" ref="md" >
        <template v-slot:left-toolbar-before>
          <button
            type="button"
            aria-hidden="true"
            title="颜色"
            class="op-icon"
            style="position: relative"
            v-if="queryParams.channel != 2"
          >
            <i class="el-icon-s-open" />
            <el-color-picker
              v-model="color1"
              style="opacity: 0; position: absolute; top: 0; left: 0; width: 100%; height: 100%;"
              @change="activeChange"
            />
          </button>
        </template>
      </mavon-editor>
    </div>
  </template>

<script>
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import { ColorPicker } from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import URL from '@/plugin/axios'
// import axios from 'axios'

export default {
  name: 'Editor',
  components: { mavonEditor, 'el-color-picker': ColorPicker },
  props: {
    value: {
      type: String,
      default: ''
    },
    queryParams: {
      type: Object,
      default: () => ({})
    }
  },
  data () {
    return {
      content: this.value, // 初始内容从父组件传入
      color1: '',
      upUrl: URL + '/upload',
      imgs: []
    }
  },
  watch: {
    content (newVal) {
      this.$emit('input', newVal)
    },
    value (newVal) {
      this.content = newVal
    }
  },
  methods: {
    activeChange (e) {
      const insertText = {
        prefix: `<font color="${e}">`,
        subfix: '</font>',
        str: ''
      }

      const $vm = this.$refs.md
      $vm.insertText($vm.getTextareaDom(), insertText)
    }

  }
}
</script>

  <style>
  .editor {
    margin: auto;
    width: 80%;
    height: 800px;
  }
  .op-icon {
    cursor: pointer;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;
    border-radius: 4px;
    border: 1px solid #dcdfe6;
    background-color: #fff;
    transition: all 0.3s;
  }
  .op-icon:hover {
    border-color: #c6e2ff;
  }
  </style>
