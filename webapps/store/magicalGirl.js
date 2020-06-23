import firebase from '~/plugins/firebase'

const ATTRIBUTE_COUNT = [
  {
    attribute: 'fire',
    count: 1
  },
  {
    attribute: 'aqua',
    count: 1
  },
  {
    attribute: 'folest',
    count: 1
  },
  {
    attribute: 'dark',
    count: 1
  },
  {
    attribute: 'light',
    count: 1
  },
]

export const state = () => ({
  magicalGirls : [],
  isLoading : false,
  attributes : [],
  magicalGirlFilter : {
    attributes: [],
    types: [],
    mental: [],
    disksRange: {
      acceleRange: [0,3],
      blastvRange: [0,3],
      blasthRange: [0,3],
      chargeRange: [0,3]
    }
  }
})

export const mutations = {
  clearMagicalGirls(state) {
    state.magicalGirls = [];
  },
  updateMagicalGirls(state, magicalGirls) {
    state.magicalGirls = magicalGirls;
  },
  updateIsLooading(state, isLoading) {
    state.isLoading = isLoading;
  },
  updateAttributes(state, attributes) {
    state.attributes = attributes;
  },
  updateAttributeFilter(state, attributes) {
    state.magicalGirlFilter.attributes = attributes
  },
  updateTypeFilter(state, types) {
    state.magicalGirlFilter.types = types;
  },
  updateDisksRangeFilter(state, disksRange) {
    state.magicalGirlFilter.disksRange = disksRange;
  }
}

export const getters = {
  attributes(state) {
    return state.attributes;
  },
  isLoading(state) {
    return state.isLoading
  },
  magicalGirls(state) {
    return state.magicalGirls.filter(magicalGirl => {
        return state.magicalGirlFilter.attributes.length === 0 || state.magicalGirlFilter.attributes.includes(magicalGirl.Attribute.Key)
      }).filter(magicalGirl => {
        return state.magicalGirlFilter.types.length === 0 || state.magicalGirlFilter.types.includes(magicalGirl.Type.Key)
      }).filter(magicalGirl => {
        let result = true;
        if (!(state.magicalGirlFilter.disksRange.acceleRange[0] == 1 && state.magicalGirlFilter.disksRange.acceleRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.acceleRange[0] <= magicalGirl.Disk.Accele 
            && magicalGirl.Disk.Accele <= state.magicalGirlFilter.disksRange.acceleRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.blastvRange[0] == 1 && state.magicalGirlFilter.disksRange.blastvRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.blastvRange[0] <= magicalGirl.Disk.Blastv 
            && magicalGirl.Disk.Blastv <= state.magicalGirlFilter.disksRange.blastvRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.blasthRange[0] == 1 && state.magicalGirlFilter.disksRange.blasthRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.blasthRange[0] <= magicalGirl.Disk.Blasth 
            && magicalGirl.Disk.Blasth <= state.magicalGirlFilter.disksRange.blasthRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.chargeRange[0] == 1 && state.magicalGirlFilter.disksRange.chargeRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.chargeRange[0] <= magicalGirl.Disk.Charge 
            && magicalGirl.Disk.Charge <= state.magicalGirlFilter.disksRange.chargeRange[1];
        }
        return result;
      })
  }
}

const db = firebase.firestore()
const fsMagicalGirls = db.collection("private/v1/magicalGirls");

export const actions = {
  fetchMagicalGirls({commit}) {
    fsMagicalGirls.get()
      .then(res => {
        commit('updateIsLooading', true);
        let mg = []
        res.forEach((doc) => {
          mg.push(doc.data())
        })
        console.log(mg.length)
        commit('updateMagicalGirls', mg);
        commit('updateIsLooading', false);
      })
    .catch(err => {
      console.log("error : " + err)
    })
  },
  fetchAttributes({commit}) {
    commit('updateAttributes', ATTRIBUTE_COUNT);
  },
  applyAttributeFilter({commit}, {attributes}) {
    commit('updateAttributeFilter', attributes);
  },
  applyTypeFilter({commit}, {types}) {
    commit('updateTypeFilter', types);
  },
  applyDisksRangeFilter({commit}, {disksRange}) {
    let dr = {
      acceleRange: [],
      blastvRange: [],
      blasthRange: [],
      chargeRange: [],
    }
    dr.acceleRange[0] = disksRange.acceleRange[0]
    dr.acceleRange[1] = disksRange.acceleRange[1]
    dr.blastvRange[0] = disksRange.blastvRange[0]
    dr.blastvRange[1] = disksRange.blastvRange[1]
    dr.blasthRange[0] = disksRange.blasthRange[0]
    dr.blasthRange[1] = disksRange.blasthRange[1]
    dr.chargeRange[0] = disksRange.chargeRange[0]
    dr.chargeRange[1] = disksRange.chargeRange[1]
    commit('updateDisksRangeFilter', dr);
  }
}
