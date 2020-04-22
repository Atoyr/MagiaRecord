<template>
    <v-container fluid>
      <v-row class="vrow" no-gutters>
        <v-col col=12 md=8 >
          <div class="label-left">
          精神強化
          </div>
        </v-col>
        <v-col col=12 md=4>
          <div class="label-right">
            <v-checkbox v-model="selectAll" :readonly="selectAll" class="check-all">
              <template v-slot:label>
                <p>all</p>
              </template>
            </v-checkbox>
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

<style>
.label-left {
  display: flex;
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
  align-content:stretch;
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

.label-right {
  display: flex;
  position: relative;
  height: 32px;/*リボンの高さ*/
  line-height: 32px;/*リボンの高さ*/
  text-align: center;
  padding: 0 30px;/*横の大きさ*/
  margin: 0;
  font-size: 18px;/*文字の大きさ*/
  color: #FFF;/*文字色*/
  box-sizing: border-box; 
  background: #bbb;
  align-content:stretch;
}
.label-right:after {
  position: absolute;
  content: '';
  width: 0px;
  height: 0px;
  top: 0;
  right: -10px;
  border-width: 16px 0px 16px 10px;
  border-color: transparent transparent transparent #bbb ;
  border-style: solid;
}

.vrow{
  height:30px;
}
.check-all{
 margin: 0;
 padding: 0;
}

</style>



