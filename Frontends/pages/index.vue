<template>
  <v-app>
    <v-data-table
      :headers="headers"
      :items="magicalGirls"
      :items-per-page="5">
    </v-data-table>
  </v-app>
</template>

<script>
import { mapState } from 'vuex'

export default {
  computed: {
    ...mapState(['magicalGirls'])
  },
  async fetch ({ store, params }) {
    await store.dispatch('fetchMagicalGirls');
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

