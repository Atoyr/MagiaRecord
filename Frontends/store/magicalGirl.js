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
  attributes : [ ],
  magicalGirlFilter : {
    attributes: [],
    types: [],
    mental: []
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
  }
}

export const getters = {
  attributes(state) {
    return state.attributes;
  },
  magicalGirls(state) {
    return state.magicalGirls.filter(magicalGirl => {
        console.log(state.magicalGirlFilter.attributes)
        return state.magicalGirlFilter.attributes.length === 0 || state.magicalGirlFilter.attributes.includes(magicalGirl.attribute)
      }).filter(magicalGirl => {
        return state.magicalGirlFilter.types.length === 0 || state.magicalGirlFilter.types.includes(magicalGirl.type)
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
  }
}
