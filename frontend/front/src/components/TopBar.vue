<template>
<div class="navv">
    <v-app-bar app color="primary" flat>
        <!-- logo -->

        <v-avatar  class="mx-12" size="40" color="white">
          <v-img @click="gobackend" src="../assets/logo.png" alt=""></v-img>
        </v-avatar>

        <v-container class="py-0 fill-height">
            <v-btn @click="$router.push('/')" text color="white">首页</v-btn>
            <v-btn v-for="item in cateList" :key="item.id" text color="white" @click="$router.push(`Category/${item.ID}`)">{{item.name}}</v-btn>
        </v-container>
        <v-spacer></v-spacer>
                   <v-responsive max-width color="white">
            <v-text-field dense flat hide-details rounded dark solo-inverted     append-icon="mdi-magnify">
      </v-text-field>
    </v-responsive>
    </v-app-bar>
</div>
</template>
<script>

export default {
  data () {
    return {
      cateList: []
    }
  },
  created () {
    this.GetCateList()
  },
  methods: {
    // 获取分类
    async GetCateList () {
      const { data: res } = await this.$http.get('category')
      this.cateList = res.data
    },
    gobackend () {
      window.location.href = 'http://localhost:3000/admin'
    }
  }
}
</script>
<style scoped>
.navv{
  z-index: 1000;
}
</style>
