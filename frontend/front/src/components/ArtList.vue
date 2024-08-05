<template>
    <v-col >
      <v-card
        class="ma-3 mx-15 mb-7 article-card"
        rounded="lg"
        outlined
        v-for="(item, index) in ArtList"
        :key="item.ID"
        @click="$router.push(`/detail/${item.ID}`)"

      >
        <v-row no-gutters>
          <v-col :cols="6" :class="index % 2 === 0 ? 'image-col' : 'content-col'">
            <div class="image-container" v-if="index % 2 === 0">
              <img :src="item.img" alt="Article Image" class="article-image" />
            </div>
            <div class="article-content" v-else>
              <div class="article-header">
                <div class="article-category">
                  <v-icon small>mdi-tag</v-icon>
                  <span>{{ item.Category.name }}</span>
                </div>
                <div class="article-date">
                  <v-icon small>mdi-calendar</v-icon>
                  <span>{{ formatDate(item.CreatedAt) }}</span>
                </div>
              </div>
              <h3 class="article-title">{{ item.title }}</h3>
              <p class="article-description">{{ item.desc }}</p>
              <div class="dots">
                <v-icon large>mdi-dots-horizontal</v-icon>
              </div>
            </div>
          </v-col>
          <v-col :cols="6" :class="index % 2 === 0 ? 'content-col' : 'image-col'">
            <div class="article-content" v-if="index % 2 === 0">
              <div class="article-header">
                <div class="article-category">
                  <v-icon small>mdi-tag</v-icon>
                  <span>{{ item.Category.name }}</span>
                </div>
                <div class="article-date">
                  <v-icon small>mdi-calendar</v-icon>
                  <span>{{ formatDate(item.CreatedAt) }}</span>
                </div>
              </div>
              <h3 class="article-title">{{ item.title }}</h3>
              <p class="article-description">{{ item.desc }}</p>
              <div class="dots">
                <v-icon large>mdi-dots-horizontal</v-icon>
              </div>
            </div>
            <div class="image-container" v-else>
              <img :src="item.img" alt="Article Image" class="article-image" />
            </div>
          </v-col>
        </v-row>
      </v-card>
      <div class="text-center" v-if="this.total">
      <v-pagination   :total-visible="7" v-model="queryParam.pagenum" :length="Math.ceil(this.total/queryParam.pagesize)" @input="getArtList()"></v-pagination>
    </div>
    </v-col>
  </template>

<script>
export default {
  data () {
    return {
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      ArtList: [],
      total: 0
    }
  },
  created () {
    this.getArtList()
  },
  methods: {
    async getArtList () {
      const { data: res } = await this.$http('article', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.ArtList = res.data
      this.total = res.total
    },
    formatDate (dateString) {
      const options = { year: 'numeric', month: 'long', day: 'numeric' }
      return new Date(dateString).toLocaleDateString(undefined, options)
    }
  }
}
</script>

  <style scoped>
  .vcol {
    display: flex;
    justify-content: center; /* Center children horizontally */

  }

  .article-card {
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    transition: transform 0.2s, box-shadow 0.2s;
    width: 85%;
  }
  .article-card:hover {
    transform: translateY(-5px);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  }
  .image-col {
    padding: 0;
  }
  .image-container {
    width: 100%;
    height: 300px; /* Fixed height */
    overflow: hidden;
    border-radius: 8px 0 0 8px;
  }
  .article-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  .content-col {
    padding: 0;
  }
  .article-content {
    padding: 16px;
  }
  .article-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: 14px;
    color: #757575;
  }
  .article-category {
    display: flex;
    align-items: center;
    gap: 4px;
  }
  .article-date {
    display: flex;
    align-items: center;
    gap: 4px;
  }
  .article-title {
    font-size: 24px; /* Increased font size */
    margin: 8px 0;
    font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif; /* Changed font family */
    color: #333; /* Darker color for better contrast */
    text-shadow: 1px 1px 4px rgba(0, 0, 0, 0.3); /* Added text shadow */
  }
  .article-description {
    font-size: 16px;
    color: #555;
    margin: 16px 0;
  }
  .dots {
    position: absolute; /* Positioned absolutely within the content */
    bottom: 16px; /* Distance from the bottom */
    right: 16px; /* Distance from the right */
  }
  </style>
