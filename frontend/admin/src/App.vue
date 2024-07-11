<template>
  <a-config-provider :locale="locale">
    <div id="app">
      <!-- 解决编辑文章和写文章用的同一个组件的缓存问题（不想编辑了跳转到写文章发现数据是之前编辑的） -->
      <router-view :key="$route.path"/>
    </div>
  </a-config-provider>
</template>
<script>
import zhCN from 'ant-design-vue/lib/locale-provider/zh_CN'
export default {
  data () {
    return {
      locale: zhCN
    }
  }
}
</script>
<style>
</style>
