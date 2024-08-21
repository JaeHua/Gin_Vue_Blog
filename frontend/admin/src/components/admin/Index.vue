<template>
    <div>
      <h3>欢迎来到Blog管理后台</h3>
      <div ref="chart" style="width: 500px; height: 400px;"></div>
    </div>
  </template>

<script>
import * as echarts from 'echarts'
import dayjs from 'dayjs'

export default {
  data () {
    return {
      chartInstance: null
    }
  },
  methods: {
    async getArtInfo () {
      try {
        const { data: res } = await this.$http.get('article')
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        if (!res.data) {
          return this.$message.error('No data returned from API.')
        }

        const dateCounts = {}
        res.data.forEach(article => {
          const date = dayjs(article.CreatedAt).format('YYYY-MM-DD')
          if (!dateCounts[date]) {
            dateCounts[date] = 0
          }
          dateCounts[date]++
        })

        const categories = Object.keys(dateCounts).sort((a, b) => new Date(a) - new Date(b))
        const data = categories.map(date => dateCounts[date])

        this.updateChart(categories, data)
      } catch (error) {
        console.error('获取文章信息出错:', error)
        this.$message.error('获取文章信息出错')
      }
    },
    updateChart (categories, data) {
      const options = {
        title: {
          text: '每日文章数量'
        },
        tooltip: {
          trigger: 'axis'
        },
        xAxis: {
          type: 'category',
          data: categories
        },
        yAxis: {
          minInterval: 1
        },
        dataZoom: [
          {
            type: 'slider',
            start: 0,
            end: Math.min((5 / categories.length) * 100, 100)
          }
        ],
        series: [
          {
            name: '文章数量',
            type: 'bar',
            data: data,
            barWidth: '40%'
          },
          {
            name: '文章数量',
            type: 'line',
            data: data
          }
        ]
      }
      this.chartInstance.setOption(options)
    }
  },
  mounted () {
    this.chartInstance = echarts.init(this.$refs.chart)
    this.getArtInfo()
  }
}
</script>
