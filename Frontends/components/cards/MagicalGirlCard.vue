<template>
  <v-col :cols=12 :md="size" class="Card">
    <v-card class="px-4 py-2" color="#f5f5f5">
      <v-card-title >
        魔法少女
      </v-card-title>
      <v-data-table
        dense
        :headers="headers"
        :items="magicalGirls"
        locale="ja-JP"
        :loading="isLoading" 
        loading-text="Loading... Please wait"
        :items-per-page="10">
        <template v-slot:item.Attribute="{ item }">
          <AttributeImage :attribute="item.Attribute.Key"></AttributeImage>
        </template>
        <template v-slot:item.Type="{ item }">
          <p>{{Type.Name}}</p>
        </template>
        <template v-slot:item.Disk="{ item }">
          <DiskImage 
            :accele="item.Disk.Accele" 
            :blastv="item.Disk.Blastv"
            :blasth="item.Disk.Blasth"
            :charge="item.Disk.Charge"></DiskImage>
        </template>
      </v-data-table>
    </v-card>
  </v-col>
</template>

<script>
import { mapGetters} from 'vuex'
import AttributeImage from '@/components/AttributeImage.vue'
import DiskImage from '@/components/DiskImage.vue'

export default {
  components : {
    AttributeImage,
    DiskImage
  },
  computed: {
    ...mapGetters("magicalGirl",["magicalGirls","isLoading"])
  },
  props: {
    size: Number
  },
  data() {
    return {
      headers: [
        {
          text: 'ID',
          align: ' d-none',
          sortable: true,
          value: 'Key',
          width: 0,
        },
        {
          text: '名前',
          align: 'start',
          sortable: true,
          value: 'Name',
        },
        {
          text: 'バージョン',
          align: 'start',
          sortable: true,
          value: 'Version',
        },
        {
          text: '限定',
          align: 'start',
          sortable: true,
          value: 'IsLimited',
        },
        { text: '属性', value: 'Attribute' },
        { text: 'タイプ', value: 'Type.Name' },
        { text: 'Disk', value: 'Disk' },
        { text: 'HP', value: 'Status.Hp' },
        { text: 'Attack', value: 'Status.Attack' },
        { text: 'Difense', value: 'Status.Difense' },
        // { text: '覚醒HP', value: 'AwakeRate.Hp' },
        // { text: '覚醒Attack', value: 'AwakeRate.Attack' },
        // { text: '覚醒Difense', value: 'AwakeRate.Difense' },
        // { text: '覚醒Accele', value: 'AwakeRate.Accele' },
        // { text: '覚醒Blast', value: 'AwakeRate.Blast' },
        // { text: '覚醒Charge', value: 'AwakeRate.Charge' },
      ],
    }
  },
}
</script>

