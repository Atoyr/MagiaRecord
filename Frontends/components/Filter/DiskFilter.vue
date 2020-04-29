<template>
    <v-container fluid>
      <v-row no-gutters>
        <v-col>
          <div class="flex">
            <div class="label-left">
            ディスク
            </div>
            <div class="label-right">
              <v-checkbox v-model="selectAll" :readonly="selectAll" class="check-all" label="全て選択">
              </v-checkbox>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <v-slider v-model="acceleSize" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :accele=1></DiskImage>
          </template>
        </v-slider>
      </v-row>
      <v-row>
        <v-slider v-model="blastvSize" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :blastv=1></DiskImage>
          </template>
        </v-slider>
      </v-row>
      <v-row>
        <v-slider v-model="blasthSize" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :blasth=1></DiskImage>
          </template>
        </v-slider>
      </v-row>
      <v-row>
        <v-slider v-model="chargeSize" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :charge=1></DiskImage>
          </template>
        </v-slider>
      </v-row>
    </v-container>
</template>

<script>
import DiskImage from '@/components/DiskImage.vue'
export default {
  components : {
    DiskImage
  },
  computed: {
    selectAll: {
      get: function() {
        if(this.acceleSize == 0 && this.blastvSize == 0 && this.blasthSize == 0 && this.chargeSize == 0){
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        if (value) {
          this.acceleSize = 0;
          this.blastvSize = 0;
          this.blasthSize = 0;
          this.chargeSize = 0;
        }
      }
    }
  },
  data()  {
    return {
      acceleSize: 0,
      blastvSize: 0,
      blasthSize: 0,
      chargeSize: 0,
    }
  },
  watch: {
    checkAttributes: {
      handler: function(newValue, oldValue) {
        this.$store.dispatch('magicalGirl/applyAttributeFilter',{attributes: this.checkAttributes});
      },
      deep: true
    }
  }
}
</script>

<style>
.check-all{
  margin: 0;
}
</style>

