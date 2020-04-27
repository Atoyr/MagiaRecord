<template>
    <v-container fluid>
      <v-row class="vrow" no-gutters>
        <v-col>
          <div class="flex">
            <div class="label-left">
            精神強化
            </div>
            <div class="label-right">
              <v-checkbox v-model="selectAll" :readonly="selectAll" class="check-all" label="全て選択">
              </v-checkbox>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <template v-for="(awake,i) in awakes">
          <v-col :key="i">
            <v-checkbox color="pink lighten-1"  v-model="checkAwakes" :value="i" :label="awake">
            </v-checkbox>
          </v-col>
        </template>
      </v-row>
    </v-container>
</template>

<script>
export default {
  computed: {
    selectAll: {
      get: function() {
        if(this.checkAwakes.length == 0 ){
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        if (value) {
          this.checkAwakes = [];
        }
      }
    }
  },
  data()  {
    return {
      checkAwakes: [],
      awakes: ["解放済み","未解放"]
    }
  }
}
</script>


