<template>
  <v-app class="app">
    <v-card
      class="toolbar-card"
      color="grey lighten-4"
      flat
      tile
    >
      <v-toolbar dense flat class="toolbar">
        <v-app-bar-nav-icon>
        </v-app-bar-nav-icon>
        <v-spacer></v-spacer>
        <v-btn icon @click.stop="dialog = true">
          <v-icon>mdi-filter-outline</v-icon>
        </v-btn>
        <v-dialog v-model="dialog">
            <MagicalGirlFilter />
        </v-dialog>
      </v-toolbar>
    </v-card>
    <v-container fluid>
      <nuxt />
    </v-container>
  </v-app>
</template>

<script>
import { mapState } from 'vuex'
import MagicalGirlFilter from  '@/components/Filter/MagicalGirlFilter.vue'

export default {
  components : {
    MagicalGirlFilter
  },
  data() {
    return {
      dialog: false
    }
  }
}

</script>


<style>
.app {
  margin: 0 auto;
  background: inherit !important;
}

.container {
  margin: 0;
}

.toolbar-card{
  margin:0;
  padding:0;
}

.toolbar {
  background: linear-gradient(#f399a3,#e95591,#e62a7f);
  margin:0;
  padding:0;
}

</style>
