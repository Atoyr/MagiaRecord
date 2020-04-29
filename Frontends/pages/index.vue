<template>
  <div>
    <h2>マギアレコード データベース</h2>
    <v-row>
    <AttributeChartCard :size=3 />
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
    ...mapState(['magicalGirls'])
  },
  async fetch ({ store, params }) {
    console.log(store)
    await store.dispatch('magicalGirl/fetchMagicalGirls');
    await store.dispatch('magicalGirl/fetchAttributes');
  },
  data() {
    return {
      headers: [
        {
          text: 'ID',
          align: ' d-none',
          sortable: true,
          value: 'id',
          width: 0,
        },
        {
          text: 'name',
          align: 'start',
          sortable: true,
          value: 'name',
        },
        { text: 'Attribute', value: 'attribute' },
        { text: 'Type', value: 'type' },
        { text: 'HP', value: 'hp' },
        { text: 'Attack', value: 'attack' },
        { text: 'Deffence', value: 'deffence' },
      ],
    }
  },
}
</script>

