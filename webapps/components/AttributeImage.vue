<template>
  <img v-if="attributeLower==='flame'" class="img" src="/Flame.png" height=32/>
  <img v-else-if="attributeLower==='aqua'" class="img" src="/Aqua.png" height=32/>
  <img v-else-if="attributeLower==='forest'" class="img" src="/Forest.png" height=32/>
  <img v-else-if="attributeLower==='dark'" class="img" src="/Dark.png" height=32/>
  <img v-else-if="attributeLower==='light'" class="img" src="/Light.png" height=32/>
  <img v-else="attributeLower==='void'" class="img" src="/Void.png" height=32/>
</template>

<script>

export default{
  props: {
    attribute: String
  },
  computed: {
    attributeLower: function() {
      return this.attribute.toLowerCase();
    }
  },
  data() {
    return {
    };
  }
}
</script>
