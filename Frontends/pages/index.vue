<template>
  <div>
    <v-row>
    <MagicalGirlCard :size=12 />
    <Filter />
    </v-row>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import MagicalGirlCard from  '@/components/cards/MagicalGirlCard.vue'
import AttributeChartCard from  '@/components/cards/AttributeChartCard.vue'

export default {
  components : {
    MagicalGirlCard,
    AttributeChartCard,
  },
  computed: {
  },
  async fetch ({ store, params }) {
    await store.dispatch('magicalGirl/fetchMagicalGirls');
    await store.dispatch('magicalGirl/fetchAttributes');
  },
  data() {
    return {
    }
  },
}
</script>

