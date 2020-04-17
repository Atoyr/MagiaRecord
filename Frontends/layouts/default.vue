<template>
  <v-app class="app">
    <v-card
      color="grey lighten-4"
      flat
      height="50px"
      tile
    >
    <v-toolbar dense class="toolbar">
      <v-app-bar-nav-icon>
      </v-app-bar-nav-icon>
      <v-spacer></v-spacer>
      <v-btn icon @click.stop="dialog = true">
        <v-icon>mdi-filter-outline</v-icon>
      </v-btn>
      <v-dialog v-model="dialog">
        <v-card>
          <h2>foooooo</h2>
        </v-card>
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

export default {
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

.toolbar {
  background: linear-gradient(#f399a3,#e95591,#e62a7f);
}

</style>
