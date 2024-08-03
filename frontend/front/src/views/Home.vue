<template>
  <v-app app>
    <TopBar></TopBar>
    <v-main class="white lighten-3">
      <v-container fluid class="full">
        <v-row v-if="!isDetailPage">
          <v-col cols="12">
            <v-sheet min-height="50vh" rounded="lg" class="pa-0">
              <Nav></Nav>
            </v-sheet>
          </v-col>
        </v-row>
        <!-- <v-divider class="mt-10"></v-divider> -->
        <v-row justify="center" >
          <v-col cols="12" :md="!isDetailPage ? 8 : 0">
            <div v-if="!isDetailPage">
           <v-icon class="ml-11 mb-0"> {{'mdi-animation-outline'}}</v-icon><span style="color:grey;">Discovery</span>
            </div>
          <v-divider ></v-divider>
            <v-sheet min-height="50vh" rounded="lg">
              <router-view></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer></Footer>
  </v-app>
</template>

<script>
import TopBar from '../components/TopBar'
import Footer from '../components/Footer'
import Nav from '../components/Nav'

export default {
  components: { TopBar, Footer, Nav },
  computed: {
    isDetailPage () {
      return this.$route.path.startsWith('/detail/')
    }
  }
}
</script>

<style scoped>
.full{
  height: 100%;
    margin: 0;
    padding: 0;
}
.v-main {
  padding-top: 0;
  padding-bottom: 0;
}
</style>
