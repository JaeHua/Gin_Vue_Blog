<template>
  <v-col>
    <h1 class="catename text-center">
      <span class="bracket">{</span>
      {{ this.category }}
      <span class="bracket">}</span>
    </h1>

    <v-card
      class="mt-12 mb-7 article-card"
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
      <v-pagination
        :total-visible="7"
        v-model="queryParam.pagenum"
        :length="Math.ceil(this.total/queryParam.pagesize)"
        @input="getArtList()"
      ></v-pagination>
    </div>
  </v-col>
</template>

<script>
export default {
  props: ['id'],
  data () {
    return {
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      ArtList: [],
      total: 0,
      category: ''
    }
  },
  created () {
    this.getArtList()
  },
  watch: {
    '$route.params.id': 'getArtList'
  },
  methods: {
    async getArtList () {
      const { data: res } = await this.$http(`category/article/${this.id}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      this.ArtList = res.data
      this.total = res.total
      this.getCate()
    },
    async getCate () {
      const { data: res } = await this.$http(`category/${this.id}`)
      this.category = res.data.name
    },
    formatDate (dateString) {
      const options = { year: 'numeric', month: 'long', day: 'numeric' }
      return new Date(dateString).toLocaleDateString(undefined, options)
    }
  }
}
</script>

<style scoped>
.v-col {
  display: flex;
  justify-content: center;
}

.article-card {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s, box-shadow 0.2s;
  width: 60%;
  margin: 0 auto;
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
  height: 300px;
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
  font-size: 24px;
  margin: 8px 0;
  font-family: "Segoe UI", Tahoma, Geneva, Verdana, sans-serif;
  color: #333;
  text-shadow: 1px 1px 4px rgba(0, 0, 0, 0.3);
}

.article-description {
  font-size: 16px;
  color: #555;
  margin: 16px 0;
}

.dots {
  position: absolute;
  bottom: 16px;
  right: 16px;
}

.catename {
  font-size: 3rem;
  font-weight: bold;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  color: #333;
  text-shadow:
    1px 1px 2px rgba(0, 0, 0, 0.15),
    2px 2px 4px rgba(0, 0, 0, 0.1),
    3px 3px 8px rgba(0, 0, 0, 0.05);
}

.bracket {
  font-size: 3rem;
  font-weight: bold;
  color: #f5be0b;
  text-shadow:
    1px 1px 2px rgba(0, 0, 0, 0.15),
    2px 2px 4px rgba(0, 0, 0, 0.1),
    3px 3px 8px rgba(0, 0, 0, 0.05);
  display: inline-block;
  margin: 0 0.5rem;
}

@media (max-width: 1024px) {
  .article-card {
    width: 70%;
  }
}

@media (max-width: 768px) {
  .article-card {
    width: 80%;
  }
  .image-container,
  .article-image {
    height: 250px;
  }
}

@media (max-width: 480px) {
  .article-card {
    width: 95%;
  }
  .image-container,
  .article-image {
    height: 200px;
  }
  .article-title {
    font-size: 20px;
  }
  .article-description {
    font-size: 14px;
  }
}
</style>
