<template>
  <div>
    <v-icon v-for="i in maxRare" color="yellow accent-3" :small="small" :middium="middium" :large="large" :xlarge="x-large">mdi-hexagram</v-icon>
    <v-icon v-for="i in minRare" color="gray darken-3" :small="small" :middium="middium" :large="large" :xlarge="x-large">mdi-hexagram</v-icon>
  </div>
</template>

<script>
export default{
  props: {
    maxRare: { type: Number, default: 0 },
    minRare: { type: Number, default: 0 },
    size: { type: String, default: "" },
  },
  computed: {
    small () {
      return (this.size == 'small');
    },
    middium () {
      return (this.size == 'middium');
    },
    large () {
      return (this.size == 'large');
    },
    xlarge () {
      return (this.size == 'x-large');
    },
  },
}
</script>


