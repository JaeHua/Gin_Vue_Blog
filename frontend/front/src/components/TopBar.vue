<template>
<div>
    <v-app-bar app color="primary" flat>
        <!-- logo -->

        <v-avatar size="40" color="white"></v-avatar>

        <v-container class="py-0 fill-height">
            <v-btn  text color="white">首页</v-btn>
            <v-btn v-for="item in cateList" :key="item.id" text color="white">{{item.name}}</v-btn>
        </v-container>
        <v-spacer></v-spacer>
    <v-responsive max-width color="white">
        <v-text-field dense flat hide-details rounded dark solo-inverted></v-text-field>
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
    }
  }
}
</script>
