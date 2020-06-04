<template>
  <div>
    <v-row>
    <MagicalGirlCard :size=12 />
    <AttributeChartCard :size=3 />
    <Filter />
    </v-row>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import MagicalGirlCard from  '@/components/cards/MagicalGirlCard.vue'
import AttributeChartCard from  '@/components/cards/AttributeChartCard.vue'
import firebase from '@/plugins/firebase'

export default {
  components : {
    MagicalGirlCard,
    AttributeChartCard,
  },
  computed: {
    ...mapState(['magicalGirls'])
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

