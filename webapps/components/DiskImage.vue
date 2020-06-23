<template>
  <div>
    <img class="img" v-for="i in accele" src="/Accele.png" :height="height" />
    <img class="img" v-for="i in blastv" src="/BlastV.png" :height="height" />
    <img class="img" v-for="i in blasth" src="/BlastH.png" :height="height" />
    <img class="img" v-for="i in charge" src="/Charge.png" :height="height" />
  </div>
</template>

<script>

export default{
  props: {
    accele: { type: Number, default: 0 },
    blastv: { type: Number, default: 0 },
    blasth: { type: Number, default: 0 },
    charge: { type: Number, default: 0 },
    size: { type: String, default: "middium" },
  },
  computed: {
    height () {
      let h = "32px"
      switch (this.size){
        case "middium":
          h = "32px";
      }
      return h;
    }
  },
}
</script>

