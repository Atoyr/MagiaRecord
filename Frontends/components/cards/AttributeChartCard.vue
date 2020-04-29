<template>
  <v-col cols=12 :md="size" >
    <v-card class="px-4 py-2" color="#f5f5f5">
      <v-card-title >
        属性割合
      </v-card-title>
      <no-ssr>
        <doughnut-chart :chart-data="attributeChartData" :options="chartOptions" :height=200 :width=200 /> 
      </no-ssr>
    </v-card>
  </v-col>
</template>

<script>
import { mapState } from 'vuex';
import colors from 'vuetify/es5/util/colors';

export default {
  methods: {
  },
  computed: {
    ...mapState({
      attributeChartData (state)  {
        let data = state.magicalGirl.attributes;
        const count = [];
        const labels = [];
        const datacolors = [];
      data.forEach(a => {
        count.push(a.count)
        labels.push(a.attribute)
        switch(a.attribute) {
          case 'fire' :
            datacolors.push(colors.red.lighten1)
            break;
          case 'aqua' :
            datacolors.push(colors.blue.lighten1)
            break;
          case 'folest' :
            datacolors.push(colors.green.lighten1)
            break;
          case 'light' :
            datacolors.push(colors.yellow.lighten1)
            break;
          case 'dark' :
            datacolors.push(colors.purple.lighten1)
            break;
        }
      });
        return {
          datasets: [
            {
              data: count,
              backgroundColor: datacolors,
              borderColor: 'transparent' 
            },
          ],
          labels: labels
        }
      }
    })},
  props: {
    size: Number
  },
  data() {
    return {
      chartOptions: {
        padding: {
          left: 50,
          right: 50,
          top: 50,
          bottom: 50
        },
        legend: {
          display: true,
          position: 'right'
        },
        title: {

        },
      }
    }
  },
}
</script>
