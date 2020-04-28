const MAGICAL_GIRLS = [
        {
          id: 1,
          name: 'tamaki iroha',
          attribute: 'light',
          type: 'heel',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 2,
          name: 'nanami yachiyo',
          attribute: 'watter',
          type: 'balance',
          hp: 2500,
          attack: 400,
          deffence: 300,
        },
        {
          id: 3,
          name: 'yui tsuruno',
          attribute: 'fire',
          type: 'magia',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 4,
          name: 'mitsuki ferishia',
          attribute: 'dark',
          type: 'attack',
          hp: 2400,
          attack: 400,
          deffence: 300,
        },
        {
          id: 5,
          name: 'futaba sana',
          attribute: 'folest',
          type: 'deffence',
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
