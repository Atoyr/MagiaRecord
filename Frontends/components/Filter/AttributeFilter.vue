<template>
    <v-container fluid>
      <v-row no-gutters>
        <v-col>
          <div class="flex">
            <div class="label-left">
            属性
            </div>
            <div class="label-right">
              <v-checkbox v-model="selectAll" :readonly="selectAll" class="check-all" label="全て選択">
              </v-checkbox>
            </div>
          </div>
        </v-col>
      </v-row>
      <v-row>
        <template v-for="(attribute,i) in attributes">
          <v-col :key="attribute">
            <v-checkbox color="pink lighten-1"  v-model="checkAttributes" :value="attribute">
              <template v-slot:label>
                <AttributeImage :attribute="attribute" />
              </template>
            </v-checkbox>
          </v-col>
        </template>
      </v-row>
    </v-container>
</template>

<script>
import AttributeImage from '@/components/AttributeImage.vue'
export default {
  components : {
    AttributeImage
  },
  computed: {
    selectAll: {
      get: function() {
        if(this.checkAttributes.length == 0 ){
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        if (value) {
          this.checkAttributes = [];
        }
      }
    }
  },
  data()  {
    return {
      checkAttributes: [],
      attributes: ["flame","aqua","forest","dark","light","void"]
    }
  }
}
</script>

<style>
.check-all{
  margin: 0;
}
</style>
