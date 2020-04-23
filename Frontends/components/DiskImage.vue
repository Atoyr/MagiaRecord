<template>
  <div>
    <img v-for="index in accele" class="img" src="/Accele.png" height=32/>
    <img class="img" v-for="i in blastv" src="/BlastV.png" height=32/>
    <img class="img" v-for="i in blasth" src="/BlastH.png" height=32/>
    <img class="img" v-for="i in charge" :key="i" src="/Charge.png" height=32/>
  </div>
</template>

<script>

export default{
  props: {
    accele: { type: Number, default: 0 },
    blastv: { type: Number, default: 0 },
    blasth: { type: Number, default: 0 },
    charge: { type: Number, default: 0 }
  },
  data() {
    return {
    };
  }
}
</script>

