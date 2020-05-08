const MAGICAL_GIRLS = [
        {
          id: 1,
          name: '環 いろは',
          attribute: 'light',
          type: '回復',
          disk: {
            accele: 2,
            blasth: 1,
            blastv: 1,
            charge: 1,
          },
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 2,
          name: '七海 やちよ',
          attribute: 'aqua',
          type: 'バランス',
          disk: {
            accele: 1,
            blasth: 3,
            blastv: 0,
            charge: 1,
          },
          hp: 2500,
          attack: 400,
          deffence: 300,
        },
        {
          id: 3,
          name: '由依 鶴乃',
          attribute: 'flame',
          type: 'マギア',
          disk: {
            accele: 2,
            blasth: 1,
            blastv: 1,
            charge: 1,
          },
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 4,
          name: '深月 フェリシア',
          attribute: 'dark',
          type: 'アタック',
          disk: {
            accele: 2,
            blasth: 1,
            blastv: 1,
            charge: 1,
          },
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 5,
          name: '二葉 さな',
          attribute: 'forest',
          type: 'ディフェンス',
          disk: {
            accele: 2,
            blasth: 1,
            blastv: 1,
            charge: 1,
          },
          hp: 2400,
          attack: 400,
          deffence: 400,
        },
      ]

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
  magicalGirls(state) {
    return state.magicalGirls.filter(magicalGirl => {
        return state.magicalGirlFilter.attributes.length === 0 || state.magicalGirlFilter.attributes.includes(magicalGirl.attribute)
      }).filter(magicalGirl => {
        return state.magicalGirlFilter.types.length === 0 || state.magicalGirlFilter.types.includes(magicalGirl.type)
      }).filter(magicalGirl => {
        let result = true;
        if (!(state.magicalGirlFilter.disksRange.acceleRange[0] == 0 && state.magicalGirlFilter.disksRange.acceleRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.acceleRange[0] <= magicalGirl.disk.accele 
            && magicalGirl.disk.accele <= state.magicalGirlFilter.disksRange.acceleRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.blastvRange[0] == 0 && state.magicalGirlFilter.disksRange.blastvRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.blastvRange[0] <= magicalGirl.disk.blastv 
            && magicalGirl.disk.blastv <= state.magicalGirlFilter.disksRange.blastvRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.blasthRange[0] == 0 && state.magicalGirlFilter.disksRange.blasthRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.blasthRange[0] <= magicalGirl.disk.blasth 
            && magicalGirl.disk.blasth <= state.magicalGirlFilter.disksRange.blasthRange[1];
        }
        if (result && !(state.magicalGirlFilter.disksRange.chargeRange[0] == 0 && state.magicalGirlFilter.disksRange.chargeRange[1] == 3)) {
          result = state.magicalGirlFilter.disksRange.chargeRange[0] <= magicalGirl.disk.charge 
            && magicalGirl.disk.charge <= state.magicalGirlFilter.disksRange.chargeRange[1];
        }
        return result;
      })
  }
}

export const actions = {
  fetchMagicalGirls({commit}) {
    commit('updateMagicalGirls', MAGICAL_GIRLS);
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
