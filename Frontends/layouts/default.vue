<template>
  <v-app class="app">
    <v-card
      color="grey lighten-4"
      flat
      height="50px"
      tile
    >
    <v-toolbar dense >
      <v-app-bar-nav-icon>
      </v-app-bar-nav-icon>
    </v-toolbar>
    </v-card>
    <v-container fluid>
      <nuxt />
    </v-container>
  </v-app>
</template>

<style>
.app {
  margin: 0 auto;
  background-color: inherit !important;
}

.container {
  margin: 0;
}

</style>
