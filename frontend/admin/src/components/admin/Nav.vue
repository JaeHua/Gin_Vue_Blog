<template>
    <a-layout-sider breakpoint="lg" v-model="collapsed">
      <div class="log" @click="navigateToBlog">
        <span>{{collapsed?'Blog':'My Blog'}}</span>
      </div>
      <a-menu theme="dark" mode="inline" @click="goToPage">
        <a-menu-item key="index">
          <a-icon type="dashboard" />
          <span>仪表盘</span>
        </a-menu-item>
        <a-sub-menu>
            <span slot="title"><a-icon type="file" /><span>文章管理</span></span>
            <a-menu-item key="addart"><a-icon type="form" /><span>写文章</span></a-menu-item>
            <a-menu-item key="artlist"><a-icon type="snippets" /><span>文章列表</span></a-menu-item>
        </a-sub-menu>
        <a-menu-item key="catelist"><a-icon type="book" /><span>分类列表</span></a-menu-item>
        <a-menu-item key="userlist"><a-icon type="user" /><span>用户列表</span></a-menu-item>
        <a-menu-item key="profile"><a-icon type="setting" /><span>用户设置</span></a-menu-item>
      </a-menu>
    </a-layout-sider>
  </template>
<script>
export default {
  data () {
    return {
      collapsed: false
    }
  },
  methods: {
    goToPage (item) {
      // 连续点击菜单就不会报错了(catch了)
      this.$router.push('/' + item.key).catch(err => err)
    },
    navigateToBlog () {
      window.location.href = 'http://localhost:8080/'
    }
  }
}
</script>
<style scoped>
.log {
    height: 32px;
    margin: 16px;
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 17px;
}
</style>
