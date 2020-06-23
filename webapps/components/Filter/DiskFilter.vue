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
        <v-range-slider v-model="diskFilter.acceleRange" min="1" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :accele=1></DiskImage>
          </template>
          <template v-slot:thumb-label="props">
            <DiskImage :accele="props.value" :height=10></DiskImage>
          </template>
        </v-range-slider>
      </v-row>
      <v-row>
        <v-range-slider v-model="diskFilter.blastvRange" min="1" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :blastv=1></DiskImage>
          </template>
          <template v-slot:thumb-label="props">
            <DiskImage :blastv="props.value" :height=10></DiskImage>
          </template>
        </v-range-slider>
      </v-row>
      <v-row>
        <v-range-slider v-model="diskFilter.blasthRange" min="1" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :blasth=1></DiskImage>
          </template>
          <template v-slot:thumb-label="props">
            <DiskImage :blasth="props.value" :height=10></DiskImage>
          </template>
        </v-range-slider>
      </v-row>
      <v-row>
        <v-range-slider v-model="diskFilter.chargeRange" min="1" max="3" step="1" ticks="always" tick-size="4">
          <template v-slot:label>
            <DiskImage :charge=1></DiskImage>
          </template>
          <template v-slot:thumb-label="props">
            <DiskImage :charge="props.value" :height=10></DiskImage>
          </template>
        </v-range-slider>
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
        if(this.diskFilter.acceleRange[0] == 1 
          && this.diskFilter.acceleRange[1] == 3
          && this.diskFilter.blastvRange[0] == 1
          && this.diskFilter.blastvRange[1] == 3
          && this.diskFilter.blasthRange[0] == 1
          && this.diskFilter.blasthRange[1] == 3
          && this.diskFilter.chargeRange[0] == 1
          && this.diskFilter.chargeRange[1] == 3 ){
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        if (value) {
          this.diskFilter.acceleRange = [1, 3];
          this.diskFilter.blastvRange = [1, 3];
          this.diskFilter.blasthRange = [1, 3];
          this.diskFilter.chargeRange = [1, 3];
        }
      }
    },
  },
  data()  {
    return {
      diskFilter: {
        acceleRange: [1, 3],
        blastvRange: [1, 3],
        blasthRange: [1, 3],
        chargeRange: [1, 3],
      }
    }
  },
  watch: {
    diskFilter: {
      handler: function(newValue, oldValue) {
        this.$store.dispatch('magicalGirl/applyDisksRangeFilter',{disksRange: this.diskFilter});
      },
      deep: true,
    }
  }
}
</script>

<style>
.check-all{
  margin: 0;
}
</style>

