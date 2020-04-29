<template>
  <v-col :cols=12 :md="size" class="Card">
    <v-card class="px-4 py-2" color="#f5f5f5">
      <v-card-title >
        魔法少女
      </v-card-title>
      <v-data-table
        dense
        :headers="headers"
        :items="magicalGirls"
        :items-per-page="5">
      </v-data-table>
    </v-card>
  </v-col>
</template>

<script>
import { mapGetters} from 'vuex'

export default {
  computed: {
    ...mapGetters("magicalGirl",["magicalGirls"])
  },
  props: {
    size: Number
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
          text: '名前',
          align: 'start',
          sortable: true,
          value: 'name',
        },
        { text: '属性', value: 'attribute' },
        { text: 'タイプ', value: 'type' },
        { text: 'HP', value: 'hp' },
        { text: 'Attack', value: 'attack' },
        { text: 'Deffence', value: 'deffence' },
      ],
    }
  },
}
</script>

