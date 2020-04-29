<template>
  <v-card class="px-4 py-2" color="#f5f5f5">
    <v-container fluid>
      <v-row>
        <v-col cols=12 md=8>
          <AttributeFilter />
        </v-col>
        <v-col cols=12 md=4>
          <Mentalstrengthening />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=12 md=8>
          <TypeFilter />
        </v-col>
        <v-col cols=12 md=4>
          <DiskFilter />
        </v-col>
      </v-row>
      <v-row>
      </v-row>
    </v-container>
  </v-card>
</template>

<script>
import AttributeFilter from  '@/components/Filter/AttributeFilter.vue'
import Mentalstrengthening from  '@/components/Filter/Mentalstrengthening.vue'
import TypeFilter from  '@/components/Filter/TypeFilter.vue'
import DiskFilter from  '@/components/Filter/DiskFilter.vue'

export default {
  components : {
    AttributeFilter,
    Mentalstrengthening,
    TypeFilter,
    DiskFilter
  },
}
</script>


