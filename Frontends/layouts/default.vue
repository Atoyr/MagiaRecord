<template>
  <v-app class="app">
      <v-app-bar 
        dense 
        flat 
        class="toolbar">
        <v-app-bar-nav-icon>
        </v-app-bar-nav-icon>
        <v-toolbar-title>マギアレコード データベース</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click.stop="dialog = true">
          <v-icon>mdi-filter-outline</v-icon>
        </v-btn>
        <v-dialog v-model="dialog">
            <MagicalGirlFilter />
        </v-dialog>
      </v-app-bar>
    <v-container fluid>
      <nuxt />
    </v-container>
  </v-app>
</template>

<script>
import { mapState } from 'vuex'
import MagicalGirlFilter from  '@/components/filter/MagicalGirlFilter.vue'

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
