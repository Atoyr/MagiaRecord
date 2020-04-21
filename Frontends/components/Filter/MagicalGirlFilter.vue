<template>
  <v-card class="px-4 py-2" color="#f5f5f5">
    <v-container fluid>
      <v-row>
        <v-col cols=12 md=8>
          <AttributeFilter />
        </v-col>
        <v-col cols=12 md=4>
          <h2>haa</h2>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import AttributeFilter from  '@/components/Filter/AttributeFilter.vue'

export default {
  components : {
    AttributeFilter
  },
}
</script>


