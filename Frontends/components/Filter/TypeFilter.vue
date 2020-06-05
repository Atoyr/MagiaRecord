<template>
    <v-container fluid>
      <v-row no-gutters>
        <v-col>
          <div class="flex">
            <div class="label-left">
            タイプ
            </div>
            <div class="label-right">
              <v-checkbox v-model="selectAll" :readonly="selectAll" class="check-all" label="全て選択">
              </v-checkbox>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <template v-for="(type,i) in types">
          <v-col :key="type.key" class="flex-wrap" cols=4  lg=3>
            <v-checkbox color="pink lighten-1"  v-model="checkTypes" :value="type.key" :label="type.name">
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
        if(this.checkTypes.length == 0 ){
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        if (value) {
          this.checkTypes = [];
        }
      }
    }
  },
  data()  {
    return {
      checkTypes: [],
      types: [
        {id:1 ,key: "balance" , name: "バランス"},
        {id:2 ,key: "attack" , name: "アタック"},
        {id:3 ,key: "difense" , name: "ディフェンス"},
        {id:4 ,key: "magia" , name: "マギア"},
        {id:5 ,key: "support" , name: "サポート"},
        {id:6 ,key: "heel" , name: "ヒール"},
        {id:7 ,key: "ultimet" , name: "アルティメット"},
        {id:8 ,key: "ring magia" , name: "円環マギア"},
        {id:9 ,key: "ring support" , name: "円環サポート"},
        ]

    }
  },
  watch: {
    checkTypes: {
      handler: function(newValue, oldValue) {
        this.$store.dispatch('magicalGirl/applyTypeFilter',{types: this.checkTypes});
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

