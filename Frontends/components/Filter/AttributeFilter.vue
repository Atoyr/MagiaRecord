<template>
    <v-container fluid>
      <v-row>
        <v-col col=12 md=8>
          <div class="label-left">
          属性
          </div>
        </v-col>
        <v-col col=12 md=4>
          <v-checkbox v-model="selectAll">
            <template v-slot:label>
              <p>all</p>
            </template>
          </v-checkbox>
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
        if(this.checkAttributes.length == 0 || this.checkAttributes.length == this.attributes.length){
          this.checkAttributes = [];
          return true;
        }else{
          return false;
        }
      },
      set: function(value) {
        let checkArray = [];
        if (value) {
            this.attributes.forEach(function (attribute) {
                checkArray.push(attribute);
            });
        }
        this.checkAttributes = checkArray;
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
.label-left {
  display: inline-block;
  position: relative;
  height: 32px;/*リボンの高さ*/
  line-height: 32px;/*リボンの高さ*/
  text-align: center;
  padding: 0 30px;/*横の大きさ*/
  margin: 0;
  font-size: 18px;/*文字の大きさ*/
  color: #FFF;/*文字色*/
  box-sizing: border-box; 
  background: #e62a7f;
}
.label-left:before {
  position: absolute;
  content: '';
  width: 0px;
  height: 0px;
  top: 0;
  left: -10px;
  border-width: 16px 10px 16px 0px;
  border-color: transparent #e62a7f transparent transparent;
  border-style: solid;
}
</style>


